// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1691753758_initial.up.sql (5.738kB)
// 1692701329_add_collectibles_and_collections_data_cache.up.sql (1.808kB)
// 1692701339_add_scope_to_pending.up.sql (576B)
// 1694540071_add_collectibles_ownership_update_timestamp.up.sql (349B)
// 1694692748_add_raw_balance_to_token_balances.up.sql (165B)
// 1695133989_add_community_id_to_collectibles_and_collections_data_cache.up.sql (275B)
// doc.go (74B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
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

var __1691753758_initialUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\xcd\x8e\xdb\x36\x10\xbe\xfb\x29\x88\x3d\xad\x01\x1d\x9a\xb4\x28\x0a\xe4\xa4\xf5\x6a\x37\x42\x5d\x39\x95\xe5\x26\x39\x11\xb4\x34\x5e\x13\x2b\x89\x0a\x49\x39\x76\x9e\xbe\xa0\x7e\x49\x89\x92\xdd\x45\x7a\x94\x38\xe4\xfc\x7f\xf3\x91\xab\xd0\x73\x23\x0f\x45\xee\xc3\xda\x43\xfb\x94\xc5\xaf\x02\xdd\x2f\x10\x42\x28\x07\xf9\x9d\xf1\x57\x4c\x13\xb4\x0b\xb6\xfe\x73\xe0\x3d\xa2\x07\xff\xd9\x0f\x22\x14\x6c\x22\x14\xec\xd6\x6b\xa7\x12\x24\x49\xc2\x41\x08\xf4\x8f\x1b\xae\x3e\xba\xe1\x60\x75\x9f\xbe\xe2\xbc\xcc\xf6\xc0\xed\xbb\xd5\xfa\x91\x88\xa3\x7d\x35\x65\x24\x81\x04\x3d\x6c\x36\x6b\xf4\xe8\x3d\xb9\xbb\x75\x84\x9e\xdc\xf5\xd6\xab\x97\x57\x9b\x60\x1b\x85\xae\xda\x57\xe6\xf4\x5b\x09\x38\x23\x45\x41\xf3\x17\x7c\x60\x1c\x93\x38\x66\x65\x2e\xb1\x64\xb8\xf2\x0c\x17\xc0\x71\xe3\x16\xda\x05\xfe\xdf\x3b\x0f\xdd\x37\xd6\x3b\xad\x1d\x4e\xef\xf7\x72\xb1\xfc\xb0\x58\x58\x22\x84\x39\xc9\x5f\xe0\x7f\x08\xd4\x81\xb3\x6c\x3a\x4c\x92\x4d\xac\x91\x94\xe4\x31\xa0\x87\xf5\xe6\xa1\xfe\x93\x33\xf5\xed\x07\x91\xf7\xec\x85\xf3\x4e\x60\x01\xdf\x4a\xc8\x25\x25\xe9\xcf\xf7\x47\x48\xc2\x65\xb3\x59\xf3\x92\x72\x21\xa7\xdd\x4c\xc9\xd4\xea\xa7\xd0\xff\xcb\x0d\xbf\xa2\x3f\xbd\xaf\xe8\xbe\x37\xd3\x69\x2d\x59\x2e\x96\xe8\xb3\x1f\x7d\xdc\xec\x22\x14\x6e\x3e\xfb\x8f\x43\xcf\x0b\xc8\x13\x55\x1d\x92\x93\x5c\x90\x58\x52\x96\xff\xe7\x2c\x56\xc5\x6a\x77\x59\xd2\x0c\x84\x24\x59\x71\xe5\x08\x95\x66\x3c\x1f\x3d\xc9\x86\x02\xf5\x7f\x71\xc9\xf6\x2c\x35\xff\xbd\x10\x81\x0b\x4e\x8d\x0a\x50\xff\x52\x9a\x51\xa9\xfd\x3b\x91\xb4\xd4\x65\x12\x22\x09\x8a\xbc\x2f\x4d\x6a\xe4\xa5\x00\xf3\x60\x92\x24\x54\xc5\x88\xa4\x78\x20\x9b\x95\xa9\xa4\x7a\x18\x55\xdc\xba\x24\x4f\xe6\x49\xc5\xee\x6a\x92\xee\x04\x39\x41\xd2\xba\x0f\xe2\xae\xc9\xd0\x7c\xc0\x72\x92\x41\x65\xe0\x30\xd6\xe4\xc4\x4a\x4e\x25\x54\x20\xe2\xb9\x41\xb7\x6e\x03\x14\x0e\x19\x3b\x35\x80\x73\x4d\xb6\x2c\x12\x22\x01\xc7\xaa\x9b\x90\x9e\xe4\x4e\xf8\x97\x5a\x30\x3e\x12\x9a\x63\x71\x64\x5c\x62\x65\x66\xef\x42\x2b\x78\x77\x57\x4b\x42\x2e\x2a\x89\x49\x01\x2a\xb0\x04\xd5\x1e\x8d\x7d\x16\xb3\x62\x0e\x44\xaa\xf8\xc9\xca\xa8\x81\x2d\x46\x6a\x5a\xec\xeb\xf4\x3a\xad\x82\xab\x49\x92\xec\x15\x72\xdc\x20\x4f\xdb\x43\xa5\x00\x7e\xb5\xae\xd5\x3e\xc3\x47\x9b\x80\x59\xe6\x56\x91\x5b\xf4\xc4\x2c\x65\x7c\x24\x31\x8a\x6a\x2d\x9c\x40\x4c\x33\x92\x0a\x34\x6e\xd8\x56\x40\xc4\x9c\x16\xaa\xdc\x6f\x3d\xb3\xe4\x63\x27\x46\xa2\x2d\x7e\xdb\x3d\xa9\xab\xa7\x6e\xaf\x39\x3c\xd4\x83\xef\x74\xbb\x1c\x23\xa0\x4b\xb4\x09\xd4\xd0\x7c\x5a\xfb\xab\x08\x85\xde\xa7\xb5\xbb\xf2\xc6\x53\xae\xda\x22\x6e\x6b\xbb\x5b\x61\x73\xaa\x3f\x67\x33\xdd\xa5\xa4\x3b\xbb\xc3\x18\x23\xb5\x73\xb5\x6d\xcc\xf3\xf9\xa2\x3e\x51\x41\xf7\x29\x60\xc3\xff\x2e\xfe\x63\x1b\xa6\x42\x33\x8e\x68\x5c\x72\x0e\x79\x7c\x51\xbc\x24\x23\x12\xc7\x24\x3e\x42\xa3\x60\x3e\x02\x54\x14\x29\xb9\xcc\x15\xa7\x90\x9c\x16\x0a\x8b\x69\xaa\x66\xdb\x0f\xe0\x0c\xc4\x08\xc1\xc6\x36\x8d\x30\xbc\xf5\xf8\x86\xf1\x54\x8b\x08\x01\x72\x56\x20\x53\xfc\xeb\xe6\x09\x67\x11\x98\xd1\xa0\x8f\xab\x37\xce\x61\xa5\xc1\x30\x51\xb3\x7d\xba\xb0\xbb\xbd\xd7\x44\x62\xce\x84\xc0\xf2\xac\x04\xa6\x00\xbd\xd2\x25\xcf\x0d\x05\xee\x46\xb3\x64\xc6\x4f\x0b\x11\xad\x31\x03\x1f\xa9\x90\x8c\x5f\x26\x6b\xf5\x0d\xc4\xad\xad\xd5\x49\x5e\xa7\xe6\x9d\x3d\x9e\x5d\xdc\x2d\xcc\x8e\x4a\x95\xcb\x79\xf6\x3a\xf6\xb3\xe2\x35\x46\xbb\x54\xcd\xf9\x36\xd3\x6b\x92\x14\x7a\xee\x7a\xa6\x2d\xfc\xa7\x6a\xd1\xfb\xe2\x6f\xa3\xad\x82\x99\x14\x62\xa9\x60\x41\x60\xf6\x3d\x07\x2e\x8e\xb4\x30\x0c\xba\x31\xe8\x31\xcb\x25\x27\xb1\xbc\x69\x70\xd1\xa4\x8a\xc7\x60\xad\xd2\x3f\xb9\xdf\x82\xe3\xaa\xb7\x0f\xc0\x7f\x26\xc7\xbd\x4e\xfa\xe7\x08\xf2\x59\xab\x71\x01\x79\x02\x03\x08\xe7\x10\x03\x2d\x74\xd2\x9a\xb2\x17\xbd\x2f\xa6\x9b\xfe\xda\x45\xf3\x56\x50\xb0\x5d\x39\xdf\xdd\xc2\x7a\xf7\x44\x00\x56\xbc\xfb\x00\x83\x49\x37\xea\x7a\x21\x89\x2c\x45\xbf\xb5\x71\x1b\x57\xee\xf5\x10\x33\x82\x85\x94\xbd\x60\x9a\x27\x70\xd6\xb4\x56\x37\xdb\x81\x5c\x5c\x66\x65\x4a\x24\x3d\xd5\x06\x95\x02\x12\x7d\x72\x0e\x0a\xb1\xa7\xf5\x63\x61\x85\x5e\xa6\x42\x79\x1e\x98\x59\x70\x26\x21\x96\x3d\x6b\x1e\xde\x40\xc6\x83\xb3\xbb\xb1\xe0\x38\x25\x59\x01\xc9\xef\xbf\x99\x8b\xb2\xea\xb1\x62\x6a\xf9\x00\x30\xb5\x5c\x03\x3a\x2e\x48\x92\x40\xf2\xee\xfd\x1f\x47\x38\x23\x55\x2d\xf7\xbf\xbe\x5f\x36\x12\xcd\xb3\x40\x77\x49\x6e\x72\x42\x7f\xe8\xb1\x37\x38\xa6\x0e\xcc\x5a\x7b\x76\x01\x31\x26\xa6\xb1\xa2\x4d\xb9\xfe\xff\xd3\x26\xf4\xfc\xe7\x40\x11\x15\xfd\x7a\x34\x7c\x8a\x58\xa2\xd0\x7b\xf2\x42\x2f\x58\x79\xdb\xe6\xd2\x3e\x2f\xbe\x51\x17\x82\xb5\x17\x79\x68\xe5\x6e\x57\xee\xe3\xe4\x2b\x49\x0b\x0c\xd5\x8b\x48\x73\x8e\xf5\x75\xa4\x7a\x11\x69\x15\x4d\xbc\x8c\xf8\xc1\xa3\xf7\x65\x38\x90\xf0\x81\xa6\x12\x38\x86\x5c\x72\x0a\x42\x99\x36\x1a\x59\x3d\x47\xed\x69\x6b\x03\xe0\x4e\xed\xaf\xd3\x77\xad\xd3\x4c\x90\x91\x62\x9a\x9c\x3b\x7f\x04\xae\x1e\x10\xea\x0e\xde\x04\x3d\x00\xde\xf7\xe8\xe0\x34\x1d\xae\x1d\xd4\x78\x5b\x9f\xa7\x4d\x1c\x4c\x13\xc8\x25\x3d\x5c\x2a\x37\x2e\xea\x44\x63\x1e\x55\xc5\xd0\x1b\x3d\x75\xe2\x30\x34\xe3\x53\xdf\x10\x9a\x29\x65\x56\xa6\x69\x51\x39\xc1\x48\x9b\x7b\xc2\x7c\x90\xeb\xdc\x2e\xf4\x00\xa3\x7b\x1b\x38\xb6\xb1\x36\xf2\x58\x83\x9f\x4e\xcc\x9d\x61\x0f\x39\x66\xeb\x38\x66\x3b\x3a\x5d\x17\x3a\xd5\x40\xb0\x5a\x6b\x7b\xea\x51\x16\xdb\x9f\x80\xec\xc6\x0f\x4c\xd2\xec\x31\x6c\xd7\x5c\x9b\x8e\xde\x98\x6d\x2b\x6b\x6c\x1c\x7c\x5a\xab\x72\xd6\xd1\xb8\xb7\xa1\xbb\xa5\xcb\x8e\xce\xbd\x9d\x9e\xe3\x8e\xac\xba\x89\xe8\x8c\xdb\x78\x9e\x11\xf5\x85\x6b\x90\x95\xe5\x87\xc5\xbf\x01\x00\x00\xff\xff\x40\xaa\x79\x12\x6a\x16\x00\x00")

