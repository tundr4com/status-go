// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 0001_accounts.down.sql (21B)
// 0001_accounts.up.sql (163B)
// 1605007189_identity_images.down.sql (29B)
// 1605007189_identity_images.up.sql (268B)
// 1606224181_drop_photo_path_from_accounts.down.sql (892B)
// 1606224181_drop_photo_path_from_accounts.up.sql (866B)
// 1648646095_image_clock.down.sql (939B)
// 1648646095_image_clock.up.sql (69B)
// 1649317600_add_color_hash.up.sql (201B)
// 1660238799_accounts_kdf.up.sql (115B)
// 1679505708_add_customization_color.up.sql (78B)
// 1687853321_add_customization_color_updated_at.up.sql (80B)
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

var __0001_accountsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x4c\x4e\xce\x2f\xcd\x2b\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\x96\x1e\x13\xa1\x15\x00\x00\x00")

func _0001_accountsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_accountsDownSql,
		"0001_accounts.down.sql",
	)
}

func _0001_accountsDownSql() (*asset, error) {
	bytes, err := _0001_accountsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_accounts.down.sql", size: 21, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd2, 0x61, 0x4c, 0x18, 0xfc, 0xc, 0xdf, 0x5c, 0x1f, 0x5e, 0xd3, 0xbd, 0xfa, 0x12, 0x5e, 0x8d, 0x8d, 0x8b, 0xb9, 0x5f, 0x99, 0x46, 0x63, 0xa5, 0xe3, 0xa6, 0x8a, 0x4, 0xf1, 0x73, 0x8a, 0xe9}}
	return a, nil
}

var __0001_accountsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xcc\xb1\x6e\x83\x30\x14\x46\xe1\xdd\x4f\xf1\x8f\xad\xe4\x37\xe8\x64\xa8\x5b\xae\x42\x00\x99\x4b\x80\xd1\x02\x04\x56\x82\x8d\xc0\x19\xf2\xf6\x51\x58\x8f\xf4\x9d\xd4\x68\xc5\x1a\xac\x92\x5c\x83\xfe\x50\x94\x0c\xdd\x51\xcd\x35\xec\x30\x84\xa7\x8f\x07\xbe\xc4\x7d\x7a\x35\x6e\xc4\x4d\x99\x34\x53\x06\x95\xa1\xab\x32\x3d\x2e\xba\x97\xc2\xdb\x75\x02\xeb\x8e\x4f\x5b\x34\x79\x2e\xc5\x23\xcc\xce\xb3\x5b\xa7\x23\xda\x75\x43\x42\xff\xa0\x82\xa5\xd8\x96\x10\x43\x65\xe3\x72\x02\xf9\xf9\x0e\x76\x1f\x2b\xeb\x76\xe7\xe7\x33\x8a\x6f\xb4\xc4\x59\xd9\x30\x4c\xd9\xd2\xef\x8f\x78\x07\x00\x00\xff\xff\xab\xcf\xa2\xbd\xa3\x00\x00\x00")

func _0001_accountsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_accountsUpSql,
		"0001_accounts.up.sql",
	)
}

func _0001_accountsUpSql() (*asset, error) {
	bytes, err := _0001_accountsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_accounts.up.sql", size: 163, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf2, 0xfa, 0x99, 0x8e, 0x96, 0xb3, 0x13, 0x6c, 0x1f, 0x6, 0x27, 0xc5, 0xd2, 0xd4, 0xe0, 0xa5, 0x26, 0x82, 0xa7, 0x26, 0xf2, 0x68, 0x9d, 0xed, 0x9c, 0x3d, 0xbb, 0xdc, 0x37, 0x28, 0xbc, 0x1}}
	return a, nil
}

var __1605007189_identity_imagesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\x4c\x49\xcd\x2b\xc9\x2c\xa9\x8c\xcf\xcc\x4d\x4c\x4f\x2d\xb6\xe6\xe5\x02\x04\x00\x00\xff\xff\xa1\x22\x72\x37\x1d\x00\x00\x00")

func _1605007189_identity_imagesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1605007189_identity_imagesDownSql,
		"1605007189_identity_images.down.sql",
	)
}

