// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1593601728_initial_schema.down.sql (200B)
// 1593601728_initial_schema.up.sql (675B)
// 1598419937_add_push_notifications_table.down.sql (51B)
// 1598419937_add_push_notifications_table.up.sql (104B)
// doc.go (382B)

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

var __1593601728_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x28\x2d\xce\x88\xcf\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x8b\x2f\x4e\x2d\x2a\x4b\x2d\x8a\x2f\x4a\x4d\xcf\x2c\x2e\x29\x02\x8b\x15\x5b\x73\x81\xb5\x78\xfa\xb9\xb8\x46\x28\x64\xa6\x54\xc4\x13\xa7\x2d\xbe\xa0\x34\x29\x27\x33\x39\x3e\x3b\xb5\x92\x72\x13\xe2\x33\xf3\x8a\x4b\x12\x73\x72\x20\x8a\x33\x53\xac\xb9\xb8\x00\x01\x00\x00\xff\xff\x90\x39\xe0\x1c\xc8\x00\x00\x00")

func _1593601728_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1593601728_initial_schemaDownSql,
		"1593601728_initial_schema.down.sql",
	)
}

func _1593601728_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1593601728_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1593601728_initial_schema.down.sql", size: 200, mode: os.FileMode(0644), modTime: time.Unix(1697447578, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x88, 0x8a, 0x61, 0x81, 0x57, 0x45, 0x9b, 0x97, 0x9b, 0x1f, 0xf6, 0x94, 0x8a, 0x20, 0xb3, 0x2b, 0xff, 0x69, 0x49, 0xf4, 0x58, 0xcc, 0xd0, 0x55, 0xcc, 0x9a, 0x8b, 0xb6, 0x7f, 0x29, 0x53, 0xc1}}
	return a, nil
}

var __1593601728_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\x31\x6b\xc3\x30\x14\x84\x77\xfd\x8a\x37\xc6\x90\xa1\xbb\x27\xd9\x91\xa9\x40\x48\xad\x23\x97\x6c\xc2\xb5\xd5\xe6\x51\x23\x07\x49\x31\xf5\xbf\x2f\x71\x86\x2a\x69\x87\x10\xb2\x1e\x8f\xbb\xf7\xdd\x95\x35\xa3\x9a\x81\xa6\x85\x60\xc0\x2b\x90\x4a\x03\xdb\xf1\xad\xde\xc2\xe1\x18\xf6\xc6\x8d\x11\x3f\xb0\x6b\x23\x8e\xce\x04\xeb\x27\xeb\x8d\xb7\x9f\x18\xa2\x5f\xb4\x00\x2b\x02\x70\x38\xbe\x0f\xd8\x99\x2f\x3b\x43\x21\x54\xb1\xb8\xc8\x46\x88\x35\x01\x40\x17\x62\x3b\x0c\x67\x07\xec\xe1\x8d\xd6\xe5\x33\xad\x2f\x6e\x26\xeb\x03\x8e\x0e\xb8\xd4\x17\x7a\x9a\xb4\x38\x9f\xc4\x46\xf2\xd7\x86\xad\x7e\x33\xd7\xd7\x19\x19\x28\x09\xa5\x92\x95\xe0\xa5\x86\x9a\xbd\x08\x5a\x32\x92\xe5\x84\xdc\x83\x8b\xbd\x75\x11\xe3\x7c\x26\xf5\x38\xb5\xd1\xfe\x8f\x1a\x66\x17\xf7\x36\x62\x77\xe2\x4c\x59\x60\xc3\x2a\xda\x08\x0d\x4f\x09\x40\x7a\x9d\xa5\xdf\x71\xb9\x61\x3b\xc0\xfe\xdb\xdc\x36\x81\x49\xea\x57\xf2\xc6\xdd\x92\xfe\xb2\xfc\x01\xc9\xe6\x7a\xe7\x7b\x3e\xf9\xbb\x64\x4e\xc8\x4f\x00\x00\x00\xff\xff\xcc\xa0\x4d\x54\xa3\x02\x00\x00")

func _1593601728_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1593601728_initial_schemaUpSql,
		"1593601728_initial_schema.up.sql",
	)
}

func _1593601728_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1593601728_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1593601728_initial_schema.up.sql", size: 675, mode: os.FileMode(0644), modTime: time.Unix(1697447578, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfd, 0x61, 0x90, 0x79, 0xd9, 0x14, 0x65, 0xe9, 0x96, 0x53, 0x17, 0x33, 0x54, 0xeb, 0x8b, 0x5d, 0x95, 0x99, 0x10, 0x36, 0x58, 0xdd, 0xb2, 0xbf, 0x45, 0xd9, 0xbb, 0xc4, 0x92, 0xe, 0xce, 0x2}}
	return a, nil
}

var __1598419937_add_push_notifications_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x28\x2d\xce\x88\xcf\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x8b\x2f\x4e\x2d\x2a\x4b\x2d\x42\x11\x2b\xb6\xe6\x02\x04\x00\x00\xff\xff\xb7\xdc\x38\x53\x33\x00\x00\x00")

func _1598419937_add_push_notifications_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1598419937_add_push_notifications_tableDownSql,
		"1598419937_add_push_notifications_table.down.sql",
	)
}

