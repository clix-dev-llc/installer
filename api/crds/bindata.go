// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// installer.kubedb.com_kubedbcatalogs.yaml
// installer.kubedb.com_kubedboperators.yaml
package crds

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

// ModTime return file modify time
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

var _installerKubedbCom_kubedbcatalogsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x8f\x1b\x37\x0c\xbd\xfb\x57\x10\xe8\x21\x97\x7a\x8c\x45\x2f\xc5\xdc\xda\x4d\x0b\x04\xe9\x17\xb2\x69\xee\x1c\x89\x1e\xb3\xab\x91\x14\x92\x72\x77\xfb\xeb\x0b\x69\x6c\xaf\xbd\xf6\x26\x5d\x03\x9d\x9b\x9e\xc8\x47\xf2\x91\xd2\x08\x33\x7f\x22\x51\x4e\xb1\x07\xcc\x4c\x0f\x46\xb1\xae\xb4\xbb\xff\x5e\x3b\x4e\xab\xed\xcd\x40\x86\x37\x8b\x7b\x8e\xbe\x87\xdb\xa2\x96\xa6\x0f\xa4\xa9\x88\xa3\xb7\xb4\xe6\xc8\xc6\x29\x2e\x26\x32\xf4\x68\xd8\x2f\x00\x9c\x10\x56\xf0\x23\x4f\xa4\x86\x53\xee\x21\x96\x10\x16\x00\x01\x07\x0a\x5a\x6d\x00\x30\xe7\xee\xbe\x0c\x24\x91\x8c\x5a\xa8\x88\x13\xf5\x50\x31\x3f\x2c\x00\x8e\x97\x0e\x0d\x43\x1a\xb5\xe3\xa8\x86\x21\x90\x74\xf3\x46\xe7\xd2\xb4\xd0\x4c\xae\x92\x8e\x92\x4a\xee\xe1\xa2\xcd\xcc\xb7\x8b\xed\xd0\x68\x4c\xc2\xfb\xf5\xf2\x29\x6a\x5d\x60\xce\xea\x92\xa7\xb6\x9c\x0b\x7f\x5f\x06\x7a\xfb\xe3\xed\x9c\x46\xc3\x03\xab\xbd\x3f\xdf\xfb\x85\xd5\xda\x7e\x0e\x45\x30\x3c\x2f\xa0\x6d\x29\xc7\xb1\x04\x94\x67\x9b\x0b\x80\x2c\xa4\x24\x5b\xfa\x33\xde\xc7\xf4\x77\xfc\x99\x29\x78\xed\x61\x8d\x41\x6b\x36\xea\x52\xa6\x1e\x7e\xab\x95\x64\x74\xe4\x17\x00\x5b\x0c\xec\x9b\xde\x73\x2d\x29\x53\xfc\xe1\x8f\x77\x9f\xbe\xbb\x73\x1b\x9a\x70\x06\x2b\x73\xca\x24\x76\x28\x79\x6e\xc1\xa1\xf9\x07\x0c\xc0\x93\x3a\xe1\xdc\x18\xe1\x4d\xa5\x9a\x6d\xc0\xd7\x76\x93\x82\x6d\x08\xb6\x33\x46\x1e\xb4\x85\x81\xb4\x06\xdb\xb0\x82\x50\xab\x21\x5a\x4b\xe9\x88\x16\xaa\x09\x46\x48\xc3\x5f\xe4\xac\x83\xbb\x5a\xa7\x28\xe8\x26\x95\xe0\xc1\xa5\xb8\x25\x31\x10\x72\x69\x8c\xfc\xcf\x81\x59\xc1\x52\x0b\x19\xd0\x68\xa7\xed\xfe\xe3\x68\x24\x11\x43\x15\xa1\xd0\xb7\x80\xd1\xc3\x84\x8f\x20\x54\x63\x40\x89\x47\x6c\xcd\x44\x3b\xf8\x35\x09\x01\xc7\x75\xea\x61\x63\x96\xb5\x5f\xad\x46\xb6\xfd\xb8\xbb\x34\x4d\x25\xb2\x3d\xae\x5c\x8a\x26\x3c\x14\x4b\xa2\x2b\x4f\x5b\x0a\x2b\xe5\x71\x89\xe2\x36\x6c\xe4\xac\x08\xad\x30\xf3\xb2\x25\x1e\xad\x9d\x99\xc9\x7f\x23\xbb\xb3\xa1\x6f\x8e\x32\xb5\xc7\xda\x36\x35\xe1\x38\x1e\xe0\x36\x58\x2f\xea\x5e\x47\x0b\x58\x01\x77\x6e\x73\xfe\x4f\xf2\x56\xa8\xaa\xf2\xe1\xa7\xbb\x8f\xb0\x0f\xda\x5a\x70\xaa\x79\x53\xfb\xc9\x4d\x9f\x84\xaf\x42\x71\x5c\x93\xcc\x8d\x5b\x4b\x9a\x1a\x23\x45\x9f\x13\x47\x6b\x0b\x17\x98\xe2\xa9\xe8\x5a\x86\x89\xad\x76\xfa\x73\x21\xb5\xda\x9f\x0e\x6e\x31\xc6\x64\x30\x10\x94\xec\xd1\xc8\x77\xf0\x2e\xc2\x2d\x4e\x14\x6e\x51\xe9\x7f\x97\xbd\x2a\xac\xcb\x2a\xe9\xd7\x85\x3f\xbe\xab\x4e\x0d\x67\xb5\x0e\xf0\xfe\x5e\xb9\xd8\xa1\x93\x53\x7f\x97\xc9\xd5\x6e\x55\xc9\xaa\x17\xac\x93\x80\x90\x67\xdd\x9f\x94\x23\x9a\x4b\x47\x71\x77\x2b\x55\xae\x53\xf0\x65\xf3\xfa\x51\x40\x35\x76\x4a\x55\x9f\xf3\xed\x7d\x5d\x43\x4a\x81\x30\x9e\xbb\x9b\xf3\xaf\xf7\x9a\x68\x72\xe8\x36\x74\x8d\x6b\x8a\x63\xba\xc2\xed\x51\x3f\x87\xd7\xbb\x65\x12\x97\x22\x3e\x98\xa0\x1f\xae\x70\x1f\x87\x54\xa2\x23\xb9\xc2\x35\xa9\x8d\x72\xa9\x63\x5f\xf5\x94\xf4\x70\x5d\xb5\x6d\xdc\x5e\xeb\x76\x71\xee\xeb\xb7\x2e\x21\xd4\x3f\xe6\xef\x5b\x12\x61\x4f\xcf\x89\x2f\x9e\xac\xfa\xf1\x84\xe3\x99\xf5\x97\x66\x58\x68\x64\x35\x79\x7c\x39\xf5\x0b\x51\xa0\x5d\x3f\x2c\xe7\x43\xb8\x3c\x10\xfe\xd7\x52\xaf\x28\xf3\x52\xf0\x25\xb8\xa3\xe7\xc1\x1e\x6b\x72\x7c\xf9\xa2\x79\x06\x6d\xf7\xef\xb1\xed\x0d\x86\xbc\xc1\x9b\x27\xac\xa9\xb7\xdc\xbd\x8c\x8e\xb6\x01\xda\xab\xc1\xf7\x60\x52\xe6\x68\x6a\x49\x6a\x23\x66\xe4\xdf\x00\x00\x00\xff\xff\x15\xd5\x80\xb0\xe7\x09\x00\x00")

func installerKubedbCom_kubedbcatalogsYamlBytes() ([]byte, error) {
	return bindataRead(
		_installerKubedbCom_kubedbcatalogsYaml,
		"installer.kubedb.com_kubedbcatalogs.yaml",
	)
}

