// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 000001_init.down.db.sql (65B)
// 000001_init.up.db.sql (2.719kB)
// 000002_add_last_ens_clock_value.down.sql (0)
// 000002_add_last_ens_clock_value.up.sql (77B)
// 1586358095_add_replace.down.sql (0)
// 1586358095_add_replace.up.sql (224B)
// 1588665364_add_image_data.down.sql (0)
// 1588665364_add_image_data.up.sql (186B)
// 1589365189_add_pow_target.down.sql (0)
// 1589365189_add_pow_target.up.sql (66B)
// 1591277220_add_index_messages.down.sql (237B)
// 1591277220_add_index_messages.up.sql (240B)
// 1593087212_add_mute_chat.down.sql (0)
// 1593087212_add_mute_chat.up.sql (58B)
// doc.go (850B)

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

var __000001_initDownDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xce\x48\x2c\x29\xb6\xe6\x42\x12\x29\x2d\x4e\x2d\x8a\xcf\x4d\x2d\x2e\x4e\x4c\x4f\x45\x95\x49\xce\xcf\x2b\x49\x4c\x06\x29\x07\x04\x00\x00\xff\xff\x61\x86\xbd\x5f\x41\x00\x00\x00")

func _000001_initDownDbSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initDownDbSql,
		"000001_init.down.db.sql",
	)
}

func _000001_initDownDbSql() (*asset, error) {
	bytes, err := _000001_initDownDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.down.db.sql", size: 65, mode: os.FileMode(0644), modTime: time.Unix(1578682784, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5e, 0xbb, 0x3f, 0x1, 0x75, 0x19, 0x70, 0x86, 0xa7, 0x34, 0x40, 0x17, 0x34, 0x3e, 0x18, 0x51, 0x79, 0xd4, 0x22, 0xad, 0x8f, 0x80, 0xcc, 0xa6, 0xcc, 0x6, 0x2b, 0x62, 0x2, 0x47, 0xba, 0xf9}}
	return a, nil
}

