package server

import (
	"net"
	"net/http"
	"os"
	"runtime"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/mattermost/focalboard/server/api"
	"github.com/mattermost/focalboard/server/app"
	"github.com/mattermost/focalboard/server/auth"
	"github.com/mattermost/focalboard/server/context"
	appModel "github.com/mattermost/focalboard/server/model"
	"github.com/mattermost/focalboard/server/services/config"
	"github.com/mattermost/focalboard/server/services/mlog"
	"github.com/mattermost/focalboard/server/services/scheduler"
	"github.com/mattermost/focalboard/server/services/store"
	"github.com/mattermost/focalboard/server/services/store/sqlstore"
	"github.com/mattermost/focalboard/server/services/telemetry"
	"github.com/mattermost/focalboard/server/services/webhook"
	"github.com/mattermost/focalboard/server/web"
	"github.com/mattermost/focalboard/server/ws"

	"github.com/mattermost/mattermost-server/v5/services/filesstore"
	"github.com/mattermost/mattermost-server/v5/utils"
)

const (
	cleanupSessionTaskFrequency = 10 * time.Minute

	//nolint:gomnd
	minSessionExpiryTime = int64(60 * 60 * 24 * 31) // 31 days
)

type Server struct {
	config              *config.Configuration
	wsServer            *ws.Server
	webServer           *web.Server
	store               store.Store
	filesBackend        filesstore.FileBackend
	telemetry           *telemetry.Service
	logger              *mlog.Logger
	cleanUpSessionsTask *scheduler.ScheduledTask

	localRouter     *mux.Router
	localModeServer *http.Server
	api             *api.API
	appBuilder      func() *app.App
}

func New(cfg *config.Configuration, singleUserToken string, logger *mlog.Logger) (*Server, error) {
	db, err := sqlstore.New(cfg.DBType, cfg.DBConfigString, cfg.DBTablePrefix, logger)
	if err != nil {
		logger.Error("Unable to start the database", mlog.Err(err))
		return nil, err
	}

	authenticator := auth.New(cfg, db)

	wsServer := ws.NewServer(authenticator, singleUserToken, logger)

	filesBackendSettings := filesstore.FileBackendSettings{}
	filesBackendSettings.DriverName = "local"
	filesBackendSettings.Directory = cfg.FilesPath
	filesBackend, appErr := filesstore.NewFileBackend(filesBackendSettings)

	if appErr != nil {
		logger.Error("Unable to initialize the files storage", mlog.Err(appErr))

		return nil, errors.New("unable to initialize the files storage")
	}

	webhookClient := webhook.NewClient(cfg, logger)

	appBuilder := func() *app.App { return app.New(cfg, db, authenticator, wsServer, filesBackend, webhookClient, logger) }
	focalboardAPI := api.NewAPI(appBuilder, singleUserToken, cfg.AuthMode, logger)

	// Local router for admin APIs
	localRouter := mux.NewRouter()
	focalboardAPI.RegisterAdminRoutes(localRouter)

	// Init workspace
	if _, err = appBuilder().GetRootWorkspace(); err != nil {
		logger.Error("Unable to get root workspace", mlog.Err(err))
		return nil, err
	}

	webServer := web.NewServer(cfg.WebPath, cfg.ServerRoot, cfg.Port, cfg.UseSSL, cfg.LocalOnly, logger)
	webServer.AddRoutes(wsServer)
	webServer.AddRoutes(focalboardAPI)

	// Init telemetry
	settings, err := db.GetSystemSettings()
	if err != nil {
		return nil, err
	}

	telemetryID := settings["TelemetryID"]

	if len(telemetryID) == 0 {
		telemetryID = uuid.New().String()
		if err = db.SetSystemSetting("TelemetryID", uuid.New().String()); err != nil {
			return nil, err
		}
	}

	registeredUserCount, err := appBuilder().GetRegisteredUserCount()
	if err != nil {
		return nil, err
	}

	dailyActiveUsers, err := appBuilder().GetDailyActiveUsers()
	if err != nil {
		return nil, err
	}

	weeklyActiveUsers, err := appBuilder().GetWeeklyActiveUsers()
	if err != nil {
		return nil, err
	}

	monthlyActiveUsers, err := appBuilder().GetMonthlyActiveUsers()
	if err != nil {
		return nil, err
	}

	telemetryService := telemetry.New(telemetryID, logger.StdLogger(mlog.Telemetry))
	telemetryService.RegisterTracker("server", func() map[string]interface{} {
		return map[string]interface{}{
			"version":          appModel.CurrentVersion,
			"build_number":     appModel.BuildNumber,
			"build_hash":       appModel.BuildHash,
			"edition":          appModel.Edition,
			"operating_system": runtime.GOOS,
		}
	})
	telemetryService.RegisterTracker("config", func() map[string]interface{} {
		return map[string]interface{}{
			"serverRoot":  cfg.ServerRoot == config.DefaultServerRoot,
			"port":        cfg.Port == config.DefaultPort,
			"useSSL":      cfg.UseSSL,
			"dbType":      cfg.DBType,
			"single_user": len(singleUserToken) > 0,
		}
	})
	telemetryService.RegisterTracker("activity", func() map[string]interface{} {
		return map[string]interface{}{
			"registered_users":     registeredUserCount,
			"daily_active_users":   dailyActiveUsers,
			"weekly_active_users":  weeklyActiveUsers,
			"monthly_active_users": monthlyActiveUsers,
		}
	})

	server := Server{
		config:       cfg,
		wsServer:     wsServer,
		webServer:    webServer,
		store:        db,
		filesBackend: filesBackend,
		telemetry:    telemetryService,
		logger:       logger,
		localRouter:  localRouter,
		api:          focalboardAPI,
		appBuilder:   appBuilder,
	}

	server.initHandlers()

	return &server, nil
}

