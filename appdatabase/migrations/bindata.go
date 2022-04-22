// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1640111208_dummy.up.sql (258B)
// 1642666031_add_removed_clock_to_bookmarks.up.sql (117B)
// 1643644541_gif_api_key_setting.up.sql (108B)
// 1644188994_recent_stickers.up.sql (79B)
// 1646659233_add_address_to_dapp_permisssion.up.sql (700B)
// 1646841105_add_emoji_account.up.sql (96B)
// 1647278782_display_name.up.sql (110B)
// 1647862838_reset_last_backup.up.sql (37B)
// 1647871652_add_settings_sync_clock_table.up.sql (1.044kB)
// 1647880168_add_torrent_config.up.sql (211B)
// 1647882837_add_communities_settings_table.up.sql (206B)
// 1647956635_add_waku_messages_table.up.sql (266B)
// 1648554928_network_test.up.sql (132B)
// 1649174829_add_visitble_token.up.sql (84B)
// 1649882262_add_derived_from_accounts.up.sql (110B)
// 1650612625_add_community_message_archive_hashes_table.up.sql (130B)
// 1650616788_add_communities_archives_info_table.up.sql (208B)
// doc.go (74B)

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

var __1640111208_dummyUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xb1\x4e\x04\x31\x0c\x44\x7b\xbe\x62\x3a\x40\x22\x7f\x71\x57\x50\xc0\x4a\xe8\x44\x9f\x4b\x26\x1b\x8b\x5d\x1b\x62\xa3\xe5\xf3\x51\x44\x45\x73\x6e\x3d\xef\xcd\xa4\x84\xd3\x82\xd7\xe5\x82\xb7\xf3\xcb\xf2\x7e\xbe\x4b\x09\xe9\xd6\xcd\xc0\xa5\x8b\xa3\xc9\x46\x54\xa3\x43\x2d\xba\xe8\x3a\x3f\xcf\x71\xef\x28\x83\x39\x58\xe1\x86\xe8\xc4\x2e\xeb\xc8\x21\xa6\xf8\x50\x3b\x1c\x6d\xd8\x8e\xa3\x4b\xe9\xf0\xaf\xed\xcf\x13\x86\x62\x1a\xa2\xdf\x84\x8b\x16\x62\xda\xfe\xd3\x3d\x3b\xae\xa4\xc2\x3f\x37\x89\x59\x20\x8a\x38\x0c\x0f\x57\x36\x1b\x7c\x42\xd6\x8a\xdc\x82\x03\x6a\x95\xa5\xad\xe0\x4f\x8c\x5c\x26\xfe\x38\x85\x27\x9b\x63\x91\x6b\x45\xb1\x4a\x74\x0e\xfe\x06\x00\x00\xff\xff\x9b\xc1\xf3\x13\x02\x01\x00\x00")

func _1640111208_dummyUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1640111208_dummyUpSql,
		"1640111208_dummy.up.sql",
	)
}

func _1640111208_dummyUpSql() (*asset, error) {
	bytes, err := _1640111208_dummyUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1640111208_dummy.up.sql", size: 258, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3e, 0xf0, 0xae, 0x20, 0x6e, 0x75, 0xd1, 0x36, 0x14, 0xf2, 0x40, 0xe5, 0xd6, 0x7a, 0xc4, 0xa5, 0x72, 0xaa, 0xb5, 0x4d, 0x71, 0x97, 0xb8, 0xe8, 0x95, 0x22, 0x95, 0xa2, 0xac, 0xaf, 0x48, 0x58}}
	return a, nil
}

var __1642666031_add_removed_clock_to_bookmarksUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xcf\xcf\xce\x4d\x2c\xca\x2e\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x4a\xcd\xcd\x2f\x4b\x4d\x51\x70\xf2\xf7\xf7\x71\x75\xf4\x53\x70\x71\x75\x73\x0c\xf5\x09\x51\x70\x73\xf4\x09\x76\xb5\xe6\x22\xa8\x3f\x39\x27\x3f\x39\x5b\xc1\xd3\x2f\x04\xae\xd3\xc0\x1a\x10\x00\x00\xff\xff\xe6\xf6\xbf\x66\x75\x00\x00\x00")

func _1642666031_add_removed_clock_to_bookmarksUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1642666031_add_removed_clock_to_bookmarksUpSql,
		"1642666031_add_removed_clock_to_bookmarks.up.sql",
	)
}

func _1642666031_add_removed_clock_to_bookmarksUpSql() (*asset, error) {
	bytes, err := _1642666031_add_removed_clock_to_bookmarksUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1642666031_add_removed_clock_to_bookmarks.up.sql", size: 117, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x84, 0x4e, 0x38, 0x99, 0x7a, 0xc, 0x90, 0x13, 0xec, 0xfe, 0x2f, 0x55, 0xff, 0xb7, 0xb6, 0xaa, 0x96, 0xc6, 0x92, 0x79, 0xcc, 0xee, 0x4e, 0x99, 0x53, 0xfe, 0x1c, 0xbb, 0x32, 0x2, 0xa4, 0x27}}
	return a, nil
}