var __000001_initUpDbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xd1\x6f\xe2\xb8\x13\x7e\xe7\xaf\x18\xe9\xf7\x40\x2b\xd1\x9f\xf6\xa4\xbd\xbd\x93\xfa\x44\xd9\xf4\x0e\x1d\x07\x2b\x9a\x9e\xba\x4f\xd6\xe0\x0c\xc4\xc2\xb1\x23\x7b\x02\x8b\xb4\x7f\xfc\xc9\x81\x40\x0c\x81\xee\xea\xfa\x50\xb5\x33\xe3\x19\xcf\x37\xf3\x7d\xce\x68\x9e\x0c\xd3\x04\xd2\xe1\xd3\x24\x81\xf1\x33\x4c\x67\x29\x24\x6f\xe3\x97\xf4\x05\x64\x8e\xec\xe1\xae\x07\xa0\x32\xf8\x67\x38\x1f\xfd\x39\x9c\xc3\x97\xf9\xf8\xef\xe1\xfc\x2b\xfc\x95\x7c\x85\xd9\x14\x46\xb3\xe9\xf3\x64\x3c\x4a\x61\x9e\x7c\x99\x0c\x47\xc9\xa0\x07\x60\xb0\xa0\x63\x7c\xc8\x37\x7d\x9d\x4c\x82\x43\x5a\x6d\xdd\x85\x07\x3e\x27\xcf\xc3\xd7\x49\x0a\xfd\xff\xe1\x2f\xbf\xff\x96\xfd\xda\x0f\xb1\xbc\x2b\x09\xc6\xd3\x34\x4a\x80\x92\xd5\x86\xe0\x69\x36\x9b\x24\xc3\xe9\x65\x86\x74\xfe\x5a\xdf\x80\x55\x41\x9e\xb1\x28\x2f\x32\x64\xa4\x89\x29\x13\xc8\x42\x6a\x2b\xd7\x62\x83\xba\x8a\x0b\x1d\xb3\x7d\x08\x07\xca\x6a\xa1\x95\x14\x6b\xda\xc1\xd3\x64\xf6\x14\x4c\x95\xd9\x28\xda\x52\x26\x0a\xf2\x1e\x57\x24\xa4\xad\x0c\xdf\xc8\xa1\xd1\xff\x68\xb9\x3a\xf4\x90\xf7\x58\xb0\xa0\x62\x41\xce\x9f\xff\x9f\xab\x52\x54\x65\x86\x4c\x7b\x57\xef\xfe\xb1\xd7\x8b\xe6\x29\xad\x61\x94\xa7\x21\xa6\xc9\x5b\xfa\x23\x13\xc4\x2c\x73\xe4\xfd\x3e\xbe\x0d\x5f\x3d\xda\x0b\x2b\x19\x2f\x36\xe4\xd4\x52\x51\x76\x1c\x4e\xd3\xd6\xf3\x70\xf2\x92\x9c\x47\x09\xbc\x85\x17\x6a\x85\x1d\xc5\x55\x46\x86\x95\xb4\xe6\xd2\x55\xe6\x96\xed\xa5\xb9\x46\x73\x0f\x51\x76\xa3\x9e\xdf\x79\xa6\x42\x30\xae\x4e\x18\x67\xb4\x51\x92\x84\x32\x4b\x7b\xb4\xb1\x53\x8b\x8a\x49\xb0\x15\x8c\x7a\x1d\xd7\xab\xd1\x7f\x78\x80\x31\xf7\x3d\xa8\xa2\xb4\x8e\xd1\x30\x70\x8e\xe1\x97\xf2\xc0\xb8\xd0\x04\x39\x7a\x70\x76\xab\x32\x40\x0f\x5b\x02\x47\x7a\x07\xd6\x80\xe2\x70\x78\x9b\x93\x09\x87\x35\x15\xa1\x57\xb3\x02\x65\x96\xca\x28\xa6\x07\x2f\x9d\xd5\xfa\xff\xbd\x1b\x84\xad\x3c\xb9\x66\x79\xf6\x33\xff\x59\xea\x02\x6c\x73\xe5\x4b\x72\x22\xa2\x50\xf2\x47\x12\x33\x19\xc0\xdb\xca\xc9\x8e\x5d\x08\xc8\x79\x56\x06\x59\x59\x73\x44\x0e\x80\xe9\x1b\x77\x8a\x02\xd4\x5b\x4a\x86\x45\x27\xe5\xa1\xee\xaa\x2d\x29\x87\x7c\x57\x29\x0e\xb5\x70\x89\x56\xe3\xb1\x57\x5b\x89\x5a\xdc\x8e\xc9\x55\x46\xd7\x37\x19\xc0\x91\x2f\xad\xf1\x61\x15\xe2\x6b\x35\x92\xd0\xf4\x72\xb8\xd0\x15\xee\x1f\xa0\x24\x32\xd7\x35\xad\x55\xd5\x56\xbc\xb2\xca\xac\x84\x67\xe4\xca\xc7\x95\x4b\x74\x9e\x32\x51\xe3\x7c\x82\xdd\xe1\x56\x94\xb8\xd3\x16\xb3\x96\xd5\xb3\x92\x6b\x72\xa2\x44\xb9\x3e\xdd\xb2\xb1\xe6\xe8\xf3\x38\xb7\xb4\x45\x81\x26\x6b\xe1\x15\xdb\xf7\x9d\x75\xba\x1a\x29\xe9\x74\x2e\x9d\x2d\xba\x3d\x61\x27\x1c\x4a\xee\xf6\xb2\x43\xe3\xc3\x63\x60\xcd\x8d\xdb\x7a\xb5\x32\xc8\x95\xa3\x56\xe7\x47\x1f\x23\xd7\xb3\x68\x8b\xe6\x78\xfa\x39\x79\x03\x95\x7d\x13\x87\xed\x9e\x4d\x63\x4e\xdd\xed\xed\xf7\x8f\x1d\x27\x08\x9d\xcc\xc5\x62\x77\xdc\xac\xd9\x14\xce\x4e\xef\x51\xae\x16\x9e\xdd\x5d\xff\xc3\x7f\xfc\xe9\xc3\xf7\xef\xed\xc5\x1a\xc0\xc3\xa7\x8f\x03\xf8\xf4\xf1\x3e\x38\x54\x36\x68\x68\x30\xa8\xb7\xf9\xf2\x71\x88\xb5\x23\x2c\x4a\x24\x1d\x3f\x27\x1c\xef\x93\xaa\xd6\x62\x4f\x67\x0f\x65\x2d\xbf\x54\x0f\xfc\xfc\x0d\x3d\xb8\xf8\x3a\x0f\x1d\xd5\x47\xb1\x62\x5b\x20\x2b\x89\x5a\xef\xae\x47\x77\x51\xd3\x91\x54\xa5\x22\xc3\x27\xe1\x6f\xb3\xe5\x1d\xcc\x42\x46\x32\xab\xa0\x96\xa7\x85\xf4\xe1\x79\xd8\xa0\x56\xe1\xd5\xa9\x91\x6c\x0a\xc7\xec\xb9\xe4\x54\xd4\xf9\xb5\x15\x6f\x0f\x63\xdf\x01\xbb\xdd\x09\xbd\x60\x5a\x2a\x57\x43\x4d\xa6\xb1\xc4\x4c\x88\xeb\xb4\x2e\x7b\x0e\x5d\xf3\x39\x75\xf6\x0d\xd4\xc9\x98\x4e\x2c\x22\x28\x02\x1f\xde\x47\xec\xae\xf5\xf7\xfd\x63\xef\xdf\x00\x00\x00\xff\xff\xd7\x95\x8b\x6b\x9f\x0a\x00\x00")

