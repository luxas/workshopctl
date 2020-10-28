// Code generated for package charts by go-bindata DO NOT EDIT. (@generated)
// sources:
// charts/core-workshop-infra/Chart.yaml
// charts/core-workshop-infra/templates/code-server.yaml
// charts/core-workshop-infra/templates/external-dns.yaml
// charts/core-workshop-infra/templates/traefik-2.yaml
// charts/podinfo/external-chart
// charts/podinfo/values-override.yaml
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

var _coreWorkshopInfraTemplatesCodeServerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x4d\x73\xdb\x36\x10\xbd\xeb\x57\xec\x38\x67\x92\x76\x7a\xf1\xf0\xe6\x5a\x71\xc7\x53\xdb\xd2\x48\x4a\x7a\xcc\xac\xc0\x95\x88\x0a\x04\x38\xc0\x52\xb6\x9a\xe6\xbf\x77\xc0\x2f\x81\xa2\xd4\x4c\x7c\x08\x4f\x24\x80\x7d\xfb\xf0\xf0\x76\x41\x2c\xe5\x17\xb2\x4e\x1a\x9d\xc2\xfe\x66\xb2\x93\x3a\x4b\xe1\x05\x0b\x72\x25\x0a\x9a\x14\xc4\x98\x21\x63\x3a\x01\xd0\x58\x50\x0a\xaf\xc6\xee\x5c\x6e\x4a\xc1\x6a\x12\x45\xd1\x24\x04\xb0\x6b\x14\x31\x56\x9c\x1b\x2b\xff\x41\x96\x46\xc7\xbb\x5b\x17\x4b\x93\xf4\xd0\xf7\xaa\x72\x4c\x76\x61\x14\xfd\x2e\x75\x26\xf5\xf6\x4c\x0e\x61\x32\x8a\x1c\xd9\x3d\xd9\x76\xac\x66\x33\x4c\x6e\x8d\xa2\x05\x6d\x7c\x14\x96\xf2\x0f\x6b\xaa\xf2\x7f\x18\x4c\x00\x46\x04\x8e\xf9\x9a\xb1\x08\xb3\x42\xea\x89\xab\xd6\x7f\x93\x60\x97\x4e\xa2\x36\x66\x49\x76\x2f\x05\xdd\x09\x61\x2a\xcd\x3f\x47\xf3\x54\xa3\x5e\x89\x13\xd0\x77\xaa\x70\x11\xfe\xde\xe8\x8d\xdc\x3e\x63\xf9\x5e\x64\x00\x85\x6b\x52\xce\x87\x79\x89\xcb\x61\x5c\x87\x28\xea\x3c\x29\xfc\x1b\xb5\xeb\x06\x5c\xfc\x50\xc8\xa7\x1e\x68\xe5\x6e\x91\xa3\xee\xbb\xf9\xac\x17\x90\x65\xb9\x91\x02\x99\xa2\xf6\x30\xf9\x90\x42\xb2\x47\x9b\xd8\x4a\x27\x8e\x84\x25\x76\xc9\xae\x5a\x93\xd5\xc4\x54\x7b\xcc\x35\x82\x62\x23\x68\x22\x30\x16\x96\x7b\xcc\x86\x76\x0a\x39\x73\xe9\xd2\x24\x8c\xcd\x68\x83\x95\xea\x96\x36\x12\x85\x63\xc2\x68\xa6\x37\x3e\xf2\x6d\xbe\x03\xbe\xed\x06\x60\x88\x34\x90\xf6\x74\xaa\x72\xe3\x80\x33\xa9\x2b\x6b\x49\x73\xd4\xa5\x1c\x4c\x7a\x88\x9e\xd4\x38\xb6\xcd\xd1\x67\x64\xb3\x23\xfd\x20\x15\xfd\xac\x92\x75\xe0\xc8\x69\x58\x96\xee\x58\xd7\x53\x2a\x95\x39\x14\xf4\x7e\x27\xff\xc0\x6f\xae\x24\xe1\xa7\x1c\x29\x12\x6c\xda\x7d\x15\xc8\x22\x7f\x0a\xe2\xce\x44\x02\x30\x15\xa5\x42\xa6\x36\x26\xe0\xe7\x1f\x35\x08\x3f\x0b\x00\xd0\xa5\xef\xac\x74\xac\xdc\x97\x33\x5b\xec\x5c\x83\x52\xf7\x47\xe4\x0f\x49\x16\xb8\xa5\x14\x54\xf5\x86\x2e\xd9\xdd\xba\xe8\x95\xd6\x91\xcc\x28\xdd\x7f\x8c\x6f\xe2\x9b\x9e\xc2\x07\x58\xcd\xa6\xb3\x14\x1e\x35\x70\x4e\xb0\xa9\xb8\xb2\x94\xc2\x36\x17\xd6\x9f\x8f\x50\xa6\xca\x22\x8d\x2c\xf7\x14\x69\x63\x33\x29\x4e\xe0\xae\xe3\x8f\xf1\x75\x0f\x57\xa7\x9d\x57\x4a\xcd\x8d\x92\xe2\x90\xc2\x9d\x7a\xc5\x83\x1b\xd8\xf4\xdc\x16\x00\x4a\x63\x39\x90\xa6\xb3\x99\x2f\xa2\x7e\x30\xd8\xea\xdc\x58\x4e\xe1\xf6\xfa\xf6\x98\x9b\xf4\x7e\x1c\xbf\xfa\xbc\x9a\x2d\x1e\xef\x9e\x96\x5f\x17\x9f\xe6\xb3\x00\x69\x8f\xaa\xa2\x07\x6b\x8a\x34\x18\xf4\x8a\x7b\x9b\xfe\x49\x87\xb6\xe1\x87\xcf\xf8\x5a\x1a\xce\xef\xe8\x70\x31\xe1\x98\xcf\xf4\x71\xf1\x4b\xe9\x84\xf9\x3a\x36\xf3\xbb\xe5\xf2\xaf\xd9\x62\xfa\x0b\x88\xdc\x3f\x7d\x5e\xae\x3e\x2d\xbe\x8e\x52\xee\x8d\xaa\x0a\x7a\xf6\x06\x3f\x73\xfe\xbe\x5d\x88\x63\x43\x6f\x9e\xc2\x2f\x9e\x23\xe7\x29\x24\xb9\x29\x28\xf1\x86\xb2\x49\xec\x17\x8f\x10\x32\x23\x76\x81\xcf\x86\xd1\x5d\x77\x6a\x16\xc5\xce\x88\xdd\x24\xa4\x15\x54\xd4\x45\x3e\xa2\xbb\xff\x42\x5d\x2e\x19\xfd\x02\xab\xdc\xb8\x86\x52\x00\x51\x5e\xa6\xf8\xa3\xab\xfe\xbd\x9d\xf1\x7c\xeb\x3b\xd3\xa6\xfa\x6a\x8d\xea\x57\x5f\x89\xf5\x52\x46\xbb\x25\x3e\x29\xce\xa0\x92\x4f\x89\x6b\x62\x9f\x5f\xea\xed\xe8\xff\xed\x51\x6f\x2d\x39\xf7\xfe\x26\x8f\x5a\x1b\xae\xff\xcc\xda\x53\xfc\x00\x0f\xc6\x0a\x02\x04\x97\x1b\xcb\xb0\x5a\x3d\x81\x33\xc0\x39\x32\x4c\x5f\x96\x60\x49\x18\x9b\x81\xc8\x51\x6f\xc9\x81\x40\x0d\xa5\x35\x25\x6e\x91\x09\x36\xe8\x2f\xde\x1a\x87\xde\x98\xac\x46\x15\x65\xda\xc5\xa8\xca\x1c\xe3\xe1\xad\xc6\xac\x52\xb8\xfa\xed\xda\x5d\xf5\x8a\xca\x66\x37\xf7\x0a\x9d\x6b\xfa\x38\x5b\xa4\x8d\xf4\x76\xb3\x95\xa2\x56\x4c\xef\x83\x14\xae\xbe\x7d\x83\xf8\x8b\x2f\x44\x17\x07\x7b\x8a\xbb\x1a\x9a\xce\x9e\xef\x1e\x5f\xe0\xfb\xf7\xab\x9a\x8f\x57\xb6\x33\x8e\x37\x4d\x60\xda\xd6\x43\x93\xd0\x53\xab\x43\xe9\x4b\xdf\xd2\x46\xbe\xf5\x13\x6b\x14\x3b\xd2\x59\xe8\xbf\xf6\xf2\x19\x56\xfb\xe5\x06\x0e\x6d\x13\x1f\x75\x87\xaa\x58\xfb\x1f\x90\xdb\xeb\xc9\x7f\x01\x00\x00\xff\xff\x6d\x42\x76\x1b\x03\x0c\x00\x00")

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

	info := bindataFileInfo{name: "core-workshop-infra/templates/code-server.yaml", size: 3075, mode: os.FileMode(436), modTime: time.Unix(1577836800, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _coreWorkshopInfraTemplatesExternalDnsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\xdf\x6f\xdb\xb6\x13\x7f\xf7\x5f\x71\xd0\xf7\x35\xf2\xb7\x45\x1f\x36\x08\xf0\x43\xd6\x18\x43\x80\xc6\x29\x9c\x2c\x1b\x36\x04\x01\x4d\x9e\x25\x2e\x14\x49\xdc\x9d\x94\xa8\x45\xfe\xf7\x81\x92\x9c\x48\x8d\xbb\x62\x88\x9e\x64\x9e\xee\xee\xf3\x8b\x56\xd1\xde\x20\xb1\x0d\xbe\x80\xf6\xfd\xe2\xde\x7a\x53\xc0\x15\x52\x6b\x35\x9e\x6a\x1d\x1a\x2f\x8b\x1a\x45\x19\x25\xaa\x58\x00\x78\x55\x63\x01\xf8\x28\x48\x5e\xb9\xdc\x78\x1e\x0f\x39\x2a\x8d\x05\x3c\x04\xba\xe7\x2a\x44\x2d\x6e\x91\xe7\xf9\x62\x3a\x9f\x76\x4a\x2f\x55\x23\x55\x20\xfb\x45\x89\x0d\x7e\x79\xff\x33\x2f\x6d\xf8\xff\xf3\xe6\x8f\xae\x61\x41\xda\x06\x87\x3f\x5a\x4b\x8d\x43\x2e\x16\x39\xa8\x68\x7f\xa5\xd0\x44\x2e\xe0\xaf\x2c\xbb\x5d\x00\x10\x72\x68\x48\x63\x7f\xc2\x03\x19\xce\x4e\x32\xf4\x26\x06\xeb\x25\xbd\xc7\x60\xb8\xff\xb8\x45\xda\xf5\x1f\x96\x28\xd9\x49\xf6\xa0\x44\x57\xd9\x49\xe6\x2c\x4b\x76\xfb\xed\xf8\x84\xc0\x27\x36\x69\x84\x47\x49\x74\xad\x2f\x47\x1e\xaf\x97\x5b\x5f\x12\x32\x23\x67\xb7\xf0\x5f\x77\xbd\x9e\xe6\x83\xc1\x6f\x40\x1f\x5a\xdf\x24\xf5\x2f\xd6\x1b\xeb\xcb\x1f\x28\x9e\xb7\x16\x1f\x90\xbe\xef\x37\x05\x87\x5b\xdc\xa7\xee\x03\x91\x7f\x41\xb2\x00\x78\xed\xf9\x51\xa7\xb9\xd9\xfd\x8d\x5a\x7a\xb3\x8f\x06\xf4\x6d\xb1\x54\x31\xf2\x8b\x2c\x67\x18\x5d\xe8\x6a\x7c\x43\xee\x39\xa2\x2e\x7a\xef\xa2\xb3\x5a\x71\x01\xef\x17\x00\x2c\xa4\x04\xcb\x2e\x55\x00\xa4\x8b\x58\xc0\x16\x35\xa1\x92\xc4\x9b\xd1\xa1\x96\x40\x43\xb9\x4e\xc9\xf8\xa4\x76\xe8\x78\x38\x48\x92\xc6\x57\x00\x04\xeb\xe8\x94\xe0\xd8\x34\xc1\x9b\x1e\x37\xeb\x3f\x3e\x01\xe0\x80\xb6\x7f\x9f\xe9\xba\x39\xc6\x39\x3d\x3a\x78\x51\xd6\x23\x3d\x0f\xcf\x8f\x0b\x34\x3c\xb6\x56\x25\x16\x90\x6c\x2f\x35\xa5\x10\x4e\x3f\x9b\xfd\x28\xda\x77\xcb\x9f\x96\x1f\x5e\x20\x53\x39\x21\xf0\x3f\xf8\x3d\xe9\x02\xe7\xc3\xad\x02\xe5\xcd\x21\x0a\x10\x86\x8c\xf4\x67\x83\xa8\x70\xb6\xb9\x02\x42\x1d\xc8\x30\xe8\x40\x84\x1c\x43\x1f\x74\xd7\xc1\xf3\xcc\x1c\xf2\x7c\xb8\x62\xab\xf1\xb2\x1e\x2b\x8d\xc2\x4c\x90\x5c\xa8\x7b\x64\x58\x8f\xd8\xd3\x2a\x46\x84\xe0\x5d\x07\x52\x21\x54\x81\x05\x0d\x7c\x09\x1e\x79\x70\xd3\xfa\x12\x22\x85\xd6\x1a\x34\x60\x42\xad\xac\x3f\x81\x50\x5b\x01\x09\xa9\xa0\x7b\x46\xce\x81\x6a\x95\x75\x6a\xe7\xe6\x43\x66\xa8\x86\xf6\x7c\x6f\x9d\x20\xad\xbe\x7e\x85\xe5\x8d\x72\x0d\xf2\x72\x12\xc3\xe5\xf6\xf2\xf2\xfa\xee\xec\xf2\xe2\xf4\x7c\x03\x4f\x4f\xcb\xd9\x80\x11\xc8\x77\x7b\xd7\x7f\x5c\xaf\xb7\x9b\xd3\x4f\x77\x67\x9b\xab\xbb\xcf\xdb\xcb\x9b\xf3\xb3\xf5\x16\x9e\x9e\x26\x02\x7c\x26\x6c\xd1\xcb\x5c\x83\x3d\x85\x1a\x0c\x3a\x94\x44\x57\xf9\xee\x60\xc0\x0b\x55\xf4\x3d\xb7\x7d\xe3\x1c\x70\xe7\x75\x45\xc1\x8f\xff\x0c\x73\x88\xc1\x59\xdd\xad\x9a\xc8\x48\x92\x27\x61\x67\x65\xc2\xd2\xb2\x50\xb7\x92\x47\x99\x15\xe4\x51\xf2\xf0\xe0\x91\x72\x6b\x56\xd3\x6b\xf9\x02\x7d\x8b\x69\x31\x84\xbd\xa0\x07\xc5\x20\x95\x65\xb0\x0c\x0a\x2a\x5b\x56\xae\x03\xd3\x79\x55\x5b\x0d\xdc\xb1\x60\x3d\x1b\x6f\xbd\x20\xb5\xca\xad\x3e\xbc\x9b\x7b\xe2\x42\x99\x3b\x6c\xd1\xad\x0c\xee\x9a\x72\xaa\x54\x60\xb6\x3b\xd7\x81\xf5\x7d\x36\xf6\x8d\x34\x84\x27\xa0\x8c\x81\x3c\x1f\x55\xe4\x26\xc6\x40\x32\x6d\x1b\x65\xcf\x3f\xfe\x76\x75\x7d\x79\x71\xfe\xe7\xba\x80\x0d\xa2\x49\x22\x32\x0a\x44\xa4\x67\x1f\x61\xbd\xb9\x81\x9b\xd3\x2d\x43\x85\x84\x8b\x7f\x02\x00\x00\xff\xff\xfa\xb6\x0f\x82\xd6\x07\x00\x00")

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

	info := bindataFileInfo{name: "core-workshop-infra/templates/external-dns.yaml", size: 2006, mode: os.FileMode(436), modTime: time.Unix(1577836800, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _coreWorkshopInfraTemplatesTraefik2Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x7d\x6b\x1b\x3d\x12\xff\xdf\x9f\x62\xb0\x1f\xee\xfa\x40\x56\x49\xda\x1e\x94\x05\xc3\x39\x8e\xdb\x9a\xe6\x0d\xdb\x29\xdc\x11\x48\x65\xed\x78\x57\xb5\x56\x12\x92\xd6\xa9\x2f\x97\xef\x7e\x68\x5f\xbc\x6b\x6f\xec\x5e\xf2\x94\x6e\x20\xec\x4a\x33\xbf\x19\xfd\x34\x33\x1a\xb9\xd7\xeb\xc1\xe4\x6c\x30\x84\x5e\xaf\xd7\xa1\x9a\x7f\x45\x63\xb9\x92\x21\x98\x39\x65\x84\x66\x2e\x51\x86\xff\x87\x3a\xae\x24\x59\x7e\xb0\x84\xab\xe3\xd5\x69\x67\xc9\x65\x14\xc2\x50\x64\xd6\xa1\x99\x28\x81\x9d\x14\x1d\x8d\xa8\xa3\x61\x07\x40\xd2\x14\x43\x70\x86\xe2\x82\x2f\x3b\x26\x13\x68\xc3\x4e\x00\x54\xf3\x4f\x46\x65\xda\x7a\x19\x80\x00\xba\xdd\x0e\x80\x41\xab\x32\xc3\x70\x33\x6a\xd1\xac\x38\x43\x5b\x7e\xa2\x8c\xb4\xe2\xd2\xd9\xcd\x34\x33\x98\x7f\xad\xd0\xcc\x37\x5a\x31\xba\xf2\x4d\x70\x5b\xbd\x3e\x50\xc7\x92\xe7\x2c\xe3\x0f\x87\xd2\xaf\xb3\x42\x95\xe8\x1e\x94\x59\x72\x19\x97\xab\x7c\xce\x33\x2e\x63\x83\xd6\x6e\x5c\x2b\xbf\x99\xa0\xe5\xe0\xaf\xf2\x68\xbf\xe5\x63\xeb\xa8\xcb\x5a\xb6\x32\x1d\x51\x87\x9d\x20\x08\xb6\xf6\x70\xb3\x53\xd3\x82\xd3\x01\x63\x2a\x93\xee\xc0\x66\x15\xdf\x56\x53\x86\x21\x78\x46\x6c\xa2\x34\x73\xa2\x05\xfd\xb2\xf0\x38\xe3\x32\xe2\x32\x3e\x14\x25\x4a\xe0\x04\x17\x7e\xa2\xe2\xe6\x80\x91\x0e\x40\x3b\x04\x77\x21\x6d\x36\xff\x8e\xcc\xe5\xb1\xf7\x2c\x0d\x2f\x5a\xbc\x4f\x94\xa1\x92\x0b\x1e\x5f\x52\x7d\x04\xe7\xa8\x85\x5a\xa7\x28\x1d\xfc\xad\xc2\x6d\xe5\x50\x4d\x45\xa5\xb8\x9f\x81\x80\x2d\xe2\xfd\x1e\x00\x08\x3a\x47\x51\x6e\x38\xd5\xba\xf6\xba\x42\xa3\x9a\x93\x35\x4d\x45\x08\xff\xcd\x85\x12\xe7\x74\x21\x0e\x60\x54\xe6\xd0\xd8\xea\x13\x20\x5d\x07\x54\xf3\xfa\x1b\xc0\xa7\x69\x08\x9f\x95\x75\x6f\xbe\x95\xd0\xe4\xf1\x11\xc8\x57\x2a\x32\xb4\xa4\xe1\x0c\x19\x5e\xdc\x4e\x67\xa3\xc9\xfd\xf9\xf5\xe5\x60\x7c\x05\x4f\x4f\xdf\xfe\x6c\x00\x95\xf9\x1b\x7a\x87\xfe\xc9\xa5\x43\x23\xa9\x68\xcc\xa3\x74\x66\x7d\x93\xa7\x74\xd3\x7e\xd0\xd8\x86\x7a\xec\x01\xe7\x16\x59\x66\xb0\x31\x9a\xf2\x28\x12\xf8\x40\x0d\xee\x00\x50\xcd\x03\x1f\x2d\x9d\x03\x72\x95\x4c\x53\x73\x4e\x2d\x67\x83\x9d\x41\x80\xcc\x6e\x51\xe6\x9f\x1e\xcc\x12\x6e\x61\xe5\x39\x01\x6e\xc1\xa0\x16\x94\x61\x04\x4a\x06\x11\xa6\x54\x46\xb0\x30\x2a\x05\x97\x20\xc4\x7c\x85\x12\x50\xae\x60\x45\xcd\x11\xb8\x84\x3a\x60\x2a\x45\x5b\x88\x50\x98\xe6\x95\x6c\x0b\x3f\x80\xee\xdd\xe3\xdd\x63\xae\xd5\xad\x68\x3e\x1b\x4c\xc7\xc3\xfb\xc1\xed\xec\xf3\xfd\xd9\x70\xf2\xaf\x9b\x59\x17\xee\x9e\xee\x9e\xba\xad\x9c\xa4\x5a\xdb\x3a\xfd\xea\x08\x7d\x4d\xbe\x1f\x08\x38\xab\x91\x85\x79\x95\xd2\x82\x33\x6a\x43\x38\xed\xf8\x6d\x17\xc8\x9c\x32\x85\x42\xea\xeb\xdc\x45\x03\x61\x07\x03\xc0\x61\xaa\x05\x75\x58\xca\x37\x3c\xf4\x8f\xd8\x52\x6d\x29\x03\x54\x4e\x34\x22\xae\x4c\xeb\xab\x9d\x05\xfa\x87\x29\xe9\x28\x97\x8d\xfd\x0c\x5a\x44\x14\x0f\x4f\x69\x5c\x0f\x87\xab\xb7\xe4\x5d\xed\x84\x89\x1b\x2e\x05\x10\xf8\x1c\xea\x3b\x93\xe1\xd6\x60\x8a\xce\x70\x66\x89\x36\x2a\x45\x97\x60\x66\xdb\x32\x94\x31\xb4\x56\xa8\xf8\xc0\x14\x59\x70\x81\x9a\xba\xa4\x7f\xbc\xa9\x11\x94\x25\x78\x5c\x48\x10\xa1\xe2\x2d\x45\xaf\x22\x70\x85\xa2\x7f\x3e\x3a\xbb\xfd\xb4\x35\xa7\x8d\x5a\xf1\x08\x8d\xcd\x41\x49\xc4\x4d\xbe\x59\xeb\x06\xf4\x22\x3e\xa4\x91\x1f\x5c\x6d\x67\x6b\xa9\x65\x36\x47\x23\xd1\xa1\x2d\x0f\xab\x17\x09\x93\xea\x84\x2b\x4f\x7b\xa2\xb3\xb9\xe0\x36\xc1\xa8\xdc\xdc\x7e\x23\x38\x8f\x77\x37\xcd\x83\xe7\x65\xa5\xe8\x14\xc8\x03\xce\x09\x8d\xa2\xdc\x8b\xf0\xc3\xc9\x41\x39\x5f\x26\x89\xc1\x82\x11\x7f\xfe\x92\x5a\x82\x38\xd5\x6f\x57\xa0\x17\xa3\x58\x96\x60\x8a\x7d\x2f\x63\x0f\xa1\x14\x76\x6a\xcf\xdf\xbf\x7f\xf7\x7f\x88\xe7\xa6\x9d\xb0\x84\xa1\x71\xbe\x73\x10\x2b\x34\x7d\x81\xce\xa2\x64\x66\xad\xdd\x5e\x8c\xaa\xcc\x37\xa8\x3a\x40\x56\x25\xfd\x32\x7b\x5e\x88\x2f\x38\xa3\x0e\x6d\x25\x6c\x49\x43\x9a\x50\x96\x22\x89\xa4\x65\x09\x15\x02\x65\x8c\xed\xc0\x79\x05\x08\xa9\x82\xad\xff\xc7\x9b\xd9\x64\x30\xfa\x38\xfe\x72\x7f\x7e\x35\xbd\xbf\x99\x5c\x7f\x1d\x9f\x8f\x26\x7f\xfe\x65\x03\x1b\xb9\xfe\x29\xc9\xff\xc2\x7f\xbc\x3b\x3a\x25\x27\xe4\x24\x7f\x7d\x0d\x3e\xa6\x94\x8b\xfe\x1f\x6f\x2e\x46\xb3\xe9\xe8\x2a\xaf\xf5\xf7\xa3\xcb\xc1\xf8\xe2\x55\xde\x5a\xa7\x0c\x8d\xb1\x5d\x3d\x52\x24\xdf\xad\x92\x1b\xcc\x1e\xdc\x08\xa4\x16\x41\x2a\x87\xc5\x41\xe5\xfc\x31\xc7\x6d\x7e\x90\x59\x47\x63\x2e\x63\xb8\x40\xf7\x77\x0b\xa3\xc2\x44\x5e\x74\xd1\x90\x06\xc8\xb5\x64\x08\x6b\x95\xf9\xce\xd7\x03\xc8\xd8\x42\xd9\x54\x1f\xe5\xe3\x36\x51\x99\x88\xc0\x60\xaa\x56\xa5\x9d\x87\x44\x09\x04\xc1\x25\x02\x15\x4e\xc5\xbe\x5e\x6e\x81\xbe\x60\xbd\x8c\x16\x3e\x15\x89\x16\x1e\xe7\x2b\x0d\x4a\xef\x83\xd5\xc9\x5b\xe2\x5b\xa4\xa6\x9a\x32\xf1\xf1\xa6\x14\x6e\x8c\x6a\x65\xdc\x56\x9d\x2f\x4e\x0a\x8f\xda\x38\xac\x37\x47\xca\x8d\x32\x2e\x84\xad\xc4\xa9\x15\xec\x7e\x8d\xed\xec\x2e\x54\x68\x94\x72\x79\xc8\x48\xc3\x0c\xca\x55\xdb\xc9\x56\xe4\x34\xc0\xf2\xae\xe5\xa3\x51\xe9\x76\x4f\x53\x5c\xa8\xbe\xe0\xba\x6c\xbf\x9b\x4f\x01\xba\xdd\x17\x34\x9f\x25\xae\x0f\xd9\xac\xbc\x7a\x2e\x01\x7f\x83\x63\x07\xcd\x96\xad\x1c\xb7\xbe\xcf\x8b\x80\xcb\x3c\xd6\xbb\xfe\x7e\xc5\x99\x27\x7e\xc1\xe3\xcc\xe4\x97\x8e\x2e\x2c\x94\xc9\xa7\x67\x45\x26\xc1\xe0\x66\xdc\x5a\xe4\xde\x76\xed\x37\xac\xf4\xe7\xb6\x7b\x50\x31\x10\x0c\x6f\xa7\xb3\xeb\xcb\xf1\xbf\x47\x21\x7c\x46\xe3\x13\xd1\xff\x97\x88\x91\x05\xa7\x60\x8e\xa0\x24\x42\xaa\x0c\x42\xb4\x96\x34\xe5\x0c\xba\xda\x1f\xff\xdd\x22\x67\x53\xba\x44\xbb\xa1\x82\xce\x05\x82\x53\x0d\x43\x45\x7b\x92\xf3\x75\x7e\x35\x05\xa7\x96\x58\xc7\xf4\x4a\x89\x2c\xc5\x4b\xdf\xa8\x3d\x93\x63\xcf\x35\x23\x00\xa9\x97\xbe\xa1\x2e\x09\x61\x4f\xbb\xb2\xa3\xed\xeb\xdc\xcf\xf4\x1b\x32\x85\x4f\xfb\x5a\xc3\x2d\x53\xac\xba\xc8\x35\xb7\x69\x9f\xf8\x61\xbf\x30\xd5\x6e\x7d\xce\x4d\x08\x8f\x4f\x3f\xbd\xb9\xff\xe2\x16\x1e\x80\x4a\xa9\x5c\x1e\xdd\xe5\x3c\xfe\x28\x6e\x6a\x41\x24\x2d\xa1\x42\x27\xb4\xd1\xa5\xf9\x7b\x7d\xa2\xac\x2b\x2c\x77\x5f\x7a\x3f\xec\x6e\xee\x0c\x9b\xda\xda\xaa\xaa\xba\x2e\xa3\xed\x02\xaa\x1b\x15\x73\xfb\xa2\xb1\x7b\xab\x58\x6b\x5f\x05\x15\x8d\xce\xa8\xa0\x92\xa1\x69\x51\xdb\xfa\xa9\xa7\xbe\x32\x8d\x8b\x26\x74\x28\xa8\xb5\x35\xe3\xf0\x0c\xe5\x2d\xfa\x9a\xbf\x05\xed\x10\xc7\x6d\x10\xe1\x82\x66\xc2\x05\xf9\x74\x4e\x60\x86\x35\x29\xbe\xca\x1b\x25\x04\x9a\x8d\x89\x5c\xaf\x80\x0c\xea\xe9\xce\xff\x02\x00\x00\xff\xff\xf0\x5d\x30\x93\xae\x13\x00\x00")

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

	info := bindataFileInfo{name: "core-workshop-infra/templates/traefik-2.yaml", size: 5038, mode: os.FileMode(436), modTime: time.Unix(1577836800, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _podinfoExternalChart = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x28\x29\x29\x28\xb6\xd2\xd7\x2f\x2e\x49\x4d\x4b\xcc\x2b\x28\xca\x4f\x49\xcc\xd3\x4b\xcf\x2c\xc9\x28\x4d\xd2\xcb\xcc\xd7\x2f\xc8\x4f\xc9\xcc\x4b\x83\xd3\x80\x00\x00\x00\xff\xff\xb0\xe7\xfd\xe6\x2e\x00\x00\x00")

func podinfoExternalChartBytes() ([]byte, error) {
	return bindataRead(
		_podinfoExternalChart,
		"podinfo/external-chart",
	)
}

func podinfoExternalChart() (*asset, error) {
	bytes, err := podinfoExternalChartBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "podinfo/external-chart", size: 46, mode: os.FileMode(436), modTime: time.Unix(1577836800, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _podinfoValuesOverrideYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xce\x4d\x4b\x03\x31\x10\x87\xf1\xfb\x7c\x8a\xbf\xf1\x6c\x14\xd6\x52\x98\x9b\x54\x41\xa1\x2a\xf8\x72\x96\xb8\x3b\xee\x06\xa7\x99\x90\x64\xed\xa1\xec\x77\x97\xed\xf5\xc7\x73\x78\x8a\x64\x8d\x7d\xd8\xd9\x9c\x1a\xa3\x23\xb5\x71\x2f\x7f\xa2\x8c\x98\x7e\x8c\x68\x8e\x4c\x40\x6f\x6a\x85\xe1\x2e\xbb\xdb\xcd\x76\xdb\x3b\x02\x0e\x52\x6b\x18\x85\xe1\x1e\x45\xd5\x70\xb4\xa2\xc3\x85\x23\x8a\x87\x95\x09\x68\x61\x64\x6c\xfc\x8d\xef\x88\x62\x1a\x8b\xd4\xba\xb2\xa4\xf0\xad\x32\x30\x5a\x99\x85\x80\x1c\xda\xc4\xb8\x26\x60\xb2\xda\xce\xc9\x15\x5c\xb6\x61\x1d\xf0\xa7\x13\xfc\xd1\xca\x6f\x9d\x2c\xf7\x4d\xfd\x6e\xff\xf9\xfe\xf1\xf0\xf6\x75\xff\xfa\x7c\xf7\xf4\x82\x65\x71\xf4\x1f\x00\x00\xff\xff\x80\x9d\x33\xb3\xc3\x00\x00\x00")

func podinfoValuesOverrideYamlBytes() ([]byte, error) {
	return bindataRead(
		_podinfoValuesOverrideYaml,
		"podinfo/values-override.yaml",
	)
}

func podinfoValuesOverrideYaml() (*asset, error) {
	bytes, err := podinfoValuesOverrideYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "podinfo/values-override.yaml", size: 195, mode: os.FileMode(436), modTime: time.Unix(1577836800, 0)}
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
	"podinfo/external-chart":                          podinfoExternalChart,
	"podinfo/values-override.yaml":                    podinfoValuesOverrideYaml,
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
	"core-workshop-infra": {nil, map[string]*bintree{
		"Chart.yaml": {coreWorkshopInfraChartYaml, map[string]*bintree{}},
		"templates": {nil, map[string]*bintree{
			"code-server.yaml":  {coreWorkshopInfraTemplatesCodeServerYaml, map[string]*bintree{}},
			"external-dns.yaml": {coreWorkshopInfraTemplatesExternalDnsYaml, map[string]*bintree{}},
			"traefik-2.yaml":    {coreWorkshopInfraTemplatesTraefik2Yaml, map[string]*bintree{}},
		}},
	}},
	"podinfo": {nil, map[string]*bintree{
		"external-chart":       {podinfoExternalChart, map[string]*bintree{}},
		"values-override.yaml": {podinfoValuesOverrideYaml, map[string]*bintree{}},
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