var __1643644541_gif_api_key_settingUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4e\x2d\x29\xc9\xcc\x4b\x2f\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xcf\x4c\x8b\x4f\x2c\xc8\x8c\xcf\x4e\xad\x54\x08\x71\x8d\x08\x51\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x51\x50\x52\xb2\xe6\x0a\x0d\x70\x71\x0c\x41\x32\x20\xd8\x35\x04\x45\xa7\x2d\x58\x15\x20\x00\x00\xff\xff\x59\x5f\x0d\x48\x6c\x00\x00\x00")

func _1643644541_gif_api_key_settingUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1643644541_gif_api_key_settingUpSql,
		"1643644541_gif_api_key_setting.up.sql",
	)
}

func _1643644541_gif_api_key_settingUpSql() (*asset, error) {
	bytes, err := _1643644541_gif_api_key_settingUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1643644541_gif_api_key_setting.up.sql", size: 108, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1b, 0x94, 0x28, 0xfb, 0x66, 0xd1, 0x7c, 0xb8, 0x89, 0xe2, 0xb4, 0x71, 0x65, 0x24, 0x57, 0x22, 0x95, 0x38, 0x97, 0x3, 0x9b, 0xc6, 0xa4, 0x41, 0x7b, 0xba, 0xf7, 0xdb, 0x70, 0xf7, 0x20, 0x3a}}
	return a, nil
}

var __1644188994_recent_stickersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\x0d\x70\x71\x0c\x71\x55\x28\x4e\x2d\x29\xc9\xcc\x4b\x2f\x56\x08\x76\x0d\x51\x28\x2e\xc9\x4c\xce\x4e\x2d\x2a\x8e\x2f\x4a\x4d\x4e\xcd\x2b\x89\x87\xf1\x15\x6c\x15\xfc\x42\x7d\x7c\x14\xc2\x3d\x5c\x83\x5c\x15\x8a\x2b\xf3\x4a\x32\x52\x4b\x32\x93\xe3\x33\x53\x14\x6c\x15\xd4\x33\x53\xd4\xad\xb9\x00\x01\x00\x00\xff\xff\x1d\x83\x1b\xca\x4f\x00\x00\x00")

func _1644188994_recent_stickersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1644188994_recent_stickersUpSql,
		"1644188994_recent_stickers.up.sql",
	)
}

func _1644188994_recent_stickersUpSql() (*asset, error) {
	bytes, err := _1644188994_recent_stickersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1644188994_recent_stickers.up.sql", size: 79, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1e, 0xad, 0xaa, 0x30, 0xbf, 0x4, 0x7, 0xf8, 0xc3, 0x3, 0xb8, 0x97, 0x23, 0x2b, 0xbd, 0x1c, 0x60, 0x69, 0xb0, 0x42, 0x5e, 0x6b, 0xd, 0xa7, 0xa3, 0x6b, 0x2e, 0xdc, 0x70, 0x13, 0x72, 0x7}}
	return a, nil
}

var __1646659233_add_address_to_dapp_permisssionUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x51\xcd\x6e\x82\x40\x18\xbc\xf3\x14\x13\x4e\x98\x10\x0f\x5e\x49\x0f\x5b\xfc\x30\x9b\xc2\xae\x5d\x96\x44\x4f\x84\x14\x4c\xb6\x51\xb0\xa2\xef\xdf\xb0\xfc\x08\x36\xe9\x0d\xe6\x9b\x19\x66\x86\x50\x11\xd3\x04\xcd\xde\x63\x02\x8f\x20\xa4\x06\x1d\x78\xaa\x53\x94\xc5\xf5\xda\xe6\x1b\x78\x0e\x00\x98\x12\x5c\x68\xda\x91\xc2\x5e\xf1\x84\xa9\x23\x3e\xe8\x08\x96\x69\xc9\x45\xa8\x28\x21\xa1\xad\x5a\x64\x71\xec\x5b\x49\x5d\x5c\x2a\x68\x3a\xe8\xfe\xb5\x28\xcb\x5b\xd5\xb6\x16\x99\x98\x28\xab\x53\xf1\x38\xdf\xe1\xba\x3d\x2b\x94\x22\xd5\x8a\x71\xa1\xf1\xa8\xcd\xcf\xa3\xca\xfb\x1c\x9d\x59\x3e\x5a\x64\x82\x7f\x66\x04\xaf\x03\xfd\x01\x5c\x39\xab\xc0\x71\xfe\xe9\x73\xad\x6e\x17\xd3\xb6\xa6\xa9\x9f\xad\x3a\xef\xdc\x94\x30\xf5\xfd\x25\xfc\x93\xbd\x0c\xdc\x5f\x23\xa9\x88\xef\x44\x37\x81\x37\x78\xac\xa0\x28\x22\x45\x22\xa4\x69\x3b\xaf\x83\xa5\xc0\x96\x62\xd2\x84\x90\xa5\x21\xdb\x92\x0d\xca\x45\x4a\x4a\x77\x9b\xca\x69\xe9\x94\x62\x0a\x87\xcf\xd8\xf5\x7c\xb8\x2e\x22\x25\x93\x9e\xf2\x22\x5b\x16\x6a\xab\x73\xf5\x75\x1f\xbd\xd6\xa6\xf4\xe7\x84\xf5\xac\xcf\xe9\xd6\x5c\xe6\x37\x7c\x37\xa6\x9e\x42\x34\xd3\xe3\xda\xfe\xc0\xb7\x85\x8d\x2d\xdb\xe1\x81\xe3\x6c\x95\xdc\x0f\x4b\xcf\x28\xc1\x1c\x1f\x63\xb3\x58\x93\xfa\xcb\xcd\x37\x50\x24\x58\x42\x58\xd6\x09\x16\x82\x31\xd9\x93\xda\xdb\xfe\x06\x00\x00\xff\xff\xb0\x41\x9a\x48\xbc\x02\x00\x00")