func _000001_initUpDbSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initUpDbSql,
		"000001_init.up.db.sql",
	)
}

func _000001_initUpDbSql() (*asset, error) {
	bytes, err := _000001_initUpDbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.up.db.sql", size: 2719, mode: os.FileMode(0644), modTime: time.Unix(1578682784, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x60, 0xdc, 0xeb, 0xe, 0xc2, 0x4f, 0x75, 0xa, 0xf6, 0x3e, 0xc7, 0xc4, 0x4, 0xe2, 0xe1, 0xa4, 0x73, 0x2f, 0x4a, 0xad, 0x1a, 0x0, 0xc3, 0x93, 0x9d, 0x77, 0x3e, 0x31, 0x91, 0x77, 0x2e, 0xc8}}
	return a, nil
}

var __000002_add_last_ens_clock_valueDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000002_add_last_ens_clock_valueDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_add_last_ens_clock_valueDownSql,
		"000002_add_last_ens_clock_value.down.sql",
	)
}

func _000002_add_last_ens_clock_valueDownSql() (*asset, error) {
	bytes, err := _000002_add_last_ens_clock_valueDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_add_last_ens_clock_value.down.sql", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1584434371, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __000002_add_last_ens_clock_valueUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0a\x85\x20\x10\x06\xe0\xfd\x3b\xc5\x7f\x84\xb7\x6f\x35\xa5\x41\x30\x8d\x10\xe3\x5a\x64\x70\x95\xe8\x42\xeb\xfc\x7d\xc4\xea\x2f\x28\xad\xec\x61\xbd\xcd\x6c\x73\x80\x9c\xc3\x16\x38\x9e\x82\x9a\xc7\x4c\xa5\x8d\x64\xb5\xdb\x9d\xde\x5c\x9f\x82\x43\x14\x12\x14\x12\x99\xe1\xfc\x4e\x91\x15\xff\xe5\xf7\x05\x00\x00\xff\xff\xd0\x66\x8a\xf7\x4d\x00\x00\x00")

func _000002_add_last_ens_clock_valueUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_add_last_ens_clock_valueUpSql,
		"000002_add_last_ens_clock_value.up.sql",
	)
}

func _000002_add_last_ens_clock_valueUpSql() (*asset, error) {
	bytes, err := _000002_add_last_ens_clock_valueUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_add_last_ens_clock_value.up.sql", size: 77, mode: os.FileMode(0644), modTime: time.Unix(1584434371, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4d, 0x3, 0x8f, 0xd5, 0x85, 0x83, 0x47, 0xbe, 0xf9, 0x82, 0x7e, 0x81, 0xa4, 0xbd, 0xaa, 0xd5, 0x98, 0x18, 0x5, 0x2d, 0x82, 0x42, 0x3b, 0x3, 0x50, 0xc3, 0x1e, 0x84, 0x35, 0xf, 0xb6, 0x2b}}
	return a, nil
}