func _1691753758_initialUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1691753758_initialUpSql,
		"1691753758_initial.up.sql",
	)
}

func _1691753758_initialUpSql() (*asset, error) {
	bytes, err := _1691753758_initialUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1691753758_initial.up.sql", size: 5738, mode: os.FileMode(0644), modTime: time.Unix(1695161107, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6b, 0x25, 0x31, 0xc8, 0x27, 0x3, 0x6b, 0x9f, 0x15, 0x42, 0x2f, 0x85, 0xfb, 0xe3, 0x6, 0xea, 0xf7, 0x97, 0x12, 0x56, 0x3c, 0x9a, 0x5b, 0x1a, 0xca, 0xb1, 0x23, 0xfa, 0xcd, 0x57, 0x25, 0x5c}}
	return a, nil
}

var __1692701329_add_collectibles_and_collections_data_cacheUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\xc1\x8e\xda\x30\x10\x86\xef\x3c\xc5\x1c\x5b\x89\x37\xe8\x29\x24\x03\x8d\x1a\x99\x36\x24\xd5\xf6\x14\xcd\xda\x2e\x3b\xc2\xb1\x23\xc7\x59\x95\xb7\xaf\x80\x6e\xbb\x85\x75\xc8\xae\x84\xf6\xea\x19\xff\xb6\xe7\xfb\x7f\xa7\x25\x26\x15\x42\x95\x2c\x0a\x84\x7c\x09\x62\x5d\x01\xde\xe5\x9b\x6a\x03\xd2\x19\xa3\x65\xe0\x7b\xa3\x1b\x45\x81\x1a\x49\xf2\x41\xc3\x87\x19\x00\x80\x7c\x20\xb6\x0d\x2b\xa8\xc5\x26\x5f\x09\xcc\x60\x91\xaf\x72\x51\x1d\x05\x44\x5d\x14\xf3\x53\x9b\xb3\xc1\x93\x0c\x0d\x29\xe5\x75\xdf\xc3\xf7\xa4\x4c\x3f\x27\xe5\x59\x5b\x70\x3b\x7d\x54\x5b\x14\xeb\xc5\x59\xad\xf3\xee\x91\x95\xf6\x91\xad\x96\x5a\x1d\x29\x29\xdd\x4b\xcf\x5d\x60\x67\x23\x1d\x9d\xf6\x2d\x19\xb6\xbb\x48\x9d\x5b\xda\xea\x66\xf0\x26\x52\x27\xcb\x2d\x1d\xf4\x27\xf5\xb4\x5a\x31\x35\x61\xdf\xc5\x2e\x7c\x4f\x72\xb7\xf5\x6e\xb0\xaa\x91\xce\xb8\xd8\x93\x4f\xd3\x1a\x3c\x5f\xd4\x67\x1f\x3f\xcd\x66\x7f\x90\xd6\x22\xff\x56\x23\xe4\x22\xc3\xbb\x6b\x64\x59\x69\x1b\xf8\xe7\xbe\xd1\x36\xf8\x3d\xac\x45\x14\xfe\x13\xf7\xf9\x05\xda\xf9\x5f\x8a\xcf\x2e\x71\xcd\x57\xc1\x13\x87\xfe\x9d\x9c\x75\x3c\x7c\x8c\xc7\xa9\xe1\x91\xcc\x10\xb5\x18\xf7\x9d\xa1\xfd\x98\x48\x4b\xbf\x46\x25\x96\xeb\x12\xf3\x95\x80\x2f\xf8\x63\xda\x74\xa1\xc4\x25\x96\x28\x52\x8c\x65\x74\x9a\xce\xf1\x74\x38\xc0\xae\xbf\x66\x07\x58\x69\xb2\x49\x93\x0c\xff\xad\x67\x58\xe0\xb3\xf5\x89\x5c\x9d\xbd\xf5\x77\xf1\xf6\x2f\xa1\x37\xc3\xf6\xd5\x59\x7f\x65\xaa\x9e\xde\x1f\x0f\xd5\xd9\x84\xe2\xb4\xa6\x4f\xfc\xf6\x41\xba\x16\x96\x96\x2d\x94\x98\x14\x97\xf6\x7f\x69\x79\x9a\xeb\x5f\xf4\xfa\x7f\xd3\x1b\xdb\xfc\x16\x83\xff\x0e\x00\x00\xff\xff\xad\x71\x76\xba\x10\x07\x00\x00")