func _1646659233_add_address_to_dapp_permisssionUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1646659233_add_address_to_dapp_permisssionUpSql,
		"1646659233_add_address_to_dapp_permisssion.up.sql",
	)
}

func _1646659233_add_address_to_dapp_permisssionUpSql() (*asset, error) {
	bytes, err := _1646659233_add_address_to_dapp_permisssionUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1646659233_add_address_to_dapp_permisssion.up.sql", size: 700, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xed, 0xb0, 0x35, 0xcc, 0x2e, 0x16, 0xe6, 0x15, 0x86, 0x2c, 0x37, 0x80, 0xae, 0xa3, 0xc5, 0x31, 0x78, 0x5, 0x9d, 0xcd, 0x7b, 0xeb, 0x5f, 0xf2, 0xb3, 0x74, 0x72, 0xdf, 0xcf, 0x88, 0xb, 0x40}}
	return a, nil
}

var __1646841105_add_emoji_accountUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4c\x4e\xce\x2f\xcd\x2b\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xcd\xcd\xcf\xca\x54\x08\x71\x8d\x08\x51\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x51\x50\x52\xb2\xe6\x0a\x0d\x70\x71\x0c\x41\xd2\x1a\xec\x1a\x02\xd5\x63\x0b\x96\x07\x04\x00\x00\xff\xff\x2b\x1c\x4e\xaa\x60\x00\x00\x00")

func _1646841105_add_emoji_accountUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1646841105_add_emoji_accountUpSql,
		"1646841105_add_emoji_account.up.sql",
	)
}

func _1646841105_add_emoji_accountUpSql() (*asset, error) {
	bytes, err := _1646841105_add_emoji_accountUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1646841105_add_emoji_account.up.sql", size: 96, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe6, 0x77, 0x29, 0x95, 0x18, 0x64, 0x82, 0x63, 0xe7, 0xaf, 0x6c, 0xa9, 0x15, 0x7d, 0x46, 0xa6, 0xbc, 0xdf, 0xa7, 0xd, 0x2b, 0xd2, 0x2d, 0x97, 0x4d, 0xa, 0x6b, 0xd, 0x6e, 0x90, 0x42, 0x5c}}
	return a, nil
}

var __1647278782_display_nameUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4e\x2d\x29\xc9\xcc\x4b\x2f\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xc9\x2c\x2e\xc8\x49\xac\x8c\xcf\x4b\xcc\x4d\x55\x08\x71\x8d\x08\x51\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x51\x50\x52\xb2\xe6\x0a\x0d\x70\x71\x0c\x41\x32\x21\xd8\x35\x04\x55\xab\x2d\x58\x19\x20\x00\x00\xff\xff\xc7\x11\xdd\x01\x6e\x00\x00\x00")

func _1647278782_display_nameUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1647278782_display_nameUpSql,
		"1647278782_display_name.up.sql",
	)
}

func _1647278782_display_nameUpSql() (*asset, error) {
	bytes, err := _1647278782_display_nameUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1647278782_display_name.up.sql", size: 110, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf4, 0xa1, 0x1f, 0x3e, 0x61, 0x65, 0x8d, 0xff, 0xee, 0xde, 0xc5, 0x91, 0xd9, 0x5c, 0xb5, 0xe2, 0xf0, 0xb7, 0xe7, 0x5c, 0x5c, 0x16, 0x25, 0x89, 0xee, 0x78, 0x12, 0xea, 0x3e, 0x48, 0x41, 0xa6}}
	return a, nil
}

var __1647862838_reset_last_backupUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\x0d\x70\x71\x0c\x71\x55\x28\x4e\x2d\x29\xc9\xcc\x4b\x2f\x56\x08\x76\x0d\x51\xc8\x49\x2c\x2e\x89\x4f\x4a\x4c\xce\x2e\x2d\x50\xb0\x55\x30\xb0\xe6\x02\x04\x00\x00\xff\xff\x2e\x14\x2f\x78\x25\x00\x00\x00")

func _1647862838_reset_last_backupUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1647862838_reset_last_backupUpSql,
		"1647862838_reset_last_backup.up.sql",
	)
}

func _1647862838_reset_last_backupUpSql() (*asset, error) {
	bytes, err := _1647862838_reset_last_backupUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1647862838_reset_last_backup.up.sql", size: 37, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x21, 0xe3, 0xd5, 0xf6, 0x5f, 0xfe, 0x65, 0xfa, 0x1d, 0x88, 0xf8, 0x5f, 0x24, 0x71, 0x34, 0x68, 0x96, 0x2a, 0x60, 0x87, 0x15, 0x82, 0x4d, 0x8a, 0x59, 0x3d, 0x1f, 0xd8, 0x56, 0xd4, 0xfb, 0xda}}
	return a, nil
}