func _1605007189_identity_imagesDownSql() (*asset, error) {
	bytes, err := _1605007189_identity_imagesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1605007189_identity_images.down.sql", size: 29, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2f, 0xcf, 0xa7, 0xae, 0xd5, 0x4f, 0xcd, 0x14, 0x63, 0x9, 0xbe, 0x39, 0x49, 0x18, 0x96, 0xb2, 0xa3, 0x8, 0x7d, 0x41, 0xdb, 0x50, 0x5d, 0xf5, 0x4d, 0xa2, 0xd, 0x8f, 0x57, 0x79, 0x77, 0x67}}
	return a, nil
}

var __1605007189_identity_imagesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xc1\x6a\xc3\x30\x10\x04\xd0\xbb\xc1\xff\x30\xc7\x04\xf2\x07\x3d\xc9\xaa\x42\x44\x55\x29\x28\x4a\xd3\x9c\x84\x40\x5b\x7b\x69\xe2\x96\x58\xa5\xb8\x5f\x5f\xea\xfa\x60\x72\xdc\xc7\xec\x30\xd2\x2b\x11\x14\x82\x68\x8c\x82\xde\xc2\xba\x00\xf5\xaa\x0f\xe1\x00\xce\xd4\x17\x2e\x63\xe4\x6b\x6a\x69\x58\xd5\x15\x00\xbc\xd3\x18\xbf\x38\xe3\x45\x78\xb9\x13\x7e\xf3\xaf\x7d\xba\xd2\x1d\x4d\x5f\xf1\x33\x8d\x97\x8f\x94\xd1\x18\xd7\x4c\xe5\xf6\x68\xcc\x9c\xf8\xe6\x5c\x3a\x70\x5f\xe6\xbb\x23\x6e\xbb\xb2\x80\x37\xbe\x50\x1c\xf8\x87\x16\x76\xa3\x3f\x88\x25\xdd\x5a\x5a\x66\xf7\x5e\x3f\x0b\x7f\xc6\x93\x3a\x63\x35\x8f\xdc\x4c\xbb\xd6\x70\x16\xd2\xd9\xad\xd1\x32\xc0\xab\xbd\x11\x52\xd5\xd5\x1a\x27\x1d\x76\xee\x18\xe0\xdd\x49\x3f\x3e\xd4\xd5\x6f\x00\x00\x00\xff\xff\x8c\x6a\x0a\x57\x0c\x01\x00\x00")

func _1605007189_identity_imagesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1605007189_identity_imagesUpSql,
		"1605007189_identity_images.up.sql",
	)
}

func _1605007189_identity_imagesUpSql() (*asset, error) {
	bytes, err := _1605007189_identity_imagesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1605007189_identity_images.up.sql", size: 268, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x50, 0xb6, 0xc1, 0x5c, 0x76, 0x72, 0x6b, 0x22, 0x34, 0xdc, 0x96, 0xdc, 0x2b, 0xfd, 0x2d, 0xbe, 0xcc, 0x1e, 0xd4, 0x5, 0x93, 0xd, 0xc2, 0x51, 0xf3, 0x1a, 0xef, 0x2b, 0x26, 0xa4, 0xeb, 0x65}}
	return a, nil
}

