// Code generated for package charts by go-bindata DO NOT EDIT. (@generated)
// sources:
// charts/core-workshop-infra/Chart.yaml
// charts/core-workshop-infra/templates/code-server.yaml
// charts/core-workshop-infra/templates/external-dns.yaml
// charts/core-workshop-infra/templates/traefik-2.yaml
package charts

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

var _coreWorkshopInfraChartYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x4b\xcc\x4d\xb5\x52\x48\xce\x2f\x4a\xd5\x2d\xcf\x2f\xca\x2e\xce\xc8\x2f\xd0\xcd\xcc\x4b\x2b\x4a\xe4\x2a\x4b\x2d\x2a\xce\xcc\xcf\xb3\x52\x30\xd0\x33\xd4\x33\xe0\x02\x04\x00\x00\xff\xff\xc3\x23\xd1\xe6\x29\x00\x00\x00")

func coreWorkshopInfraChartYamlBytes() ([]byte, error) {
	return bindataRead(
		_coreWorkshopInfraChartYaml,
		"core-workshop-infra/Chart.yaml",
	)
}

func coreWorkshopInfraChartYaml() (*asset, error) {
	bytes, err := coreWorkshopInfraChartYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "core-workshop-infra/Chart.yaml", size: 41, mode: os.FileMode(436), modTime: time.Unix(1577836800, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _coreWorkshopInfraTemplatesCodeServerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x4d\x8f\xdb\x36\x10\xbd\xeb\x57\x10\xe9\x59\xd2\x26\xa7\x05\x6f\xee\x3a\x29\x8c\xee\xae\x0d\x7b\x93\x1e\x03\x9a\x1a\x5b\xac\x28\x52\x20\x87\xda\xb8\x69\xfe\x7b\x41\x7d\x99\xb2\xe4\x06\x71\x80\xe8\x64\x91\x33\x6f\xde\xbc\xf9\x90\x59\x25\x3e\x81\xb1\x42\x2b\x4a\xea\xb7\x51\x21\x54\x46\xc9\x33\x2b\xc1\x56\x8c\x43\x54\x02\xb2\x8c\x21\xa3\x11\x21\x8a\x95\x40\xc9\xab\x36\x85\xcd\x75\xc5\x51\x46\x71\x1c\x47\x21\x80\xd9\x33\x9e\x30\x87\xb9\x36\xe2\x1f\x86\x42\xab\xa4\xb8\xb7\x89\xd0\xe9\x00\xfd\x20\x9d\x45\x30\x5b\x2d\xe1\x77\xa1\x32\xa1\x8e\x33\x31\xb8\xce\x20\xb6\x60\x6a\x30\xdd\x59\xc3\x66\x1c\xdc\x68\x09\x5b\x38\x78\x2f\x56\x89\x3f\x8c\x76\xd5\xff\x30\x88\x08\x99\x10\x38\xc7\x6b\xcf\x62\x96\x95\x42\x45\xd6\xed\xff\x06\x8e\x96\x46\x71\xe7\xb3\x03\x53\x0b\x0e\x0b\xce\xb5\x53\xf8\x63\x34\x2f\x35\x1a\x94\xb8\x00\xbd\x51\x85\xab\xf0\x0f\x5a\x1d\xc4\xf1\x89\x55\xb7\x22\x13\x22\xd9\x1e\xa4\xf5\x6e\x5e\xe2\x6a\xec\xd7\x23\xf2\x26\x0e\x25\xff\xc6\x9d\xdd\x88\x8b\x3f\x0a\xf9\x34\x07\x9d\xdc\x1d\x72\xdc\xbf\xb7\xaf\x8d\x01\x18\x14\x07\xc1\x19\x42\xdc\x15\x13\x4f\x94\xa4\x35\x33\xa9\x71\x2a\xb5\xc0\x0d\xa0\x4d\x0b\xb7\x07\xa3\x00\xa1\xe9\x31\xdb\x0a\xca\x5a\x41\x53\xce\x12\x6e\x70\xc0\x6c\x69\x53\x92\x23\x56\x96\xa6\xa1\x6f\x06\x07\xe6\x64\x6f\xda\x4a\x14\x9e\x71\xad\x10\xbe\xe0\x99\x6f\xfb\x1e\xf0\xed\x12\x20\x63\xa4\x91\xb4\x97\x57\xce\x4e\x1d\x66\x42\x3b\x63\x40\x61\xdc\x87\x1c\x5d\x7a\x88\x81\xd4\xd4\xb7\x8b\x31\x44\x44\x5d\x80\xfa\x20\x24\xfc\xa8\x92\x8d\xe3\xa4\xd3\x58\x55\xd9\xf3\x5c\x2f\xa1\x92\xfa\x54\xc2\xed\x9d\xfc\x9d\x7e\xb3\x15\x70\x7f\x65\x41\x02\x47\xdd\xe5\x55\x32\xe4\xf9\x63\xe0\x37\xe3\x49\x08\x42\x59\x49\x86\xd0\xf9\x04\xfc\xfc\x23\x47\xee\xb3\x00\x84\xf4\xe1\xfb\x56\x3a\x4f\xee\xf3\x4c\x8a\x7d\xd7\x30\xa1\x86\x12\xf9\x22\x89\x92\x1d\x81\x12\xe9\xbe\x30\x9b\x16\xf7\x36\x7e\x85\x7d\x2c\x32\xa0\xf5\xbb\xe4\x6d\x72\x37\x50\xf8\x8d\xbc\xac\x97\x6b\x4a\x56\x8a\x60\x0e\xe4\xe0\xd0\x19\xa0\xe4\x98\x73\xe3\xeb\xc3\xa5\x76\x59\xac\x18\x8a\x1a\x62\xa5\x4d\x26\xf8\x05\xdc\x5d\xf2\x2e\x80\x6b\xc2\x6e\x9c\x94\x1b\x2d\x05\x3f\x51\xb2\x90\xaf\xec\x64\x47\x6d\x3a\x97\x02\x21\x95\x36\x18\x48\xd3\xb7\x99\x1f\xa2\xe1\x30\x48\x75\xa3\x0d\x52\x72\x7f\x77\x7f\x8e\x0d\xaa\x9e\xfa\xbf\x7c\x7c\x59\x6f\x57\x8b\xc7\xdd\xe7\xed\xfb\xcd\x3a\x40\xaa\x99\x74\xf0\xc1\xe8\x92\x06\x87\x5e\x71\xdf\xa6\x7f\xc2\xa9\x5b\xf8\xe1\x33\xfd\x2c\x8d\xef\x0b\x38\x5d\x0d\x38\xe5\xb3\x5c\x6d\x7f\x29\x9d\x30\x5e\xcf\x66\xb3\xd8\xed\xfe\x5a\x6f\x97\xbf\x80\xc8\xc3\xe3\xc7\xdd\xcb\xfb\xed\xe7\x49\xc8\x5a\x4b\x57\xc2\x93\x6f\xf0\x99\xfa\xfb\x75\xc1\xcf\x0b\xbd\x7d\x4a\x6f\xbc\x61\x98\x53\x92\xe6\xba\x84\xd4\x37\x94\x49\x13\x6f\x3c\x41\xc8\x34\x2f\x82\x3e\x1b\x7b\xf7\xdb\xa9\x35\x4a\xac\xe6\x45\x14\xd2\x0a\x26\xea\x2a\x1f\xde\x7f\xff\x42\x5d\xae\x35\xfa\x15\x56\xb9\xb6\x2d\xa5\x00\xa2\xba\x4e\xf1\x7b\x9f\xfa\x5b\x37\xe3\xfc\xea\x9b\x59\x53\xc3\xb4\xc6\xcd\x4f\x3f\x89\x8d\x29\x32\x73\x04\xbc\x18\xce\x60\x92\x2f\x89\x2b\x40\x1f\x5f\xa8\xe3\xf9\xff\xdb\x1e\x90\xf5\xf9\xac\xd4\xd1\x80\xb5\x3f\x9b\x8f\x68\x61\x1e\x24\xb3\xb6\xdd\xa2\x68\x18\x1c\x84\x2f\xb6\x71\x12\xba\x54\x3c\xc5\xbe\x02\x5e\xfd\xa0\xfa\x5d\x31\x86\xf2\xec\x19\x2f\x40\x65\x61\xbd\xba\x65\x7d\x6d\x4b\x07\x26\x9b\x50\x32\x5f\x7a\x4a\xde\x7c\xfd\x4a\x92\x4f\x7e\xf6\x6c\x12\x64\x90\xf4\x63\xb3\x5c\x3f\x2d\x56\xcf\xe4\xdb\xb7\x37\xd1\x7f\x01\x00\x00\xff\xff\x4c\xeb\x62\xbe\x49\x0b\x00\x00")

func coreWorkshopInfraTemplatesCodeServerYamlBytes() ([]byte, error) {
	return bindataRead(
		_coreWorkshopInfraTemplatesCodeServerYaml,
		"core-workshop-infra/templates/code-server.yaml",
	)
}

func coreWorkshopInfraTemplatesCodeServerYaml() (*asset, error) {
	bytes, err := coreWorkshopInfraTemplatesCodeServerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "core-workshop-infra/templates/code-server.yaml", size: 2889, mode: os.FileMode(436), modTime: time.Unix(1577836800, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _coreWorkshopInfraTemplatesExternalDnsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x53\x51\x6f\xdb\x36\x10\x7e\xd7\xaf\x38\x68\xaf\x91\xd7\xa2\x0f\x1b\x04\xf8\x21\x6b\x8c\x21\x40\xe2\x14\x76\xea\x0d\x1b\x82\x80\x26\xcf\x12\x17\x8a\x24\xee\x4e\x4a\xd4\x22\xff\x7d\xa0\xe5\x24\x52\xe3\xae\x18\xc2\x27\x89\xc7\xfb\xee\xfb\xbe\xbb\x53\xd1\x6e\x90\xd8\x06\x5f\x42\xf7\x3e\xbb\xb3\xde\x94\xb0\x46\xea\xac\xc6\x53\xad\x43\xeb\x25\x6b\x50\x94\x51\xa2\xca\x0c\xc0\xab\x06\x4b\xc0\x07\x41\xf2\xca\x15\xc6\xf3\xe1\x92\xa3\xd2\x58\xc2\x7d\xa0\x3b\xae\x43\xd4\xe2\xb2\xa2\x28\xb2\x31\x3e\x6d\x95\x9e\xa9\x56\xea\x40\xf6\x8b\x12\x1b\xfc\xec\xee\x57\x9e\xd9\xf0\xf3\x73\xe5\x8f\xae\x65\x41\x5a\x05\x87\x3f\x2a\x4b\xad\x43\x2e\xb3\x02\x54\xb4\xbf\x53\x68\x23\x97\xf0\x77\x9e\xdf\x64\x00\x84\x1c\x5a\xd2\xb8\xbf\xe1\x41\x0c\xe7\x27\x39\x7a\x13\x83\xf5\x92\xbe\x63\x30\xbc\x7f\xdc\x21\x6d\xf7\x0f\x2b\x94\xfc\x24\xbf\x57\xa2\xeb\xfc\x24\x77\x96\x25\xbf\xf9\x16\x3e\x31\xf0\x49\x4d\x82\xf0\x28\x49\xae\xf5\xd5\x41\xc7\xeb\xe2\xd6\x57\x84\xcc\xc8\xf9\x0d\xfc\xdf\x5a\xaf\xd1\x7c\x30\xf8\x0d\xe9\xa7\xd4\x37\x59\xfd\x9b\xf5\xc6\xfa\xea\x07\x8e\x17\x9d\xc5\x7b\xa4\xef\xf7\x9b\x82\xc3\x15\xee\x52\xf6\x93\x90\xff\x60\x92\x01\xbc\xee\xf9\xd1\x4e\x73\xbb\xfd\x07\xb5\xec\x9b\x7d\x74\x40\xdf\x36\x96\x2a\x46\x7e\xb1\xe5\x0c\xa3\x0b\x7d\x83\x6f\x98\x7b\x8e\xa8\xcb\x7d\xef\xa2\xb3\x5a\x71\x09\xef\x33\x00\x16\x52\x82\x55\x9f\x22\x00\xd2\x47\x2c\x61\x85\x9a\x50\x49\xd2\xcd\xe8\x50\x4b\xa0\x21\xdc\xa4\xc9\xb8\x50\x5b\x74\x3c\x5c\x24\x4b\xe3\x2b\x02\x82\x4d\x74\x4a\xf0\x90\x34\xe2\x9b\x8e\x9b\xe4\x1f\x47\x00\x78\x62\xbb\xff\x9e\xf8\xba\x3c\xa6\x39\x1d\x1d\xbc\x28\xeb\x91\x9e\xc1\x8b\xe3\x06\x0d\xc7\x36\xaa\xc2\x12\x52\xdb\x2b\x4d\x69\x08\xc7\xcf\x26\x3f\x65\xf7\x6e\xf6\xcb\xec\xc3\x0b\x65\xaa\x46\x02\x7e\x82\x3f\x92\x2f\x70\x3e\x6c\x15\x84\x61\x2e\x40\x79\x03\x83\x91\x70\xb6\x5c\x03\xa1\x0e\x64\x18\x74\x20\x42\x8e\x61\x3f\xdc\xae\x87\x67\x9c\x02\x8a\x62\x58\xab\xf9\x61\x41\x47\x25\x2e\xd5\x1d\x32\x2c\x0e\xa4\x12\x1e\x23\x42\xf0\xae\x07\xa9\x11\xea\xc0\x82\x06\xbe\x04\x8f\x3c\xb4\xc9\xfa\x0a\x22\x85\xce\x1a\x34\x60\x42\xa3\xac\x3f\x81\xd0\x58\x01\x09\x29\xa0\x13\x55\xe5\x1c\xa8\x4e\x59\xa7\xb6\x6e\x0a\x32\x61\x35\xa4\x17\x3b\xeb\x04\x69\xfe\xf5\x2b\xcc\x36\xca\xb5\xc8\xb3\xd1\x7c\xcd\x3e\x5e\x7c\x5e\x5f\x2f\x56\xb7\x67\x57\x97\xa7\xe7\x4b\x78\x7c\x9c\x4d\x30\x0e\x5c\xbe\x9b\xbe\xf8\xf3\x7a\xb1\x5a\x9e\x5e\xdc\x9e\x2d\xd7\xb7\x9f\x56\x57\x9b\xf3\xb3\xc5\x0a\x1e\x1f\x47\x1e\x7c\x22\xec\xd0\xcb\xd4\x86\x1d\x85\x06\x0c\x3a\x94\xa4\x58\xf9\xfe\xc9\xe8\x17\xb5\xe8\xf7\xf2\x76\xad\x73\xc0\xbd\xd7\x35\x05\x7f\xd8\xfa\x29\xc5\xe0\xac\xee\xe7\x6d\x64\x24\x29\x92\xb7\x93\x30\x61\x65\x59\xa8\x9f\xcb\x83\x4c\x02\xf2\x20\x45\xb8\xf7\x48\x85\x35\xf3\xf1\xca\xbd\x50\x5f\x61\x2a\x0c\x61\x27\xe8\x41\x31\x48\x6d\x19\x2c\x83\x82\xda\x56\xb5\xeb\xc1\xf4\x5e\x35\x56\x03\xf7\x2c\xd8\x4c\xe0\xad\x17\xa4\x4e\xb9\xf9\x87\x77\xd3\xb6\xb8\x50\x15\x0e\x3b\x74\x73\x83\xdb\xb6\x1a\x3b\x15\x98\xed\xd6\xf5\x60\xfd\x7e\x3c\x76\xad\xb4\x84\x27\xa0\x8c\x81\xa2\x38\xb8\xc8\x6d\x8c\x81\x64\x9c\x76\xb0\xbd\xf8\xf8\x79\x7d\x7d\x75\x79\xfe\xd7\xa2\x84\x25\xa2\x49\x26\x32\x0a\x44\xa4\xe7\x3e\xc2\x62\xb9\x81\xcd\xe9\x8a\xa1\x46\xc2\xec\xdf\x00\x00\x00\xff\xff\x5c\xbe\x30\x6d\xb2\x07\x00\x00")

func coreWorkshopInfraTemplatesExternalDnsYamlBytes() ([]byte, error) {
	return bindataRead(
		_coreWorkshopInfraTemplatesExternalDnsYaml,
		"core-workshop-infra/templates/external-dns.yaml",
	)
}

func coreWorkshopInfraTemplatesExternalDnsYaml() (*asset, error) {
	bytes, err := coreWorkshopInfraTemplatesExternalDnsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "core-workshop-infra/templates/external-dns.yaml", size: 1970, mode: os.FileMode(436), modTime: time.Unix(1577836800, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _coreWorkshopInfraTemplatesTraefik2Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\xff\x6e\xe3\xb8\x11\xfe\xdf\x4f\x31\xb0\x0f\xed\x1e\xb0\xa2\x37\xb7\x57\xe0\x20\xc0\x40\x1d\xc7\x7b\x6b\x5c\x92\x0d\x6c\x67\x81\x16\x0b\xe4\x68\x6a\x2c\xf1\x4c\x91\x04\x49\x29\x75\xd3\xbc\x7b\x41\xfd\xb2\x6c\xd9\x4e\x93\xb6\xab\x00\x81\x44\xce\x7c\x33\x1c\xce\x7c\x1c\x7a\x30\x18\xc0\xfc\x72\x3c\x81\xc1\x60\xd0\xa3\x9a\x7f\x45\x63\xb9\x92\x21\x98\x15\x65\x84\x66\x2e\x51\x86\xff\x93\x3a\xae\x24\xd9\xfc\x62\x09\x57\xc3\xfc\xa2\xb7\xe1\x32\x0a\x61\x22\x32\xeb\xd0\xcc\x95\xc0\x5e\x8a\x8e\x46\xd4\xd1\xb0\x07\x20\x69\x8a\x21\x38\x43\x71\xcd\x37\x3d\x93\x09\xb4\x61\x2f\x00\xaa\xf9\xaf\x46\x65\xda\x7a\x19\x80\x00\xfa\xfd\x1e\x80\x41\xab\x32\xc3\xb0\x19\xb5\x68\x72\xce\xd0\x56\x9f\x28\x23\xad\xb8\x74\xb6\x99\x66\x06\x8b\xaf\x1c\xcd\xaa\xd1\x8a\xd1\x55\x6f\x82\xdb\xfa\xf5\x91\x3a\x96\x1c\xb3\x8c\xff\x70\x28\xfd\x3a\x6b\x54\x89\xee\x51\x99\x0d\x97\x71\xb5\xca\x63\x9e\x71\x19\x1b\xb4\xb6\x71\xad\xfa\x66\x82\x56\x83\xff\x2b\x8f\x4e\x5b\x1e\x5a\x47\x5d\xd6\xb1\x95\xe9\x88\x3a\xec\x05\x41\xb0\xb7\x87\xcd\x4e\x2d\xca\x98\x8e\x19\x53\x99\x74\x67\x36\xab\xfc\xb6\x9a\x32\x0c\xc1\x47\xc4\x26\x4a\x33\x27\x3a\xd0\xaf\x4b\x8f\x4b\x2e\x23\x2e\xe3\x73\x59\xa2\x04\xce\x71\xed\x27\xea\xd8\x9c\x31\xd2\x03\xe8\xa6\xe0\x21\xa4\xcd\x56\x7f\x20\x73\x45\xee\x1d\x0d\xc3\xab\x16\xef\x0b\x65\xa2\xe4\x9a\xc7\x37\x54\xbf\x87\x2b\xd4\x42\x6d\x53\x94\x0e\xfe\x54\xe3\x76\x6a\x68\x17\x8a\x5a\xf1\x74\x04\x02\xb6\x8e\x4f\x7b\x00\x20\xe8\x0a\x45\xb5\xe1\x54\xeb\x9d\xd7\x35\x1a\xd5\x9c\x6c\x69\x2a\x42\xf8\x57\x21\x94\x38\xa7\x4b\x71\x00\xa3\x32\x87\xc6\xd6\x9f\x00\xe9\x36\xa0\x9a\xef\xbe\x01\x7c\x99\x86\xf0\x59\x59\xf7\xee\xf7\x0a\x9a\x3c\x3d\x01\xf9\x4a\x45\x86\x96\xb4\x9c\x21\x93\xeb\xfb\xc5\x72\x3a\x7f\xb8\xfa\x72\x33\x9e\xdd\xc2\xf3\xf3\xef\x3f\xb6\x80\xaa\xfa\x0d\xbd\x43\x7f\xe5\xd2\xa1\x91\x54\xb4\xe6\x51\x3a\xb3\xbd\x2b\x4a\xba\x6d\x3f\x68\x6d\x43\xe3\x25\x8f\x22\x81\x8f\xd4\xe0\x81\x28\xd5\x3c\xf0\x79\xd1\x3b\x23\x57\xcb\xb4\x35\x57\xd4\x72\x36\x3e\x18\x04\xc8\xec\x5e\x70\xfc\x33\x80\x65\xc2\x2d\xe4\x7e\xf5\xc0\x2d\x18\xd4\x82\x32\x8c\x40\xc9\x20\xc2\x94\xca\x08\xd6\x46\xa5\xe0\x12\x84\x98\xe7\x28\x01\x65\x0e\x39\x35\xef\xc1\x25\xd4\x01\x53\x29\xda\x52\x84\xc2\xa2\xe0\xac\x3d\xfc\x00\xfa\xdf\x9e\xbe\x3d\x15\x5a\xfd\x3a\xa0\x97\xe3\xc5\x6c\xf2\x30\xbe\x5f\x7e\x7e\xb8\x9c\xcc\xff\x76\xb7\xec\xc3\xb7\xe7\x6f\xcf\xfd\x4e\xf5\x51\xad\xed\xae\xd0\x76\xb9\xf8\x96\xca\x3e\x93\x5a\x56\x23\x0b\x0b\x3e\xd2\x82\x33\x6a\x43\xb8\xe8\xf9\x0d\x16\xc8\x9c\x32\xa5\x42\xea\x19\xed\xba\x85\x70\x80\x01\xe0\x30\xd5\x82\x3a\xac\xe4\x5b\x1e\xfa\x47\xec\xa9\x76\x94\x01\x6a\x27\x5a\xb9\x55\x15\xf0\xed\xc1\x02\xfd\xc3\x94\x74\x94\xcb\xd6\x7e\x06\x9d\x40\x94\x0f\x4f\x69\xbc\x1b\x0e\xf3\x9f\xc8\xc7\x9d\x13\x26\x6e\xb9\x14\x40\xe0\xab\x65\xe4\x4c\x86\x7b\x83\x29\x3a\xc3\x99\x25\xda\xa8\x14\x5d\x82\x99\xed\xca\x50\xc6\xd0\x5a\xa1\xe2\x33\x53\x64\xcd\x05\x6a\xea\x92\xd1\xb0\x61\x03\xca\x12\x1c\x96\x12\x44\xa8\x78\x4f\xd1\xab\x08\xcc\x51\x8c\xae\xa6\x97\xf7\xbf\xee\xcd\x69\xa3\x72\x1e\xa1\xb1\x05\x28\x89\xb8\x29\x36\x6b\xdb\x82\x5e\xc7\xe7\x34\x8a\x23\xaa\xeb\xec\x4e\x6a\x93\xad\xd0\x48\x74\x68\xab\x63\xe9\x55\xc2\xa4\x3e\xcb\xaa\x73\x9d\xe8\x6c\x25\xb8\x4d\x30\xaa\x36\x77\xd4\x4a\xce\xe1\xe1\xa6\x79\xf0\x82\x40\xca\x9e\x80\x3c\xe2\x8a\xd0\x28\x2a\xbc\x08\x7f\xf9\x70\x56\xce\x13\x22\x31\x58\x46\xc4\x9f\xb4\x64\x27\x41\x9c\x1a\x3d\xe2\xca\x22\xcb\x0c\xbe\x1d\xc5\xb2\x04\x53\x1c\x79\x19\x7b\x0e\xa5\xb4\xb3\xf3\xfc\xe7\x9f\x3f\xfe\x07\xe2\x85\x69\x27\x2c\x61\x68\x9c\xef\x11\x44\x8e\x66\x24\xd0\x59\x94\xcc\x6c\xb5\x3b\x89\x51\x13\x7a\x2b\x54\x67\x82\x55\x4b\xbf\xce\x9e\x17\xe2\x6b\xce\xa8\x43\x5b\x0b\x5b\xd2\x92\x26\x94\xa5\x48\x22\x69\x59\x42\x85\x40\x19\x63\x37\x71\xde\x00\x42\xea\x64\x1b\xfd\xf0\x6e\x39\x1f\x4f\x3f\xcd\x7e\x7b\xb8\xba\x5d\x3c\xdc\xcd\xbf\x7c\x9d\x5d\x4d\xe7\x3f\xfe\xd7\x06\x1a\xb9\xd1\x05\x29\xfe\xc2\xbf\x7c\x7c\x7f\x41\x3e\x90\x0f\xc5\xeb\x5b\xf0\x31\xa5\x5c\x8c\x7e\x78\x77\x3d\x5d\x2e\xa6\xb7\x05\xd7\x3f\x4c\x6f\xc6\xb3\xeb\x37\x79\x6b\x9d\x32\x34\xc6\x2e\x7b\xa4\x48\xfe\xb0\x4a\x36\x98\x03\xb8\x13\x48\x2d\x82\x54\x0e\xcb\x83\xca\xf9\x63\x8e\xdb\xe2\x20\xb3\x8e\xc6\x5c\xc6\x70\x8d\xee\xcf\x16\xa6\xa5\x89\x82\x74\xd1\x90\x16\xc8\x17\xc9\x10\xb6\x2a\xf3\x3d\xae\x07\x90\xb1\x85\xaa\x7d\x7e\x5f\x8c\xdb\x44\x65\x22\x02\x83\xa9\xca\x2b\x3b\x8f\x89\x12\x08\x82\x4b\x04\x2a\x9c\x8a\x3d\x5f\xb6\x40\x5f\xb1\x5a\x46\x4b\x8f\xca\x32\x0b\x87\xc5\x3a\x83\xca\xf7\x20\xff\xf0\x13\xf1\xad\x50\x5b\x4d\x99\x78\xd8\x10\x61\x63\x52\x2b\xe3\xf6\x58\xbe\x3c\x27\x3c\x6a\xeb\xa8\x6e\x0e\x94\x3b\x65\x5c\x08\x7b\x65\xb3\x53\xb0\xa7\x35\xf6\x6b\xbb\x54\xa1\x51\xca\xe5\x39\x23\x47\xcc\x14\x3a\xc1\x4b\xde\xb5\x14\x51\xe6\xdd\xd5\x75\x12\xae\x05\x56\x34\x3b\x9f\x8c\x4a\xf7\x5b\xa1\xf2\xc6\xf5\x1b\x6e\xab\xfe\xbc\xfd\x94\xa0\xfb\xed\x44\xfb\xd9\xe0\xf6\x9c\xcd\xda\xab\x63\x75\xfb\x1d\x1c\x3b\x6b\xb6\xea\x00\xb9\xf5\xed\x61\x04\x5c\x16\x25\xd2\xf7\x17\x30\xce\x7c\xe0\xd7\x3c\xce\x4c\x71\x2b\xe9\xc3\x5a\x99\x62\x7a\x59\x16\x20\x8c\xef\x66\x9d\x45\x9e\xec\xf2\xbe\xc3\x4a\x5f\xb6\x3d\x80\x3a\x02\xc1\xe4\x7e\xb1\xfc\x72\x33\xfb\xfb\x34\x84\xcf\x68\x7c\xfd\xfa\xff\x12\x31\xb2\xe0\x14\xac\x10\x94\x44\x48\x95\x41\x88\xb6\x92\xa6\x9c\x41\x5f\xfb\xae\xa1\x5f\x96\x7a\x4a\x37\x68\x9b\x50\xd0\x95\x40\x70\xaa\x65\xa8\xec\x6a\x8a\x78\x5d\xdd\x2e\xc0\xa9\x0d\xee\x8a\x21\x57\x22\x4b\xf1\xc6\xf7\x77\x47\x8a\xf3\x58\x0f\x03\x90\x7a\xe9\x3b\xea\x92\x10\x4e\x74\x39\x07\xda\x9e\x1e\x5f\xd2\x6f\xc9\x94\x3e\x9d\xea\x28\xf7\x4c\xb1\xfa\xa6\xd7\xde\xa6\x53\xe2\xe7\xfd\xc2\x54\xbb\xed\x15\x37\x21\x3c\x3d\xbf\x78\xb5\xff\x3f\x75\xfe\x0d\x47\x76\xd8\x51\xef\xe8\xb0\x4b\x84\xba\xc5\x7c\x03\x18\x0b\xa1\x1e\x6d\xb3\xed\x6a\xaf\x52\xee\x67\x04\x3e\x29\x03\x11\xa6\x0a\x32\x4b\x63\x9f\x5d\x62\x4b\x7a\xc7\xd8\x52\xb7\x59\x6e\xff\x1e\x72\x78\xe9\xd8\x6a\xcf\x76\x8a\x46\x97\x54\x50\xc9\xd0\x74\x42\xd8\xf9\xcd\x67\x98\x5f\xac\xd0\xd1\x3a\xb2\xb3\xb2\x51\x9d\x08\x6a\xed\x2e\xbc\x70\x24\xbe\x54\x4a\xe5\x0a\x26\xa8\x62\xd9\xfe\x65\xa8\xd5\x02\x7b\x1b\xdc\x06\x11\xae\x69\x26\x5c\x50\x4c\x87\xd0\xf7\x5d\x50\xbf\x09\xb9\xa7\x74\xa3\x84\x40\xd3\x98\x28\xf4\x4a\xc8\x60\x37\xdd\xfb\x77\x00\x00\x00\xff\xff\x9c\x61\x27\x85\xbc\x13\x00\x00")

func coreWorkshopInfraTemplatesTraefik2YamlBytes() ([]byte, error) {
	return bindataRead(
		_coreWorkshopInfraTemplatesTraefik2Yaml,
		"core-workshop-infra/templates/traefik-2.yaml",
	)
}

func coreWorkshopInfraTemplatesTraefik2Yaml() (*asset, error) {
	bytes, err := coreWorkshopInfraTemplatesTraefik2YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "core-workshop-infra/templates/traefik-2.yaml", size: 5052, mode: os.FileMode(436), modTime: time.Unix(1577836800, 0)}
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
	"core-workshop-infra/Chart.yaml":                  coreWorkshopInfraChartYaml,
	"core-workshop-infra/templates/code-server.yaml":  coreWorkshopInfraTemplatesCodeServerYaml,
	"core-workshop-infra/templates/external-dns.yaml": coreWorkshopInfraTemplatesExternalDnsYaml,
	"core-workshop-infra/templates/traefik-2.yaml":    coreWorkshopInfraTemplatesTraefik2Yaml,
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
	"core-workshop-infra": &bintree{nil, map[string]*bintree{
		"Chart.yaml": &bintree{coreWorkshopInfraChartYaml, map[string]*bintree{}},
		"templates": &bintree{nil, map[string]*bintree{
			"code-server.yaml":  &bintree{coreWorkshopInfraTemplatesCodeServerYaml, map[string]*bintree{}},
			"external-dns.yaml": &bintree{coreWorkshopInfraTemplatesExternalDnsYaml, map[string]*bintree{}},
			"traefik-2.yaml":    &bintree{coreWorkshopInfraTemplatesTraefik2Yaml, map[string]*bintree{}},
		}},
	}},
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