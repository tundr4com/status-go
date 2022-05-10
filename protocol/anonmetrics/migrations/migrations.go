// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1619446565_postgres_make_anon_metrics_table.down.sql (24B)
// 1619446565_postgres_make_anon_metrics_table.up.sql (443B)
// doc.go (380B)

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

var __1619446565_postgres_make_anon_metrics_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x2c\x28\x88\xcf\x4d\x2d\x29\xca\x4c\x2e\xb6\xe6\x02\x04\x00\x00\xff\xff\x99\xa7\x42\x7d\x18\x00\x00\x00")

func _1619446565_postgres_make_anon_metrics_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1619446565_postgres_make_anon_metrics_tableDownSql,
		"1619446565_postgres_make_anon_metrics_table.down.sql",
	)
}

func _1619446565_postgres_make_anon_metrics_tableDownSql() (*asset, error) {
	bytes, err := _1619446565_postgres_make_anon_metrics_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1619446565_postgres_make_anon_metrics_table.down.sql", size: 24, mode: os.FileMode(0664), modTime: time.Unix(1650871814, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x75, 0xea, 0x1, 0x74, 0xe6, 0xa3, 0x11, 0xd0, 0x86, 0x87, 0x7e, 0x31, 0xb4, 0x1a, 0x27, 0x5d, 0xda, 0x77, 0xa3, 0xf5, 0x1d, 0x88, 0x79, 0xcf, 0xd5, 0x95, 0x75, 0xd, 0x47, 0xa1, 0x90, 0x5}}
	return a, nil
}

var __1619446565_postgres_make_anon_metrics_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x4d\x6a\xc3\x30\x14\x84\xf7\x39\xc5\x2c\xdb\x10\xc8\x01\xba\x52\xdc\x17\xea\x56\xb6\x53\x59\x2e\x64\x65\x84\xfd\x30\x82\xfa\x07\x49\x75\xe9\xed\x8b\x1d\x92\x38\x81\xac\xbf\xef\x69\x46\x13\x29\x12\x9a\xa0\xc5\x4e\x12\xcc\x30\x94\x2d\x07\x67\x2b\x8f\xa7\x15\x00\xd8\x1a\x39\xa9\x58\x48\x1c\x54\x9c\x08\x75\xc4\x07\x1d\x37\xab\x99\x6d\xd7\x88\xbb\xaa\x6f\x6d\xd7\xe0\x74\x85\xda\x04\x83\xf5\x76\xc6\x2d\x7b\x6f\x1a\x2e\x6d\x8d\x2f\xa1\xa2\x37\xa1\x50\xa4\xf1\x67\x41\x48\x33\x8d\xb4\x90\x72\x33\x7b\x3c\x72\x17\x2e\xca\x2d\x1b\xcd\xf7\x0f\xe3\x3d\xcf\xd2\x3b\x30\x35\x1d\xd9\x79\xdb\x77\x0f\x4e\xfb\x81\x9d\x09\xb6\x6b\x4a\xff\xe7\x03\xb7\x0f\x34\xcf\x7e\x7a\x64\xd9\xf2\x56\xa8\x1c\x9b\xc0\x75\x69\x02\x74\x9c\x50\xae\x45\x72\x58\x28\xe7\x25\x54\xff\x3b\x8d\x60\x96\x0b\x0c\xae\xaf\xd8\x7b\xae\xb1\xcb\x32\x49\xe2\xfa\x09\xbc\xd2\x5e\x14\x52\x63\x2f\x64\x4e\xa7\x20\xc7\x15\xdb\xf1\x3e\xe9\x2c\x46\x85\x52\x94\xea\xf2\x42\x9e\x5f\xfe\x03\x00\x00\xff\xff\xee\x42\x32\x03\xbb\x01\x00\x00")

func _1619446565_postgres_make_anon_metrics_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1619446565_postgres_make_anon_metrics_tableUpSql,
		"1619446565_postgres_make_anon_metrics_table.up.sql",
	)
}