var __1606224181_drop_photo_path_from_accountsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x93\xc1\x6e\x9c\x30\x18\x84\xef\x3c\xc5\x1c\xdb\x95\x25\x1e\x60\x4f\x04\x9c\xc4\x2a\x0b\xc8\x98\x66\x73\x6a\x1c\xb0\x82\x15\xb0\x11\x6b\x54\xed\xdb\x57\x6b\x58\x76\x8b\xd4\xaa\xd7\x5e\xff\xf9\x7f\xcf\x68\x3e\x39\xdc\x21\xb6\xc3\x19\xae\x55\x90\x75\x6d\x27\xe3\x4e\x70\xf2\xbd\x53\xd0\xc6\x59\x48\x38\xd5\x0f\xf3\x84\x80\x1e\xe3\xb4\x4a\xa8\xdf\x7e\xd3\x8d\x32\x4e\xd7\xd6\xbc\xa1\xb6\xdd\xd4\x1b\x48\xd3\x80\x65\x77\x2b\x43\x6b\x9d\x2d\xa4\x6b\xd7\x95\x5d\x18\xc4\x9c\x46\x82\x42\xd0\x43\x91\xf3\x88\xbf\x42\x44\x0f\x29\x5d\xdd\x7f\xbc\xcb\xfa\x73\x1a\xbe\x04\x00\xf0\xa9\xce\x95\x6e\xf0\x3d\xe2\xf1\x73\xc4\x51\x70\x76\xb8\x5c\x7c\xa3\xaf\xc4\xeb\x46\xf6\x0a\x82\x1e\x05\xb2\x5c\x20\xab\xd2\x74\x9e\x77\xf6\x43\x1b\xa1\x7b\x75\x72\xb2\x1f\xf0\xc0\x9e\xc0\x32\x31\x6b\x6b\x28\x7f\x48\xae\x3e\xb5\x1c\x9b\x42\xea\x51\x9b\x0f\x2f\x04\x5f\xf1\xc2\xc4\x73\x5e\x09\xf0\xfc\x85\x25\xfb\x80\x65\x25\xe5\xe2\xf2\x50\xbe\x4d\x8b\x92\xa6\x34\x16\x4b\x5e\xe2\x73\x91\x4d\x0a\x82\xb5\x31\xb2\x35\x7c\xe4\xf9\x61\x7d\x73\x1f\x04\xe1\x0e\xc9\x68\x07\xdf\xa2\xed\x9a\x2d\x9a\x4b\xd1\xa3\xaa\x47\x25\x9d\xc2\x4f\xed\x5a\xc8\xae\x5b\x2a\x3e\x2d\x94\x58\xf6\xf4\x1b\xa3\x1b\x1c\xaf\xdc\xa1\xd9\x85\x41\xc2\xf3\x62\xc3\x61\xbf\x82\xf2\x63\xf6\xe8\x2b\xa6\x47\x56\x8a\xf2\x96\xe7\x3f\xc1\xf4\x6f\x7c\x56\xcb\xbf\xf3\x59\x98\xcf\x98\x84\x6e\xce\x98\x06\x82\xe6\xca\xeb\xf6\x61\xfe\xd0\xec\xf5\xfe\x57\x00\x00\x00\xff\xff\x72\x67\x20\xf9\x7c\x03\x00\x00")

func _1606224181_drop_photo_path_from_accountsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1606224181_drop_photo_path_from_accountsDownSql,
		"1606224181_drop_photo_path_from_accounts.down.sql",
	)
}

func _1606224181_drop_photo_path_from_accountsDownSql() (*asset, error) {
	bytes, err := _1606224181_drop_photo_path_from_accountsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1606224181_drop_photo_path_from_accounts.down.sql", size: 892, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x90, 0x24, 0x17, 0x7, 0x80, 0x93, 0x6f, 0x8d, 0x5d, 0xaa, 0x8c, 0x79, 0x15, 0x5d, 0xb3, 0x19, 0xd7, 0xd8, 0x39, 0xf9, 0x3a, 0x63, 0x8f, 0x81, 0x15, 0xb6, 0xd6, 0x9a, 0x37, 0xa8, 0x8e, 0x9b}}
	return a, nil
}

