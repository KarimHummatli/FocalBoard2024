// Code generated by go-bindata. DO NOT EDIT.
// sources:
// migrations_files/000001_init.down.sql (30B)
// migrations_files/000001_init.up.sql (513B)
// migrations_files/000002_system_settings_table.down.sql (39B)
// migrations_files/000002_system_settings_table.up.sql (108B)
// migrations_files/000003_blocks_rootid.down.sql (51B)
// migrations_files/000003_blocks_rootid.up.sql (62B)
// migrations_files/000004_auth_table.down.sql (61B)
// migrations_files/000004_auth_table.up.sql (583B)
// migrations_files/000005_blocks_modifiedby.down.sql (55B)
// migrations_files/000005_blocks_modifiedby.up.sql (66B)
// migrations_files/000006_sharing_table.down.sql (31B)
// migrations_files/000006_sharing_table.up.sql (170B)
// migrations_files/000007_workspaces_table.down.sql (34B)
// migrations_files/000007_workspaces_table.up.sql (222B)
// migrations_files/000008_teams.down.sql (173B)
// migrations_files/000008_teams.up.sql (304B)
// migrations_files/000009_blocks_history.down.sql (97B)
// migrations_files/000009_blocks_history.up.sql (1.053kB)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __000001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xb6\xe6\x02\x04\x00\x00\xff\xff\x2d\x73\xd0\xe1\x1e\x00\x00\x00")

func _000001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initDownSql,
		"000001_init.down.sql",
	)
}

func _000001_initDownSql() (*asset, error) {
	bytes, err := _000001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.down.sql", size: 30, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc3, 0xa, 0x75, 0x74, 0x21, 0x4f, 0xed, 0x29, 0xf4, 0xfb, 0x14, 0x9a, 0xda, 0x7a, 0x6c, 0x3b, 0x58, 0x34, 0x7c, 0xcd, 0x48, 0xf4, 0x9, 0x5c, 0x96, 0xa1, 0xb9, 0xb2, 0x43, 0xfd, 0x76, 0xa2}}
	return a, nil
}

var __000001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xcd\x6e\xc2\x30\x10\x84\xcf\xc9\x53\xec\x25\x4a\x22\x05\x2e\x95\x38\xd0\x93\x01\x53\xd2\xe6\x07\x25\xa6\x40\x2f\x90\x92\x4d\x6b\xd5\x40\x88\x5d\xa9\xc8\xf2\xbb\x57\x41\x15\x45\xc0\xcd\x9e\x5d\x7d\x33\x3b\xc3\x8c\x12\x46\x81\x91\x41\x44\x21\x1c\x43\x92\x32\xa0\x8b\x30\x67\x39\x68\xdd\xad\x1b\xac\xf8\x8f\x31\xef\x62\xbf\xf9\x92\xe0\xd9\x16\x2f\xe1\x95\x64\xc3\x09\xc9\xbc\x87\x9e\x1f\xd8\x00\x00\x5a\xf3\x0a\xba\xf5\x5e\xaa\x8f\x06\xa5\x31\x7c\x27\xb1\x51\xab\x42\x01\x0b\x63\x9a\x33\x12\x4f\xd9\xdb\x89\x9c\xcc\xa2\x08\x46\x74\x4c\x66\x11\x83\x24\x9d\x7b\x7e\xa0\x35\xee\x4a\x63\x2e\x40\xf2\x20\xb8\xc2\x4b\xcc\x88\x30\xda\xa2\x6e\x18\x5e\xce\xb2\x71\x3b\xf1\x5c\x67\xd9\x71\xb6\x1d\xa7\x04\x67\xd2\x77\xe2\xbe\x53\xb9\x01\xb8\x49\x3a\x77\xfd\x7b\x1e\xdb\xa3\x3c\x88\x7b\x16\x5e\xcf\xbf\x9f\xb4\xf7\x8f\xb1\xea\xa2\xc1\x9d\x5a\x5d\x77\x61\xad\xe5\xe6\x13\xb7\xc5\x1a\x06\xe1\x53\x98\xb0\xc0\xb6\xd4\xb1\x46\x60\x74\x71\x7a\x73\x25\xce\x9f\x8a\xa3\x28\xe5\x4d\x75\xcf\x79\x9a\x68\x8d\x42\xa2\x31\xed\xe6\x9f\x65\x60\x5b\x9b\x06\x0b\x85\x6d\xd4\x33\xfc\xbb\x2e\xaf\xa5\x12\x05\x5e\x49\xd3\x2c\x8c\x49\xb6\x84\x17\xba\x04\x8f\x97\x01\x9c\x8f\xf6\x6d\xff\xd1\xfe\x0d\x00\x00\xff\xff\xe0\xe8\xf1\x7f\x01\x02\x00\x00")

