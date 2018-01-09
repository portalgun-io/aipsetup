// Code generated by go-bindata.
// sources:
// ui/main.glade
// DO NOT EDIT!

package ui

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

var _uiMainGlade = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x59\xcd\x72\xdb\x36\x10\xbe\xfb\x29\x50\x5c\x3b\xb2\x94\xf8\xd0\x1e\x48\x66\x9a\xb4\xf6\x64\xf2\xe7\x46\x6a\x7b\xe4\x2c\x81\x15\x85\x0a\x02\x68\x00\x94\xac\xb7\xef\x90\x74\x6c\x47\x02\x2c\x92\x52\x2c\xa7\xce\xcd\x1e\xee\x2e\xf6\xdb\x5f\x7c\x50\xf4\xea\x7a\x21\xc9\x12\x8d\x15\x5a\xc5\xf4\xc5\xe9\x88\x12\x54\x4c\x73\xa1\xf2\x98\xfe\x35\x39\x1f\xfc\x4a\x5f\x25\x27\xd1\x4f\x83\x01\xb9\x40\x85\x06\x1c\x72\xb2\x12\x6e\x46\x72\x09\x1c\xc9\xd9\xe9\xcb\xd1\xe9\x88\x0c\x06\xc9\x49\x24\x94\x43\x33\x05\x86\xc9\x09\x21\x91\xc1\xab\x52\x18\xb4\x44\x8a\x2c\xa6\xb9\x9b\xff\x4c\xef\x0e\xaa\xd4\xe8\xb0\x96\xd3\xd9\xbf\xc8\x1c\x61\x12\xac\x8d\xe9\x85\x9b\xff\x23\x14\xd7\x2b\x4a\x04\x8f\xa9\xd1\xda\xd1\x4a\x8c\x90\xa8\x30\xba\x40\xe3\xd6\x44\xc1\x02\x63\xca\x40\xa5\x53\xcd\x4a\x4b\x93\x73\x90\x16\xa3\xe1\x17\x81\x1b\x79\x36\x13\x92\x37\x7f\xfb\x8e\x79\xad\xaf\xe9\x97\xaf\xdb\xd6\x97\xc2\x8a\x4c\x22\x4d\x26\xa6\xdc\x32\xdd\xc7\x1d\x9f\x8e\x36\x02\x95\x03\x27\xb4\xa2\xc9\x12\x8d\x13\x0c\xa4\x57\xf1\x2b\x2c\x7e\x3c\x1f\x50\x95\xaf\xc1\x34\x71\x5b\x64\xf4\xbe\x74\x0f\x7c\x7d\x31\xfa\xdd\x0d\xbb\xfc\xd6\xe1\x82\x6e\x8a\xf6\xf4\x77\x1f\x9f\x7d\xba\x12\x32\x94\x94\x38\x03\xca\x4a\x70\x90\x49\x8c\xe9\x1a\x2d\x4d\xd2\x73\x21\x3b\x99\x2a\x2d\xa6\xa5\xe2\x68\xa4\x50\x6d\x50\xd4\x11\x24\x6e\x5d\x60\x4c\x6d\x99\x2d\x50\x95\x9e\x28\x85\x82\xea\x15\xdd\x23\xa8\xfb\x06\x96\x04\x8b\xe2\x61\x1c\x75\x71\x34\x05\x2d\x52\x83\x52\x03\x0f\x40\xdb\x1b\xde\x21\x20\xfa\x6c\x84\x6b\xe8\x73\x8d\xa7\x8f\xcd\x8e\xc5\x74\x6b\x06\x18\x43\x59\xcd\x70\x6d\xc8\x1c\xd7\x31\x35\x94\x58\x91\x2b\x90\x31\x05\xe6\xc4\x12\x1c\x52\xb2\xd0\x5c\x4c\x05\x9a\x2a\x0d\xbf\xbf\x4b\xdf\x7c\xfa\x38\xf9\xfc\xe9\x7d\xfa\xe1\xb7\xf1\xbb\x66\x64\x7b\x6d\x0f\x9b\x04\x06\x72\x3f\x7c\x20\xf9\x5d\x0b\x63\x8c\x05\xd4\x18\x1e\x18\x1f\xa1\xc8\x1d\xa1\x24\x1e\x31\x32\x5b\x2d\x63\x61\x89\x4f\x3b\x3a\x5d\x1a\x66\x0c\xcb\x5e\x16\x0f\xd6\x2e\xf6\x47\xbb\xfc\xaf\xdb\xe5\xaa\x14\xee\x69\x47\xa7\x4b\xbb\xfc\x59\x0a\x77\xd4\x76\xb9\x3a\x7e\xbb\x84\x15\x03\x4a\x7e\x05\x8f\xb0\x4f\x30\x2a\x80\xcd\x85\xca\x1f\xbe\x45\xe3\x75\x01\x8a\xef\xba\x42\x6f\x28\x4d\x85\x94\xdd\x2e\xeb\x85\xb6\xa2\x61\x16\x23\xbf\x4a\x34\xdc\x72\x77\x0b\x67\x1b\xe2\x71\x69\x74\x6e\xd0\xda\x5b\xf2\x51\xec\x22\x1f\xad\x79\xc4\x86\x9e\xd2\xa9\x9d\xe9\x55\x0a\x9d\x63\x51\x4a\x8b\xa9\x75\x58\xd0\x64\x74\x3a\x7a\x11\x0a\xc8\x77\x93\xd2\x20\x82\xc3\xa4\xf4\xa3\x76\x98\x69\x3d\x6f\xf2\xa9\xbe\x35\x99\xec\xa2\x56\x17\x40\xa6\x0d\x47\x73\x20\x16\x3a\x66\x46\x4b\x89\xfc\xe6\xdd\xe1\x71\xb8\x68\x57\xd5\x99\xad\xbd\xcc\xc0\xa4\x85\x96\x82\xad\x69\x02\x72\x05\x6b\xdb\xc5\xc8\xf2\x10\x46\xf4\x12\x8d\x84\x75\xda\xd8\x12\x2a\x6f\x43\xac\xc3\x03\x7a\x2b\x1d\x13\x83\xf8\xb7\xc0\x9b\x07\x20\xa1\xa6\x3a\xad\xb7\xda\x23\x73\xda\x3e\xea\xa8\x2a\x47\xd3\xdc\x08\x9e\x56\xcb\xd3\xd2\x24\xd3\x6e\xd6\x8e\x19\x93\xfa\xf9\x4c\x81\x1c\xd4\xff\xc6\xd4\xa2\x44\xd6\xb4\x7b\xeb\x2b\x4d\x15\xbc\xf1\xad\x5e\x60\x9f\x1e\x65\x5f\x6e\xbc\x69\x38\xd8\x1c\x2a\x3e\x38\xef\xeb\xcb\xcd\x77\xf7\x3a\xf4\x46\x2f\x0a\x89\x0e\xc9\x25\xb0\x39\xe4\x48\xde\xaa\xa9\x26\x93\x4a\x22\x7c\x40\x28\xf4\xfe\x0d\xe4\x73\xc8\x41\x96\x36\x4b\x65\x07\x16\xcf\x9a\xd8\x91\xb5\xdd\xa9\xba\x04\x85\xbe\xc7\x99\xa7\x30\x3c\x5b\xbf\xb3\x86\x0c\xdc\x2d\xde\xb3\xb3\x5f\xba\x28\xae\x04\xc7\x74\x06\x8a\xb7\x42\x1c\x6e\xcb\x42\x02\xc3\x99\x96\x1c\x8d\xa7\xa9\x83\x0d\x7d\x68\x8b\xfb\xd7\xe8\xce\x2b\x0c\xe9\x55\x9f\xcf\x62\xaa\x80\xc3\xbc\x2a\x65\x4b\xfe\xe0\xc2\x69\xf3\x2d\x67\x49\xab\x3c\x1d\x77\x04\x85\x2b\xf8\x99\x17\xca\x85\xd1\x65\xf1\xb8\x45\xf2\xf2\xc8\x45\xb2\x3f\x75\xeb\xc2\x41\xf6\x61\x6e\x81\x48\xed\x66\x6e\x5f\x43\xbc\xf7\xf1\xee\x43\x34\xbc\xf7\x13\xec\x7f\x01\x00\x00\xff\xff\x29\xdc\x79\x88\xdb\x1d\x00\x00")

func uiMainGladeBytes() ([]byte, error) {
	return bindataRead(
		_uiMainGlade,
		"ui/main.glade",
	)
}

func uiMainGlade() (*asset, error) {
	bytes, err := uiMainGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui/main.glade", size: 7643, mode: os.FileMode(420), modTime: time.Unix(1515418806, 0)}
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
	"ui/main.glade": uiMainGlade,
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
	"ui": &bintree{nil, map[string]*bintree{
		"main.glade": &bintree{uiMainGlade, map[string]*bintree{}},
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