var __1606224181_drop_photo_path_from_accountsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x93\xc1\x8e\x9b\x30\x18\x84\xef\x3c\xc5\x1c\x5b\x64\x89\x07\xc8\x89\x05\x77\xd7\x2a\x01\xe4\x35\xdd\xdd\x53\x71\xc0\x0a\x56\xc0\x46\xc4\xa8\xca\xdb\x57\x01\x42\x54\xd4\x56\xb9\xed\xf5\x1f\xff\xff\x8c\xe6\x93\x03\x1f\x91\xed\x2f\x70\x8d\x82\xac\x2a\x3b\x1a\x77\x86\x93\x87\x56\x41\x1b\x67\x21\xe1\x54\xd7\xcf\x13\x02\xfa\x1e\x25\x45\x4c\xa7\xd7\x65\xdf\x58\x67\x73\xe9\x9a\x12\xd2\xd4\x60\xe9\xac\x95\xba\x56\xc6\xe9\xca\x9a\x12\x95\x6d\xc7\xce\xc0\x0f\xbc\x88\xd3\x50\x50\x08\xba\xcf\x33\x1e\xf2\x0f\x88\xf0\x29\xa1\xab\xe5\xcf\x83\xac\x4e\x63\xff\xc5\x03\x80\x93\xba\x14\xba\xc6\x8f\x90\x47\x2f\x21\x47\xce\xd9\xfe\xba\xf1\x9d\x7e\x90\x49\x37\xb2\x53\x10\xf4\x5d\x20\xcd\x04\xd2\x22\x49\xe6\x79\x6b\x8f\xda\x08\xdd\xa9\xb3\x93\x5d\x8f\x27\xf6\x0c\x96\x8a\x59\x5b\x43\x4d\x8b\xe4\xe6\x53\xc9\xa1\xce\xa5\x1e\xb4\x39\x4e\x82\xf7\x15\x6f\x4c\xbc\x64\x85\x00\xcf\xde\x58\xbc\xf3\x58\xfa\x4a\xb9\xb8\x1e\xca\xb6\x69\xf1\x4a\x13\x1a\x89\x25\x2f\x99\x72\x91\x4d\x0a\x82\xb5\x26\xb2\x35\xfc\xc6\xb3\xfd\x7a\x73\xe7\x79\x81\x8f\x78\xb0\xfd\xd4\xae\x6d\xeb\x2d\x8f\x6b\xc9\x83\xaa\x06\x25\x9d\xc2\x2f\xed\x1a\xc8\xb6\x5d\x2a\x3e\x2f\x68\x58\xfa\xfc\x0f\x30\x93\x72\x47\xe3\x07\x5e\xcc\xb3\x7c\x83\x61\xb7\x72\xfa\x63\xfc\xf9\x58\xf0\x10\x97\xc7\x80\xac\x9e\xff\x07\xb2\x40\x9e\xb9\x08\x5d\x5f\x30\xf6\x04\xf5\x0d\xd0\xfd\x5b\xe0\xef\x5d\xde\xf6\x7f\x07\x00\x00\xff\xff\xc4\x05\x28\x49\x62\x03\x00\x00")

func _1606224181_drop_photo_path_from_accountsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1606224181_drop_photo_path_from_accountsUpSql,
		"1606224181_drop_photo_path_from_accounts.up.sql",
	)
}

func _1606224181_drop_photo_path_from_accountsUpSql() (*asset, error) {
	bytes, err := _1606224181_drop_photo_path_from_accountsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1606224181_drop_photo_path_from_accounts.up.sql", size: 866, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xff, 0x4c, 0x97, 0xee, 0xef, 0x82, 0xb8, 0x6c, 0x71, 0xbb, 0x50, 0x7b, 0xe6, 0xd9, 0x22, 0x31, 0x7c, 0x1a, 0xfe, 0x91, 0x28, 0xf6, 0x6, 0x36, 0xe, 0xb1, 0xf1, 0xc8, 0x25, 0xac, 0x7e, 0xd6}}
	return a, nil
}