func _000001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initUpSql,
		"000001_init.up.sql",
	)
}

func _000001_initUpSql() (*asset, error) {
	bytes, err := _000001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.up.sql", size: 513, mode: os.FileMode(0644), modTime: time.Unix(1618834938, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf8, 0x7a, 0x40, 0x6, 0x3d, 0x76, 0xd3, 0xfe, 0xb5, 0x2d, 0xdc, 0x8b, 0xf5, 0x11, 0x2, 0xd7, 0x77, 0x14, 0x2f, 0x7, 0xf0, 0xe9, 0x10, 0xe, 0xcb, 0xd2, 0x86, 0x83, 0x38, 0xac, 0x14, 0xd5}}
	return a, nil
}

var __000002_system_settings_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x2d\xae\x2c\x2e\x49\xcd\x8d\x2f\x4e\x2d\x29\xc9\xcc\x4b\x2f\xb6\xe6\x02\x04\x00\x00\xff\xff\xd2\x63\x5d\x39\x27\x00\x00\x00")

func _000002_system_settings_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_system_settings_tableDownSql,
		"000002_system_settings_table.down.sql",
	)
}

func _000002_system_settings_tableDownSql() (*asset, error) {
	bytes, err := _000002_system_settings_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_system_settings_table.down.sql", size: 39, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x28, 0x3, 0x29, 0x5d, 0x28, 0xad, 0x1c, 0x15, 0xd, 0x2e, 0x84, 0xc8, 0xda, 0x3a, 0xd, 0x2f, 0xd8, 0x1e, 0xd1, 0x7f, 0xa1, 0xa1, 0x8d, 0x8b, 0xf3, 0x71, 0xa5, 0xc, 0x2d, 0x3a, 0x61, 0x62}}
	return a, nil
}

var __000002_system_settings_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x2d\xae\x2c\x2e\x49\xcd\x8d\x2f\x4e\x2d\x29\xc9\xcc\x4b\x2f\x56\xd0\xe0\xe2\xcc\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x34\x30\xd0\xd4\xe1\xe2\x2c\x4b\xcc\x29\x4d\x55\x08\x71\x8d\x08\xd1\xe1\xe2\x0c\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\x54\xd0\xc8\x4c\xd1\xe4\xd2\xb4\xe6\x02\x04\x00\x00\xff\xff\xe4\x3d\xdb\x86\x6c\x00\x00\x00")

func _000002_system_settings_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_system_settings_tableUpSql,
		"000002_system_settings_table.up.sql",
	)
}

func _000002_system_settings_tableUpSql() (*asset, error) {
	bytes, err := _000002_system_settings_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_system_settings_table.up.sql", size: 108, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd3, 0xc5, 0xab, 0xe8, 0x30, 0x86, 0xbd, 0x6d, 0x70, 0x20, 0x8b, 0xc1, 0x7d, 0xc, 0xfa, 0xba, 0x3a, 0x71, 0x4b, 0xb7, 0x97, 0x5f, 0x39, 0x46, 0xa8, 0x50, 0x2f, 0x90, 0xfb, 0x1d, 0x45, 0xc}}
	return a, nil
}

var __000003_blocks_rootidDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xca\xcf\x2f\x89\xcf\x4c\xb1\xe6\x02\x04\x00\x00\xff\xff\x51\xe5\xe2\x3a\x33\x00\x00\x00")

func _000003_blocks_rootidDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_blocks_rootidDownSql,
		"000003_blocks_rootid.down.sql",
	)
}

func _000003_blocks_rootidDownSql() (*asset, error) {
	bytes, err := _000003_blocks_rootidDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_blocks_rootid.down.sql", size: 51, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdb, 0xfc, 0x77, 0x1b, 0xf7, 0x2e, 0x8c, 0xe9, 0x96, 0x14, 0x46, 0xde, 0xdc, 0x63, 0x52, 0x28, 0x40, 0xa6, 0x92, 0xda, 0x8c, 0x1, 0x31, 0x7, 0xa6, 0x61, 0x8a, 0x57, 0x6c, 0x58, 0x88, 0xe3}}
	return a, nil
}