var __1586358095_add_replaceDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1586358095_add_replaceDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1586358095_add_replaceDownSql,
		"1586358095_add_replace.down.sql",
	)
}

func _1586358095_add_replaceDownSql() (*asset, error) {
	bytes, err := _1586358095_add_replaceDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1586358095_add_replace.down.sql", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1589265610, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __1586358095_add_replaceUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\x31\x0a\xc2\x40\x10\x46\xe1\x3e\xa7\xf8\xc9\x09\xec\x53\x4d\xdc\x09\x08\xe3\x2c\xe8\x2c\xd8\x85\x10\x06\x11\xd6\x28\xd9\xe4\xfe\x56\x56\x5a\x6c\xfd\xe0\xe3\x91\x18\x5f\x60\xd4\x0b\x63\x2f\xbe\x8e\x4f\x2f\x65\xba\x7b\x01\x85\x80\x63\x94\x74\x56\xac\xfe\xce\xd3\xec\xdf\x06\xe3\x9b\x41\xa3\x41\x93\x08\x02\x0f\x94\xc4\xd0\xb6\x5d\x53\xc7\x6d\x19\x7d\x8c\xc2\xa4\xbf\xca\x40\x72\xe5\x4a\x28\x3f\x16\x1f\xe7\xd7\xbe\x6c\x38\xe9\x9f\xa3\x43\xd7\x7c\x02\x00\x00\xff\xff\xca\x94\x3f\xe0\xe0\x00\x00\x00")

func _1586358095_add_replaceUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1586358095_add_replaceUpSql,
		"1586358095_add_replace.up.sql",
	)
}

func _1586358095_add_replaceUpSql() (*asset, error) {
	bytes, err := _1586358095_add_replaceUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1586358095_add_replace.up.sql", size: 224, mode: os.FileMode(0644), modTime: time.Unix(1589265610, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd2, 0xb3, 0xa9, 0xc7, 0x7f, 0x9d, 0x8f, 0x43, 0x8c, 0x9e, 0x58, 0x8d, 0x44, 0xbc, 0xfa, 0x6b, 0x5f, 0x3f, 0x5a, 0xbe, 0xe8, 0xb1, 0x16, 0xf, 0x91, 0x2a, 0xa0, 0x71, 0xbb, 0x8d, 0x6b, 0xcb}}
	return a, nil
}

var __1588665364_add_image_dataDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1588665364_add_image_dataDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1588665364_add_image_dataDownSql,
		"1588665364_add_image_data.down.sql",
	)
}

func _1588665364_add_image_dataDownSql() (*asset, error) {
	bytes, err := _1588665364_add_image_dataDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1588665364_add_image_data.down.sql", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1591690523, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __1588665364_add_image_dataUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xcb\x41\x0a\xc2\x40\x0c\x05\xd0\x7d\x4f\xf1\xe9\x19\xc4\x4d\x57\x33\x4e\x04\x21\x66\x40\x32\xe0\xae\x44\x0c\x45\xb0\x58\x8c\x2e\x7a\xfb\x9e\xa1\x07\x78\x89\x95\x6e\xd0\x94\x99\xf0\x0f\xff\x8e\xb3\x47\xd8\xe4\x81\x54\x0a\x4e\x95\xdb\x55\xf0\x9a\x6d\xf2\x71\xb1\xf5\xfd\xb1\x27\x32\xd7\x3c\x74\x3b\xe0\x6f\x5d\x1c\x17\xd1\x5d\xe8\x61\xe1\xc7\x03\x94\xee\x0a\xa9\x0a\x69\xcc\x28\x74\x4e\x8d\x15\x7d\x3f\x74\x5b\x00\x00\x00\xff\xff\xf8\x4b\xbd\xbe\xba\x00\x00\x00")

func _1588665364_add_image_dataUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1588665364_add_image_dataUpSql,
		"1588665364_add_image_data.up.sql",
	)
}