var __1648646095_image_clockDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x92\xbf\x6e\xdb\x30\x18\xc4\x77\x3e\xc5\x8d\x49\x40\x20\x0f\xe0\x49\x56\x68\x84\x28\x2d\xaa\x14\xd3\x34\x93\xc0\x48\x5f\x25\x22\xfa\x07\x89\x42\xe0\x3e\x7d\x61\xd9\xb0\x6b\xa3\xf5\xd6\xa5\xe3\x77\x3c\x92\x77\xfc\xf1\xf1\x01\xd9\x57\xe5\x03\xa1\xec\x69\x42\xd7\x07\x4c\xf3\x30\xf4\x63\x40\x39\xf6\xc3\xe0\xbb\x0a\x45\xdf\xcc\x6d\x37\x71\xd4\xd4\x15\x84\x4f\x42\x3b\x4f\x01\xc5\x48\x2e\x10\x1c\x02\xb5\x03\x82\x7b\x6f\x08\x0f\x8f\x2c\x36\x22\xb2\x02\x56\x6c\x53\x6d\x22\xf3\x06\x1b\xad\x95\x80\x2f\xa9\x0b\x3e\xec\x72\xdf\xba\x8a\xa6\xfc\xdd\x15\x1f\xf3\x70\xc7\x00\xe0\x83\x76\xf9\xec\x4b\x7c\x8b\x4c\xfc\x1c\x19\xbe\x88\x9d\x6b\xe9\x52\x59\x76\xe6\x83\xdb\x35\xbd\x2b\xb1\x56\x7a\x8d\x44\x5b\x24\x2f\x4a\x1d\x0c\x9f\xbe\x0c\x35\x7c\x17\x0e\x63\x4d\xbe\xaa\xc3\x79\xfe\xe1\x1b\xca\x27\xff\x93\xce\xd2\x48\xfb\x39\x0f\x6e\xac\xe8\x37\x67\x6a\xe4\x76\x1f\xfd\x8b\x78\xc3\xdd\x31\x1d\x5f\x12\xdd\x43\x27\x88\x75\xb2\x51\x32\xb6\x30\x22\x55\x51\x2c\xd8\x3d\x5e\xa5\x7d\xd6\x2f\x16\x46\xbf\xca\xa7\x15\x63\x32\xc9\x84\xb1\x90\x89\xd5\x7f\xa9\x8e\x4c\x28\x11\x5b\x5c\x1c\xcf\x2f\x4b\xf2\x43\x25\x7e\xac\xc2\xcf\x15\xf8\x55\xf4\x8d\xd1\xdb\xeb\x8b\x56\x8c\x3d\x19\x9d\xfe\x19\xc0\x8a\xb1\x13\xaa\x65\x5d\x6e\x96\xd7\x14\xdf\x65\x66\xb3\x6b\xf7\x7f\xcb\xe9\x16\xa8\x7f\x4f\xe8\xf8\x15\x6e\x82\x3a\x79\x7e\x05\x00\x00\xff\xff\xce\x2f\xe3\x37\xab\x03\x00\x00")

func _1648646095_image_clockDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1648646095_image_clockDownSql,
		"1648646095_image_clock.down.sql",
	)
}

func _1648646095_image_clockDownSql() (*asset, error) {
	bytes, err := _1648646095_image_clockDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1648646095_image_clock.down.sql", size: 939, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4d, 0xa8, 0x1f, 0xf, 0xe0, 0xd7, 0xc9, 0x68, 0x98, 0xd8, 0x37, 0xb8, 0xba, 0x9e, 0xb2, 0x19, 0xf3, 0xc4, 0x73, 0x80, 0x3, 0x17, 0x2a, 0x53, 0x68, 0x10, 0x13, 0x54, 0x99, 0xb1, 0xf5, 0x1c}}
	return a, nil
}

var __1648646095_image_clockUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\x4c\x49\xcd\x2b\xc9\x2c\xa9\x8c\xcf\xcc\x4d\x4c\x4f\x2d\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xce\xc9\x4f\xce\x56\xf0\xf4\x0b\x51\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x51\x30\xb0\xe6\x02\x04\x00\x00\xff\xff\x22\x35\x20\xbf\x45\x00\x00\x00")

func _1648646095_image_clockUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1648646095_image_clockUpSql,
		"1648646095_image_clock.up.sql",
	)
}

func _1648646095_image_clockUpSql() (*asset, error) {
	bytes, err := _1648646095_image_clockUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1648646095_image_clock.up.sql", size: 69, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x98, 0xa6, 0xa4, 0x4e, 0x4e, 0xca, 0x17, 0x56, 0xea, 0xfb, 0xf0, 0xa9, 0x81, 0x95, 0xe, 0x80, 0x52, 0x1, 0x47, 0x9b, 0xde, 0x14, 0xfa, 0x72, 0xc9, 0x62, 0x6f, 0x24, 0xa2, 0xc, 0x32, 0x50}}
	return a, nil
}

var __1649317600_add_color_hashUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4c\x4e\xce\x2f\xcd\x2b\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xce\xcf\xc9\x2f\xf2\x48\x2c\xce\x50\x08\x71\x8d\x08\x51\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x51\x50\x52\xb2\xe6\x22\xca\x0c\xcf\x14\x05\x4f\x3f\x2c\x06\x18\x58\x73\x85\x06\xb8\x38\x86\x20\x69\x0d\x76\x0d\x41\xb2\xd7\x16\x6c\x07\x4e\x35\x9e\x29\x0a\xb6\x20\x43\x00\x01\x00\x00\xff\xff\xfa\xaf\xaf\xd9\xc9\x00\x00\x00")

