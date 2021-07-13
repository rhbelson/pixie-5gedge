// Code generated for package complete by go-bindata DO NOT EDIT. (@generated)
// sources:
// 01_base_schema.graphql
// 02_unauth_schema.graphql
// 03_auth_schema.graphql
package complete

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __01_base_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\xc5\xa0\x14\x54\x5c\x2a\x10\x4a\x49\x4f\x81\xe0\x07\x1c\x7b\x38\x47\x72\xbc\x3e\xef\x46\x47\x84\xf8\x77\x94\xcb\x5d\x77\xd5\x6c\x31\xf3\xb4\x4f\x43\xe2\xec\xf1\xeb\x80\xd3\xc2\xb6\x0e\xf8\xd8\xc2\x01\xf3\x62\xde\x26\x29\x03\xde\xaf\x97\xfb\x73\xae\xc3\x57\x22\xb4\x32\x20\x0a\xb5\x3c\x1a\x7c\xce\x72\x06\xe7\x6a\x2b\x6c\xad\xd4\x83\xeb\xf0\x29\x38\x13\xa1\xd1\x1b\x51\x7d\x0e\x4c\x92\x23\x9b\x22\xb1\x11\xbe\xc4\xeb\xce\x12\x95\xfb\x0e\x26\x18\xe9\x3a\xf0\xc7\x58\x22\x23\xc6\x15\x62\x89\x0d\xdf\x53\xde\xb9\xc9\xac\xea\xd0\xf7\xc7\xc9\xd2\x32\x1e\x82\xcc\xfd\xb1\xf9\x9a\x4e\xf9\x96\x4f\xdb\x73\xfd\xa4\xba\x50\xfb\xe7\x97\x57\xe7\x36\xf8\xae\x75\xf1\x2c\x22\x75\xc0\x9b\x48\xa6\x2f\x0f\x9b\xd4\xa5\x70\xb3\xbc\xdf\xf9\x0f\x00\x00\xff\xff\x6f\xc4\xb8\xef\x28\x01\x00\x00")

func _01_base_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__01_base_schemaGraphql,
		"01_base_schema.graphql",
	)
}

func _01_base_schemaGraphql() (*asset, error) {
	bytes, err := _01_base_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_base_schema.graphql", size: 296, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_unauth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\x31\x0a\x02\x31\x10\x45\xfb\x39\xc5\x4f\xa7\x57\x48\x67\x23\x58\x28\x88\xa5\x58\x0c\xeb\xec\x1a\xd8\x24\x4b\x66\x14\x17\xd9\xbb\x0b\x81\x88\x62\x37\xbc\x79\xbc\x2f\x4f\x93\x74\x85\xcd\x93\xe0\x78\x97\x32\xe3\x45\x00\x17\x0b\x3d\x77\xa6\xab\x76\x1d\x38\x8a\xc7\xc9\x4a\x48\x83\x5b\x7b\x6c\x9a\xb1\x4b\x7d\x76\xb4\x10\xd5\xc4\x0f\xae\xa9\x60\x12\xd5\xe3\xdc\x3e\xee\xf2\x6f\x57\xf1\x21\x45\x43\x4e\x9f\x11\x02\xba\x1b\xa7\x41\xc6\x3c\x7c\x43\x0b\x51\xd4\x38\x4e\x7b\xf5\xd8\x8e\x99\xcd\xd1\xf2\x0e\x00\x00\xff\xff\xc2\xab\x64\xad\xc7\x00\x00\x00")

func _02_unauth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__02_unauth_schemaGraphql,
		"02_unauth_schema.graphql",
	)
}