func _1692701329_add_collectibles_and_collections_data_cacheUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1692701329_add_collectibles_and_collections_data_cacheUpSql,
		"1692701329_add_collectibles_and_collections_data_cache.up.sql",
	)
}

func _1692701329_add_collectibles_and_collections_data_cacheUpSql() (*asset, error) {
	bytes, err := _1692701329_add_collectibles_and_collections_data_cacheUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1692701329_add_collectibles_and_collections_data_cache.up.sql", size: 1808, mode: os.FileMode(0644), modTime: time.Unix(1695161107, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1, 0x51, 0xf4, 0x2b, 0x92, 0xde, 0x59, 0x65, 0xd8, 0x9b, 0x57, 0xe0, 0xfd, 0x7b, 0x12, 0xb, 0x29, 0x6e, 0x9d, 0xb5, 0x90, 0xe, 0xfa, 0x12, 0x97, 0xd, 0x61, 0x60, 0x7f, 0x32, 0x1d, 0xc3}}
	return a, nil
}

var __1692701339_add_scope_to_pendingUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcd\x8a\xdb\x30\x14\x85\xf7\x7a\x8a\xc3\xac\x5a\xb0\x5b\xba\xf6\xca\x33\x56\xa1\xe0\xda\x43\x46\x81\xd9\x19\xd5\xba\x89\xc5\xe8\x27\x48\xd7\x8d\xf3\xf6\x25\x89\x4b\x53\x1a\xe8\xec\x84\x74\xbf\x4f\xe7\x70\xcb\x12\x2f\x44\x50\xcb\x0b\x6b\x9e\x33\x6c\x00\x27\x1d\xb2\x1e\xd9\xc6\x90\x3f\x1f\x28\x18\x1b\xf6\xbc\x70\xd2\xe3\x1b\xa5\x4f\xfb\x28\xea\x56\xc9\x0d\x54\xfd\xd8\x4a\xac\xef\xc3\x2d\x84\xba\x69\x90\xaf\x3e\x25\x5f\x15\xba\x5e\xa1\xdb\xb6\x2d\x1a\xf9\xb5\xde\xb6\x0a\x0f\xcf\x57\xec\xa1\x12\xa2\x2c\xa1\x26\xc2\x51\xf3\x38\x51\xc2\xd1\x3a\x07\x3d\x73\x84\x21\x47\x4c\xe0\x89\x7e\xff\x02\x5e\x32\x78\xd2\x0c\x9d\x08\x63\x0c\x3b\x9b\x3c\x19\xc4\x84\x9d\xb6\x8e\xcc\x59\x26\x5d\x5e\xa1\x14\xcd\x3c\x52\x82\x8f\x66\x76\x74\x35\x4f\xfa\x27\xe1\x2f\xb9\x87\xd7\x61\xd6\xce\x9d\x10\xc3\x19\x1a\x29\x67\x1b\xf6\xef\xab\xf9\xd4\xb7\xdb\xef\xdd\x25\xf0\xb0\x3a\x1f\xfb\xbe\x95\x75\xf7\x6f\xeb\x2f\x95\x10\xcd\xa6\x7f\xc6\xb7\xae\x91\xaf\xb0\x66\x19\xee\x79\x2b\x21\x9e\x36\xb2\x56\xf2\x3f\x73\xa2\xef\xee\xe7\xfa\xe0\x67\xc7\xf6\xf6\x6e\xb0\xa6\xc0\x2e\x45\x3f\x68\x63\x12\xe5\x5c\x80\xe3\x9f\x73\x20\x3e\xc6\xf4\x76\x99\x62\xeb\x29\xb3\xf6\x87\x02\xf9\xe4\x7f\x44\x57\x80\x4f\x07\x2a\xd6\x85\x16\xb7\x55\x3f\x56\xe2\x57\x00\x00\x00\xff\xff\x91\x52\x81\x7e\x40\x02\x00\x00")