var __000003_blocks_rootidUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xca\xcf\x2f\x89\xcf\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\xc2\x68\x66\x83\x3e\x00\x00\x00")

func _000003_blocks_rootidUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_blocks_rootidUpSql,
		"000003_blocks_rootid.up.sql",
	)
}

func _000003_blocks_rootidUpSql() (*asset, error) {
	bytes, err := _000003_blocks_rootidUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_blocks_rootid.up.sql", size: 62, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe1, 0x4c, 0xe5, 0xb, 0xa4, 0x83, 0xfc, 0x49, 0x39, 0x1d, 0x5b, 0xd0, 0xcf, 0xa3, 0x5e, 0x1f, 0x5c, 0x90, 0x97, 0x13, 0x1c, 0xcc, 0xd3, 0x6f, 0x3f, 0xa5, 0x6, 0x67, 0xfd, 0x4d, 0xc0, 0x55}}
	return a, nil
}

var __000004_auth_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x2d\x2d\x4e\x2d\x2a\xb6\xe6\xc2\x2e\x59\x9c\x5a\x5c\x9c\x99\x9f\x57\x6c\xcd\x05\x08\x00\x00\xff\xff\xb6\xc1\x44\xa1\x3d\x00\x00\x00")

func _000004_auth_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_auth_tableDownSql,
		"000004_auth_table.down.sql",
	)
}

func _000004_auth_tableDownSql() (*asset, error) {
	bytes, err := _000004_auth_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_auth_table.down.sql", size: 61, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7a, 0xd8, 0xba, 0x1, 0x9f, 0x51, 0xfb, 0x45, 0xd8, 0x3, 0xd8, 0xf7, 0x73, 0xee, 0x74, 0x38, 0x32, 0x52, 0x74, 0x99, 0xb7, 0xda, 0xc6, 0x7c, 0xcb, 0xe1, 0x68, 0x6a, 0x88, 0xea, 0x82, 0x70}}
	return a, nil
}

var __000004_auth_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x90\x4d\x6b\xf3\x30\x10\x84\xcf\xd6\xaf\xd8\xa3\x0d\x21\xe4\x0d\xe4\xf4\x9e\x9c\xa0\xb6\xee\x87\x53\x6c\x51\x92\x93\x11\xd6\xba\x15\xf5\x87\xd0\x2a\x6d\x41\xe8\xbf\x17\x53\x68\x20\x76\xa1\x97\xee\xf1\x99\xdd\x1d\x66\x76\x05\x4f\x05\x07\x91\x6e\xef\x39\x64\x57\x90\xef\x05\xf0\x43\x56\x8a\x12\xbc\x5f\x1a\x8b\x8d\xfe\x08\xe1\x44\x68\x09\x62\x16\x69\x05\x4f\x69\xb1\xbb\x49\x8b\xf8\xdf\x6a\x95\x2c\x58\x34\x4a\xbd\xec\xf0\x92\x63\x27\x75\xfb\x0d\xd7\x9b\xcd\x08\x8d\x24\x7a\x1f\xec\xe4\x49\xd7\xc8\x8a\xb0\xb6\xe8\x2e\x15\x79\x72\x2f\x15\xa1\x7d\xd3\xf5\xd9\x62\x7d\x96\x94\x74\x72\xe2\x62\x07\x43\xf0\x35\xde\xeb\x06\x96\x66\x20\xf7\x6c\x91\x42\xb8\x2d\xf7\xb9\xf7\xd8\x12\x86\x20\xf8\x41\x78\x8f\xbd\x0a\x61\xc1\xa2\xda\xa2\x74\x58\x49\x37\x9e\x6d\xb3\xeb\x2c\x17\x63\x3c\xa3\x66\xa8\xc2\x16\xa7\xf4\xb1\xc8\x1e\xd2\xe2\x08\x77\xfc\x08\xb1\x56\x09\x4b\xfe\x33\xf6\xbb\x82\x09\x89\xf4\xd0\xff\xd0\xb1\x1b\x5e\xb1\x9f\x2b\xbe\x9a\xee\xfe\x7d\xf8\xb9\x98\x9f\x01\x00\x00\xff\xff\x2e\x3b\xa6\xa7\x47\x02\x00\x00")

func _000004_auth_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_auth_tableUpSql,
		"000004_auth_table.up.sql",
	)
}