func _02_unauth_schemaGraphql() (*asset, error) {
	bytes, err := _02_unauth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_unauth_schema.graphql", size: 199, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_auth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x5d\x6f\xe3\x36\xd6\xbe\xf7\xaf\x38\xd3\xb9\x68\x02\x04\x83\xe2\xc5\xdb\x62\xe1\xab\x55\x6d\xb5\xa3\x4d\xe2\x78\x63\x67\x66\x8b\xc5\x60\x40\x8b\xc7\x16\x11\x89\x54\x49\x2a\x89\x77\x31\xff\x7d\x71\xf8\x21\x89\xb6\xd2\xd9\x29\xf6\x4e\xe2\xc7\xc3\xe7\x7c\xf2\x1c\xe2\x8b\x45\xc9\xc1\x1e\x5b\x84\xbf\x77\xa8\x8f\xf0\xef\x19\x40\x67\x50\xcf\xe1\xc1\xa0\x2e\xe4\x5e\xbd\x99\x01\x28\x7d\x98\xc3\x9d\x3e\xc4\x7f\x5a\xb1\x41\x6b\x85\x3c\x18\xbf\x32\xfe\xc5\xd9\xcc\x5a\x2d\x76\x9d\xc5\x30\x3f\xfc\x07\x3c\x1a\x34\x73\xf8\x67\x7f\xcc\x27\x9a\x28\xeb\xce\x58\xd4\x17\x82\xcf\xa1\x58\xbe\xb9\x9c\xc3\xc2\x8f\xc4\x93\xc3\x82\x9f\x8f\x2b\xd6\xe0\x85\x64\x0d\xce\x61\x63\xb5\x90\x87\xd7\x17\xd3\x31\xe3\x99\xf1\x49\x0b\x25\x25\x96\x56\x28\x79\x7e\xe6\x30\x37\x00\x8a\x4c\x5b\xb1\x67\xa5\xbd\x60\xe1\x63\x7b\x6c\x71\x0e\xd9\xe8\xcf\x41\xdc\x14\x71\x88\x36\xb2\xce\xaa\x52\x35\x6d\x8d\x16\x2f\x84\x6c\x3b\x1b\x69\x5f\x41\xd9\x69\xa3\xf4\x5a\x99\x39\x14\xd2\x5e\x01\x73\x47\xce\x21\x1b\xed\xc9\xdc\x18\x81\x5f\x45\xe6\x0f\xc5\x32\x62\x5c\xa6\x8b\xef\xd1\x74\xf5\xd9\xb1\xbf\x08\xac\xf9\xe9\xd9\x7b\x1a\x0c\x12\x8c\xd6\xe6\xd2\x0a\x7b\xbc\x16\x92\x5f\xcd\x00\x00\x34\xfe\xde\x09\x8d\x3c\xd3\x07\x5a\x4c\x0a\x9d\x5e\xfe\xe9\x15\x7a\xc9\xf2\x4d\x77\x38\xa0\x21\x81\x3e\xcd\x66\x00\x6f\x61\x53\x6a\xd1\xda\xe6\xa0\x01\x25\x6f\x95\x90\xd6\x5c\x81\xc6\x3d\x6a\xb0\x0a\xb8\x2a\x0d\x08\x09\x65\xad\x3a\xce\x5a\xf1\xae\xd5\xca\xaa\x19\x40\x2d\x9e\xf0\x83\xc0\x67\xa2\x73\x13\xbe\x6f\xd1\x32\xce\x2c\xf3\x46\x8e\x2b\x16\x4a\x5a\x94\xd6\x8c\x6c\x7c\x73\x32\x45\xcb\x8d\xe3\x41\x70\x9e\x51\x0a\xe6\x67\x27\xa0\x36\xc9\xc4\x1b\x2f\xd3\x12\xdb\x5a\x1d\xe1\x11\x8f\x66\x06\xc0\xdd\x5f\x83\xd2\x5e\xe3\x91\x0e\x58\x8e\x07\x3c\x7e\xb2\x66\x04\x9f\x2e\xf5\xe8\xd9\xba\x88\xd0\xac\x15\x01\x33\x5b\x17\x3d\x98\x1f\x1d\xa1\x84\xc9\xd9\x97\xd9\x6c\x1c\xf5\xb7\x9d\x65\x64\x09\x17\xf8\x0b\x8d\xcc\x62\xf0\xfe\x24\x9a\xe0\xaf\x1c\x5b\x8d\x25\xb3\xc8\x2f\x34\x32\x43\x0e\xfa\x5d\x58\x60\x80\x69\x04\xa9\x9e\xa1\x74\x00\x1c\x9e\x04\x83\xf6\x25\x48\xf4\xdd\xe5\x0c\xe0\xa1\xe5\xcc\xe2\x07\xf1\x2f\xe1\xe2\x6a\x2f\x0e\x17\xc1\x51\xc8\x4f\x8a\xe5\x9b\x2b\x78\x1a\x4d\xce\x21\xe7\xc2\xb2\x5d\x9d\x6c\x99\x08\x71\x4f\x39\x51\xd1\x99\xc6\x00\x96\x48\x7e\xb7\x7c\x45\xc1\x3f\x2b\x55\x23\x93\x03\x9c\xd7\xd5\xa0\xb3\x08\xe0\xff\xa7\x77\x7a\x01\xc7\xa9\xf0\xc2\xf4\x19\x32\x0a\x93\x64\xca\xcb\xf3\xcc\xb9\x41\x9b\x26\xcb\x0b\x36\xca\xa3\x63\x94\x51\x3e\xbd\x9c\xca\xb0\x85\x7c\x12\x9e\xce\x05\x36\x4c\xd4\x7d\x96\xa4\x98\xd7\xc6\xae\xc6\x99\xf3\x0a\x6a\x76\x32\x74\x19\x2f\x00\x82\x49\xe5\x5b\xa3\x6e\x84\x31\x42\x49\x73\x41\xa9\xbe\x37\x60\x97\x4e\xa6\x84\x47\x13\x03\xb8\xb7\xa1\x87\xbe\xd3\x87\x5e\x73\x4a\x1f\x7a\x54\x35\x8c\x0f\x88\xa3\xc5\x84\xd6\x5f\x4d\x5f\x66\x33\xe7\xd6\x11\xde\xb9\x75\xb0\xd7\x0c\x20\xb9\x2f\x66\x00\xa9\x6a\x66\x00\xad\x28\x6d\xa7\x93\x35\x4a\x1f\x56\x27\xdb\x02\xbd\x61\x40\x98\xac\x6d\xb5\x7a\x42\x3e\xf2\x89\xc8\x25\x90\xfb\x1a\x15\x49\x72\x79\x18\x56\x9b\x09\x9c\xb1\xb3\x38\x30\x26\x59\x7d\xb4\xa2\x34\x77\xad\x55\x94\xd5\x27\xb7\x0c\x7e\xe1\x36\x59\xd5\xe9\x0d\xa2\x4c\x57\xbb\x6b\xe1\x15\x07\x9b\xde\x36\xbd\xeb\xbf\xe2\xd7\xd3\xf3\xf1\x74\xa2\x98\x47\x8a\xbc\x41\x2f\x21\xa5\x64\xf6\xd6\xcc\xe1\x97\x5a\x31\xeb\x73\xa5\x29\x87\x55\x11\x2f\x09\xf0\xff\x09\x2c\xca\xae\x99\xb8\x19\x37\x96\x59\x74\x07\x64\xf9\xe6\xf3\xc3\xea\x7a\x75\xf7\x71\x15\xfe\xd6\xf9\x6a\x59\xac\x7e\x0d\x7f\xf7\x0f\xab\xd5\xf0\xf7\x4b\x56\xdc\xe4\xcb\xf0\xb3\xcd\xef\x6f\x8b\x55\xb6\xcd\x97\x93\x27\x0d\x57\xbe\x3f\x28\xdb\x8e\x0e\x7a\x0b\x99\x04\xe4\xc2\x86\x6a\x01\x54\x49\x65\x04\x88\x3d\x30\x17\x87\x50\x31\x03\x8d\xe2\x62\x2f\x90\x83\xad\x10\xbc\xb1\x2c\xbe\x58\xd8\x1d\x41\x48\x83\x9a\x4c\x05\x4a\x03\xa7\xec\x46\xdf\x65\xc5\x34\x2b\x29\xa5\xbf\x73\x87\x6c\x2b\x41\x57\x6f\x59\x77\x1c\x0d\x5d\x18\x6e\x83\x74\x78\x8f\x78\xdc\x29\xa6\x39\x30\xc9\xa1\x65\xc6\x03\xa8\xa6\x61\x92\xbb\xed\xc4\x38\x5f\x16\x5b\x4f\x17\x0c\xd6\x58\x0e\x7c\x65\x7d\x9c\x26\x5d\x56\xca\xa0\x04\x26\x93\xea\x05\x4c\x5f\x34\xbc\x8b\xb4\xb8\xa0\xfb\xc8\x80\x2b\x06\xde\x3a\x52\xc9\x16\x5b\x31\x0b\xc2\x82\xa9\x54\x57\x73\x68\xd4\x13\xba\x45\x74\xd4\xf7\x26\xd4\x5d\x54\x61\xd0\xa0\x24\xc5\x30\x0a\xbf\x56\x0b\xb2\xae\x65\xbb\x28\xc5\x26\xbf\xc9\x17\xdb\x3f\xf0\x07\x2a\x7d\x82\x3b\x5c\x27\xee\x70\xfd\x79\x7d\xb7\x0c\x5f\x9b\x0f\x8b\xf8\xb5\xb8\x2f\xd6\xdb\xf0\xb3\xca\x6e\xf3\xcd\x3a\x5b\xe4\x43\x58\x4c\xd6\x4a\x0e\xff\x51\x48\xfe\x5a\xa9\x76\x92\x54\x82\x3b\x53\x69\xe2\xca\xc9\x7e\xb4\x61\xb6\xac\x90\x17\x92\xe3\x8b\x2b\xe5\x0a\x69\x3f\x51\x7d\x43\x4e\x3d\x05\xee\xbc\xbd\x67\xb7\x65\xbb\x13\x52\xe4\x27\xe4\x5f\x1c\x5f\x40\xed\x9d\x36\x2d\xdb\x79\xf5\xdb\x0a\xcd\xd8\x78\xbe\x56\xd8\x2b\x4d\xba\xb5\x6c\xe7\x58\xb8\xc2\xd7\x01\x7d\xac\xd0\x56\xa8\x83\xb3\x90\x47\xb1\xd1\x66\xda\x07\x96\x8c\x4f\xf8\xfe\xc0\x67\x51\xd7\xd0\xb0\x47\x6f\xda\xe0\x7f\x80\x2f\x58\x76\x2e\x2b\xd1\x39\xc3\x5f\xb6\xb7\x94\xa4\x08\x7c\x48\x47\x30\xe6\xf7\x07\xb5\xea\x94\x7d\x7c\xad\x3d\x52\xc3\x5e\xe9\x86\x59\x2a\x82\x7c\xc0\x11\xd9\x3e\xfa\x4c\x28\xbb\x9f\x2b\x51\x56\xce\xdb\x77\x88\x12\x5a\xa6\x0d\x72\x0a\xcb\x73\x1f\x56\xbd\xa3\x7b\x27\x67\xbb\x8d\x55\x2d\xb4\xca\x08\xc7\x97\xe4\xeb\xcf\x2c\xc6\xd5\x7d\xa2\xd0\x53\x0e\xc4\x8b\xc1\x13\xab\x05\xbf\x1a\xe9\x27\x2a\xf0\x9d\xbb\xd4\xf2\x7e\x7c\xac\xac\xb7\x90\xd5\x75\x62\x52\x32\x0b\xb2\xb2\x1a\x59\x9f\x48\x9a\x60\xe3\x4d\xa2\xdd\xc4\x7f\x06\xa5\x52\xf9\xcc\x84\x44\x4d\xde\xd6\xf9\x0b\xe4\xf4\x8e\x9c\x4e\xda\xc1\x6f\x87\x65\x0d\x1a\xc3\x0e\xc9\x50\x2c\x59\x4f\x6f\x8c\xeb\xbf\x98\xfc\x09\xa5\x37\xe0\xc4\x3e\x57\x2b\x6d\x45\x83\xc9\x89\x54\x2d\x9d\x0c\x46\xc0\xb5\xe2\x7f\x4a\x80\xce\x7c\xa3\x04\x00\x65\xd4\x98\xeb\x6f\x53\xf5\xf9\x26\x00\x49\x34\x9a\x8d\x62\xd2\x70\xcc\x62\xa1\x90\x1e\xb1\xf5\x0e\xcc\x71\xcf\xc8\xa5\x9d\x5a\x29\x37\x4b\x65\xab\xe0\x31\x8f\x52\x3d\x4b\xb2\xea\x62\x93\x5c\x46\xb4\x2f\xac\x37\x50\x21\xab\x6d\x75\xa4\xad\x15\x32\x6d\x77\xc8\xac\x0f\x7b\x8d\x25\x8a\x27\xe4\x74\x85\x68\x3c\x74\x35\xd3\x20\xa4\x45\x4d\x15\x8f\xbb\x47\x6c\xe5\xdd\x3c\xf4\x06\x04\xa7\xd1\xb4\x4a\x72\x62\x60\x95\xeb\x44\xd1\x58\x13\x48\xbc\xcf\xb3\x9b\xed\xfb\xdf\xce\x49\x74\x72\x44\xc3\x65\x86\x01\xb1\xf4\x7d\x3d\xdd\x8b\x0a\xd6\xe2\x45\x20\x2c\xa8\xb7\x74\x0c\x84\x01\x2a\xc1\x04\x8f\x11\x34\xc8\x70\x05\x3b\x17\xd0\xf2\x7b\x0b\xbf\x77\xa8\x8f\x2e\x62\xc8\xf9\x8d\x6a\x30\x58\x28\xdc\x4e\x1a\x0d\x36\xbb\x1a\x0d\xbc\xdf\x6e\xd7\xdf\x1b\xf8\xf1\x87\x1f\x82\xa1\x7b\xfd\x4d\x93\x77\x09\xed\xa0\x5c\xe7\x2b\xcc\xc0\x35\xc8\xf1\xeb\xfd\x7a\x11\x25\xa0\x94\xb8\xd3\xc8\x1e\xcd\x3b\x07\x50\xa9\x16\x7d\xc2\x61\xb6\xbf\x12\xa3\xe0\x0e\xb7\x24\xa2\x3b\x56\x3e\xd2\x05\x2c\x24\x3a\x91\x35\x9a\xae\xa1\xf4\x00\x81\x91\x67\x12\x78\x2e\x8b\xcd\xe2\x6e\xb5\xca\x17\x5b\x57\xb9\x9c\xea\x99\x6a\x77\xb2\xcd\x73\x85\xf2\x54\xd1\xc2\x8f\xb4\x5a\x95\x68\x0c\x65\x87\xb8\x3c\xea\x60\xbd\xcc\xb6\xbe\x3c\xf2\xb8\xbe\x07\xf4\x75\x40\x94\xdc\xab\x9d\x86\xa4\xb2\x60\x28\x5a\x99\x3c\x82\x72\x79\x6d\xdf\x69\x7f\x61\x78\x37\x76\xf8\x68\x80\xed\x54\xe7\x55\xf0\x1c\x12\xa0\xb0\x63\xdf\x54\xfa\x94\xca\xb9\x8c\x81\xcb\x33\x33\x60\xf5\x31\xf8\x9f\x3f\xc0\x53\xda\x33\x51\x63\xef\x35\xd4\xff\x0a\x09\x0c\x76\x8c\x27\x0a\x74\x42\xe6\xb1\xf6\x8b\x89\x62\xdc\xd7\xba\xe8\x6b\x99\x31\xb6\xd2\xaa\x3b\x54\xb9\xeb\x05\xa6\x7a\x89\x71\x4b\x9e\xd6\xb7\x31\x89\x24\x61\x1d\x93\xd5\xfb\xe8\xc3\x49\xde\x49\x1b\xee\xa4\xd1\xee\x67\x3f\xa0\x36\xe2\x24\xef\xf8\x13\x5e\x9f\x39\x6d\x97\x5a\x8d\xd6\x1e\x17\xd3\x93\xe7\xcf\x45\x31\xb7\x69\x55\xaf\x6b\x26\xb1\x4f\xa9\xae\x58\xe9\xff\x7c\x8a\x93\x5d\xb3\x52\x1c\xfd\xdb\x59\x18\x28\xa4\xb1\xba\xa3\x36\x00\xf9\x78\xf2\x44\x89\xe9\xf3\x9e\x57\x67\x9b\x71\xae\xd1\x24\xc9\xd8\xaa\x47\x9c\xb8\x3a\x86\xf6\xd8\x6d\x3d\x6b\x25\x85\x9b\xbb\x11\xf2\x31\xd9\xfb\x16\xee\xbf\xf2\xb0\xe5\xd0\x4f\xdf\xb3\xbe\xd2\x3d\x9e\x35\x2e\xdf\x78\x4c\x7c\xbc\x0a\xf7\xa0\x3f\x73\x7e\xc6\xc2\x59\xf3\xa5\x8e\xab\xc7\x0c\x9e\x84\xf9\xdb\xe6\x6e\xf5\x67\x48\xa4\x8f\x6d\xdf\x24\x29\x50\x5e\x88\x2c\xd3\x78\xf9\xa6\xc3\x5f\x91\xff\xe4\x19\x30\x38\x66\x2a\x7a\xdf\x15\x8c\x5e\x80\x1d\x0c\x40\xd2\xb2\xb9\xdf\x9b\x62\xf5\xf0\x8f\xcf\xd9\xed\xf2\xa7\xff\x8f\x43\xcb\xec\xfe\x63\xb1\x4a\xc7\x16\x77\xab\x6d\x56\xac\xf2\xfb\xcf\x9b\x7c\xfb\xf9\xb7\xec\xf6\x66\x33\x3d\x35\x81\x97\x2e\xd8\xe6\xb7\xeb\x1b\x4a\x3f\x1e\xa4\x0f\x81\xe1\x79\xda\x3f\xf9\xeb\xc4\x77\x4d\xc5\xfe\xef\xc7\x9f\x12\x19\xd3\x5e\xff\x5b\xb2\xd7\xf4\x4b\xc1\xe8\x3d\xc8\x5b\xfc\xfc\x09\xe5\x7c\xe3\xe8\xd9\xc7\x07\xdd\x2b\x8f\x26\xb3\x2f\xff\x09\x00\x00\xff\xff\x4a\x78\x3e\x3a\xda\x18\x00\x00")

func _03_auth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__03_auth_schemaGraphql,
		"03_auth_schema.graphql",
	)
}

func _03_auth_schemaGraphql() (*asset, error) {
	bytes, err := _03_auth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_auth_schema.graphql", size: 6362, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
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
	"01_base_schema.graphql":   _01_base_schemaGraphql,
	"02_unauth_schema.graphql": _02_unauth_schemaGraphql,
	"03_auth_schema.graphql":   _03_auth_schemaGraphql,
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
	"01_base_schema.graphql":   &bintree{_01_base_schemaGraphql, map[string]*bintree{}},
	"02_unauth_schema.graphql": &bintree{_02_unauth_schemaGraphql, map[string]*bintree{}},
	"03_auth_schema.graphql":   &bintree{_03_auth_schemaGraphql, map[string]*bintree{}},
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