func _1692701339_add_scope_to_pendingUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1692701339_add_scope_to_pendingUpSql,
		"1692701339_add_scope_to_pending.up.sql",
	)
}

func _1692701339_add_scope_to_pendingUpSql() (*asset, error) {
	bytes, err := _1692701339_add_scope_to_pendingUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1692701339_add_scope_to_pending.up.sql", size: 576, mode: os.FileMode(0644), modTime: time.Unix(1695161107, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x36, 0x8a, 0x5e, 0xe2, 0x63, 0x15, 0x37, 0xba, 0x55, 0x18, 0xf3, 0xcc, 0xe0, 0x5, 0x84, 0xe1, 0x5b, 0xe8, 0x1, 0x32, 0x6b, 0x9f, 0x7d, 0x9f, 0xd9, 0x23, 0x6c, 0xa9, 0xb5, 0xdc, 0xf4, 0x93}}
	return a, nil
}

var __1694540071_add_collectibles_ownership_update_timestampUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcf\x41\x4b\x80\x30\x18\xc6\xf1\xfb\x3e\xc5\x7b\x4c\xf0\x1b\x78\x9a\xba\x6c\x20\xaf\xa4\x5b\x78\x1b\xcb\xbd\xe1\x40\xa7\xb8\x45\xf8\xed\x03\x09\xa1\xa0\x43\xdd\xff\x3c\x3c\xbf\xaa\x17\x5c\x09\x50\xbc\x6c\x05\xc8\x47\xc0\x4e\x81\x18\xe5\xa0\x06\x98\xb6\x65\xa1\x29\xf9\xd7\x85\xa2\xd9\x3e\x02\x1d\x71\xf6\xbb\x79\xdf\x9d\x4d\x64\x92\x5f\x29\x26\xbb\xee\x11\x1e\x18\x00\xc0\x55\x18\xeb\xdc\x41\x31\xc2\x0b\xef\xab\x27\xde\x5f\x7b\xa8\xdb\x36\xbf\x9a\x69\xb6\x3e\x18\xef\x40\xe3\x20\x1b\x14\x35\x94\xb2\x91\xa8\x7e\x64\xf7\xf6\xaf\x1d\xcb\x0a\xc6\xbe\xbe\x6b\x94\xcf\x5a\x80\xc4\x5a\x8c\xff\x24\x18\xef\x28\x24\xff\x76\x1a\x0a\xe9\x38\xa1\xc3\x3f\xe8\xbf\xc1\xf3\xdb\x98\x15\xec\x33\x00\x00\xff\xff\x02\x64\xbc\xba\x5d\x01\x00\x00")