var __1647871652_add_settings_sync_clock_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xcd\x6e\xdb\x30\x10\x84\xef\x7e\x8a\xbd\x25\x06\x78\xc8\x3d\x27\x35\x61\x1b\xa1\xaa\x5d\x30\x74\x82\x9c\x16\x2c\xb5\x52\x16\x96\x49\x81\x4b\x29\xf0\xdb\x17\xfd\x49\x7f\x52\xb9\x92\xc1\x13\xb1\x33\x58\x72\xe6\xbb\x31\xba\xb0\x1a\x6c\xf1\xae\xd2\x20\x94\x33\x87\x56\x50\x8e\xc1\xa3\xef\xa2\xdf\xc3\xe5\x0a\x00\xc0\x0f\x29\x51\xf0\x47\x28\x37\x56\x7f\xd0\x06\x36\x5b\x0b\x9b\x5d\x55\xc1\xad\x7e\x5f\xec\x2a\x0b\x57\xea\xbb\xb0\xe5\x06\x13\x79\x0a\x59\x16\x69\x1b\x37\xc6\xc4\x99\xe6\xd5\x07\x12\x71\x2d\x09\x36\x29\x1e\xd0\xc7\x90\x9d\xcf\x82\x31\x74\xf3\xaf\xea\x13\x35\x94\x12\xd5\x18\xdc\x81\x96\xc8\x47\xa6\x17\xec\x13\x8f\x6e\xc1\xa7\xfb\x14\x1b\xee\x08\x7b\xf6\x79\x48\x24\x28\xcf\xf1\x05\x73\x3c\xdf\x38\xb2\xf0\x17\xee\x38\xcf\x2f\x15\x0a\x35\x4a\x76\x79\x10\x1c\xfa\xda\x2d\xc9\x50\x32\xfb\x3d\x25\xc1\xde\xf9\xbd\x20\x07\xc9\xae\xeb\xa8\x3e\xd7\xd8\x53\xa8\x39\xb4\xcb\x6d\x3f\x90\xc0\xd7\xfb\xbc\xf1\x18\xf2\x33\x65\xf6\xc8\x35\x3c\x14\xe6\xe6\xae\x30\xbf\x34\x17\x5c\x5f\xc0\x67\x53\x7e\x2a\xcc\x13\x7c\xd4\x4f\xab\x35\x3c\x96\xf6\x6e\xbb\xb3\x60\xb6\x8f\xe5\xed\xf5\xaa\xdc\xdc\x6b\x63\xbf\x6d\xd9\x4e\x43\xfd\xca\xb3\xfa\x13\x58\xf5\x37\x91\xea\x7f\xc8\xa9\x37\x4c\xa9\xb7\xd0\xa8\x93\x54\x4c\x4c\x7e\xd7\xae\xa6\x7a\x55\x27\x8b\xfb\x67\xf2\xb3\x19\x75\x32\xfa\x35\x3c\x14\xd5\x4e\xdf\xc3\xe5\x95\x9a\x3c\xeb\xeb\xaf\x01\x00\x00\xff\xff\xf3\x53\x0f\x1c\x14\x04\x00\x00")

func _1647871652_add_settings_sync_clock_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1647871652_add_settings_sync_clock_tableUpSql,
		"1647871652_add_settings_sync_clock_table.up.sql",
	)
}

func _1647871652_add_settings_sync_clock_tableUpSql() (*asset, error) {
	bytes, err := _1647871652_add_settings_sync_clock_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1647871652_add_settings_sync_clock_table.up.sql", size: 1044, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd8, 0x58, 0xec, 0x85, 0x90, 0xfa, 0x30, 0x98, 0x98, 0x9a, 0xa6, 0xa8, 0x96, 0x2b, 0x38, 0x93, 0xf3, 0xae, 0x46, 0x74, 0xa4, 0x41, 0x62, 0x9b, 0x2, 0x86, 0xbf, 0xe5, 0x2a, 0xce, 0xe2, 0xc0}}
	return a, nil
}

var __1647880168_add_torrent_configUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcd\x31\x6e\xc3\x20\x00\x85\xe1\x9d\x53\xbc\xcd\xad\xd4\x1b\x74\xc2\x36\xad\x51\x29\x54\x14\x62\x79\xb2\x88\xc1\x09\x92\x85\x23\xcc\x92\xdb\x47\x1e\x9c\x2d\xeb\xfb\x9f\xf4\x35\x9a\x51\xc3\x60\x68\x2d\x18\xca\x9a\x73\x48\x65\x9c\xd6\x34\xc7\x0b\xde\x08\x10\x92\x3b\x2f\xc1\xa3\x56\x4a\x30\x2a\xd1\xb2\x2f\x6a\x85\xc1\xec\x96\x2d\x7c\x10\xe0\xb6\xe6\x02\x2b\xff\xf9\xb7\x64\x2d\xb8\x34\xfb\xe8\x5d\x71\xa3\x8f\x19\x27\xaa\x9b\x8e\x6a\x48\x65\x20\xad\x10\x7b\x3c\x94\x57\x7d\xbb\xa7\x72\x0d\x25\x4e\x63\xf4\xcf\xc3\xe1\x56\xd1\x57\xf8\xd3\xfc\x97\xea\x01\x3f\x6c\x20\xef\xe8\xb9\xe9\x94\x35\xd0\xaa\xe7\xed\x27\x21\x8f\x00\x00\x00\xff\xff\x9f\xe4\xb4\x2e\xd3\x00\x00\x00")