func _1598419937_add_push_notifications_tableDownSql() (*asset, error) {
	bytes, err := _1598419937_add_push_notifications_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1598419937_add_push_notifications_table.down.sql", size: 51, mode: os.FileMode(0644), modTime: time.Unix(1697447578, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0x98, 0xc8, 0x30, 0x45, 0x5b, 0xc5, 0x7d, 0x13, 0x5d, 0xe7, 0xc8, 0x23, 0x43, 0xf7, 0xdc, 0x9c, 0xe2, 0xdd, 0x63, 0xf0, 0xb7, 0x16, 0x40, 0xc, 0xda, 0xb9, 0x16, 0x70, 0x2b, 0x5a, 0x7e}}
	return a, nil
}

var __1598419937_add_push_notifications_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x28\x2d\xce\x88\xcf\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x8b\x2f\x4e\x2d\x2a\x4b\x2d\x42\x11\x2b\x56\xd0\xe0\x52\x50\xc8\x4c\x51\x70\xf2\xf1\x77\x02\xeb\xf6\x0b\xf5\xf1\xd1\xe1\x52\x50\x08\xf5\xf3\x0c\x0c\x75\xd5\xc8\x4c\xd1\xe4\xd2\xb4\xe6\x02\x04\x00\x00\xff\xff\xf3\xc8\x52\x6b\x68\x00\x00\x00")

func _1598419937_add_push_notifications_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1598419937_add_push_notifications_tableUpSql,
		"1598419937_add_push_notifications_table.up.sql",
	)
}

func _1598419937_add_push_notifications_tableUpSql() (*asset, error) {
	bytes, err := _1598419937_add_push_notifications_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1598419937_add_push_notifications_table.up.sql", size: 104, mode: os.FileMode(0644), modTime: time.Unix(1697447578, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2, 0x3e, 0xef, 0xf, 0xc2, 0xdf, 0xbc, 0x99, 0x7a, 0xc2, 0xd3, 0x64, 0x4f, 0x4c, 0x7e, 0xfc, 0x2e, 0x8c, 0xa7, 0x54, 0xd3, 0x4d, 0x25, 0x98, 0x41, 0xbc, 0xea, 0xd7, 0x2, 0xc1, 0xd0, 0x52}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x3d\x6e\xec\x30\x0c\x84\x7b\x9d\x62\xb0\xcd\x36\xcf\x52\xf3\xaa\x74\x29\xd3\xe7\x02\x5c\x89\x96\x88\xb5\x24\x43\xa4\xf7\xe7\xf6\x81\x37\x01\xe2\x2e\xed\x87\xf9\x86\xc3\x10\xf0\x59\x44\x31\xcb\xc2\x10\x45\xe3\xc8\xaa\x34\x9e\xb8\x70\xa4\x4d\x19\xa7\x2c\x56\xb6\x8b\x8f\xbd\x06\x35\xb2\x4d\x27\xa9\xa1\x4a\x1e\x64\x1c\x6e\xff\x4f\x2e\x04\x44\x6a\x67\x43\xa1\x96\x16\x7e\x75\x29\xd4\x68\x98\xb4\x8c\xbb\x58\x01\x61\x1d\x3c\xcb\xc3\xe3\xdd\xb0\x30\xa9\xc1\x0a\xd9\x59\x61\x85\x11\x49\x79\xaf\x99\xfb\x40\xee\xd3\x45\x5a\x22\x23\xbf\xa3\x8f\xf9\x40\xf6\x85\x91\x96\x85\x13\xe6\xd1\xeb\xcb\x55\xaa\x8c\x24\x83\xa3\xf5\xf1\xfc\x07\x52\x65\x43\xa3\xca\xba\xfb\x85\x6e\x8c\xd6\x7f\xce\x83\x5a\xfa\xfb\x23\xdc\xfb\xb8\x2a\x48\xc1\x8f\x95\xa3\x71\xf2\xce\xad\x14\xaf\x94\x19\xdf\x39\xe9\x4d\x9d\x0b\x21\xf7\xb7\xcc\x8d\x77\xf3\xb8\x73\x5a\xaf\xf9\x90\xc4\xd4\xe1\x7d\xf8\x05\x3e\x77\xf8\xe0\xbe\x02\x00\x00\xff\xff\x4d\x1d\x5d\x50\x7e\x01\x00\x00")

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

	info := bindataFileInfo{name: "doc.go", size: 382, mode: os.FileMode(0644), modTime: time.Unix(1697447578, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc0, 0x2f, 0x1e, 0x64, 0x9, 0x93, 0xe4, 0x8b, 0xf2, 0x98, 0x5a, 0x45, 0xe2, 0x80, 0x88, 0x67, 0x7a, 0x2d, 0xd7, 0x4b, 0xd1, 0x73, 0xb6, 0x6d, 0x15, 0xc2, 0x0, 0x34, 0xcd, 0xa0, 0xdb, 0x20}}
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
	"1593601728_initial_schema.down.sql": _1593601728_initial_schemaDownSql,

	"1593601728_initial_schema.up.sql": _1593601728_initial_schemaUpSql,

	"1598419937_add_push_notifications_table.down.sql": _1598419937_add_push_notifications_tableDownSql,

	"1598419937_add_push_notifications_table.up.sql": _1598419937_add_push_notifications_tableUpSql,

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
	"1593601728_initial_schema.down.sql":               &bintree{_1593601728_initial_schemaDownSql, map[string]*bintree{}},
	"1593601728_initial_schema.up.sql":                 &bintree{_1593601728_initial_schemaUpSql, map[string]*bintree{}},
	"1598419937_add_push_notifications_table.down.sql": &bintree{_1598419937_add_push_notifications_tableDownSql, map[string]*bintree{}},
	"1598419937_add_push_notifications_table.up.sql":   &bintree{_1598419937_add_push_notifications_tableUpSql, map[string]*bintree{}},
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
