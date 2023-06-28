// Code generated by go-bindata.
// sources:
// 1593601729_initial_schema.down.sql
// 1593601729_initial_schema.up.sql
// 1597909626_add_server_type.down.sql
// 1597909626_add_server_type.up.sql
// 1599053776_add_chat_id_and_type.down.sql
// 1599053776_add_chat_id_and_type.up.sql
// doc.go
// DO NOT EDIT!

package migrations

import (
	"bytes"
	"compress/gzip"
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
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

var __1593601729_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x28\x2d\xce\x88\xcf\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x8b\x4f\xce\xc9\x4c\xcd\x2b\x89\x2f\x4e\x2d\x2a\x4b\x2d\x2a\xb6\xe6\x22\x46\x71\x66\x5e\x5a\x3e\x54\xa5\xa7\x9f\x8b\x6b\x84\x42\x66\x4a\x45\x3c\x5e\xd5\xf1\x05\xa5\x49\x39\x99\xc9\xf1\xd9\xa9\x95\xd6\x5c\x80\x00\x00\x00\xff\xff\x6d\xb4\xf8\x65\x90\x00\x00\x00")

func _1593601729_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1593601729_initial_schemaDownSql,
		"1593601729_initial_schema.down.sql",
	)
}

func _1593601729_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1593601729_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1593601729_initial_schema.down.sql", size: 144, mode: os.FileMode(436), modTime: time.Unix(1687862441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1593601729_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\xc1\x8e\xdb\x20\x10\xbd\xfb\x2b\xe6\xb8\x91\x72\xe8\x7d\x4f\x4e\x96\x54\x96\x10\x6e\x13\x22\xe5\x86\x28\x9e\x5d\xa3\xb8\x78\x0b\x78\xd5\xfc\x7d\x85\x9d\x78\x93\xc5\xc5\x55\xba\x17\x4b\x1e\x1e\xa3\x79\x6f\x1e\x6f\xbd\x25\x39\x27\xc0\xf3\x15\x25\x50\x6c\x80\x95\x1c\xc8\xa1\xd8\xf1\x1d\xbc\x76\xae\x16\xa6\xf5\xfa\x59\x2b\xe9\x75\x6b\x84\x6a\x34\x1a\x2f\x1c\xda\x37\xb4\x0e\x1e\x32\x80\xd7\xee\x47\xa3\x95\x38\xe2\x09\x56\xb4\x5c\xf5\xf7\xd9\x9e\xd2\x65\x06\x60\xf1\x45\x3b\x8f\x16\x2b\x58\x95\x25\x25\x39\x83\x27\xb2\xc9\xf7\x94\xc3\x26\xa7\x3b\x72\x8b\x11\xd2\x43\xc1\xf8\xd8\x61\xc4\x7e\x09\xb8\x46\x3a\x2f\x2c\x7a\xab\xe7\x90\x01\x74\x12\xaa\xed\x4c\x0a\x25\x95\x42\xe7\x84\x6f\x8f\x68\x80\x93\x03\x0f\xc5\x3d\x2b\xbe\xef\xc9\xc3\x3b\xa7\x05\x94\x0c\xd6\x25\xdb\xd0\x62\xcd\x61\x4b\xbe\xd1\x7c\x4d\xb2\xc5\x63\x96\xdd\xa3\xdb\xaf\x0e\xad\xc6\x79\xdd\x06\x5c\x44\xf3\x72\x74\x12\xba\x8a\x2f\x45\xb3\x2f\x2f\xd8\xcf\x25\xa1\xcd\x73\x3b\xcb\x60\x70\x88\x48\x41\xb4\x71\x5e\x36\xcd\xd0\x5b\x57\xfd\x0e\x6e\x00\xd1\x86\x3e\x78\x2b\x58\xe1\x6d\x5a\xa5\xe0\x4e\xdd\x9a\xa8\x1e\x6b\xf4\x71\x8c\x65\x3c\xfa\xe7\xca\xe7\xad\x54\x47\xac\xc4\x4f\x74\x4e\xbe\x9c\xcd\x70\xfe\x99\xdc\xab\xaa\xa5\x9f\xd4\xe7\xd2\x69\x82\xff\x99\xe7\x7b\xdb\x5b\x0e\xc5\x57\x56\x6e\x49\x06\x70\x2f\x09\x17\x3e\xd7\x07\xf3\x34\x52\x56\xa8\xa5\xab\xb1\xfa\x3f\xb7\xf4\xf9\x30\x99\x0e\xff\x9e\x09\xae\xeb\x2d\x37\x86\x55\x84\x1a\x53\x0b\xad\x6d\x6d\xa2\x53\xb4\x80\x25\x24\x4c\x37\xed\xb0\xfb\xd7\x33\x64\xaa\xbd\xda\xcc\x25\x67\x87\x5a\x2c\x2f\x80\x6a\x8d\x97\x2a\x38\xcd\xf5\xc7\x43\xd5\x9d\x8c\xaf\xd1\x6b\x15\x34\xff\x3b\xdd\x91\xf0\x35\x7e\xf6\xdd\x14\xec\x89\x1c\x40\x57\xbf\x45\x32\x6c\xae\x7d\x51\xb2\x74\x30\xa5\x9e\xf6\xe2\x31\xfb\x13\x00\x00\xff\xff\xfb\x06\xc2\x3d\xed\x06\x00\x00")