func _1647880168_add_torrent_configUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1647880168_add_torrent_configUpSql,
		"1647880168_add_torrent_config.up.sql",
	)
}

func _1647880168_add_torrent_configUpSql() (*asset, error) {
	bytes, err := _1647880168_add_torrent_configUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1647880168_add_torrent_config.up.sql", size: 211, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1, 0x92, 0x22, 0x37, 0x96, 0xf3, 0xb5, 0x5b, 0x27, 0xd0, 0x7d, 0x43, 0x5, 0x4e, 0x9d, 0xe2, 0x49, 0xbe, 0x86, 0x31, 0xa1, 0x89, 0xff, 0xd6, 0x51, 0xe0, 0x9c, 0xb, 0xda, 0xfc, 0xf2, 0x93}}
	return a, nil
}

var __1647882837_add_communities_settings_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xcb\xb1\xaa\xc2\x30\x14\x06\xe0\x3d\x4f\xf1\x8f\xf7\x82\x2f\x71\x1a\x4f\xa1\x18\x9b\x12\x23\xd8\x29\xd4\xf6\xd8\x06\x6c\x04\x13\x05\xdf\xde\xcd\x49\x70\xfe\xf8\xb4\x63\xf2\x0c\x4f\x95\x61\x8c\xb7\x75\x7d\xa4\x58\xa2\xe4\x90\xa5\x94\x98\xe6\x8c\x3f\x85\x0f\xbc\x42\x9c\xe0\xf9\xe4\xd1\xb9\x66\x4f\xae\xc7\x8e\x7b\xd8\x16\xda\xb6\xb5\x69\xb4\x87\xe3\xce\x90\xe6\x8d\x02\x56\xc9\x79\x98\x25\x0c\xf7\x71\x89\x4f\x09\x59\x64\x8a\x69\x0e\x92\x86\xf3\x55\x26\x54\xd6\x1a\xa6\x16\x5b\xae\xe9\x68\x3c\x6a\x32\x87\xaf\xf1\x22\x65\x5c\x7e\x4e\xf5\xaf\xd4\x3b\x00\x00\xff\xff\x01\x6b\xfa\x19\xce\x00\x00\x00")

func _1647882837_add_communities_settings_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1647882837_add_communities_settings_tableUpSql,
		"1647882837_add_communities_settings_table.up.sql",
	)
}

func _1647882837_add_communities_settings_tableUpSql() (*asset, error) {
	bytes, err := _1647882837_add_communities_settings_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1647882837_add_communities_settings_table.up.sql", size: 206, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbd, 0x87, 0x78, 0x99, 0xd9, 0x5d, 0xbd, 0xf7, 0x57, 0x9c, 0xca, 0x97, 0xbd, 0xb3, 0xe9, 0xb5, 0x89, 0x31, 0x3f, 0xf6, 0x5c, 0x13, 0xb, 0xc3, 0x54, 0x93, 0x18, 0x40, 0x7, 0x82, 0xfe, 0x7e}}
	return a, nil
}

var __1647956635_add_waku_messages_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcf\xc1\xca\x82\x50\x10\x05\xe0\xfd\x3c\xc5\x59\xfe\x82\x6f\xe0\x4a\xff\x66\x21\xdd\xae\x21\x37\xd0\x95\x0c\x29\x7a\x29\xf3\xc2\x18\xd1\xdb\xb7\x30\x04\xab\xed\xcc\xe1\xcc\x37\xff\x25\xa7\x8e\xe1\xd2\xcc\x30\x1e\x72\xb9\x37\x63\xa7\x2a\x7d\xa7\xf8\x23\x40\x7d\x8f\xcc\x14\x19\x6c\xe1\x60\x4f\xc6\xc4\x04\xcc\x7e\xec\x74\x96\x31\x20\xb7\x6e\xbb\x99\x82\x3f\xc3\x71\xb5\x1d\x07\x79\x5e\x27\x69\xbf\x9b\x82\xb4\xad\xbf\xfd\x38\x31\x88\x0e\x4b\xcf\xb1\xcc\x0f\x69\x59\x63\xcf\xf5\x9a\xa0\x28\x21\x7a\xcb\x73\xbb\xe3\x6a\x2b\x6f\x56\x60\xb3\x80\x0a\xfb\xf9\xda\x9a\x88\x17\x73\x94\xd0\x2b\x00\x00\xff\xff\xde\xa7\x3f\x4b\x0a\x01\x00\x00")