func _1694540071_add_collectibles_ownership_update_timestampUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1694540071_add_collectibles_ownership_update_timestampUpSql,
		"1694540071_add_collectibles_ownership_update_timestamp.up.sql",
	)
}

func _1694540071_add_collectibles_ownership_update_timestampUpSql() (*asset, error) {
	bytes, err := _1694540071_add_collectibles_ownership_update_timestampUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1694540071_add_collectibles_ownership_update_timestamp.up.sql", size: 349, mode: os.FileMode(0644), modTime: time.Unix(1695161107, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7f, 0x45, 0xc7, 0xce, 0x79, 0x63, 0xbc, 0x6f, 0x83, 0x5f, 0xe2, 0x3, 0x56, 0xcc, 0x5, 0x2f, 0x85, 0xda, 0x7e, 0xea, 0xf5, 0xd2, 0xac, 0x19, 0xd4, 0xd8, 0x5e, 0xdd, 0xed, 0xe2, 0xa9, 0x97}}
	return a, nil
}

var __1694692748_add_raw_balance_to_token_balancesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcd\xc1\x0a\x82\x40\x10\x87\xf1\xbb\x4f\xf1\xc7\xbb\xd0\x2b\x4c\x6a\x74\xd8\x14\x64\xed\x1a\xab\x3b\xc9\x92\xce\xc6\xce\x4a\x3d\x7e\x04\x75\xfd\x0e\xdf\xaf\xaa\x40\xde\x63\x8e\xeb\xbe\x09\xee\x31\x41\x73\x4c\x41\x16\x24\xf7\xc2\xe4\x56\x27\x33\x83\xdf\xcf\xc4\xaa\xec\x11\x04\x93\x53\xc6\x2e\x41\xb2\xc2\x29\xa6\xb0\x20\x48\xe6\x85\x13\x3c\xcf\x61\x73\x6b\x41\xc6\xb6\x03\x2c\x1d\x4d\x8b\x1c\x1f\x2c\xb7\xdf\x49\x41\x4d\x83\xba\x37\xe3\xa5\xfb\x0a\xff\x8e\x2b\x0d\xf5\x99\x06\x74\xbd\x45\x37\x1a\x83\xa6\x3d\xd1\x68\x2c\xca\x43\x59\x7c\x02\x00\x00\xff\xff\x88\x21\xef\xa1\xa5\x00\x00\x00")

