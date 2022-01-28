package mattermostauthlayer

import (
	"database/sql"
	"encoding/json"

	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"

	"github.com/mattermost/focalboard/server/model"
	"github.com/mattermost/focalboard/server/services/store"
	"github.com/mattermost/focalboard/server/utils"

	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

const (
	sqliteDBType   = "sqlite3"
	postgresDBType = "postgres"
	mysqlDBType    = "mysql"

	directChannelType = "D"
)

var (
	errUnsupportedDatabaseError = errors.New("method is unsupported on current database. Supported databases are - MySQL and PostgreSQL")
)

type NotSupportedError struct {
	msg string
}

func (pe NotSupportedError) Error() string {
	return pe.msg
}

// Store represents the abstraction of the data storage.
type MattermostAuthLayer struct {
	store.Store
	dbType string
	mmDB   *sql.DB
	logger *mlog.Logger
}

// New creates a new SQL implementation of the store.
func New(dbType string, db *sql.DB, store store.Store, logger *mlog.Logger) (*MattermostAuthLayer, error) {
	layer := &MattermostAuthLayer{
		Store:  store,
		dbType: dbType,
		mmDB:   db,
		logger: logger,
	}

	return layer, nil
}

// Shutdown close the connection with the store.
func (s *MattermostAuthLayer) Shutdown() error {
	return s.Store.Shutdown()
}

func (s *MattermostAuthLayer) GetRegisteredUserCount() (int, error) {
	query := s.getQueryBuilder().
		Select("count(*)").
		From("Users").
		Where(sq.Eq{"deleteAt": 0})
	row := query.QueryRow()

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *MattermostAuthLayer) getUserByCondition(condition sq.Eq) (*model.User, error) {
	users, err := s.getUsersByCondition(condition)
	if err != nil {
		return nil, err
	}

	var user *model.User
	for _, u := range users {
		user = u
		break
	}

	return user, nil
}

func (s *MattermostAuthLayer) getUsersByCondition(condition sq.Eq) (map[string]*model.User, error) {
	query := s.getQueryBuilder().
		Select("id", "username", "email", "password", "MFASecret as mfa_secret", "AuthService as auth_service", "COALESCE(AuthData, '') as auth_data",
			"props", "CreateAt as create_at", "UpdateAt as update_at", "DeleteAt as delete_at").
		From("Users").
		Where(sq.Eq{"deleteAt": 0}).
		Where(condition)
	row, err := query.Query()
	if err != nil {
		return nil, err
	}

	users := map[string]*model.User{}

	for row.Next() {
		user := model.User{}

		var propsBytes []byte
		err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.MfaSecret, &user.AuthService,
			&user.AuthData, &propsBytes, &user.CreateAt, &user.UpdateAt, &user.DeleteAt)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(propsBytes, &user.Props)
		if err != nil {
			return nil, err
		}

		users[user.ID] = &user
	}

	return users, nil
}

func (s *MattermostAuthLayer) GetUserByID(userID string) (*model.User, error) {
	return s.getUserByCondition(sq.Eq{"id": userID})
}

func (s *MattermostAuthLayer) GetUserByEmail(email string) (*model.User, error) {
	return s.getUserByCondition(sq.Eq{"email": email})
}

func (s *MattermostAuthLayer) GetUserByUsername(username string) (*model.User, error) {
	return s.getUserByCondition(sq.Eq{"username": username})
}

func (s *MattermostAuthLayer) CreateUser(user *model.User) error {
	return NotSupportedError{"no user creation allowed from focalboard, create it using mattermost"}
}

func (s *MattermostAuthLayer) UpdateUser(user *model.User) error {
	return NotSupportedError{"no update allowed from focalboard, update it using mattermost"}
}

func (s *MattermostAuthLayer) UpdateUserPassword(username, password string) error {
	return NotSupportedError{"no update allowed from focalboard, update it using mattermost"}
}

func (s *MattermostAuthLayer) UpdateUserPasswordByID(userID, password string) error {
	return NotSupportedError{"no update allowed from focalboard, update it using mattermost"}
}