func _1647956635_add_waku_messages_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1647956635_add_waku_messages_tableUpSql,
		"1647956635_add_waku_messages_table.up.sql",
	)
}

func _1647956635_add_waku_messages_tableUpSql() (*asset, error) {
	bytes, err := _1647956635_add_waku_messages_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1647956635_add_waku_messages_table.up.sql", size: 266, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd1, 0xe, 0xe1, 0xdc, 0xda, 0x2e, 0x89, 0x8d, 0xdc, 0x2a, 0x1c, 0x13, 0xa1, 0xfc, 0xfe, 0xf, 0xb2, 0xb9, 0x85, 0xc8, 0x45, 0xd6, 0xd1, 0x7, 0x5c, 0xa3, 0x8, 0x47, 0x44, 0x6d, 0x96, 0xe0}}
	return a, nil
}

var __1648554928_network_testUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcc\x31\x0a\x02\x31\x10\x05\xd0\xde\x53\xfc\x23\xd8\x2f\x16\xb3\x66\xb6\xfa\x26\xe2\x4e\xea\xa0\x38\x88\x28\x11\xcc\x80\xd7\xb7\xb5\xf1\x02\x4f\x68\x7a\x82\xc9\x4c\xc5\xf0\x88\x7b\xbf\x0d\x48\x4a\xd8\x17\xd6\x43\x46\xf8\x88\xd6\x3d\x3e\xaf\xf7\x63\x34\xef\xe7\xcb\xd3\xaf\x98\x4b\xa1\x4a\x46\x2e\x86\x5c\x49\x24\x5d\xa4\xd2\xb0\x08\x57\x9d\x36\xf5\x98\xc4\x7e\xc0\x55\xed\x8f\xb4\xc3\x76\xfa\x06\x00\x00\xff\xff\x36\x2f\x0a\x06\x84\x00\x00\x00")

func _1648554928_network_testUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1648554928_network_testUpSql,
		"1648554928_network_test.up.sql",
	)
}

func _1648554928_network_testUpSql() (*asset, error) {
	bytes, err := _1648554928_network_testUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1648554928_network_test.up.sql", size: 132, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9a, 0xc5, 0x7f, 0x87, 0xf3, 0x2c, 0xf7, 0xbb, 0xd3, 0x3a, 0x4e, 0x76, 0x88, 0xca, 0xaf, 0x73, 0xce, 0x8f, 0xa1, 0xf6, 0x3d, 0x4d, 0xed, 0x6f, 0x49, 0xf2, 0xfe, 0x56, 0x2a, 0x60, 0x68, 0xca}}
	return a, nil
}

var __1649174829_add_visitble_tokenUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x28\xcb\x2c\xce\x4c\xca\x49\x8d\x2f\xc9\xcf\x4e\xcd\x2b\x56\xd0\xe0\x52\x50\x48\xce\x48\xcc\xcc\x8b\xcf\x4c\x51\x08\xf5\x0b\xf6\x74\xf7\x73\x75\x51\xf0\xf4\x0b\xd1\xe1\x52\x50\x48\x4c\x49\x29\x4a\x2d\x2e\x56\x08\x73\x0c\x72\xf6\x70\x0c\x52\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\xe1\xd2\xb4\x06\x04\x00\x00\xff\xff\xa0\x5f\x37\x13\x54\x00\x00\x00")

func _1649174829_add_visitble_tokenUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1649174829_add_visitble_tokenUpSql,
		"1649174829_add_visitble_token.up.sql",
	)
}

func _1649174829_add_visitble_tokenUpSql() (*asset, error) {
	bytes, err := _1649174829_add_visitble_tokenUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1649174829_add_visitble_token.up.sql", size: 84, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa3, 0x22, 0xc0, 0x2b, 0x3f, 0x4f, 0x3d, 0x5e, 0x4c, 0x68, 0x7c, 0xd0, 0x15, 0x36, 0x9f, 0xec, 0xa1, 0x2a, 0x7b, 0xb4, 0xe3, 0xc6, 0xc9, 0xb4, 0x81, 0x50, 0x4a, 0x11, 0x3b, 0x35, 0x7, 0xcf}}
	return a, nil
}

var __1649882262_add_derived_from_accountsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4c\x4e\xce\x2f\xcd\x2b\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x49\x2d\xca\x2c\x4b\x4d\x89\x4f\x2b\xca\xcf\x55\x08\x71\x8d\x08\x51\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x51\x50\x52\xb2\xe6\x0a\x0d\x70\x71\x0c\x41\x32\x21\xd8\x35\x04\x55\xab\x2d\x58\x19\x20\x00\x00\xff\xff\x3e\x48\xa7\x03\x6e\x00\x00\x00")

func _1649882262_add_derived_from_accountsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1649882262_add_derived_from_accountsUpSql,
		"1649882262_add_derived_from_accounts.up.sql",
	)
}

func _1649882262_add_derived_from_accountsUpSql() (*asset, error) {
	bytes, err := _1649882262_add_derived_from_accountsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1649882262_add_derived_from_accounts.up.sql", size: 110, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x11, 0xb9, 0x44, 0x4d, 0x85, 0x8d, 0x7f, 0xb4, 0xae, 0x4f, 0x5c, 0x66, 0x64, 0xb6, 0xe2, 0xe, 0x3d, 0xad, 0x9d, 0x8, 0x4f, 0xab, 0x6e, 0xa8, 0x7d, 0x76, 0x3, 0xad, 0x96, 0x1, 0xee, 0x5c}}
	return a, nil
}

