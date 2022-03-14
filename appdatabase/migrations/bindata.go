// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1640111208_dummy.up.sql (258B)
// 1642666031_add_removed_clock_to_bookmarks.up.sql (117B)
// 1643644541_gif_api_key_setting.up.sql (108B)
// 1644188994_recent_stickers.up.sql (79B)
// 1646659233_add_address_to_dapp_permisssion.up.sql (700B)
// 1646841105_add_emoji_account.up.sql (96B)
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

	info := bindataFileInfo{name: "1640111208_dummy.up.sql", size: 258, mode: os.FileMode(0664), modTime: time.Unix(1647246959, 0)}
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

	info := bindataFileInfo{name: "1642666031_add_removed_clock_to_bookmarks.up.sql", size: 117, mode: os.FileMode(0664), modTime: time.Unix(1647246959, 0)}
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

	info := bindataFileInfo{name: "1643644541_gif_api_key_setting.up.sql", size: 108, mode: os.FileMode(0664), modTime: time.Unix(1647246959, 0)}
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

	info := bindataFileInfo{name: "1644188994_recent_stickers.up.sql", size: 79, mode: os.FileMode(0664), modTime: time.Unix(1647246959, 0)}
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

	info := bindataFileInfo{name: "1646659233_add_address_to_dapp_permisssion.up.sql", size: 700, mode: os.FileMode(0664), modTime: time.Unix(1647247370, 0)}
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

	info := bindataFileInfo{name: "1646841105_add_emoji_account.up.sql", size: 96, mode: os.FileMode(0664), modTime: time.Unix(1647247370, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe6, 0x77, 0x29, 0x95, 0x18, 0x64, 0x82, 0x63, 0xe7, 0xaf, 0x6c, 0xa9, 0x15, 0x7d, 0x46, 0xa6, 0xbc, 0xdf, 0xa7, 0xd, 0x2b, 0xd2, 0x2d, 0x97, 0x4d, 0xa, 0x6b, 0xd, 0x6e, 0x90, 0x42, 0x5c}}
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

	info := bindataFileInfo{name: "doc.go", size: 74, mode: os.FileMode(0664), modTime: time.Unix(1647246959, 0)}
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
	"1640111208_dummy.up.sql":                           _1640111208_dummyUpSql,
	"1642666031_add_removed_clock_to_bookmarks.up.sql":  _1642666031_add_removed_clock_to_bookmarksUpSql,
	"1643644541_gif_api_key_setting.up.sql":             _1643644541_gif_api_key_settingUpSql,
	"1644188994_recent_stickers.up.sql":                 _1644188994_recent_stickersUpSql,
	"1646659233_add_address_to_dapp_permisssion.up.sql": _1646659233_add_address_to_dapp_permisssionUpSql,
	"1646841105_add_emoji_account.up.sql":               _1646841105_add_emoji_accountUpSql,
	"doc.go":                                            docGo,
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
	"1640111208_dummy.up.sql": {_1640111208_dummyUpSql, map[string]*bintree{}},
	"1642666031_add_removed_clock_to_bookmarks.up.sql": {_1642666031_add_removed_clock_to_bookmarksUpSql, map[string]*bintree{}},
	"1643644541_gif_api_key_setting.up.sql": {_1643644541_gif_api_key_settingUpSql, map[string]*bintree{}},
	"1644188994_recent_stickers.up.sql": {_1644188994_recent_stickersUpSql, map[string]*bintree{}},
	"1646659233_add_address_to_dapp_permisssion.up.sql": {_1646659233_add_address_to_dapp_permisssionUpSql, map[string]*bintree{}},
	"1646841105_add_emoji_account.up.sql": {_1646841105_add_emoji_accountUpSql, map[string]*bintree{}},
	"doc.go": {docGo, map[string]*bintree{}},
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