// GetActiveUserCount returns the number of users with active sessions within N seconds ago.
func (s *MattermostAuthLayer) GetActiveUserCount(updatedSecondsAgo int64) (int, error) {
	query := s.getQueryBuilder().
		Select("count(distinct userId)").
		From("Sessions").
		Where(sq.Gt{"LastActivityAt": utils.GetMillis() - utils.SecondsToMillis(updatedSecondsAgo)})

	row := query.QueryRow()

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *MattermostAuthLayer) GetSession(token string, expireTime int64) (*model.Session, error) {
	return nil, NotSupportedError{"sessions not used when using mattermost"}
}

func (s *MattermostAuthLayer) CreateSession(session *model.Session) error {
	return NotSupportedError{"no update allowed from focalboard, update it using mattermost"}
}

func (s *MattermostAuthLayer) RefreshSession(session *model.Session) error {
	return NotSupportedError{"no update allowed from focalboard, update it using mattermost"}
}

func (s *MattermostAuthLayer) UpdateSession(session *model.Session) error {
	return NotSupportedError{"no update allowed from focalboard, update it using mattermost"}
}

func (s *MattermostAuthLayer) DeleteSession(sessionID string) error {
	return NotSupportedError{"no update allowed from focalboard, update it using mattermost"}
}

func (s *MattermostAuthLayer) CleanUpSessions(expireTime int64) error {
	return NotSupportedError{"no update allowed from focalboard, update it using mattermost"}
}

func (s *MattermostAuthLayer) GetTeam(id string) (*model.Team, error) {
	if id == "0" {
		team := model.Team{
			ID:    id,
			Title: "",
		}

		return &team, nil
	}

	query := s.getQueryBuilder().
		Select("DisplayName").
		From("Teams").
		Where(sq.Eq{"ID": id})

	row := query.QueryRow()
	var displayName string
	err := row.Scan(&displayName)
	if err != nil {
		s.logger.Error("GetTeam scan error", mlog.Err(err))
		return nil, err
	}

	return &model.Team{ID: id, Title: displayName}, nil
}

// GetTeamsForUser retrieves all the teams that the user is a member of
func (s *MattermostAuthLayer) GetTeamsForUser(userID string) ([]*model.Team, error) {
	query := s.getQueryBuilder().
		Select("t.Id", "t.DisplayName").
		From("Teams as t").
		Join("TeamMembers as tm on t.Id=tm.TeamId").
		Where(sq.Eq{"tm.UserId": userID})

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer s.CloseRows(rows)

	teams := []*model.Team{}
	for rows.Next() {
		var team model.Team

		err := rows.Scan(
			&team.ID,
			&team.Title,
		)
		if err != nil {
			return nil, err
		}

		teams = append(teams, &team)
	}

	return teams, nil
}

func (s *MattermostAuthLayer) getQueryBuilder() sq.StatementBuilderType {
	builder := sq.StatementBuilder
	if s.dbType == postgresDBType || s.dbType == sqliteDBType {
		builder = builder.PlaceholderFormat(sq.Dollar)
	}

	return builder.RunWith(s.mmDB)
}

func (s *MattermostAuthLayer) GetUsersByTeam(teamID string) ([]*model.User, error) {
	query := s.getQueryBuilder().
		Select("u.id", "u.username", "u.email", "u.password", "u.MFASecret as mfa_secret", "u.AuthService as auth_service", "COALESCE(u.AuthData, '') as auth_data",
			"u.props", "u.CreateAt as create_at", "u.UpdateAt as update_at", "u.DeleteAt as delete_at").
		From("Users as u").
		Join("TeamMembers as tm ON tm.UserID = u.ID").
		Where(sq.Eq{"u.deleteAt": 0}).
		Where(sq.Eq{"tm.TeamId": teamID})

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer s.CloseRows(rows)

	users, err := s.usersFromRows(rows)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *MattermostAuthLayer) usersFromRows(rows *sql.Rows) ([]*model.User, error) {
	users := []*model.User{}

	for rows.Next() {
		var user model.User
		var propsBytes []byte

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.MfaSecret,
			&user.AuthService,
			&user.AuthData,
			&propsBytes,
			&user.CreateAt,
			&user.UpdateAt,
			&user.DeleteAt,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(propsBytes, &user.Props)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (s *MattermostAuthLayer) CloseRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		s.logger.Error("error closing MattermostAuthLayer row set", mlog.Err(err))
	}
}