var __1650612625_add_community_message_archive_hashes_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x48\xce\xcf\xcd\x2d\xcd\xcb\x2c\xa9\x8c\xcf\x4d\x2d\x2e\x4e\x4c\x4f\x8d\x4f\x2c\x4a\xce\xc8\x2c\x4b\x8d\xcf\x48\x2c\xce\x48\x2d\x56\xd0\xe0\x52\x40\x52\x94\x99\xa2\x10\xe2\x1a\x11\x02\x36\xc1\x2f\xd4\xc7\x47\x87\x4b\x41\x01\xa4\x10\x22\x1a\x10\xe4\xe9\xeb\x18\x14\xa9\xe0\xed\x1a\x09\x57\xc1\xa5\x69\xcd\xc5\xc5\x05\x08\x00\x00\xff\xff\xc1\xb0\x9f\xaa\x82\x00\x00\x00")

func _1650612625_add_community_message_archive_hashes_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1650612625_add_community_message_archive_hashes_tableUpSql,
		"1650612625_add_community_message_archive_hashes_table.up.sql",
	)
}

func _1650612625_add_community_message_archive_hashes_tableUpSql() (*asset, error) {
	bytes, err := _1650612625_add_community_message_archive_hashes_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1650612625_add_community_message_archive_hashes_table.up.sql", size: 130, mode: os.FileMode(0664), modTime: time.Unix(1652164735, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x48, 0x31, 0xb3, 0x75, 0x23, 0xe2, 0x45, 0xe, 0x47, 0x1b, 0x35, 0xa5, 0x6e, 0x83, 0x4e, 0x64, 0x7d, 0xd7, 0xa2, 0xda, 0xe9, 0x53, 0xf1, 0x16, 0x86, 0x2c, 0x57, 0xad, 0xfa, 0xca, 0x39, 0xde}}
	return a, nil
}

var __1650616788_add_communities_archives_info_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcd\xb1\x4a\xc5\x30\x18\xc5\xf1\x3d\x4f\x71\x46\x05\x07\x5f\x21\xc6\xaf\x10\x8c\xe9\x25\xf7\xbb\xd0\x4e\x21\xa4\xb1\x86\x36\x29\x98\x28\xf8\xf6\xa2\x43\x37\xe7\x73\xfe\xfc\x94\x23\xc9\x04\x96\x4f\x86\xa0\x07\xd8\x91\x41\x93\xbe\xf2\x15\xf1\x28\xe5\xb3\xe6\x9e\x53\xf3\xe1\x23\xbe\xe7\xaf\xe4\x73\x7d\x3b\x70\x27\x70\x8e\xdf\x3e\x2f\x60\x9a\x18\x17\xa7\x5f\xa5\x9b\xf1\x42\x33\x46\x0b\x35\xda\xc1\x68\xc5\x70\x74\x31\x52\xd1\x83\x00\x4a\x58\x6b\xea\x7b\xae\x9b\x8f\xfb\x11\x37\x68\xcb\x7f\xa2\xbd\x19\x83\x67\x1a\xe4\xcd\x30\x1e\x7f\xaf\x7b\x68\xdd\x97\xd4\x5a\x58\xd3\xa9\xa7\xba\xf8\x25\xf4\xf4\x4f\x27\xee\x85\xf8\x09\x00\x00\xff\xff\x74\x80\x98\x81\xd0\x00\x00\x00")

func _1650616788_add_communities_archives_info_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1650616788_add_communities_archives_info_tableUpSql,
		"1650616788_add_communities_archives_info_table.up.sql",
	)
}