func _1588665364_add_image_dataUpSql() (*asset, error) {
	bytes, err := _1588665364_add_image_dataUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1588665364_add_image_data.up.sql", size: 186, mode: os.FileMode(0644), modTime: time.Unix(1591690523, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd6, 0xc6, 0x35, 0xb4, 0x4c, 0x39, 0x96, 0x29, 0x30, 0xda, 0xf4, 0x8f, 0xcb, 0xf1, 0x9f, 0x84, 0xdc, 0x88, 0xd4, 0xd5, 0xbc, 0xb6, 0x5b, 0x46, 0x78, 0x67, 0x76, 0x1a, 0x5, 0x36, 0xdc, 0xe5}}
	return a, nil
}

var __1589365189_add_pow_targetDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1589365189_add_pow_targetDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1589365189_add_pow_targetDownSql,
		"1589365189_add_pow_target.down.sql",
	)
}

func _1589365189_add_pow_targetDownSql() (*asset, error) {
	bytes, err := _1589365189_add_pow_targetDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1589365189_add_pow_target.down.sql", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1591690523, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __1589365189_add_pow_targetUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4a\x2c\x8f\xcf\x4d\x2d\x2e\x4e\x4c\x4f\x2d\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xc8\x2f\x8f\x2f\x49\x2c\x4a\x4f\x2d\x51\x08\x72\x75\xf4\x51\x48\x49\x4d\x4b\x2c\xcd\x29\x51\x30\xd0\x33\x30\xb2\xe6\x02\x04\x00\x00\xff\xff\x49\xd6\x04\x23\x42\x00\x00\x00")

func _1589365189_add_pow_targetUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1589365189_add_pow_targetUpSql,
		"1589365189_add_pow_target.up.sql",
	)
}

func _1589365189_add_pow_targetUpSql() (*asset, error) {
	bytes, err := _1589365189_add_pow_targetUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1589365189_add_pow_target.up.sql", size: 66, mode: os.FileMode(0644), modTime: time.Unix(1591690523, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4e, 0x3a, 0xe2, 0x2e, 0x7d, 0xaf, 0xbb, 0xcc, 0x21, 0xa1, 0x7a, 0x41, 0x9a, 0xd0, 0xbb, 0xa9, 0xc8, 0x35, 0xf9, 0x32, 0x34, 0x46, 0x44, 0x9a, 0x86, 0x40, 0x7c, 0xb9, 0x23, 0xc7, 0x3, 0x3f}}
	return a, nil
}

var __1591277220_add_index_messagesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xce\xb1\xaa\xc2\x30\x18\xc5\xf1\x3d\x4f\x71\xb6\xb6\x90\x0b\x77\x28\x5d\x3a\x89\xed\xe0\xd2\x4a\x71\x70\xfb\x48\x93\x60\x83\xd1\x40\xbe\x44\x14\xfa\xf0\x22\x74\x74\xf3\xac\x87\x1f\xfc\xbb\x69\x3c\xe2\x30\x74\xfd\x19\xce\x3c\x89\xad\x8a\x7a\xa1\xf9\x45\x3e\x68\xe5\x49\x2f\x2a\x91\x33\xc4\x21\x26\x0a\x77\xd2\x39\x72\x88\xad\x10\xfb\xa9\xdf\x9d\xfa\xaf\x72\x33\x18\x07\x20\xb3\x8d\x74\xb3\xcc\xea\x62\xb9\x14\x00\xc0\x79\xe6\x14\xcb\xe2\xff\xc7\x15\x58\x57\x68\x1f\xf4\x95\x1e\xca\x67\x2b\xf1\xd7\xd4\x12\x4d\x5d\x7d\x0e\x67\x24\xb6\x10\x89\xc5\x19\x2b\xaa\x56\xbc\x03\x00\x00\xff\xff\xaf\xf2\xdc\xc0\xed\x00\x00\x00")

func _1591277220_add_index_messagesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1591277220_add_index_messagesDownSql,
		"1591277220_add_index_messages.down.sql",
	)
}