func _1694692748_add_raw_balance_to_token_balancesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1694692748_add_raw_balance_to_token_balancesUpSql,
		"1694692748_add_raw_balance_to_token_balances.up.sql",
	)
}

func _1694692748_add_raw_balance_to_token_balancesUpSql() (*asset, error) {
	bytes, err := _1694692748_add_raw_balance_to_token_balancesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1694692748_add_raw_balance_to_token_balances.up.sql", size: 165, mode: os.FileMode(0644), modTime: time.Unix(1695211597, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd4, 0xe0, 0x5b, 0x42, 0xf0, 0x96, 0xa5, 0xf5, 0xed, 0xc0, 0x97, 0x88, 0xb0, 0x6d, 0xfe, 0x7d, 0x97, 0x2e, 0x17, 0xd2, 0x16, 0xbc, 0x2a, 0xf2, 0xcc, 0x67, 0x9e, 0xc5, 0x47, 0xf6, 0x69, 0x1}}
	return a, nil
}

var __1695133989_add_community_id_to_collectibles_and_collections_data_cacheUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xce\xcf\xc9\x49\x4d\x2e\xc9\x4c\xca\x49\x8d\x4f\x49\x2c\x49\x8c\x4f\x4e\x4c\xce\x48\x55\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xce\xcf\xcd\x2d\xcd\xcb\x2c\xa9\x8c\xcf\x4c\x51\x08\x71\x8d\x08\x51\xf0\xf3\x0f\x51\xf0\x0b\xf5\xf1\x51\x70\x71\x75\x73\x0c\xf5\x09\x51\x50\x52\xb2\xe6\x0a\x0d\x70\x71\x0c\xc1\x69\x5e\xb0\x6b\x08\xaa\x41\xb6\x60\x4d\x5c\xd8\x9c\x92\x9f\x47\x55\x97\xa0\x1a\x87\xc3\x21\x80\x00\x00\x00\xff\xff\x2e\x30\x6f\xa7\x13\x01\x00\x00")