func _1650616788_add_communities_archives_info_tableUpSql() (*asset, error) {
	bytes, err := _1650616788_add_communities_archives_info_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1650616788_add_communities_archives_info_table.up.sql", size: 208, mode: os.FileMode(0664), modTime: time.Unix(1652164735, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd1, 0x4f, 0x80, 0x45, 0xb9, 0xd9, 0x15, 0xe2, 0x78, 0xd0, 0xcb, 0x71, 0xc1, 0x1b, 0xb7, 0x1b, 0x1b, 0x97, 0xfe, 0x47, 0x53, 0x3c, 0x62, 0xbc, 0xdd, 0x3a, 0x94, 0x1a, 0xc, 0x48, 0x76, 0xe}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x0d\xc4\x20\x0c\x05\xd0\x9e\x29\xfe\x02\xd8\xfd\x6d\xe3\x4b\xac\x2f\x44\x82\x09\x78\x7f\xa5\x49\xfd\xa6\x1d\xdd\xe8\xd8\xcf\x55\x8a\x2a\xe3\x47\x1f\xbe\x2c\x1d\x8c\xfa\x6f\xe3\xb4\x34\xd4\xd9\x89\xbb\x71\x59\xb6\x18\x1b\x35\x20\xa2\x9f\x0a\x03\xa2\xe5\x0d\x00\x00\xff\xff\x60\xcd\x06\xbe\x4a\x00\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 74, mode: os.FileMode(0664), modTime: time.Unix(1651740732, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xde, 0x7c, 0x28, 0xcd, 0x47, 0xf2, 0xfa, 0x7c, 0x51, 0x2d, 0xd8, 0x38, 0xb, 0xb0, 0x34, 0x9d, 0x4c, 0x62, 0xa, 0x9e, 0x28, 0xc3, 0x31, 0x23, 0xd9, 0xbb, 0x89, 0x9f, 0xa0, 0x89, 0x1f, 0xe8}}
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
	"1640111208_dummy.up.sql": _1640111208_dummyUpSql,

	"1642666031_add_removed_clock_to_bookmarks.up.sql": _1642666031_add_removed_clock_to_bookmarksUpSql,

	"1643644541_gif_api_key_setting.up.sql": _1643644541_gif_api_key_settingUpSql,

	"1644188994_recent_stickers.up.sql": _1644188994_recent_stickersUpSql,

	"1646659233_add_address_to_dapp_permisssion.up.sql": _1646659233_add_address_to_dapp_permisssionUpSql,

	"1646841105_add_emoji_account.up.sql": _1646841105_add_emoji_accountUpSql,

	"1647278782_display_name.up.sql": _1647278782_display_nameUpSql,

	"1647862838_reset_last_backup.up.sql": _1647862838_reset_last_backupUpSql,

	"1647871652_add_settings_sync_clock_table.up.sql": _1647871652_add_settings_sync_clock_tableUpSql,

	"1647880168_add_torrent_config.up.sql": _1647880168_add_torrent_configUpSql,

	"1647882837_add_communities_settings_table.up.sql": _1647882837_add_communities_settings_tableUpSql,

	"1647956635_add_waku_messages_table.up.sql": _1647956635_add_waku_messages_tableUpSql,

	"1648554928_network_test.up.sql": _1648554928_network_testUpSql,

	"1649174829_add_visitble_token.up.sql": _1649174829_add_visitble_tokenUpSql,

	"1649882262_add_derived_from_accounts.up.sql": _1649882262_add_derived_from_accountsUpSql,

	"1650612625_add_community_message_archive_hashes_table.up.sql": _1650612625_add_community_message_archive_hashes_tableUpSql,

	"1650616788_add_communities_archives_info_table.up.sql": _1650616788_add_communities_archives_info_tableUpSql,

	"doc.go": docGo,
}

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
	"1640111208_dummy.up.sql":                                      &bintree{_1640111208_dummyUpSql, map[string]*bintree{}},
	"1642666031_add_removed_clock_to_bookmarks.up.sql":             &bintree{_1642666031_add_removed_clock_to_bookmarksUpSql, map[string]*bintree{}},
	"1643644541_gif_api_key_setting.up.sql":                        &bintree{_1643644541_gif_api_key_settingUpSql, map[string]*bintree{}},
	"1644188994_recent_stickers.up.sql":                            &bintree{_1644188994_recent_stickersUpSql, map[string]*bintree{}},
	"1646659233_add_address_to_dapp_permisssion.up.sql":            &bintree{_1646659233_add_address_to_dapp_permisssionUpSql, map[string]*bintree{}},
	"1646841105_add_emoji_account.up.sql":                          &bintree{_1646841105_add_emoji_accountUpSql, map[string]*bintree{}},
	"1647278782_display_name.up.sql":                               &bintree{_1647278782_display_nameUpSql, map[string]*bintree{}},
	"1647862838_reset_last_backup.up.sql":                          &bintree{_1647862838_reset_last_backupUpSql, map[string]*bintree{}},
	"1647871652_add_settings_sync_clock_table.up.sql":              &bintree{_1647871652_add_settings_sync_clock_tableUpSql, map[string]*bintree{}},
	"1647880168_add_torrent_config.up.sql":                         &bintree{_1647880168_add_torrent_configUpSql, map[string]*bintree{}},
	"1647882837_add_communities_settings_table.up.sql":             &bintree{_1647882837_add_communities_settings_tableUpSql, map[string]*bintree{}},
	"1647956635_add_waku_messages_table.up.sql":                    &bintree{_1647956635_add_waku_messages_tableUpSql, map[string]*bintree{}},
	"1648554928_network_test.up.sql":                               &bintree{_1648554928_network_testUpSql, map[string]*bintree{}},
	"1649174829_add_visitble_token.up.sql":                         &bintree{_1649174829_add_visitble_tokenUpSql, map[string]*bintree{}},
	"1649882262_add_derived_from_accounts.up.sql":                  &bintree{_1649882262_add_derived_from_accountsUpSql, map[string]*bintree{}},
	"1650612625_add_community_message_archive_hashes_table.up.sql": &bintree{_1650612625_add_community_message_archive_hashes_tableUpSql, map[string]*bintree{}},
	"1650616788_add_communities_archives_info_table.up.sql":        &bintree{_1650616788_add_communities_archives_info_tableUpSql, map[string]*bintree{}},
	"doc.go": &bintree{docGo, map[string]*bintree{}},
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