func _1591277220_add_index_messagesDownSql() (*asset, error) {
	bytes, err := _1591277220_add_index_messagesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1591277220_add_index_messages.down.sql", size: 237, mode: os.FileMode(0644), modTime: time.Unix(1591690523, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x79, 0xe5, 0x42, 0x56, 0x64, 0x1d, 0xb7, 0x8a, 0x1b, 0x0, 0x99, 0xf0, 0x18, 0x8c, 0x69, 0xe3, 0x14, 0x3a, 0x7f, 0x78, 0xfe, 0xe3, 0x2e, 0xcb, 0x6e, 0x5c, 0x8c, 0x1f, 0x7b, 0xfc, 0x21, 0xc7}}
	return a, nil
}

var __1591277220_add_index_messagesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xca\xb1\xaa\x83\x30\x14\x87\xf1\xdd\xa7\xf8\x6f\x2a\x78\xe1\x0e\xe2\xe2\x24\x9a\xa1\x8b\x16\xed\xd0\xed\x10\x63\xa8\xd2\xb4\x81\x73\x4c\x69\xc1\x87\x2f\x1d\x3a\x14\xba\xf5\x5b\xbf\x5f\xd3\x77\x7b\xec\xda\x46\x1d\xb1\x4c\x77\x12\xab\xd9\xcc\x34\x3e\xc8\xcc\x7a\xa5\x65\x2a\xa3\xa8\xee\x55\x75\x50\x5f\x91\xf3\x46\xbb\x37\x25\xf1\xbc\x92\xbf\x92\x09\x2c\x9e\xd1\xb5\x08\x62\x99\x2e\x56\x44\x9f\xac\x20\xf9\xe0\xa8\x86\x3a\x83\x84\x51\x56\x4e\xe2\xff\x1f\x8b\xb1\x6d\x30\xce\x9b\x33\xdd\xb4\x0b\x36\xc3\x5f\x91\x67\x28\xf2\xf4\x35\x96\x09\x8d\x1a\xea\xb4\x8c\x9e\x01\x00\x00\xff\xff\xdb\x1d\x3d\x0b\xf0\x00\x00\x00")

func _1591277220_add_index_messagesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1591277220_add_index_messagesUpSql,
		"1591277220_add_index_messages.up.sql",
	)
}

func _1591277220_add_index_messagesUpSql() (*asset, error) {
	bytes, err := _1591277220_add_index_messagesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1591277220_add_index_messages.up.sql", size: 240, mode: os.FileMode(0644), modTime: time.Unix(1591690523, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0xfe, 0xbe, 0xd5, 0xb8, 0x8f, 0xdd, 0xef, 0xbb, 0xa8, 0xad, 0x7f, 0xed, 0x5b, 0x5b, 0x2f, 0xe6, 0x82, 0x27, 0x78, 0x1f, 0xb9, 0x57, 0xdc, 0x8, 0xc2, 0xb2, 0xa9, 0x9a, 0x4, 0xe1, 0x7a}}
	return a, nil
}

var __1593087212_add_mute_chatDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1593087212_add_mute_chatDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1593087212_add_mute_chatDownSql,
		"1593087212_add_mute_chat.down.sql",
	)
}

func _1593087212_add_mute_chatDownSql() (*asset, error) {
	bytes, err := _1593087212_add_mute_chatDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1593087212_add_mute_chat.down.sql", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1593087310, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __1593087212_add_mute_chatUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xce\x48\x2c\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\x2d\x2d\x49\x4d\x51\x70\xf2\xf7\xf7\x71\x75\xf4\x53\x70\x71\x75\x73\x0c\xf5\x09\x51\x70\x73\xf4\x09\x76\xb5\xe6\x02\x04\x00\x00\xff\xff\x59\x4c\x4b\xec\x3a\x00\x00\x00")

func _1593087212_add_mute_chatUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1593087212_add_mute_chatUpSql,
		"1593087212_add_mute_chat.up.sql",
	)
}