func _1649317600_add_color_hashUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1649317600_add_color_hashUpSql,
		"1649317600_add_color_hash.up.sql",
	)
}

func _1649317600_add_color_hashUpSql() (*asset, error) {
	bytes, err := _1649317600_add_color_hashUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1649317600_add_color_hash.up.sql", size: 201, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1a, 0xf, 0x37, 0x6d, 0xcf, 0x99, 0xc9, 0x2e, 0xdc, 0x70, 0x11, 0xb4, 0x36, 0x26, 0x4f, 0x39, 0xa8, 0x44, 0xf, 0xcb, 0xcc, 0x81, 0x74, 0x7a, 0x88, 0xaa, 0x54, 0x8c, 0xc4, 0xe, 0x56, 0x4f}}
	return a, nil
}

var __1660238799_accounts_kdfUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4c\x4e\xce\x2f\xcd\x2b\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\x4e\x49\xf3\x2c\x49\x2d\x4a\x2c\xc9\xcc\xcf\x2b\x56\xf0\xf4\x0b\x51\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x51\x30\x36\x32\x30\xb0\xe6\x0a\x0d\x70\x71\x0c\x41\x32\x23\xd8\x35\x04\x4d\xb3\x2d\x54\x25\x20\x00\x00\xff\xff\x37\x9c\xbc\xd5\x73\x00\x00\x00")

func _1660238799_accounts_kdfUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1660238799_accounts_kdfUpSql,
		"1660238799_accounts_kdf.up.sql",
	)
}

func _1660238799_accounts_kdfUpSql() (*asset, error) {
	bytes, err := _1660238799_accounts_kdfUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1660238799_accounts_kdf.up.sql", size: 115, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdf, 0xe6, 0x7a, 0x69, 0x25, 0x42, 0x3b, 0x9c, 0x20, 0xf5, 0xcb, 0xae, 0xb0, 0xb3, 0x1b, 0x66, 0xc2, 0x5d, 0xd0, 0xc1, 0x59, 0xe8, 0xa9, 0xc5, 0x69, 0x58, 0x8f, 0xae, 0xe6, 0xd1, 0x4c, 0x53}}
	return a, nil
}

var __1679505708_add_customization_colorUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4c\x4e\xce\x2f\xcd\x2b\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2e\x2d\x2e\xc9\xcf\xcd\xac\x4a\x2c\xc9\xcc\xcf\x73\xce\xcf\xc9\x2f\x52\x08\x73\x0c\x72\xf6\x70\x0c\x52\x70\x71\x75\x73\x0c\xf5\x09\x51\x50\x2a\x28\xca\xcc\x4d\x2c\xaa\x54\xb2\xe6\x02\x04\x00\x00\xff\xff\x08\xb6\x89\xf4\x4e\x00\x00\x00")

func _1679505708_add_customization_colorUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1679505708_add_customization_colorUpSql,
		"1679505708_add_customization_color.up.sql",
	)
}

func _1679505708_add_customization_colorUpSql() (*asset, error) {
	bytes, err := _1679505708_add_customization_colorUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1679505708_add_customization_color.up.sql", size: 78, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa9, 0xe1, 0x3d, 0xaa, 0x5d, 0x35, 0x87, 0x8a, 0x8b, 0xe9, 0x4a, 0xa6, 0x7b, 0x85, 0xbc, 0x33, 0x11, 0xc7, 0x7d, 0x61, 0xac, 0x65, 0x59, 0xda, 0x32, 0x59, 0x68, 0x9d, 0xa1, 0x10, 0x7b, 0xa9}}
	return a, nil
}

var __1687853321_add_customization_color_updated_atUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x31\x0a\x43\x21\x0c\x06\xe0\xbd\xa7\xf8\x8f\xd0\xbd\x53\xaa\x16\x0a\x69\x84\x12\x0f\x20\x99\xa4\xd6\xc0\x53\x97\x77\xfa\xf7\x11\x6b\xfa\x42\xe9\xc9\x09\xd5\xcc\xf7\x58\x13\x14\x23\x42\xe6\xf2\x11\xd8\x9e\xcb\xff\xed\xac\xab\xf9\x08\xde\xfd\x08\xdd\xed\x87\xb7\x28\x24\x2b\xa4\x30\x23\xa6\x17\x15\x56\xdc\x1f\xb7\x2b\x00\x00\xff\xff\xfd\x48\x7a\xa4\x50\x00\x00\x00")