func (s *Server) Start() error {
	s.logger.Info("Server.Start")

	s.webServer.Start()

	if s.config.EnableLocalMode {
		if err := s.startLocalModeServer(); err != nil {
			return err
		}
	}

	s.cleanUpSessionsTask = scheduler.CreateRecurringTask("cleanUpSessions", func() {
		secondsAgo := minSessionExpiryTime
		if secondsAgo < s.config.SessionExpireTime {
			secondsAgo = s.config.SessionExpireTime
		}

		if err := s.store.CleanUpSessions(secondsAgo); err != nil {
			s.logger.Error("Unable to clean up the sessions", mlog.Err(err))
		}
	}, cleanupSessionTaskFrequency)

	if s.config.Telemetry {
		firstRun := utils.MillisFromTime(time.Now())
		s.telemetry.RunTelemetryJob(firstRun)
	}

	return nil
}

func (s *Server) Shutdown() error {
	if err := s.webServer.Shutdown(); err != nil {
		return err
	}

	s.stopLocalModeServer()

	if s.cleanUpSessionsTask != nil {
		s.cleanUpSessionsTask.Cancel()
	}

	if err := s.telemetry.Shutdown(); err != nil {
		s.logger.Warn("Error occurred when shutting down telemetry", mlog.Err(err))
	}

	defer s.logger.Info("Server.Shutdown")

	return s.store.Shutdown()
}

func (s *Server) Config() *config.Configuration {
	return s.config
}

func (s *Server) Logger() *mlog.Logger {
	return s.logger
}

// Local server

func (s *Server) startLocalModeServer() error {
	s.localModeServer = &http.Server{
		Handler:     s.localRouter,
		ConnContext: context.SetContextConn,
	}

	// TODO: Close and delete socket file on shutdown
	if err := syscall.Unlink(s.config.LocalModeSocketLocation); err != nil {
		s.logger.Error("Unable to unlink socket.", mlog.Err(err))
	}

	socket := s.config.LocalModeSocketLocation
	unixListener, err := net.Listen("unix", socket)
	if err != nil {
		return err
	}
	if err = os.Chmod(socket, 0600); err != nil {
		return err
	}

	go func() {
		s.logger.Info("Starting unix socket server")
		err = s.localModeServer.Serve(unixListener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("Error starting unix socket server", mlog.Err(err))
		}
	}()

	return nil
}

func (s *Server) stopLocalModeServer() {
	if s.localModeServer != nil {
		_ = s.localModeServer.Close()
		s.localModeServer = nil
	}
}

func (s *Server) GetRootRouter() *mux.Router {
	return s.webServer.Router()
}