func _1593087212_add_mute_chatUpSql() (*asset, error) {
	bytes, err := _1593087212_add_mute_chatUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1593087212_add_mute_chat.up.sql", size: 58, mode: os.FileMode(0644), modTime: time.Unix(1593091669, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0x9, 0xa0, 0x4a, 0x8e, 0x23, 0xe4, 0xce, 0xbc, 0xd4, 0x9, 0xeb, 0xf9, 0x67, 0x90, 0xc0, 0x4b, 0x67, 0x84, 0xe4, 0x42, 0x8d, 0x0, 0x17, 0x29, 0x7f, 0x12, 0xbf, 0x7d, 0x4e, 0x78, 0xec}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\x3f\x8f\xdb\x3e\x0c\xdd\xf3\x29\x1e\x6e\xb9\xe5\x22\x07\xf8\xfd\xa6\xdb\x3a\x74\xe8\xd2\x2e\xd9\x0b\x46\xa6\x6d\x22\x32\xe5\x8a\xf4\x39\xf9\xf6\x85\x74\x17\x9c\x51\x14\xe8\x4a\x89\x8f\xef\x5f\xd7\xe1\x3c\x89\x61\x90\xc4\x10\x83\x72\x64\x33\x2a\x77\x5c\x38\xd2\x6a\x8c\xa7\x51\x7c\x5a\x2f\x21\xe6\xb9\x33\x27\x5f\xed\x28\x73\x37\xcb\x58\xc8\xb9\x7b\xfb\xff\xe9\xd0\x75\x88\xa4\xcf\x8e\x89\xb4\x4f\xdc\xb0\x0c\xe6\x54\x5c\x74\xc4\x26\x3e\x81\xb0\x14\x1e\xe4\x16\xf0\xc5\x91\x98\xcc\xe1\x13\xf9\xb3\xc1\x27\x46\x24\xe3\x0a\x33\xe4\x82\x31\x1f\x2f\xa2\x3d\x39\x85\x3a\xfa\x36\xec\x26\x95\x61\xa4\x94\xb8\xc7\x50\xf2\xdc\x76\x8d\x66\x46\x2f\x85\xa3\xe7\x72\x7f\x01\x99\xb1\x43\x69\x66\xab\xfb\x13\xbd\x31\x34\x7f\x9c\x07\x69\xff\x6f\x45\xd8\x72\xb9\x1a\xc8\xc0\xb7\x85\xa3\x73\x1f\x0e\x15\xeb\xfb\x8f\xf3\xd7\x57\x9c\x27\xae\xf0\x55\x5a\x1e\x1a\x85\x66\x9e\x32\xf7\x06\xcf\x18\x72\x4a\x79\x6b\x0f\xab\xca\x0d\x2e\x33\x9b\xd3\xbc\x20\x66\x7d\x63\x75\xc9\x5a\xd1\x56\x4d\x72\xe5\xf6\xcf\xb7\x0c\x51\x71\xa1\xf4\xee\x5e\x93\x7e\x7e\x37\xe8\x11\x44\x5c\x4b\x61\xf5\x74\x6f\x2b\xac\xb1\xdc\x97\x8a\x85\x77\xe6\x92\xd5\x9a\xbc\xa5\x64\xcf\x31\xa7\xdd\xbc\xa2\xd9\x44\x85\x3f\x1d\x73\xba\x24\x7e\xc1\x36\x49\x9c\x30\x33\xa9\xb5\x40\xda\x87\x44\xce\xe6\x9f\xfb\x10\x85\x73\x99\xad\x0a\xae\xfc\xaa\xbb\x15\xb3\x16\xe7\x91\xc3\x8e\x50\x33\x7f\xa1\xf8\x51\x85\xc7\x95\xd5\xd8\x40\x7f\x98\xf2\x08\x79\x63\x50\xdf\xe3\x74\x3a\x9d\xfe\xfb\x19\x42\x68\x5d\xe0\x1b\xcd\x4b\xa5\xe9\xb5\xa3\x9b\xa4\x84\x0b\x43\x46\xcd\x85\xfb\xca\x8a\x6f\x62\xad\x64\x31\x09\xab\xd7\xcc\x2a\x5e\x4e\x3d\x97\xaa\x47\xf7\x7a\xfe\x66\x59\x38\x1c\x16\x8a\x57\x1a\x19\xf6\x2b\x89\x73\x0d\x7a\xcc\xaf\x23\x2b\xd7\x3a\xec\xcb\x77\x5c\xae\xe3\xde\xec\x63\x46\x08\xdd\xe7\x20\x8c\x19\xe1\xf0\x3b\x00\x00\xff\xff\x12\xcd\x7f\xc4\x52\x03\x00\x00")

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

	info := bindataFileInfo{name: "doc.go", size: 850, mode: os.FileMode(0644), modTime: time.Unix(1589265610, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa0, 0xcc, 0x41, 0xe1, 0x61, 0x12, 0x97, 0xe, 0x36, 0x8c, 0xa7, 0x9e, 0xe0, 0x6e, 0x59, 0x9e, 0xee, 0xd5, 0x4a, 0xcf, 0x1e, 0x60, 0xd6, 0xc3, 0x3a, 0xc9, 0x6c, 0xf2, 0x86, 0x5a, 0xb4, 0x1e}}
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
	"000001_init.down.db.sql": _000001_initDownDbSql,

	"000001_init.up.db.sql": _000001_initUpDbSql,

	"000002_add_last_ens_clock_value.down.sql": _000002_add_last_ens_clock_valueDownSql,

	"000002_add_last_ens_clock_value.up.sql": _000002_add_last_ens_clock_valueUpSql,

	"1586358095_add_replace.down.sql": _1586358095_add_replaceDownSql,

	"1586358095_add_replace.up.sql": _1586358095_add_replaceUpSql,

	"1588665364_add_image_data.down.sql": _1588665364_add_image_dataDownSql,

	"1588665364_add_image_data.up.sql": _1588665364_add_image_dataUpSql,

	"1589365189_add_pow_target.down.sql": _1589365189_add_pow_targetDownSql,

	"1589365189_add_pow_target.up.sql": _1589365189_add_pow_targetUpSql,

	"1591277220_add_index_messages.down.sql": _1591277220_add_index_messagesDownSql,

	"1591277220_add_index_messages.up.sql": _1591277220_add_index_messagesUpSql,

	"1593087212_add_mute_chat.down.sql": _1593087212_add_mute_chatDownSql,

	"1593087212_add_mute_chat.up.sql": _1593087212_add_mute_chatUpSql,

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
	"000001_init.down.db.sql":                  &bintree{_000001_initDownDbSql, map[string]*bintree{}},
	"000001_init.up.db.sql":                    &bintree{_000001_initUpDbSql, map[string]*bintree{}},
	"000002_add_last_ens_clock_value.down.sql": &bintree{_000002_add_last_ens_clock_valueDownSql, map[string]*bintree{}},
	"000002_add_last_ens_clock_value.up.sql":   &bintree{_000002_add_last_ens_clock_valueUpSql, map[string]*bintree{}},
	"1586358095_add_replace.down.sql":          &bintree{_1586358095_add_replaceDownSql, map[string]*bintree{}},
	"1586358095_add_replace.up.sql":            &bintree{_1586358095_add_replaceUpSql, map[string]*bintree{}},
	"1588665364_add_image_data.down.sql":       &bintree{_1588665364_add_image_dataDownSql, map[string]*bintree{}},
	"1588665364_add_image_data.up.sql":         &bintree{_1588665364_add_image_dataUpSql, map[string]*bintree{}},
	"1589365189_add_pow_target.down.sql":       &bintree{_1589365189_add_pow_targetDownSql, map[string]*bintree{}},
	"1589365189_add_pow_target.up.sql":         &bintree{_1589365189_add_pow_targetUpSql, map[string]*bintree{}},
	"1591277220_add_index_messages.down.sql":   &bintree{_1591277220_add_index_messagesDownSql, map[string]*bintree{}},
	"1591277220_add_index_messages.up.sql":     &bintree{_1591277220_add_index_messagesUpSql, map[string]*bintree{}},
	"1593087212_add_mute_chat.down.sql":        &bintree{_1593087212_add_mute_chatDownSql, map[string]*bintree{}},
	"1593087212_add_mute_chat.up.sql":          &bintree{_1593087212_add_mute_chatUpSql, map[string]*bintree{}},
	"doc.go":                                   &bintree{docGo, map[string]*bintree{}},
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