func _1687853321_add_customization_color_updated_atUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1687853321_add_customization_color_updated_atUpSql,
		"1687853321_add_customization_color_updated_at.up.sql",
	)
}

func _1687853321_add_customization_color_updated_atUpSql() (*asset, error) {
	bytes, err := _1687853321_add_customization_color_updated_atUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1687853321_add_customization_color_updated_at.up.sql", size: 80, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa8, 0xc2, 0x9, 0xec, 0xf4, 0xd1, 0x46, 0x29, 0xc5, 0xce, 0x4d, 0xd4, 0xf, 0x9c, 0xfa, 0x62, 0x1, 0x29, 0xe6, 0xd2, 0xd5, 0xe, 0xf0, 0x27, 0x81, 0x4a, 0x82, 0x25, 0x5f, 0x67, 0xff, 0xd1}}
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

	info := bindataFileInfo{name: "doc.go", size: 74, mode: os.FileMode(0644), modTime: time.Unix(1699005551, 0)}
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
	"0001_accounts.down.sql": _0001_accountsDownSql,

	"0001_accounts.up.sql": _0001_accountsUpSql,

	"1605007189_identity_images.down.sql": _1605007189_identity_imagesDownSql,

	"1605007189_identity_images.up.sql": _1605007189_identity_imagesUpSql,

	"1606224181_drop_photo_path_from_accounts.down.sql": _1606224181_drop_photo_path_from_accountsDownSql,

	"1606224181_drop_photo_path_from_accounts.up.sql": _1606224181_drop_photo_path_from_accountsUpSql,

	"1648646095_image_clock.down.sql": _1648646095_image_clockDownSql,

	"1648646095_image_clock.up.sql": _1648646095_image_clockUpSql,

	"1649317600_add_color_hash.up.sql": _1649317600_add_color_hashUpSql,

	"1660238799_accounts_kdf.up.sql": _1660238799_accounts_kdfUpSql,

	"1679505708_add_customization_color.up.sql": _1679505708_add_customization_colorUpSql,

	"1687853321_add_customization_color_updated_at.up.sql": _1687853321_add_customization_color_updated_atUpSql,

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
	"0001_accounts.down.sql":                               &bintree{_0001_accountsDownSql, map[string]*bintree{}},
	"0001_accounts.up.sql":                                 &bintree{_0001_accountsUpSql, map[string]*bintree{}},
	"1605007189_identity_images.down.sql":                  &bintree{_1605007189_identity_imagesDownSql, map[string]*bintree{}},
	"1605007189_identity_images.up.sql":                    &bintree{_1605007189_identity_imagesUpSql, map[string]*bintree{}},
	"1606224181_drop_photo_path_from_accounts.down.sql":    &bintree{_1606224181_drop_photo_path_from_accountsDownSql, map[string]*bintree{}},
	"1606224181_drop_photo_path_from_accounts.up.sql":      &bintree{_1606224181_drop_photo_path_from_accountsUpSql, map[string]*bintree{}},
	"1648646095_image_clock.down.sql":                      &bintree{_1648646095_image_clockDownSql, map[string]*bintree{}},
	"1648646095_image_clock.up.sql":                        &bintree{_1648646095_image_clockUpSql, map[string]*bintree{}},
	"1649317600_add_color_hash.up.sql":                     &bintree{_1649317600_add_color_hashUpSql, map[string]*bintree{}},
	"1660238799_accounts_kdf.up.sql":                       &bintree{_1660238799_accounts_kdfUpSql, map[string]*bintree{}},
	"1679505708_add_customization_color.up.sql":            &bintree{_1679505708_add_customization_colorUpSql, map[string]*bintree{}},
	"1687853321_add_customization_color_updated_at.up.sql": &bintree{_1687853321_add_customization_color_updated_atUpSql, map[string]*bintree{}},
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