func _1619446565_postgres_make_anon_metrics_tableUpSql() (*asset, error) {
	bytes, err := _1619446565_postgres_make_anon_metrics_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1619446565_postgres_make_anon_metrics_table.up.sql", size: 443, mode: os.FileMode(0664), modTime: time.Unix(1650871814, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd5, 0xdc, 0x72, 0x28, 0x3c, 0xf6, 0x94, 0xb0, 0x47, 0x3d, 0xca, 0x55, 0x3d, 0xf7, 0x83, 0xb8, 0x7d, 0x2f, 0x1e, 0x98, 0xb7, 0xde, 0xa, 0xff, 0xa0, 0x52, 0x60, 0x83, 0x56, 0xc5, 0xd1, 0xa2}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\xbd\x6e\xf3\x30\x0c\x45\x77\x3f\xc5\x45\x96\x2c\x9f\xa5\xe5\x9b\xba\x75\xec\xde\x17\x60\xe4\x6b\x49\x88\x2d\x1a\x22\xf3\xf7\xf6\x85\xd3\x02\xcd\xd6\xf5\x00\xe7\xf0\x32\x46\x7c\x96\x6a\x98\xeb\x42\x54\x43\x63\xa2\x99\xf4\x07\x4e\x4c\x72\x31\xe2\x90\xab\x97\xcb\x29\x24\x5d\xa3\xb9\xf8\xc5\xc6\xba\xc6\xb5\xe6\x2e\xce\x78\xfd\x7f\x18\x62\x44\x92\x76\x74\x14\x69\xd3\xc2\x67\xcb\x60\x2e\xdd\x6b\xcb\xb8\x55\x2f\x10\x6c\x9d\x73\xbd\x07\xbc\x3b\x16\x8a\x39\xbc\x88\x1f\x0d\x5e\x88\x24\xc6\x3d\x33\x6b\x47\xd6\xf1\x54\xdb\x24\x2e\x61\x47\x1f\xf3\x0b\xd9\x17\x26\x59\x16\x4e\x98\xbb\xae\x4f\xd7\x64\x25\xa6\xda\x99\x5c\xfb\xe3\x1f\xc4\x8c\x8e\x26\x2b\x6d\xf7\x8b\x5c\x89\xa6\x3f\xe7\x21\x6d\xfa\xfb\x23\xdc\xb4\x9f\x0d\x62\xe0\x7d\x63\x72\x4e\x61\x18\x36\x49\x67\xc9\xc4\xa6\xe6\xb9\xd3\x86\x21\xc6\xac\x6f\x99\x8d\xbb\xf7\xba\x72\xdc\xce\x19\xdf\xbd\xaa\xcd\x30\x2a\x42\x88\xbf\x20\x64\x45\x88\xc3\x57\x00\x00\x00\xff\xff\xa9\xf1\x73\x83\x7c\x01\x00\x00")

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

	info := bindataFileInfo{name: "doc.go", size: 380, mode: os.FileMode(0664), modTime: time.Unix(1650871814, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x49, 0x1, 0xd4, 0xd6, 0xc7, 0x44, 0xd4, 0xfd, 0x7b, 0x69, 0x1f, 0xe3, 0xe, 0x48, 0x14, 0x99, 0xf0, 0x8e, 0x43, 0xae, 0x54, 0x64, 0xa2, 0x8b, 0x82, 0x1c, 0x2b, 0xb, 0xec, 0xf5, 0xb3, 0xfc}}
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
	"1619446565_postgres_make_anon_metrics_table.down.sql": _1619446565_postgres_make_anon_metrics_tableDownSql,

	"1619446565_postgres_make_anon_metrics_table.up.sql": _1619446565_postgres_make_anon_metrics_tableUpSql,

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
	"1619446565_postgres_make_anon_metrics_table.down.sql": &bintree{_1619446565_postgres_make_anon_metrics_tableDownSql, map[string]*bintree{}},
	"1619446565_postgres_make_anon_metrics_table.up.sql":   &bintree{_1619446565_postgres_make_anon_metrics_tableUpSql, map[string]*bintree{}},
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