func _000004_auth_tableUpSql() (*asset, error) {
	bytes, err := _000004_auth_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_auth_table.up.sql", size: 583, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x59, 0xb2, 0x63, 0x84, 0xf4, 0xd1, 0x42, 0xca, 0x2, 0xe5, 0xa0, 0x87, 0x7e, 0xa2, 0x80, 0xb0, 0x81, 0xe1, 0x5b, 0x35, 0xa0, 0x8b, 0x6c, 0xfe, 0xd8, 0xb0, 0x59, 0xff, 0xac, 0x35, 0xc, 0xbe}}
	return a, nil
}

var __000005_blocks_modifiedbyDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\xcd\x4f\xc9\x4c\xcb\x4c\x4d\x89\x4f\xaa\xb4\xe6\x02\x04\x00\x00\xff\xff\x6a\xfe\x38\x0a\x37\x00\x00\x00")

func _000005_blocks_modifiedbyDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000005_blocks_modifiedbyDownSql,
		"000005_blocks_modifiedby.down.sql",
	)
}

func _000005_blocks_modifiedbyDownSql() (*asset, error) {
	bytes, err := _000005_blocks_modifiedbyDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000005_blocks_modifiedby.down.sql", size: 55, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x76, 0x4b, 0x62, 0xaa, 0x59, 0xd0, 0x49, 0xb4, 0x9f, 0x2e, 0xfd, 0xe7, 0xad, 0x59, 0x4b, 0xb7, 0x8d, 0x94, 0xa2, 0x87, 0x42, 0xd3, 0x68, 0xc9, 0x61, 0x59, 0x8d, 0x68, 0xff, 0x3b, 0xd5, 0xdb}}
	return a, nil
}

var __000005_blocks_modifiedbyUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\xcd\x4f\xc9\x4c\xcb\x4c\x4d\x89\x4f\xaa\x54\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\x30\x55\xd2\xd8\x42\x00\x00\x00")

func _000005_blocks_modifiedbyUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000005_blocks_modifiedbyUpSql,
		"000005_blocks_modifiedby.up.sql",
	)
}

func _000005_blocks_modifiedbyUpSql() (*asset, error) {
	bytes, err := _000005_blocks_modifiedbyUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000005_blocks_modifiedby.up.sql", size: 66, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8e, 0x1, 0xd8, 0x3, 0x3a, 0xc0, 0x92, 0x34, 0xa8, 0xd, 0x85, 0x3, 0x86, 0x9c, 0x16, 0x13, 0x2f, 0x83, 0xc7, 0x70, 0xed, 0xcb, 0x63, 0xca, 0xba, 0xd5, 0x94, 0x99, 0x39, 0xfd, 0xf8, 0x35}}
	return a, nil
}

var __000006_sharing_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x2d\xce\x48\x2c\xca\xcc\x4b\xb7\xe6\x02\x04\x00\x00\xff\xff\x7a\x74\xe5\xab\x1f\x00\x00\x00")

func _000006_sharing_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000006_sharing_tableDownSql,
		"000006_sharing_table.down.sql",
	)
}

func _000006_sharing_tableDownSql() (*asset, error) {
	bytes, err := _000006_sharing_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000006_sharing_table.down.sql", size: 31, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x31, 0x98, 0x81, 0xe, 0xc8, 0x13, 0xd0, 0x6a, 0x21, 0xec, 0x11, 0x2b, 0x65, 0x78, 0x8b, 0x69, 0x2e, 0x94, 0xa2, 0xe7, 0xf9, 0xc4, 0xbd, 0x18, 0xfc, 0x2, 0x2, 0x5b, 0xc5, 0xdc, 0x78, 0x38}}
	return a, nil
}