func installerKubedbCom_kubedbcatalogsYaml() (*asset, error) {
	bytes, err := installerKubedbCom_kubedbcatalogsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "installer.kubedb.com_kubedbcatalogs.yaml", size: 2535, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _installerKubedbCom_kubedboperatorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x6d\x73\xdc\x38\x72\xfe\xee\x5f\x81\x72\x52\x65\x29\xa5\x19\xad\x6f\x53\x57\x89\xee\xea\x52\x3a\xdb\xbb\xe5\x5a\xbf\xa8\x2c\xef\xee\x25\x97\xcb\x16\x86\xec\x99\xc1\x89\x04\xb8\x00\x28\x79\x36\x95\xff\x9e\x42\x03\xe0\x90\x43\x02\x04\x47\xd2\xd9\x7b\x47\x7c\xb1\x35\x24\x1b\x8d\x46\xa3\xd1\x2f\x0f\x09\x5a\xb1\x1f\x40\x2a\x26\xf8\x05\xa1\x15\x83\x4f\x1a\xb8\xf9\x4b\x2d\x6f\xfe\x4d\x2d\x99\x38\xbf\x7d\xbe\x02\x4d\x9f\x3f\xb9\x61\x3c\xbf\x20\x2f\x6a\xa5\x45\xf9\x01\x94\xa8\x65\x06\x2f\x61\xcd\x38\xd3\x4c\xf0\x27\x25\x68\x9a\x53\x4d\x2f\x9e\x10\x92\x49\xa0\xe6\xc7\x8f\xac\x04\xa5\x69\x59\x5d\x10\x5e\x17\xc5\x13\x42\x0a\xba\x82\x42\x99\x7b\x08\xa1\x55\xb5\xbc\xa9\x57\x20\x39\x68\xc0\xae\x38\x2d\xe1\x82\x98\xdf\xf2\xd5\x13\x42\xda\x7f\x8a\x0a\x24\xd5\x42\xaa\x25\xe3\x4a\xd3\xa2\x00\xb9\xb4\x57\x96\x99\x28\x9f\xa8\x0a\x32\x43\x75\x23\x45\x5d\x5d\x90\xc1\x7b\x2c\x41\xd7\x79\x46\x35\x6c\x84\x64\xfe\xef\xc5\xbe\x5b\xf3\x07\xad\x2a\x95\x89\x1c\xf0\x4f\x3b\xf2\xef\xea\x15\xbc\xfc\xe3\x7b\xc7\x07\x5e\x28\x98\xd2\xdf\x0d\x5c\x7c\xc3\x94\xc6\x1b\xaa\xa2\x96\xb4\xe8\x8d\x01\xaf\x29\xc6\x37\x75\x41\xe5\xe1\xd5\x27\x84\x54\x12\x14\xc8\x5b\xf8\x9e\xdf\x70\x71\xc7\xbf\x61\x50\xe4\xea\x82\xac\x69\xa1\x0c\x47\x2a\x13\x15\x5c\x90\x77\x66\x34\x15\xcd\x20\x7f\x42\xc8\x2d\x2d\x58\x8e\x42\xb7\xe3\x11\x15\xf0\xcb\xab\xd7\x3f\x7c\x7d\x9d\x6d\xa1\xa4\xf6\x47\x43\xd9\x74\xa3\x9b\x61\xdb\x79\x68\x34\xa0\xf9\x8d\x90\x1c\x54\x26\x59\x85\x14\xc9\x33\x43\xca\xde\x43\x72\x33\xe7\xa0\x88\xde\x02\xb9\xb5\xbf\x41\x4e\x14\x76\x43\xc4\x9a\xe8\x2d\x53\x44\x02\x8e\x81\x6b\x64\xa9\x45\x96\x98\x5b\x28\x27\x62\xf5\x57\xc8\xf4\x92\x5c\x9b\x71\x4a\x45\xd4\x56\xd4\x45\x4e\x32\xc1\x6f\x41\x6a\x22\x21\x13\x1b\xce\x7e\x69\x28\x2b\xa2\x05\x76\x59\x50\x0d\x4e\xba\xbe\x31\xae\x41\x72\x5a\x18\x21\xd4\x70\x46\x28\xcf\x49\x49\x77\x44\x82\xe9\x83\xd4\xbc\x45\x0d\x6f\x51\x4b\xf2\x56\x48\x20\x8c\xaf\xc5\x05\xd9\x6a\x5d\xa9\x8b\xf3\xf3\x0d\xd3\x5e\xe7\x33\x51\x96\x35\x67\x7a\x77\x9e\x09\xae\x25\x5b\xd5\x66\xda\xce\x73\xb8\x85\xe2\x5c\xb1\xcd\x82\xca\x6c\xcb\x34\x64\xba\x96\x70\x4e\x2b\xb6\x40\xc6\xb9\xc6\x85\x53\xe6\xff\x24\xdd\x02\x51\xcf\x5a\x9c\xea\x9d\x99\x36\xa5\x25\xe3\x9b\xe6\x67\x54\xae\xa0\xdc\x8d\x76\x11\xa6\x08\x75\x8f\x59\xfe\xf7\xe2\x35\x3f\x19\xa9\x7c\x78\x75\xfd\x91\xf8\x4e\x71\x0a\xba\x32\x47\x69\xef\x1f\x53\x7b\xc1\x1b\x41\x31\xbe\x06\x69\x27\x6e\x2d\x45\x89\x14\x81\xe7\x95\x60\x5c\xe3\x1f\x59\xc1\x80\x77\x85\xae\xea\x55\xc9\xb4\x99\xe9\x9f\x6b\x50\xda\xcc\xcf\x92\xbc\xa0\x9c\x0b\x4d\x56\x40\xea\x2a\xa7\x1a\xf2\x25\x79\xcd\xc9\x0b\x5a\x42\xf1\x82\x2a\x78\x74\xb1\x1b\x09\xab\x85\x11\xe9\xb8\xe0\xdb\x06\xcb\xb7\xa1\xe5\x61\x1a\x5a\xa2\xce\x2f\x84\x94\xf4\xd3\x1b\xe0\x1b\xbd\xbd\x20\xbf\xfd\xfa\xe0\x5a\x45\xb5\x51\xc9\x0b\xf2\x3f\x7f\xa6\x8b\x5f\xfe\x72\xf2\xe7\x05\x5d\xfc\xf2\xd5\xe2\xdf\xff\xf2\x2f\x7f\x76\xff\x39\xfd\x8f\x7f\x3e\x78\x66\x90\x49\xff\xb3\x9d\xc0\xe6\x67\x6f\xee\x06\x95\xa6\x6b\x8b\xae\x2b\xc8\x8c\x06\x99\x69\x34\x8f\x91\xb5\x90\x44\x42\xce\x94\x5f\xbd\x09\xe3\xa7\x79\x8e\x56\x9e\x16\x57\x22\xbf\x86\xac\x96\x4c\xef\xae\x44\xc1\xb2\xde\xad\x84\x30\x0d\x65\xef\xc7\xe0\xf8\xf6\x97\xa8\x94\x74\xd7\xed\x76\x8d\xdb\xcb\xee\x90\x58\x67\xb8\xaf\xd7\x38\x2e\xb6\x66\x90\x9f\xe1\x30\x2b\x91\x3f\x53\x68\x37\xf2\xba\x30\x2b\x24\x13\x5c\x69\x49\x19\xd7\xea\x70\xa2\x02\x03\x36\x8d\x8b\x1c\x2e\x03\x1c\xf4\xb8\x78\x89\x7f\xac\x40\xe1\x63\x0d\xe7\x6d\x2e\x64\x5d\x80\x42\xf1\x3b\x26\x97\x03\x44\x63\x0c\xd9\xeb\xb0\x06\x29\x21\x7f\x59\x1b\x41\x5e\x37\xe4\x5f\x6f\xb8\x68\x7e\x7e\xf5\x09\xb2\x5a\x1f\x58\xf4\x20\xef\x1f\x8d\x6a\x58\x42\x20\xc9\x1d\x2b\x0a\xd7\x8d\xb1\xb9\xfe\x82\x61\x18\x8d\xb0\x19\xdf\xa1\x18\xf7\x4d\x6f\xa9\x26\x8a\x6a\xa6\xd6\x3b\x1c\x67\x23\x09\xf8\x64\x8c\x0f\xba\x16\xfb\x09\x23\xab\x9d\xb3\x3b\x66\x8f\x3b\x0b\x92\x5d\xd5\x9a\x30\x8d\xc6\x2a\xdb\x0a\xa1\x80\x50\x2b\x68\xec\xef\x96\x09\xdc\x16\x88\xe0\x40\x84\x24\xa5\xb1\x32\xb8\x15\x41\x90\x62\x8b\x9d\x25\x4a\x60\x4f\x8e\x29\x52\x0a\xa5\xf7\xb2\xf6\xeb\xc7\x90\xbf\x63\x7a\x1b\x19\x3d\x90\x8d\x71\x7e\x40\x69\xa2\xea\xd2\x30\x71\x07\x6c\xb3\xd5\xea\x8c\xb0\x25\x2c\x71\xfa\x81\x66\xdb\x56\x77\x25\x40\x4f\x2f\xf7\x8d\x16\x85\x1b\x4a\x47\x97\xe0\xe7\x9a\x49\x28\x8d\x2d\x27\x27\x8d\xe1\x77\xc6\xf8\xcc\x5f\xef\x69\x49\xb8\x9b\x81\x69\x3a\x23\xa0\xb3\xe5\xe9\x19\xc9\x44\x59\xd5\xda\xc8\xdc\x8c\x69\xb5\x33\x4b\x5c\x52\xb7\xf9\x48\x51\x6f\xe2\x12\x81\xc2\x31\xea\xbd\x03\x9c\x6c\xdc\xa6\x8d\x61\xe1\x1b\xf2\xd4\x0a\xe9\xa9\xdf\xe4\x55\x5d\x06\x29\x32\x2b\x0c\x94\x5f\x49\x75\xb6\x75\xbe\x48\x26\xa4\x04\x55\x09\x8e\x14\xf1\xca\xab\xfd\x58\x7e\x17\x55\x06\x43\xec\x44\x9d\xe2\xe4\x22\xb1\x2d\xdb\x6c\xfd\x1c\x52\x09\xf8\x5b\x57\x27\x86\x16\x2f\x09\x5b\x3f\xdf\x3a\x0b\xef\x92\x13\x28\x2b\xbd\x6b\x69\x5a\x6b\x8e\x35\xc8\xb2\x19\x21\x45\xf7\x39\xd4\xec\xf6\xa0\x2c\xff\xac\xac\x8c\x61\xd6\x4e\xf3\xc8\x57\xe4\x04\x55\x8f\xe9\x67\x0a\x97\xcd\x42\x54\xa7\x4b\x72\xe9\x7d\xf2\x50\x1b\x67\x8a\x8b\xa6\x67\xd7\x85\x61\x54\x89\x08\xd1\xa6\xff\xe0\x3d\x63\x16\xb0\xcd\x1c\xf0\xac\xb7\x2f\x77\x5b\x57\xde\x56\x6b\x14\x14\x90\x69\x63\x87\x41\x96\x67\x84\x2a\x25\x32\x66\xbc\x95\x66\xfe\xa3\x24\xc9\x81\xaa\x59\x31\x87\x07\x94\x3e\x28\x82\x6e\x45\x57\x71\xc7\xee\xef\x0d\xd1\x04\x25\x66\xa5\x75\x87\xda\x36\x18\xa3\x14\x89\x59\xe3\xe6\xf9\x67\xca\x45\x6c\xf1\xd1\x91\x71\xbd\x0f\xb2\x1b\x64\xd3\xb9\xbd\xee\x4a\x02\x61\xb7\xf9\x18\xd7\x91\x32\xae\x9c\xab\x7f\x46\x28\xb9\x81\x9d\x8d\x0a\x4c\xe0\xe1\xfc\x22\xbc\x39\x89\xaa\x04\xbb\xb9\x18\x1b\x70\x03\x3b\x24\xe4\xc2\x88\x84\xe7\xd3\x67\xde\xb6\x1b\x18\x74\x36\x86\x5a\x6f\x13\xc7\xb9\x42\x1e\x51\x12\x68\x49\xa7\xc8\x8f\xd8\xa8\xbc\x60\x80\xee\x7c\xe2\x33\x11\xc7\x6e\xb8\xf9\x29\x38\x6a\x9c\x1f\x9a\x18\xc6\x4e\xec\x33\x65\x27\xc8\xac\x95\x2d\xab\x92\xc7\xa9\x05\x6a\x17\x2e\x15\x1f\x14\xfe\x60\x82\xe8\x86\x3d\x85\x96\xff\x35\x0f\x7b\x25\x87\xed\x9d\xd0\xaf\xf9\x19\x79\xf5\x89\x29\xb3\xe1\xbf\x14\xa0\xde\x09\x8d\x7f\x2e\xc9\xb7\xda\xea\xe0\x9b\x11\x53\xd1\x62\x71\xaa\x60\xed\x38\x8e\x12\xeb\x25\xb7\xfe\xb7\x11\x47\x3b\xd4\x54\x4b\xe3\x60\x8f\x9b\xc4\x7d\x6b\x16\x18\x53\x26\xf8\x13\xd2\x8b\x05\x13\x06\x48\x73\xc0\xd5\x8f\xb5\xb2\x56\x18\x53\x72\xc1\x17\xb8\x5f\x7a\x9e\x3a\x7d\x59\xa9\xa7\xb3\x29\x3b\xf3\xd3\x67\xcf\x77\x9b\x4c\x31\xcc\xda\xb7\xda\x74\xf7\xa6\xd3\x49\xfa\x82\xdc\x33\xb3\xa5\xb7\xe8\x84\x31\xbe\x29\x1a\xb7\xea\x8c\xdc\x6d\x59\xb6\x45\xbf\x3d\x99\xe8\x0a\x6c\xd6\xa4\x92\x60\xf6\x3d\xaa\x8c\x69\x34\xbf\x6c\x40\x1a\x77\x98\x79\x21\xb0\x74\x46\x25\x54\x05\xcd\x20\x27\x39\x3a\x9d\x36\x67\x41\x35\x6c\x58\x46\x4a\x90\x1b\x30\x61\x71\xb6\x4d\xd5\xfe\xe4\x0d\xc5\xb6\xc9\x8b\x25\x1c\x76\x0e\x37\xef\x52\xa7\xb0\xb4\x30\x96\x29\xe9\x3e\xd1\xce\x27\xa6\xb0\x7b\x90\x09\x88\xdf\x9c\x32\x36\x74\x38\x5c\x8a\xf1\x33\xfb\x1a\x18\x17\xcc\xbe\xc6\xec\x6b\x04\xdb\xec\x6b\xf8\x36\xfb\x1a\xb3\xaf\x31\xfb\x1a\xb3\xaf\xf1\x2b\xf2\x35\x12\x89\xda\x7c\xca\x84\xb4\xce\x8f\x36\xcf\x75\x98\xc7\x41\xc7\xc6\x17\xc8\x3a\x29\x9b\x91\x11\x19\x37\xe1\xda\xed\x65\x1f\x31\x45\xc4\x38\x12\x91\x94\x6f\x80\x3c\x5f\x3c\xff\xea\xab\xb8\x66\xad\x85\x2c\xa9\xbe\x30\x5a\xfe\xf5\x6f\x12\x64\xe2\x56\x43\xf0\xce\x71\x7d\x58\xb4\x32\x62\x91\x9b\xac\x6c\xc3\xd9\xda\xf1\x19\x1a\x9b\xec\x50\xe6\xf9\x1e\xf5\x09\x67\xe5\x9a\x14\x75\x27\xf9\xdd\x2b\x25\x04\x07\xe7\xb2\xce\xd2\x18\x77\x4d\x4a\xd0\x84\xea\x4e\x6a\x93\x95\xd0\x14\x90\x6c\x19\xc4\x16\x33\x83\x14\x7d\x6d\x24\x27\x82\xbb\xcc\xb5\xd1\x9d\x65\x22\xc7\xe1\x6a\x47\xbb\x28\x42\x32\xa0\x0a\x8c\x0f\xb1\x82\x86\x6b\x51\x1a\x2e\x19\xd7\xde\x00\x1a\x96\xc1\x4b\x35\x48\xf8\x04\x96\x9b\x25\xc9\x6b\x24\x47\xb9\xab\xd2\x9e\xda\x51\xab\x9d\xd2\x50\x62\x8d\x45\x48\xfc\xc7\x0c\x5f\xcb\x1d\xd1\xe1\x8c\x2e\xdc\x02\xd7\x35\x2d\x8a\x1d\x81\x5b\x96\xe9\x46\x7e\x58\x48\x66\xda\xd6\xc3\x42\xab\x25\xc5\x61\x3d\x5c\x8d\x51\x3b\x7d\xe0\xbe\x59\x55\x5c\x06\x23\x15\x6d\xe8\x61\xf9\x27\xbe\x48\xcd\x6d\xa8\x39\xef\x3f\x84\x33\xff\x24\x6d\x23\x39\x8c\x49\xea\xa2\x30\xf2\xb6\x85\x80\x3e\x7b\x3e\xd9\x3e\x6a\xb3\x7c\x2a\xde\x56\xb3\x3a\x1a\x67\xeb\x47\xb6\x92\x71\xf9\xee\xa5\x91\xc8\xd8\x90\x09\xf9\x28\x2a\x51\x88\xcd\xae\x2d\x7b\x5c\xfd\x58\x60\x70\x94\x29\x51\xf5\xca\x79\xb6\xe3\x8e\xdb\xbb\x83\xa9\x9c\x73\xe6\x73\x1c\x3b\xd4\xe6\x38\xb6\xd7\xe6\x38\x36\x91\xc5\x39\x8e\xc5\x36\xc7\xb1\x73\x1c\x3b\xda\xe6\x38\x76\xe0\xe6\x39\x67\x3e\xfb\x1a\x91\x36\xfb\x1a\xbd\x36\xfb\x1a\xb3\xaf\x31\xfb\x1a\xb3\xaf\x11\x6d\xb3\xaf\x31\x70\xf3\x83\xe5\xcc\xc7\xc9\x8d\x89\x67\xd1\x4f\xb4\x45\x33\xc0\x41\x96\xa2\x97\x2b\x91\x1f\x01\xa9\xaf\x44\x1e\x41\xd4\xdb\xa4\x66\x26\x16\x85\xc8\xa8\x1e\xb6\x07\x98\x4e\x35\x64\x5c\x26\x5f\xd1\xd2\xe6\x6a\xcf\xc8\x2f\x82\x83\x45\x3a\x9b\x65\x86\x99\x55\xa1\xb7\x20\xcd\xed\x27\xea\x74\x10\xa9\x3a\xa3\xf4\x07\xdb\x8c\xd2\x9f\x51\xfa\xae\xb5\x51\xfa\x5b\xaa\xac\x5e\xda\x8d\x30\x0c\xda\x6f\x59\x07\x63\x80\x7e\x17\xe5\xf7\x33\x61\xf6\x8d\x12\x3a\x65\xc1\x57\x19\xf7\x13\x6f\xc7\x95\xbb\x72\x24\xe4\x57\xdd\xd1\x44\xac\xb7\x8d\xe1\x90\x69\x9a\xe7\x90\x93\x0a\xe4\xc2\xaa\x9e\x20\x6b\xc6\xf3\x81\xb1\xf8\xf1\x07\xc9\x26\xe2\xe8\xbb\x4c\x4e\x28\x5d\xb4\xab\x2b\x1d\x03\x7d\x88\xaa\x1f\xd9\x0b\x9b\xf9\x7b\x4c\x54\x3d\x46\x5e\x7e\x73\x9b\x1e\xb2\x63\xdc\xf6\x73\x0d\x72\x47\xc4\x2d\xc8\x7d\x64\xd2\xbc\xe7\x99\x12\x84\xe0\xde\xc3\x14\xc9\xa8\xb2\x86\x7a\xdc\xd5\x9a\x16\x9d\x4e\xaf\x83\xf4\x06\x7b\x48\xc2\x46\xf9\x3e\x67\x81\x82\x48\xf4\xde\x06\x53\x1b\x03\xc5\x29\x2a\x53\x5d\x78\x5b\xba\x4a\xba\x79\x92\x73\x3a\x38\xdb\x81\x94\x47\x7a\x58\xd0\x2a\xe3\x8d\xa4\x3d\xd2\x69\x1e\xa4\x47\xee\x99\xfa\x20\x47\xa4\x3f\xc8\xb4\x14\x08\x39\x14\xaf\xe1\xd2\xed\xd3\xfd\x6c\xc8\x04\xa2\x2d\xfd\x9a\x9e\x11\x21\xc7\xc5\x23\xd3\x33\x23\xe4\x70\xf8\xcd\xf4\xc9\x5e\x9a\x64\xd2\xe0\xdb\x29\x95\x70\xaa\x64\x12\xc9\x5e\x5a\xa5\x9b\x2e\x41\xdd\xea\x64\x4c\x1e\x5b\xd8\xd3\xb2\x25\xe4\x50\xd4\x2e\x57\xc0\x30\x74\x3e\xc8\x9d\x4c\x12\x4c\x37\xcf\x12\xcc\x9f\x4c\xa2\x19\x4a\x66\x74\x73\x28\x93\x49\xf6\xf3\x2d\xbd\x3c\xca\xc3\xb0\xe9\x58\xdc\x27\x22\x26\x91\xb5\x1f\x88\x78\xc8\x64\x04\x99\x9e\x90\x20\xc7\xea\xe5\xd4\xc4\x04\x99\x98\x9c\x20\x13\x12\x14\x64\x6a\x92\x82\x4c\x4d\x54\x90\xc9\xe3\x45\x17\xe2\x4d\xeb\x43\x2f\x63\xad\xf5\x75\x81\xc9\xbb\xd1\xe4\x19\xec\x7b\x3b\x96\x55\xeb\xe8\x94\xb4\x32\x56\xe2\x7f\xcd\xd6\x8c\x8a\xff\x7f\xa9\xfb\x28\x65\x52\x19\x57\xd8\x25\xff\x5a\x14\x7c\xce\xa1\xd5\x59\x22\x51\xc3\x0d\x53\xc4\xe8\xce\x2d\x2d\x8c\x03\x62\x61\x5b\x2e\x54\x33\x9c\x1e\xfa\x6b\xa9\xeb\xfb\x6e\x6b\xc2\x73\xb3\xf9\xda\x30\x8f\x29\xf2\xf4\x06\x76\x4f\xcf\x7a\x76\xe4\xe9\x6b\xfe\x34\x95\x2a\x75\xa1\x4a\xc7\x66\x34\x9e\x8f\xe0\xc5\x8e\x3c\xc5\x6b\x4f\x53\x17\xf6\x90\xbb\x38\xc5\x11\x3c\x22\x29\x97\x74\x33\xf7\x1f\xdf\x99\x5a\x00\xdc\x3f\xd8\xe4\x57\x7c\x60\xbc\xbf\x94\x92\x6d\xf4\x1e\xd4\x75\xdf\x0f\x22\x27\xcd\x6b\xe3\x1b\x23\x79\x7d\x1a\x0e\xa5\x5b\x43\xea\x20\xd1\xd0\xe5\x2f\x81\x72\x45\x9e\xfa\xec\xd9\x33\xb5\xe7\xf1\xe9\xc3\x55\x1c\x27\xad\xe1\x74\x5b\xa4\x1d\x80\xed\xbb\x14\x77\xf5\x20\xc6\x77\xd9\x42\xf7\x55\xa2\x15\xec\xd3\x8b\x39\x39\xf1\x91\x6e\x38\xf6\xde\x37\x21\x11\x45\xd9\x79\x9c\x6b\xb6\x68\x68\xec\xe3\x5f\x13\x11\xa6\x9a\x57\x0f\x6b\xee\x6a\x80\x4f\x6e\x36\x79\xbb\xbd\x46\xa5\xac\xe0\xbb\x2d\xc8\xce\x48\x99\x72\x5f\x7b\xc2\x0a\x84\xac\x39\x37\xfd\x0a\xee\xd2\x7a\x49\x24\x8d\x99\xb1\x1f\x2d\x72\x69\x12\xeb\xf6\xe3\xa8\xd1\xf7\xdf\xcf\x52\x22\xd4\x91\xf8\x04\x26\x7e\x49\xca\x61\x26\x05\x77\x8b\xc8\xfc\xe2\x33\x71\x28\x17\xc8\x53\x25\xcb\x9a\x31\x2e\xc9\x2b\x5c\x04\x6d\xe6\x98\xc2\x99\xa4\x45\x21\xee\x52\xac\x4f\xb2\x56\xa7\xf9\x06\x8b\x36\x33\x0f\x51\x32\x98\x0c\xb3\xbf\x7b\x60\x98\xfd\x41\xea\xe9\x57\x82\xb2\x4f\x4c\xea\xcd\x50\xfb\x19\x6a\xdf\x82\xda\xe3\x43\xd6\xf2\x8d\x63\xee\xc3\x3a\x83\x58\xfc\x54\xcc\x3d\xf9\x71\x0b\xb8\xa2\x22\x09\x36\x33\x45\x65\x5d\x68\x56\xed\x0b\xd6\xca\xb2\x56\xd8\xf0\xd1\x02\x95\xd4\x41\x76\x36\xf6\x46\x00\xcd\xb6\x87\xcb\x04\xfb\xc1\x82\xb6\x42\x8b\xec\xca\x2c\xb4\x28\x1c\xb6\xde\xc4\x95\xe1\x39\x02\x57\xab\x62\x0f\x93\xc2\x7f\xe9\xbe\x60\xd8\x24\x4d\xb0\x38\x71\x62\x36\xcb\xc2\xa8\x83\xd9\xb2\xbc\x55\x8b\xd5\x5c\x7b\xfb\xaf\xcd\xca\xdc\x82\x2f\x90\x6c\xd8\x2d\xf0\xfd\x26\x7c\xa2\x4e\x4f\xc7\x60\x4d\x3a\xd1\xf5\xe8\x3b\x16\x11\xa2\x43\x2e\xc7\x59\xe2\x76\x1f\x21\xdb\x38\x02\x09\xdb\xfc\xef\x5b\xbb\xd7\x1f\x22\x34\xf7\xc5\xa1\xe0\x06\x8f\xe2\x69\xb6\xf8\x66\x02\x23\x44\xd9\xf8\x68\xd2\xf2\xa0\x13\xca\x08\x47\x94\x10\x08\x0b\x9b\x13\xdb\xa6\x94\x0f\xfe\x66\xaf\x4f\x24\x94\x0c\xa6\xc0\xdc\xc6\xcb\x05\xa9\xf1\xdf\xb1\x90\xc7\x68\x01\x60\xc6\x3c\x46\x5b\x7a\xb2\xff\xef\x0f\xfa\x18\x49\xee\x7f\xa1\x18\xc8\xa3\x93\xfa\x7f\x4b\xe8\x63\x2c\x91\x3f\xb1\xda\x45\xc6\x92\xf8\xf7\x04\x00\x8e\x81\x20\x93\x69\x06\x92\xf7\xc3\x09\xf9\x64\xaa\x43\x89\xfb\xc1\x64\x7c\x32\xc5\x19\x41\x38\x7a\xdf\xe7\x46\x10\x4e\x4c\xc8\x1f\x9b\x8c\x9f\x34\x3b\x53\x93\xf0\x2e\xbd\x9e\xc0\x46\x62\x02\xbe\x9f\x5a\x4f\x19\xe2\x68\xf2\xfd\x30\xad\x9e\x96\x74\x8a\x25\xde\x07\x53\xea\x09\x64\x87\x93\xee\xf7\x72\xa7\x92\xb5\x33\xf1\xc6\xd4\x14\xfa\xf4\xf4\x79\x02\x96\x60\x42\xea\xdc\x27\xc6\x47\x28\x3e\x44\xda\x3c\xc9\x22\x26\xaf\xb4\x34\x0b\x91\x9c\x26\x7f\x8c\x14\xf9\xc4\xf4\x78\x4a\x58\x4e\x06\x43\xf3\x58\x6a\xdc\x46\xc2\x23\x24\xd3\xd3\xe2\xed\x68\x78\x6c\xf8\xa9\x29\xf1\x76\x3c\x3c\x56\x99\x4a\x4a\x87\xf7\x93\xdd\xe9\xd5\x94\x49\xa9\xf0\x24\x6d\x4d\xc9\xbc\xa6\xa4\xbf\xef\x9d\x54\x1d\x05\xaf\x73\xcd\x8e\x05\xb0\xb7\xf5\x3a\x80\x62\x1f\xe4\x99\xde\x0a\x96\x93\xaa\xd6\x0e\xca\x3b\x19\xc9\x3e\x48\xf5\x1f\x0a\xdd\xde\x11\x7d\x14\xe2\x1e\x4f\x69\x9f\x1d\x01\x71\x0f\x52\x74\xcb\xf2\x08\x88\x7b\x98\xa4\x83\xbe\x1f\x05\x71\x0f\x52\x45\xe8\xfb\x71\x10\xf7\xd1\x15\x7f\xa8\x42\xe1\xb9\xf2\x38\xf7\x20\xc9\x71\xfc\x7b\x04\xe7\x1e\xce\x90\x47\xf1\xef\x11\x9c\x7b\x58\x9c\xc9\xf8\xf7\x1e\xce\x3d\xa2\xf2\x33\xfe\xfd\xa0\xcd\xf8\xf7\x56\x9b\xf1\xef\x89\x83\x9d\xf1\xef\x33\xfe\x7d\xac\xcd\xf8\xf7\x19\xff\x3e\xe3\xdf\x67\xfc\xfb\x8c\x7f\x9f\xf1\xef\x03\x6d\xc6\xbf\xcf\xf8\xf7\x19\xff\xde\x6a\x33\xfe\x7d\x64\x28\x33\xfe\x7d\xc6\xbf\xcf\xf8\xf7\x19\xff\x3e\xe3\xdf\x07\x6e\xf9\x2c\xf8\xf7\x4e\x12\x3a\x08\x82\x8f\xa4\x63\xf7\xdf\x4f\x99\x08\x82\x0f\xd2\x5c\xc1\x38\x08\x3e\xc8\x76\x90\x6a\xe0\x1b\x3f\x49\x48\xf8\x70\xea\xb5\x8d\x90\x9f\x84\x84\x8f\x24\xcd\x07\xbe\x4a\x7f\xcf\xaf\xcf\x93\x16\x42\xfe\x58\x24\x7c\x58\x05\xc4\x8c\x84\x9f\x91\xf0\x33\x12\x7e\x46\xc2\xcf\x48\x78\xdb\x66\x24\xfc\x8c\x84\x9f\x91\xf0\x33\x12\x7e\x46\xc2\xf7\xda\x8c\x84\x1f\x64\x77\x46\xc2\xcf\x48\xf8\x19\x09\xbf\x6f\x33\x12\xbe\xdb\x66\x24\xfc\x8c\x84\x8f\xb4\x19\x09\xff\x38\x48\xf8\xe0\x25\xca\xb9\xd0\xd6\xb9\x3f\xe4\x3f\x6d\x33\x8d\x88\x28\xdc\x69\xc5\x14\xc8\x5b\xe8\x45\x2a\xb1\xd8\x6e\xb5\xab\xa8\x52\x18\x41\x20\x42\xf8\x47\x58\x6d\x85\xb8\xf9\x93\xa4\x83\x4b\xdf\x76\xbe\x12\xa2\x00\xda\xcf\x4c\x64\x34\xfc\x4c\x60\xba\x81\xd3\x55\x01\x6f\x6b\xdd\xee\x7d\x7a\xcf\x96\x4c\x6f\x18\xd3\x09\x6d\xa4\xa8\xab\x2b\xc9\x84\x64\x7a\xf7\x96\x71\x56\xd6\x83\x58\xd8\xb1\x92\x43\xbc\xd0\xb0\x05\x5a\xe8\x6d\xb6\x85\x6c\x90\xc5\xb1\x60\xdc\x8e\x36\xb8\x34\xe2\x23\x1c\x7d\xb7\x43\x0e\xd6\x82\xee\x37\x60\xa3\x98\x8c\x6f\x5e\x80\xd4\x83\x63\x1a\x1b\x71\x46\x5f\x0c\xb3\x45\x52\xec\xc9\x06\xb8\xf1\xa1\xe0\x58\x81\x59\xfe\x41\xde\x87\x07\x4b\x21\xb2\xa3\x8e\x50\x88\xd9\xc3\x45\x33\xc2\xa9\xb3\x5d\x2b\xf8\xae\x5e\x41\x63\x3a\xbe\xf9\x39\xe7\xdf\x08\x79\x79\x33\x38\x0f\x71\x39\xdd\x82\x34\x2e\xaf\x5f\x3c\x0f\xad\x44\x21\x01\x2c\x48\x76\x18\x9c\x2e\x86\xcd\x4a\xe0\xae\x9e\xd5\xe8\xdd\x37\x64\x14\x7a\x37\xb5\xd6\x74\xef\x9a\x59\x54\xbd\x1f\xdb\x6b\xa2\x77\x31\x3c\x31\xbd\x5b\x0f\xe4\x9e\xba\x5d\x64\x66\x12\xa7\x6d\x16\x12\x36\x4c\x69\x19\xd9\x19\x02\xea\x2b\xa1\x12\x8a\x69\x71\xc4\xa3\x9a\x6e\x26\x3e\x13\x56\x14\xcf\xff\xc0\x05\xcf\x5f\xef\x92\xa6\xc9\x1b\x70\x26\x99\x66\x19\x2d\x2e\xf3\xbc\x5f\x65\x0d\xaf\x1d\xab\x85\x97\x9c\x16\x3b\xcd\xb2\x9e\xd8\x63\x0f\xe2\xc1\x48\x4c\xf5\x0c\x5b\x6c\x12\x23\x9b\x47\x7c\x7d\x87\x76\x86\xb8\xe5\xff\x1c\x3a\xd3\xd4\x4f\x46\xdf\x4c\x7c\xe1\x5e\x95\xfa\xe0\x9f\x68\xb4\xc7\x96\x81\x81\x28\x96\x43\x46\xa5\x8f\x9e\x41\x1e\xf3\xa2\x60\xc1\x4a\x36\xbc\xf1\x91\xc9\xf9\x95\x04\xef\xb9\x33\xc4\x67\x6f\xb0\x73\xf7\xe3\xca\x55\x08\x4a\xfa\xc9\x58\x31\x42\x4b\x51\xdb\xcc\x85\x7b\x6b\x2c\xe2\x90\x7b\x11\x79\x27\x9f\xbc\x15\x58\xc3\x5d\x8b\x0b\xb2\xd5\xba\x52\x17\xe7\xe7\x37\xf5\x0a\x24\x07\x0d\x6a\xc9\xc4\x79\x2e\x32\x75\x9e\x09\x9e\x41\xa5\xf1\x3f\x6b\xb6\xa9\x25\x3a\xc6\xe7\x25\xe5\x74\x03\x0b\xd7\xed\xa2\x21\xbf\x68\x24\x7d\xfe\x2c\xba\x55\x46\x7c\x7a\xf7\xd6\xdd\xe7\x92\xf8\x07\xd7\xfd\xa1\xcc\xed\xce\x71\x94\xcc\x65\xf3\x92\xd5\xeb\x35\x69\xe8\x33\x45\x44\xc9\xb4\x89\xe0\xd6\x42\x12\xba\xd7\xd2\x70\xf6\x9f\x69\x13\xaa\xd2\xba\xd0\x98\xdc\x70\xda\x81\xaf\xef\xd9\x37\x2d\xe1\x53\x55\xb0\x8c\xe9\x62\xb7\x8f\x8d\xcf\xec\x0b\xb4\x77\x4c\x85\x99\xb5\xb9\xb0\xe6\x4c\x74\x9c\xe5\x85\x8f\x8a\x31\xf4\xfd\x62\x35\x26\x7a\x59\x41\x56\x9b\x8d\xf5\x85\xe0\x1a\x3e\x0d\x9a\xc0\xce\xf4\x5f\xbb\xfb\x89\xc0\x1f\x54\x83\xc5\x70\x79\x11\x59\x73\x0c\xec\x8f\x31\x24\xb8\xf4\xae\x24\xbb\x65\x05\x6c\xe0\x95\xca\xa8\xad\x21\x25\x61\x7a\x9e\x5d\x06\x9e\x46\xb5\x91\xa2\x50\xe4\x6e\x0b\x78\xe4\x17\x35\x9c\x64\xa0\xc2\xe5\x89\x8c\x72\xb2\xa1\x8c\xdb\xd3\xab\x2a\x4f\x14\xd3\x12\x1c\x11\x27\x15\x95\xc0\xb5\x27\xe4\x0a\x0c\x66\x73\x09\xd2\xcc\x99\x84\xcc\xe8\x5d\xc3\x4f\xf3\x56\xe9\x4f\x1c\xee\x7e\x32\xbd\x28\xb2\x2e\xe8\xc6\xa2\x84\x56\xae\xda\x1f\xae\x91\x5b\x28\x9a\xd3\x8e\x3d\x2b\x41\x41\x30\x45\xb4\xac\x81\xd0\xe2\x8e\xee\xc2\x83\xbf\x73\x70\x99\x16\x6d\xa6\x2e\xc8\xf3\x53\x9c\x5c\xaa\x48\x43\x3b\x27\xbf\x39\xc5\xf7\x61\x5f\x5c\x5e\xfd\x74\xfd\x9f\xd7\x3f\x5d\xbe\x7c\xfb\xfa\x5d\x5c\x4d\x63\x81\x48\x46\x2b\xba\x62\x05\x8b\x59\xac\xde\x6b\xaa\xed\x87\x70\x99\xe6\xf9\x79\x2e\x45\x65\xc7\xe1\xb3\x55\xcd\x58\x22\x59\xf5\x97\x2d\xcb\x61\xc6\xef\x2c\x89\xaf\x48\x76\x3a\xda\x48\xca\x75\xb3\x91\x86\x15\xa9\x11\xa1\xac\xb9\x66\x65\x10\xa5\x94\x52\xa2\xa6\x79\xb4\x54\xd3\xad\xea\xe3\x1b\xb6\x6d\x96\x23\x4f\x26\xa4\x61\xbb\x8e\x85\x27\xbb\xdb\x97\x7e\xc9\xd5\xfb\xeb\xd7\x7f\x3a\x98\x8d\x5d\x15\xcf\x08\x26\xa6\x76\x53\x12\xbb\x66\xca\x93\xa5\xf3\x01\x4a\x71\xfb\x8f\x24\x9f\x51\xa7\xa2\xb1\x71\x41\x15\xeb\x0a\xb0\xe6\x6d\xf3\xc0\x5b\xcf\x93\x12\xc1\x8a\x57\xd6\x1c\x81\x8a\xa1\x7c\x5a\x4f\xed\x17\x28\xd6\x68\xcc\xa3\x5c\x33\x0b\xfa\xeb\xbc\x19\x22\x85\x18\xb5\x8a\x5b\xa1\xf4\xb2\xb3\x9e\xd7\xb4\x50\xc1\xc5\x37\x6e\x99\x8c\x71\x7d\x6b\x1c\x9b\x24\xe9\x34\x77\x93\x1c\xb8\xf0\xb8\x15\xd3\x0b\x82\xb7\xa4\xc8\x88\xf5\x92\xb4\x30\xb1\x70\x70\x28\x6b\x84\xc8\x40\xdb\x78\xa1\xc9\xf3\x86\x89\x29\x3f\xc6\xab\xa6\xc7\xf8\x57\x09\x6a\xd5\x7c\x92\xe0\xc0\x30\xed\xfd\xa6\x35\x22\x3a\x68\x8e\x95\xb4\x8a\xea\xad\x8a\xbe\xbe\x5b\x52\x75\x03\xb9\xbd\xd1\xed\x83\xce\x9f\xb3\x3d\x35\xac\x7d\x34\xe3\x5f\x03\xd5\xb5\x04\xdc\xe7\x62\xce\xd6\x0a\x7c\x2c\x17\x9f\xb4\xc8\xda\x30\x63\x78\xcf\x8b\xdd\x07\x21\xf4\x37\xac\x00\x8b\x3d\x4d\x9a\xc0\x1f\x9d\xa7\x60\xf1\x67\x8d\xa8\xcc\x56\x47\x91\xee\x02\x85\x83\xaa\xb8\x6e\x48\x8f\xee\x2c\x66\xc2\xee\xa9\x88\xb2\xe6\x97\xea\x5b\x29\xea\xa0\xb1\xeb\x6d\x90\xdf\xbe\x7e\x89\xeb\xa6\xb6\xbb\x3a\x70\x2d\x77\x16\xe0\xeb\x0a\x25\xcd\x00\x23\xeb\xd4\xf9\x16\xdf\x1b\xfd\x39\xd0\x18\xe3\xc7\xd4\x5c\x81\x5e\x92\xb7\x74\x47\x68\xa1\x84\x77\x5e\x22\x4b\xff\x4a\xe4\xd7\x5d\xdf\x73\x89\x68\x15\xfb\x18\x59\x09\xbd\x25\x07\x37\x60\x6d\xb8\xff\x5c\x38\x1a\x68\xea\xc8\xad\x3a\x18\xe3\x3d\xb2\x9a\xde\x80\x22\x95\x84\x0c\x72\xe0\x59\x70\x76\x5a\x09\xbe\xdf\xfe\x6b\x74\x06\x63\x18\x7c\x9c\xc1\x77\x82\x1b\xb5\x4c\x43\xab\xf3\x9c\x65\x0e\xfd\xe6\xa0\x64\x7b\x95\x44\x18\x8e\xf3\xcb\x28\x82\x71\x8c\x52\xc6\xd6\xbf\xb4\x40\x1d\x59\x3b\x58\xfa\x77\xf5\x0a\x0a\xd0\xd6\xe9\xbc\xb5\xe9\x42\xfb\x59\x11\x56\xd2\x0d\x10\xaa\xfd\x84\xc7\xd6\x2b\x70\x55\x4b\xff\x41\x1b\x4d\x72\x01\xb6\x66\xe6\x58\xfb\xfe\xf5\x4b\xf2\x15\x39\x31\xbc\x9d\xe2\x34\xae\x29\x2b\x62\xdf\x17\x57\x9a\xca\xc3\xb1\xb2\xb5\x27\x8d\x43\x40\x9d\x23\x42\xda\x25\x75\x46\xb8\x20\xaa\x8e\xd8\x3e\x37\x36\xe3\x09\x7b\x07\xbb\x02\x69\x26\x15\xc3\xfd\x9e\xea\x86\x54\x34\xcc\xf3\x74\xd5\xdd\xab\x68\x98\xea\x43\xa8\x6e\xa2\x61\xf9\x5e\xf5\x73\xa6\xbe\xf5\xec\xca\xf7\x0f\x68\x57\xda\x5b\xb5\xd1\xd1\xee\xa8\xad\x22\x96\xa0\x69\x4e\x35\x75\xf6\xc6\xdf\x10\xb6\xba\xe9\x53\x1a\x9b\xba\xb0\x3b\x3e\x36\xa5\xd1\xa9\x0b\x2f\xa6\xbf\xa9\x35\x52\xf0\x86\xf1\xfa\xd3\xfb\x6a\xb0\x9e\xeb\x5b\x6f\xee\xaf\x5f\xe1\x63\x38\xc5\xa8\x87\x28\x64\x0b\x2b\xc9\x7d\xfc\x14\xcd\x2a\xda\xf6\xba\x33\x95\x67\x01\xd7\x04\x97\x2b\x2d\x2c\x1e\xc1\xec\xc0\x94\xe7\x22\xfc\x16\xc9\x21\x73\xcd\x47\xb0\xf6\x0c\x4d\x51\x8e\x5f\xe3\x7a\x4f\x09\x27\x0b\xb8\x85\x22\x39\x64\x7a\x63\xee\x36\x0e\x8c\x97\x2e\x3e\xee\xa0\x1c\x68\xf6\xf7\xa0\xa2\x78\x4c\x93\xa6\x19\xa9\x48\x0a\x51\x04\x2b\x9f\xbd\x31\x7c\x10\x05\x58\xc8\x9d\x1f\x84\x79\xfc\xb3\x8f\x01\x6f\x4a\x1d\x03\x7a\xd1\x9d\x31\x60\x5c\xf1\xb9\xc7\x50\x47\x76\x8e\xde\x18\xcc\x36\xd3\x1d\x03\xda\xfc\xcf\x3b\x86\xd1\x10\xf9\x8e\xf1\x5c\xdc\xa9\xa9\xa6\xf2\x47\xfb\x98\x5f\xd7\x99\x31\x1b\x9a\xf1\x8d\x6a\x9b\x4b\x5a\x14\x49\x29\xaa\x21\x7b\xe9\x33\xb1\xf8\x36\x1c\x46\x5c\x3d\xbb\x83\x16\x34\x48\x74\x05\x46\xfe\x36\xfb\xfe\x05\xbb\xdf\x29\x36\x6d\x53\x2a\xfa\x42\x1a\x3a\x9a\xd1\xe2\xba\x82\x2c\x59\x29\xbf\x7d\x7b\x7d\xd9\x7d\xd4\xa8\xa8\x7d\x69\xcc\x8c\xc4\x5c\x27\x34\x2f\x19\x62\x5f\xa3\x6a\x79\x67\x4b\xed\xe4\xc4\x97\x01\x36\x4c\x6f\xeb\xd5\x32\x13\x65\xab\x22\xb0\x50\x6c\xa3\xce\x9d\x56\x2d\x0c\xe7\x71\xf4\x20\xe3\x05\xbe\xc6\xe7\x95\x7e\xff\x79\x43\xc7\x5c\xd6\x70\x8f\x02\x47\xc8\x5f\x1c\x92\xeb\xca\x80\xfd\xa1\xbf\xa3\x25\x58\x48\xaf\x0b\xe9\x9b\xef\x67\xd0\xa2\xda\xd2\x05\x1a\xff\x28\x69\xa3\x2c\xcc\xc1\x71\xb7\x02\x5f\xd5\x35\xdd\xd9\x82\xbf\x0b\x65\x6c\x84\x8f\x2c\xb8\x55\x62\x38\x89\x92\x6d\xe7\x0f\xee\x6d\xb4\xfa\xda\x62\xc6\x7d\x0f\x8d\x41\xb1\xb9\x37\x80\x8c\xf4\xdb\xd3\x13\x1d\xd6\xe1\xd4\x59\x37\x38\x22\xfb\xd1\x2f\xaa\x7d\xe9\xb2\x6f\xe2\x8d\x49\x22\xc7\xb8\xc3\x3d\x64\x6c\x89\x37\xae\x83\x71\x48\x74\x30\x87\x31\xca\x70\x2c\x62\x6e\xe9\xc6\x23\x23\x4b\x74\x24\x56\x49\x74\x3b\xa3\x9d\x3c\xb0\x95\x26\x0f\x6f\xa9\x6d\x8b\xea\xae\x89\xe4\xc3\x2a\x3a\xc2\xec\xa0\xfa\x7e\x68\x2b\xd4\x43\xea\xea\x7d\xca\xab\x0f\x09\xf0\x19\xc4\x59\x3d\x32\xea\x67\x5d\x17\x85\x31\x64\xef\x6f\x41\x4a\x96\xf7\x16\x6a\x70\x20\xb8\x0c\xae\xea\xa2\xb8\x12\x05\xcb\x7a\x28\x97\xf1\xe7\xae\x21\x93\xd0\x07\x39\x04\x4a\x31\xa3\x90\xe2\x7e\x71\xa4\x10\x9b\x37\x43\x11\x50\x0c\xc6\x17\x8e\xa7\x4b\xc1\x8d\xb0\x19\xef\xcd\x77\xcc\x6d\xa1\x1b\x18\x2e\x29\x24\xc0\x8a\x8f\x01\x39\x49\x51\x82\xde\x42\x7d\x14\x4c\xb5\x79\x83\xe0\x48\x90\xe7\x08\x0e\x41\xde\xb2\x0c\xde\x5a\x31\x1e\xc3\x5e\x11\x7d\xe1\xeb\xc1\x41\x31\xc7\x5b\x85\xf0\x0a\x47\x75\xe8\xaf\xfb\x66\xd6\x7a\x97\xba\x52\x4b\x5d\xd4\x47\x2d\x68\x2e\x72\x08\x7d\x09\xe0\xf1\x60\xff\xa1\xf7\x93\xbf\x34\x20\xe7\x0c\xca\x9b\x41\x79\x33\x28\x2f\xd0\x66\x50\xde\xc4\xcb\x33\x28\x6f\xa8\xcd\xa0\xbc\x19\x94\x37\x83\xf2\x82\xc4\xbf\x3c\xd0\xd9\x0c\xca\x9b\x41\x79\x7e\xa8\x33\x28\x6f\x06\xe5\x91\x19\x94\x37\x83\xf2\x66\x50\x5e\xaf\xcd\xa0\xbc\x19\x94\x37\x83\xf2\x66\x50\x5e\xd3\x66\x50\xde\xdf\xe5\x7a\x9f\x41\x79\x43\x63\x98\x41\x79\x8f\x36\x86\x19\x94\x37\x34\xd0\x19\x94\x37\x83\xf2\x66\x50\xde\x17\x0b\x0c\x9b\x41\x79\x33\x28\x6f\x06\xe5\xcd\xa0\xbc\x5f\x07\x28\xef\x91\xf1\x77\x95\xc8\x2f\x3f\xc7\xe7\x36\xab\x9e\xae\x1e\xd2\xed\xd6\x79\x07\xbc\x9b\xad\x28\x72\xfc\x64\xae\xd3\x2f\x5f\xd4\x26\x54\x6b\xc9\x56\xb5\x1e\xa8\xee\x18\x1d\xcc\x44\x59\x8a\x76\x21\xc3\xbb\x66\x4b\x62\xbd\x3c\x5a\x5c\x74\xcc\x81\xfb\xc4\x3b\xb9\x06\x18\x2e\xde\xb4\x58\xc5\xb0\xd3\xa7\x48\xdd\x27\xa6\xc5\xda\x06\xa2\x76\x67\x3d\x2c\x94\xc6\x1c\x9c\x75\x38\xe9\xdb\x11\xcf\xd3\x4b\xbb\x84\xcd\x46\x52\x57\x1e\xad\x50\xd8\x4f\xef\x1d\xfa\xd6\x07\x6e\xe7\xa0\xfe\x33\x6e\x8f\x85\x59\x92\x6b\x51\x02\xb9\x15\x45\x5d\xda\xc1\x3b\xac\x4c\x27\x89\xa8\x05\xc9\xb6\x78\xdc\x18\x7a\xa6\x77\x86\x6c\xe8\x58\x06\xe1\x50\x19\x9e\x24\xda\x44\xf3\x48\x03\x4f\xaa\x44\x7e\x41\xfe\x9b\x93\xe7\xb6\xec\x21\xee\xb0\x94\xfb\xed\xeb\x97\x61\x7f\x76\x65\x7b\xfe\xe6\x1a\xc5\x45\x7e\x63\x9f\x54\xa0\x37\x2c\x27\x2b\x6b\x74\x8c\xf5\x3c\xe1\x70\x67\x53\xf7\x66\xef\xb5\xdf\x1e\x1e\x76\xea\xd0\x36\x5a\x16\x7d\xda\xb0\x61\xd2\x75\x73\x4a\xbe\xb6\xfd\x54\x20\x9d\x7f\x68\xfa\x0a\x1f\xc5\xfe\xfe\xc3\x33\x77\xd8\x9b\xbc\x5b\xc8\xbb\xc5\x62\xb1\x30\xe3\xf4\x49\xcd\x81\xc4\x2c\x9e\xf8\x25\x72\xb6\x0e\xd7\x9b\x1b\x69\xa3\x6e\xef\x59\x51\xfe\xb0\x1f\x3b\x8a\xe5\xd0\x77\xb2\xc7\x92\x49\x23\xdf\xa6\x8b\x16\x25\xee\x51\x90\x68\x36\xe5\xc1\x01\x4f\x2e\x46\x1c\x6e\x68\xe1\xf4\xce\xc3\xa7\x76\x26\x6d\xac\xae\xf2\xb7\x3f\xde\x66\x38\xba\x7d\x80\x59\x8b\x14\x22\x1e\xa5\x08\xf1\xf0\x05\x88\xe3\x8b\x0f\xb6\xc8\x10\x5c\xf4\x53\x0b\x0f\xad\x02\x43\x20\x7f\x90\x52\x74\x08\xa5\xa7\x43\xc6\xf9\x38\x15\x1d\xf1\x67\x8f\xf4\xfd\xe2\x55\x86\x68\x85\xe1\x1e\xd5\x85\xb8\x91\xb8\x4f\x65\xc1\xcc\xcf\x20\xd1\xc4\x39\x9b\x54\x52\x78\x84\x72\xc2\x67\x31\x2b\xe3\x15\x85\xe9\xd5\x84\x94\xf4\xd8\x7d\x4a\x09\x9e\x83\x41\xc2\x53\xcb\x08\x7f\x67\x9b\xcc\x28\xfc\x3b\x56\x46\x38\xba\x84\x90\x86\xaa\x3b\x1e\x1b\x12\x29\x1b\x1c\x5d\x32\x78\x64\x9e\x63\x65\x82\xa3\x4b\x04\x8f\xcc\x73\xac\x2c\x70\x74\x49\xe0\x51\x79\x8e\xa3\xa5\x5b\x21\x15\xfa\xbb\xe3\xf6\xed\xb2\x39\xae\x11\x43\x30\x75\x58\x24\x5d\x33\xa9\x1a\x1c\x31\x6e\x77\x81\x38\xa4\x6b\x7a\xf0\x70\x67\x1f\x94\xf7\x0a\xae\xcf\xcc\x3a\x67\x25\x95\x3b\xe3\x6c\x87\x2d\x50\xc7\x60\x72\xe1\x59\xf4\x9e\x0a\x45\x10\x29\x42\xdf\x77\x71\xc1\x46\x10\x92\xe3\x75\xea\xb1\x2a\x75\x0c\xde\xa8\x76\x2a\xd3\xc3\xef\x62\x75\x71\xeb\xf6\x3e\x4c\x15\xb4\x8e\xd0\x6c\x5e\x34\xcb\x3d\x25\x2c\x85\x18\xa6\x83\x7e\x21\x06\xc3\x57\x22\x57\x36\x82\xab\xb9\xd1\x0a\x21\x75\x8b\xc6\x89\x0b\x60\x7b\xfb\xcf\x70\x86\xbf\xc4\x73\xc0\xbd\x57\x5a\xd0\x9a\x0f\x9f\xf4\x16\x91\xf2\xc0\x60\xdd\x49\x3a\xf6\xa0\x44\xc9\xa1\x20\x15\x95\xb4\x04\x0d\xd2\xed\xb0\x21\xb7\x72\xbc\xd2\xc2\xa3\x49\xdb\x0e\x33\xef\x5c\x1e\x9c\x7a\xb2\x78\xd0\x72\xa8\x6b\x92\x96\x93\xc3\x2d\x30\x91\x81\x1f\xfc\xd9\xb9\x0f\xc8\x41\xfc\x54\x9d\x05\xca\x27\x70\x29\xbc\x7b\x27\x25\x1a\x87\x97\xc1\x78\x6d\xf2\xf1\xea\x92\xb1\x9a\xa4\x59\x21\x98\x41\x6a\x9b\xa6\x14\xd7\xd5\x9b\xa0\x07\x2f\x4c\x3e\x60\xaa\x7b\x6c\x99\xa4\x17\x23\x1f\xa3\x10\xf9\xf0\x45\xc8\xe3\x0a\x90\xf1\x63\xed\x1f\xbe\xf8\xf8\x08\x85\xc7\x94\x62\xc2\xf8\xe1\x39\x93\x8a\x8d\x8f\x51\x68\x3c\xaa\xc8\xe8\x65\x19\xa4\x3a\x4d\xc6\x0f\x23\xcb\xa4\xe2\xe1\xf1\x85\x43\x22\xc2\xa0\xaf\xa3\x8a\x86\x4d\xaa\x21\x48\xf6\xb1\x0a\x86\xc7\x59\xce\xa8\x87\x7d\x8c\xf5\x44\x1d\x0b\xaf\xaf\xa3\x8a\x84\xf1\x23\xbd\x1f\xa2\x40\x78\x7c\xa8\x10\xbc\x24\xa1\x2a\x58\x46\x5f\x0c\xbd\x4b\x72\xdc\xe7\x20\xdc\x8b\xf9\x97\x59\x36\x44\x33\xfa\x49\x88\x70\xb9\x8f\x4c\x7a\x8f\xf8\x3e\x21\x95\xad\xbe\x4c\xff\xc0\x44\xc8\x03\x3d\xaa\xb4\x6a\x99\x48\x9d\x43\x2d\x0a\x90\xc3\x72\xeb\xe6\xcc\xd7\xe4\xc0\x2d\xb2\x27\x96\xb6\x9e\x3f\xd4\xbc\x80\x87\xdf\x33\x65\x95\xc8\xed\xcb\x29\x1f\x1b\x5a\xb8\x7e\xb4\xa6\xd9\xd6\x05\x97\xf6\x8a\xf1\xff\x07\x0f\xb6\x34\x16\x4c\x5b\x53\xbd\x3f\x09\x13\x88\x96\xac\x2a\x80\xfc\xbe\x39\xe3\xf7\x0c\xd6\x6b\xc8\xf4\x1f\x48\xad\x18\xdf\xb8\xf7\xeb\x83\x47\x85\x36\xa7\xec\xfe\xde\xff\xef\x0f\xfd\xd5\x15\x77\x9c\x6c\x7f\x09\x61\xce\x2b\xbc\x91\xb0\x56\x69\x02\xdc\xb0\x2c\x0d\x23\x06\xe4\x35\x7e\x0e\xa7\x3d\x53\x16\x6f\x44\x87\xb7\x4d\x42\x2d\xc9\x8f\x5b\xe0\xed\x89\x74\x5f\x08\x88\x9f\x65\x4e\x25\x90\x77\xe2\xda\x4c\x46\x5d\xc0\x19\xb9\x92\xb0\x06\xb9\xff\x05\xcd\xdb\x3b\xf1\xea\x13\x64\xb5\x0e\x00\x29\x46\x96\xd5\x4d\xe8\x0c\xbc\x8e\x90\xdc\xa1\xa2\x7b\xd1\xe0\x31\xa8\xbe\x86\xb3\x57\x45\x9f\x5f\x09\x99\x42\xe1\x64\x18\x90\xd6\x0d\xec\x54\x73\x2c\xfb\x8d\xed\x13\xab\xe5\x67\x63\xa7\x74\xfb\xd3\xdd\xed\xb9\xec\xbf\xf3\xef\x5c\x95\x2b\xc6\x2d\x63\xb6\x43\x3f\x95\xd8\xa7\x3f\xb1\x39\x88\x81\x32\x37\x21\x4b\xc7\x08\x36\xf4\x25\x90\x01\xe9\xbe\xf7\xea\xde\xbc\xe7\x69\xc3\xed\xdd\x33\x45\x24\xd8\xf7\xae\xb1\x24\xeb\xdc\x02\xfb\xa5\x80\x00\xd3\x78\x80\x5e\xd3\xbb\x7d\x0f\xd2\x9d\x56\x6f\xb4\xe5\xd5\xcf\x35\x2d\xba\x9e\x86\xfb\xc9\xde\x14\xa0\xda\x39\x8a\xdb\x3c\x74\xc7\x8a\x3c\xa3\xd2\x7e\x55\xc1\x2e\x71\xa2\x84\x4b\xb2\xa1\x65\xc9\x28\x6f\xcc\x47\x44\xc0\x38\xf3\xca\x45\xd6\x54\x6a\x96\xd5\x05\x95\xc4\xac\xc5\x8d\x90\xbb\xa3\x64\xbf\x57\xc8\x6b\xc8\x04\xcf\x53\xd2\x1d\x1f\x0f\x9f\x69\xcf\x86\xb6\x65\x79\x26\x72\xf4\x91\x59\x09\x11\xe7\xa6\xb5\x1c\x4e\xec\x21\xbf\x5e\x3b\xc5\xda\xdb\x94\x66\xd1\xb6\x3e\x1b\x81\x4a\x1b\xa0\xd9\x38\x37\x6c\x83\x1e\xcc\x69\xcb\x32\x37\xab\x72\x49\xfe\xb8\xf3\x25\xec\x33\xe7\xf6\xf0\xe0\x7b\x5c\x08\x11\x70\xfc\xb9\xc5\x61\x29\xb6\x96\xf9\x5a\x48\xb8\x05\x49\x4e\x72\x81\xb5\x50\xb8\x65\x99\x3e\x0d\xa9\xde\x7f\x81\x14\xa8\x64\x1c\x36\x54\xb3\xdb\xe6\x50\x74\x1f\x87\x6b\x07\x92\xa0\x8a\x7c\x45\x4e\x90\x18\x61\x65\x09\x39\xa3\x1a\x8a\xe0\x61\xd4\xfe\x3b\x33\x91\xf7\x23\xef\x9f\x1f\x8c\x24\x83\x06\x12\x41\x1d\x63\x68\x7d\xd9\x03\x4b\xd8\x6c\x87\x22\x24\x2e\x67\xe7\xda\x47\xca\xdb\x35\xd8\x29\xd7\x35\xc7\x77\x7b\x43\x38\xf2\x95\x91\xbf\x1a\x5d\xa3\x44\xc2\x06\xd7\x91\x5d\x23\x47\xac\xa2\x51\xbf\xf4\x30\x87\x34\xe4\x18\x2d\xf6\xe7\x20\x77\x7e\x75\xc7\x5d\x76\x7e\xdb\x1f\xa0\xd8\xf9\xf9\xe0\xc3\x6c\x9d\x6b\xfb\x8f\x98\x75\x7e\x1e\xd8\x2b\x16\x1d\x9f\xb9\x73\xa1\xeb\xfa\x3e\x89\x4a\xe0\xe0\x27\x77\xd8\xe7\x05\xb9\x7d\x8e\x41\xc7\xf3\xfd\x6f\x68\x72\x6c\x0a\xaf\x73\xd9\x9d\x3d\x9b\x5f\x20\x76\xc1\xfe\xa0\x85\xa4\x1b\x70\xbf\xfc\x7f\x00\x00\x00\xff\xff\x55\x8a\xaa\x32\xfa\x0e\x01\x00")

func installerKubedbCom_kubedboperatorsYamlBytes() ([]byte, error) {
	return bindataRead(
		_installerKubedbCom_kubedboperatorsYaml,
		"installer.kubedb.com_kubedboperators.yaml",
	)
}

func installerKubedbCom_kubedboperatorsYaml() (*asset, error) {
	bytes, err := installerKubedbCom_kubedboperatorsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "installer.kubedb.com_kubedboperators.yaml", size: 69370, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"installer.kubedb.com_kubedbcatalogs.yaml":  installerKubedbCom_kubedbcatalogsYaml,
	"installer.kubedb.com_kubedboperators.yaml": installerKubedbCom_kubedboperatorsYaml,
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
	"installer.kubedb.com_kubedbcatalogs.yaml":  {installerKubedbCom_kubedbcatalogsYaml, map[string]*bintree{}},
	"installer.kubedb.com_kubedboperators.yaml": {installerKubedbCom_kubedboperatorsYaml, map[string]*bintree{}},
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