func _1695133989_add_community_id_to_collectibles_and_collections_data_cacheUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1695133989_add_community_id_to_collectibles_and_collections_data_cacheUpSql,
		"1695133989_add_community_id_to_collectibles_and_collections_data_cache.up.sql",
	)
}

func _1695133989_add_community_id_to_collectibles_and_collections_data_cacheUpSql() (*asset, error) {
	bytes, err := _1695133989_add_community_id_to_collectibles_and_collections_data_cacheUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1695133989_add_community_id_to_collectibles_and_collections_data_cache.up.sql", size: 275, mode: os.FileMode(0644), modTime: time.Unix(1695211597, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfa, 0x2, 0xa, 0x7f, 0x4b, 0xd1, 0x3, 0xd0, 0x3, 0x29, 0x84, 0x31, 0xed, 0x49, 0x4f, 0xb1, 0x2d, 0xd7, 0x80, 0x41, 0x5b, 0xfa, 0x6, 0xae, 0xb4, 0xf6, 0x6b, 0x49, 0xee, 0x57, 0x33, 0x76}}
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

	info := bindataFileInfo{name: "doc.go", size: 74, mode: os.FileMode(0644), modTime: time.Unix(1695161107, 0)}
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
	"1691753758_initial.up.sql":                                                     _1691753758_initialUpSql,
	"1692701329_add_collectibles_and_collections_data_cache.up.sql":                 _1692701329_add_collectibles_and_collections_data_cacheUpSql,
	"1692701339_add_scope_to_pending.up.sql":                                        _1692701339_add_scope_to_pendingUpSql,
	"1694540071_add_collectibles_ownership_update_timestamp.up.sql":                 _1694540071_add_collectibles_ownership_update_timestampUpSql,
	"1694692748_add_raw_balance_to_token_balances.up.sql":                           _1694692748_add_raw_balance_to_token_balancesUpSql,
	"1695133989_add_community_id_to_collectibles_and_collections_data_cache.up.sql": _1695133989_add_community_id_to_collectibles_and_collections_data_cacheUpSql,
	"doc.go": docGo,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"1691753758_initial.up.sql":                                                     {_1691753758_initialUpSql, map[string]*bintree{}},
	"1692701329_add_collectibles_and_collections_data_cache.up.sql":                 {_1692701329_add_collectibles_and_collections_data_cacheUpSql, map[string]*bintree{}},
	"1692701339_add_scope_to_pending.up.sql":                                        {_1692701339_add_scope_to_pendingUpSql, map[string]*bintree{}},
	"1694540071_add_collectibles_ownership_update_timestamp.up.sql":                 {_1694540071_add_collectibles_ownership_update_timestampUpSql, map[string]*bintree{}},
	"1694692748_add_raw_balance_to_token_balances.up.sql":                           {_1694692748_add_raw_balance_to_token_balancesUpSql, map[string]*bintree{}},
	"1695133989_add_community_id_to_collectibles_and_collections_data_cache.up.sql": {_1695133989_add_community_id_to_collectibles_and_collections_data_cacheUpSql, map[string]*bintree{}},
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
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