var __000006_sharing_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcc\x41\x0b\x82\x30\x14\x00\xe0\xf3\xf6\x2b\xde\x51\x41\xc2\x08\xba\x74\x9a\xb2\x6a\x64\x1a\x73\x44\x9e\x64\xb2\x67\x3d\x2a\x15\x33\x28\xc4\xff\x1e\x5d\x3a\x74\xff\xf8\x62\x2d\x85\x91\x60\x44\x94\x48\x50\x6b\x48\x33\x03\xf2\xa4\x72\x93\xc3\x38\xce\xba\x1e\x6b\x7a\x4d\xd3\xe3\x62\x7b\x6a\xce\xe0\x71\x46\x0e\x8e\x42\xc7\x5b\xa1\xbd\xc5\xd2\x0f\x38\xc3\xc6\x56\x37\x74\x10\x65\x59\x22\x45\x1a\x70\x36\xb4\x57\x6c\x7e\x6a\x1e\x86\x5f\x76\x6f\x1d\xd5\x84\xae\xac\xde\x7f\xc1\xb3\x73\x76\xc0\xd2\x0e\x10\xa9\x8d\x4a\x4d\xc0\xd9\x41\xab\xbd\xd0\x05\xec\x64\x01\x1e\x39\x9f\xfb\x2b\xfe\x09\x00\x00\xff\xff\x7b\xf0\x53\xc5\xaa\x00\x00\x00")

func _000006_sharing_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000006_sharing_tableUpSql,
		"000006_sharing_table.up.sql",
	)
}

func _000006_sharing_tableUpSql() (*asset, error) {
	bytes, err := _000006_sharing_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000006_sharing_table.up.sql", size: 170, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x35, 0xdc, 0x29, 0xae, 0x1f, 0xe1, 0x7f, 0xf5, 0x9, 0xa0, 0xb3, 0x72, 0x3c, 0xbe, 0x7e, 0x40, 0x9c, 0x8d, 0xab, 0x6c, 0x1a, 0x71, 0xf1, 0xaa, 0x4d, 0x17, 0x7c, 0x23, 0xf1, 0x52, 0x78, 0x88}}
	return a, nil
}

var __000007_workspaces_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x2d\xcf\x2f\xca\x2e\x2e\x48\x4c\x4e\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\x1a\xe4\xe6\x36\x22\x00\x00\x00")

func _000007_workspaces_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000007_workspaces_tableDownSql,
		"000007_workspaces_table.down.sql",
	)
}

func _000007_workspaces_tableDownSql() (*asset, error) {
	bytes, err := _000007_workspaces_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000007_workspaces_table.down.sql", size: 34, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfc, 0xb1, 0x2b, 0x90, 0x8a, 0xcb, 0xe0, 0xd8, 0x87, 0x62, 0xcf, 0x86, 0x6b, 0xc9, 0x9c, 0x86, 0x21, 0xa4, 0x87, 0xad, 0x47, 0x49, 0xc5, 0x49, 0x34, 0xe2, 0x24, 0x49, 0x4e, 0x9a, 0x3d, 0x5a}}
	return a, nil
}

var __000007_workspaces_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8d\x41\x4b\xc3\x30\x14\x80\xcf\xcd\xaf\x78\xc7\x16\xca\x98\x08\x5e\x3c\x65\x23\x6a\xb4\x76\x92\x46\xd9\x4e\xa5\x9a\xd7\x12\xa6\x4d\x68\x32\x54\x1e\xef\xbf\x8b\x1e\x3c\xec\xfc\x7d\x7c\xdf\xd6\x28\x69\x15\x58\xb9\x69\x14\xe8\x1b\x68\x77\x16\xd4\x5e\x77\xb6\x03\xa2\x55\x5c\x70\xf4\x5f\xcc\x9f\x61\x39\xa6\x38\xbc\x61\x82\x52\x14\xde\xc1\x8b\x34\xdb\x3b\x69\xca\xcb\xab\xaa\x16\x45\xf2\xd3\x7c\x8a\x7d\x0e\x47\x9c\xff\xd1\xc5\x7a\x5d\xfd\xe5\xda\xe7\xa6\xf9\x95\x30\x67\x3f\x4f\x09\x88\xfc\x08\xab\x18\x52\x9e\x16\x4c\xcc\xf7\xdd\xae\x25\xc2\xf7\x84\xcc\x56\xed\x2d\x11\xce\x8e\xb9\x16\xc5\x47\x70\x7e\xf4\xe8\xfa\xd7\xef\xb3\xe3\x29\xba\x21\x63\x3f\x64\xd8\xe8\x5b\xdd\xda\x5a\x14\x4f\x46\x3f\x4a\x73\x80\x07\x75\x80\xd2\xbb\x4a\x54\xd7\xe2\x27\x00\x00\xff\xff\x96\x66\x5c\xc2\xde\x00\x00\x00")

func _000007_workspaces_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000007_workspaces_tableUpSql,
		"000007_workspaces_table.up.sql",
	)
}