func _1593601729_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1593601729_initial_schemaUpSql,
		"1593601729_initial_schema.up.sql",
	)
}

func _1593601729_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1593601729_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1593601729_initial_schema.up.sql", size: 1773, mode: os.FileMode(436), modTime: time.Unix(1687862441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1597909626_add_server_typeDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1597909626_add_server_typeDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1597909626_add_server_typeDownSql,
		"1597909626_add_server_type.down.sql",
	)
}

func _1597909626_add_server_typeDownSql() (*asset, error) {
	bytes, err := _1597909626_add_server_typeDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1597909626_add_server_type.down.sql", size: 0, mode: os.FileMode(436), modTime: time.Unix(1687862441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1597909626_add_server_typeUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x28\x2d\xce\x88\xcf\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x8b\x4f\xce\xc9\x4c\xcd\x2b\x89\x2f\x4e\x2d\x2a\x4b\x2d\x2a\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x80\x08\xc5\x97\x54\x16\xa4\x2a\x78\xfa\x85\x28\xb8\xb8\xba\x39\x86\xfa\x84\x28\x18\x59\x73\x85\x06\xb8\x38\x86\x10\x61\x5a\xb0\x6b\x08\x8a\x31\xb6\x20\xcd\x5c\x80\x00\x00\x00\xff\xff\x98\x88\x1e\xcd\x91\x00\x00\x00")

func _1597909626_add_server_typeUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1597909626_add_server_typeUpSql,
		"1597909626_add_server_type.up.sql",
	)
}

func _1597909626_add_server_typeUpSql() (*asset, error) {
	bytes, err := _1597909626_add_server_typeUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1597909626_add_server_type.up.sql", size: 145, mode: os.FileMode(436), modTime: time.Unix(1687862441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1599053776_add_chat_id_and_typeDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1599053776_add_chat_id_and_typeDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1599053776_add_chat_id_and_typeDownSql,
		"1599053776_add_chat_id_and_type.down.sql",
	)
}

func _1599053776_add_chat_id_and_typeDownSql() (*asset, error) {
	bytes, err := _1599053776_add_chat_id_and_typeDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1599053776_add_chat_id_and_type.down.sql", size: 0, mode: os.FileMode(436), modTime: time.Unix(1687862441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1599053776_add_chat_id_and_typeUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x28\x2d\xce\x88\xcf\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x8b\x4f\xce\xc9\x4c\xcd\x2b\x89\x2f\x06\x11\xc8\x12\xc5\x0a\x8e\x2e\x2e\x0a\xce\xfe\x3e\xa1\xbe\x7e\x0a\xc9\x19\x89\x25\xf1\x99\x29\x0a\x21\xae\x11\x21\xd6\x5c\x54\x30\x10\x45\x47\x49\x65\x41\xaa\x82\xa7\x5f\x88\x35\x17\x57\x68\x80\x8b\x63\x08\x69\xa6\x06\xbb\x86\xc0\xdd\x67\xab\xa0\xa4\xa4\x83\xc5\x70\x5b\x05\x43\x6b\x2e\x40\x00\x00\x00\xff\xff\x22\xaf\x2b\x87\x08\x01\x00\x00")

func _1599053776_add_chat_id_and_typeUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1599053776_add_chat_id_and_typeUpSql,
		"1599053776_add_chat_id_and_type.up.sql",
	)
}

func _1599053776_add_chat_id_and_typeUpSql() (*asset, error) {
	bytes, err := _1599053776_add_chat_id_and_typeUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1599053776_add_chat_id_and_type.up.sql", size: 264, mode: os.FileMode(436), modTime: time.Unix(1687862441, 0)}
	a := &asset{bytes: bytes, info: info}
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

	info := bindataFileInfo{name: "doc.go", size: 382, mode: os.FileMode(436), modTime: time.Unix(1687862441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"1593601729_initial_schema.down.sql": _1593601729_initial_schemaDownSql,
	"1593601729_initial_schema.up.sql": _1593601729_initial_schemaUpSql,
	"1597909626_add_server_type.down.sql": _1597909626_add_server_typeDownSql,
	"1597909626_add_server_type.up.sql": _1597909626_add_server_typeUpSql,
	"1599053776_add_chat_id_and_type.down.sql": _1599053776_add_chat_id_and_typeDownSql,
	"1599053776_add_chat_id_and_type.up.sql": _1599053776_add_chat_id_and_typeUpSql,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"1593601729_initial_schema.down.sql": &bintree{_1593601729_initial_schemaDownSql, map[string]*bintree{}},
	"1593601729_initial_schema.up.sql": &bintree{_1593601729_initial_schemaUpSql, map[string]*bintree{}},
	"1597909626_add_server_type.down.sql": &bintree{_1597909626_add_server_typeDownSql, map[string]*bintree{}},
	"1597909626_add_server_type.up.sql": &bintree{_1597909626_add_server_typeUpSql, map[string]*bintree{}},
	"1599053776_add_chat_id_and_type.down.sql": &bintree{_1599053776_add_chat_id_and_typeDownSql, map[string]*bintree{}},
	"1599053776_add_chat_id_and_type.up.sql": &bintree{_1599053776_add_chat_id_and_typeUpSql, map[string]*bintree{}},
	"doc.go": &bintree{docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