func _000007_workspaces_tableUpSql() (*asset, error) {
	bytes, err := _000007_workspaces_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000007_workspaces_table.up.sql", size: 222, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe7, 0x5c, 0xc9, 0xf9, 0x6d, 0x97, 0x74, 0x41, 0xc4, 0x98, 0x1b, 0x90, 0xb9, 0x9f, 0x62, 0x9a, 0x86, 0x2f, 0xcc, 0x87, 0xc6, 0xb, 0xb4, 0x93, 0xa0, 0xb8, 0xa2, 0xf4, 0x26, 0xa2, 0xb, 0xed}}
	return a, nil
}

var __000008_teamsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xcf\x2f\xca\x2e\x2e\x48\x4c\x4e\x8d\xcf\x4c\xb1\xe6\xe2\xc2\xa1\xb1\x38\x23\xb1\x28\x33\x2f\x9d\x1c\x9d\xa9\xc5\xc5\x99\xf9\x79\xa8\x96\x26\x96\x96\x64\xc4\x17\xa7\x16\x95\x65\x26\xa7\x5a\x73\x01\x02\x00\x00\xff\xff\x24\x48\xc4\xb6\xad\x00\x00\x00")

func _000008_teamsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000008_teamsDownSql,
		"000008_teams.down.sql",
	)
}

func _000008_teamsDownSql() (*asset, error) {
	bytes, err := _000008_teamsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000008_teams.down.sql", size: 173, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x71, 0xde, 0xad, 0x3b, 0xef, 0xf1, 0xd1, 0x17, 0x44, 0xee, 0x1d, 0x30, 0x33, 0x15, 0xa9, 0x84, 0x99, 0xe7, 0x6e, 0xbf, 0x8c, 0xa4, 0x4, 0xd6, 0x68, 0xab, 0x77, 0x68, 0x13, 0x74, 0x1, 0x5d}}
	return a, nil
}

var __000008_teamsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xe6\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xcf\x2f\xca\x2e\x2e\x48\x4c\x4e\x8d\xcf\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd3\xb4\xe6\xe2\xc2\x61\x46\x71\x46\x62\x51\x66\x5e\x3a\x85\x86\xa4\x16\x17\x67\xe6\xe7\xa1\x38\x25\xb1\xb4\x24\x23\xbe\x38\xb5\xa8\x2c\x33\x39\x15\x6e\x8a\x91\x01\xc8\x94\xd0\x00\x17\xc7\x10\x2c\x3e\x51\x08\x76\x0d\x41\xb5\xdd\x56\x41\xdd\x40\x5d\x21\xdc\xc3\x35\xc8\x15\x43\x42\x5d\xc1\x3f\x08\x55\xd0\x33\x58\xc1\x2f\xd4\xc7\xc7\x9a\x0b\x10\x00\x00\xff\xff\xab\x8d\x48\xa9\x30\x01\x00\x00")

func _000008_teamsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000008_teamsUpSql,
		"000008_teams.up.sql",
	)
}

func _000008_teamsUpSql() (*asset, error) {
	bytes, err := _000008_teamsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000008_teams.up.sql", size: 304, mode: os.FileMode(0644), modTime: time.Unix(1618834577, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xad, 0x91, 0xe6, 0xdf, 0x66, 0x76, 0x63, 0x68, 0x70, 0x8c, 0x30, 0x66, 0xf0, 0xc3, 0xa8, 0x76, 0xff, 0xe3, 0x59, 0x99, 0x49, 0xd, 0x90, 0xf5, 0xf4, 0x10, 0xeb, 0x6e, 0x0, 0x2c, 0x67, 0xeb}}
	return a, nil
}

var __000009_blocks_historyDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xa8\xae\xd6\x2b\x28\x4a\x4d\xcb\xac\xa8\xad\x4d\xca\xc9\x4f\xce\x2e\xb6\xe6\x72\xf4\x09\x71\x0d\xc2\x25\x1d\x9f\x91\x59\x5c\x92\x5f\x54\xa9\x10\xe4\xea\xe7\xe8\xeb\xaa\x10\xe2\x8f\xcd\x08\x40\x00\x00\x00\xff\xff\x38\xe5\xec\x7a\x61\x00\x00\x00")

func _000009_blocks_historyDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000009_blocks_historyDownSql,
		"000009_blocks_history.down.sql",
	)
}

func _000009_blocks_historyDownSql() (*asset, error) {
	bytes, err := _000009_blocks_historyDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000009_blocks_history.down.sql", size: 97, mode: os.FileMode(0644), modTime: time.Unix(1618836045, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x78, 0x56, 0x8b, 0xdc, 0x49, 0x6, 0xe2, 0x3a, 0x1f, 0x77, 0xd4, 0xb3, 0xd0, 0x8e, 0xac, 0xb9, 0xfe, 0xa0, 0x69, 0x42, 0x7c, 0xb2, 0x25, 0xa7, 0x1e, 0x29, 0xce, 0xe0, 0x27, 0x38, 0x2, 0x45}}
	return a, nil
}

var __000009_blocks_historyUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\xc1\x6f\x9b\x30\x14\xc6\xcf\xf8\xaf\x78\x17\x14\x98\x68\x2f\x93\x72\x68\x4e\x0e\x31\x89\x37\xb0\x2b\xe3\xae\xcd\x2e\x29\x0d\x66\xb5\x4a\x02\xc5\x9e\xb6\x08\xf1\xbf\x4f\x54\x4d\x9a\x26\x4c\x3b\x4c\xbd\xc1\x7b\xf0\xfb\xbe\xf7\xde\x87\x63\x49\x04\x48\x3c\x8d\x09\xb4\xed\x65\xdd\xa8\x42\xff\xee\xba\x87\xb2\x5a\x3f\x19\x10\x84\xe1\x84\x80\xe4\xe7\xbd\xd5\xa3\x36\xb6\x6a\x76\x13\x14\x0a\x82\x25\x79\x65\xd0\x08\x18\x97\x40\xee\x68\x2a\xd3\x01\xa2\x87\x1c\x9d\xc3\x37\x2c\xc2\x05\x16\xde\xe7\xb1\x1f\x20\x00\x80\xb6\xd5\x05\x5c\xd6\x95\xb1\x3f\x1a\x65\xba\x4e\x6f\x8d\x6a\xec\x2a\xb3\x20\x69\x42\x52\x89\x93\x6b\xf9\xfd\x85\xcc\x6e\xe2\x18\x66\x24\xc2\x37\xb1\x04\xc6\x6f\x3d\x3f\x68\x5b\xb5\xcd\xbb\xee\x08\x64\x9e\x4b\x6d\xd5\x31\x66\x86\x25\xe9\x51\x67\x0c\x2f\x95\x22\xea\x3b\xde\xc8\x5d\x5e\xb8\x9b\x0b\x37\x07\x77\x71\xe5\x26\x57\x6e\x31\x0a\x60\xc4\xf8\xed\xc8\x1f\xd2\xd8\xec\xcc\x73\x39\x24\xe1\x8d\xfd\x61\xa7\xe3\x37\x8c\x53\x67\x8d\xda\xda\xd5\xd0\x2e\x7e\x55\xcd\x93\xa9\xb3\xb5\x1a\xec\x36\x55\x35\xfc\xdb\xa6\xca\x75\xa1\x55\xbe\x7a\xd8\xbd\x6f\x3a\xf7\x66\xfd\xa8\x36\xd9\x3d\x4c\xe9\x9c\x32\x19\x20\xc7\xee\x6a\x05\x92\xdc\xbd\x3c\x6b\x5b\x1e\x5e\x0a\xad\xca\xdc\x9c\x9d\xe3\x4b\xca\x59\xdb\xaa\xd2\xa8\xae\xeb\xbf\x7c\x1d\x23\x40\xce\xba\x51\x99\x55\xfd\xf8\x07\xf8\xcf\x3a\x3f\x2d\xe5\xaa\x54\x27\xa5\x6b\x41\x13\x2c\x96\xf0\x95\x2c\xc1\xd3\xb9\x8f\xfc\x09\x42\xef\x36\x8b\x28\x4b\x89\x90\x40\xe7\x8c\x0b\x02\x94\x0d\xc5\x10\xbc\x94\xc4\x24\x94\xf0\x09\x22\xc1\x93\xbf\xe7\x14\xb8\x98\x11\x01\xd3\x25\x1c\x5d\x8c\xa4\xa1\x3f\x41\xfb\xa3\x9c\x4e\x7d\x30\xf0\x41\xca\xc0\x19\x84\x9c\x45\x31\x0d\x25\xcc\x78\x1f\x9a\x05\x65\xf3\x53\x43\xfb\x30\xef\xed\x70\xf1\x8f\x95\xfc\xa7\xaf\x37\xfd\x3f\x01\x00\x00\xff\xff\x3c\x60\x66\xd4\x1d\x04\x00\x00")

func _000009_blocks_historyUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000009_blocks_historyUpSql,
		"000009_blocks_history.up.sql",
	)
}

func _000009_blocks_historyUpSql() (*asset, error) {
	bytes, err := _000009_blocks_historyUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000009_blocks_history.up.sql", size: 1053, mode: os.FileMode(0644), modTime: time.Unix(1618855271, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x89, 0x9c, 0xfe, 0x80, 0x15, 0xeb, 0xe4, 0x35, 0xe3, 0x63, 0x47, 0xaf, 0x73, 0xab, 0xbe, 0x87, 0x7c, 0x8e, 0xb1, 0x10, 0x92, 0x8f, 0xcd, 0xe2, 0x1c, 0x74, 0x32, 0x1d, 0x68, 0x48, 0xbd, 0x97}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"000001_init.down.sql":                  _000001_initDownSql,
	"000001_init.up.sql":                    _000001_initUpSql,
	"000002_system_settings_table.down.sql": _000002_system_settings_tableDownSql,
	"000002_system_settings_table.up.sql":   _000002_system_settings_tableUpSql,
	"000003_blocks_rootid.down.sql":         _000003_blocks_rootidDownSql,
	"000003_blocks_rootid.up.sql":           _000003_blocks_rootidUpSql,
	"000004_auth_table.down.sql":            _000004_auth_tableDownSql,
	"000004_auth_table.up.sql":              _000004_auth_tableUpSql,
	"000005_blocks_modifiedby.down.sql":     _000005_blocks_modifiedbyDownSql,
	"000005_blocks_modifiedby.up.sql":       _000005_blocks_modifiedbyUpSql,
	"000006_sharing_table.down.sql":         _000006_sharing_tableDownSql,
	"000006_sharing_table.up.sql":           _000006_sharing_tableUpSql,
	"000007_workspaces_table.down.sql":      _000007_workspaces_tableDownSql,
	"000007_workspaces_table.up.sql":        _000007_workspaces_tableUpSql,
	"000008_teams.down.sql":                 _000008_teamsDownSql,
	"000008_teams.up.sql":                   _000008_teamsUpSql,
	"000009_blocks_history.down.sql":        _000009_blocks_historyDownSql,
	"000009_blocks_history.up.sql":          _000009_blocks_historyUpSql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"000001_init.down.sql": {_000001_initDownSql, map[string]*bintree{}},
	"000001_init.up.sql": {_000001_initUpSql, map[string]*bintree{}},
	"000002_system_settings_table.down.sql": {_000002_system_settings_tableDownSql, map[string]*bintree{}},
	"000002_system_settings_table.up.sql": {_000002_system_settings_tableUpSql, map[string]*bintree{}},
	"000003_blocks_rootid.down.sql": {_000003_blocks_rootidDownSql, map[string]*bintree{}},
	"000003_blocks_rootid.up.sql": {_000003_blocks_rootidUpSql, map[string]*bintree{}},
	"000004_auth_table.down.sql": {_000004_auth_tableDownSql, map[string]*bintree{}},
	"000004_auth_table.up.sql": {_000004_auth_tableUpSql, map[string]*bintree{}},
	"000005_blocks_modifiedby.down.sql": {_000005_blocks_modifiedbyDownSql, map[string]*bintree{}},
	"000005_blocks_modifiedby.up.sql": {_000005_blocks_modifiedbyUpSql, map[string]*bintree{}},
	"000006_sharing_table.down.sql": {_000006_sharing_tableDownSql, map[string]*bintree{}},
	"000006_sharing_table.up.sql": {_000006_sharing_tableUpSql, map[string]*bintree{}},
	"000007_workspaces_table.down.sql": {_000007_workspaces_tableDownSql, map[string]*bintree{}},
	"000007_workspaces_table.up.sql": {_000007_workspaces_tableUpSql, map[string]*bintree{}},
	"000008_teams.down.sql": {_000008_teamsDownSql, map[string]*bintree{}},
	"000008_teams.up.sql": {_000008_teamsUpSql, map[string]*bintree{}},
	"000009_blocks_history.down.sql": {_000009_blocks_historyDownSql, map[string]*bintree{}},
	"000009_blocks_history.up.sql": {_000009_blocks_historyUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
