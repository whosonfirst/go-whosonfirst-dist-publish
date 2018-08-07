// Code generated by go-bindata.
// sources:
// templates/html/inventory.html
// DO NOT EDIT!

package html

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

var _templatesHtmlInventoryHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\x6f\x6f\xe3\xb6\x0f\x7e\xdd\x7c\x0a\xfe\xfc\x3b\x60\x1b\x50\xdb\x77\x1d\xae\x5b\x53\xdb\xc0\xd0\x6b\xb1\x02\xbb\xdd\x80\xb5\xd8\x0d\xc3\x5e\xd0\x16\x13\x73\x93\x25\x4f\xa2\x93\xcb\x05\xfd\xee\x83\x9c\xd8\x71\xfa\x67\xb8\x57\x32\xa5\x87\x0f\x29\xf2\xa1\x9c\xfd\xef\xdd\x87\xab\xbb\xdf\x7f\xb9\x86\x5a\x1a\x5d\xcc\xb2\xb0\x80\x46\xb3\xcc\x23\x32\x51\x31\x03\xc8\x6a\x42\x15\x3e\x00\x32\x61\xd1\x54\xfc\x56\xdb\xaf\x3c\x7c\x30\x70\xc3\xce\x0b\xfc\x71\x73\xfb\xf1\xfd\xf5\x9f\x70\x6b\x56\x64\xc4\xba\x4d\x96\xee\x70\x3b\x9f\x86\x04\xa1\x16\x69\x63\xfa\xa7\xe3\x55\x1e\x5d\x59\x23\x64\x24\xbe\xdb\xb4\x14\x41\xb5\xb3\xf2\x48\xe8\x93\xa4\x21\xfc\x25\x54\x35\x3a\x4f\x92\xdf\xdf\xdd\xc4\xdf\x47\x53\x1e\x83\x0d\xe5\x91\xa3\x05\x39\x47\x6e\xe2\x6d\x1d\x2f\xd9\x44\x2f\xc5\xfc\x18\xdf\xff\x10\x5f\xd9\xa6\x45\xe1\x52\x4f\xc3\xde\x5e\xe7\x17\x11\xa4\x4f\x83\x60\xdb\x6a\x8a\x1b\x5b\xb2\xa6\x78\x4d\x65\x8c\x6d\x1b\x57\xd8\xe2\xb1\xff\x86\xfc\x97\xbb\x7b\x41\xe9\x7c\x5c\xa2\x8b\xbd\x6c\x8e\x78\x4a\x8d\xd5\xdf\xcf\x32\xfd\x88\x46\xd5\xa4\xd5\x8d\x63\x32\x4a\x6f\xa6\x35\x73\x1d\x3d\xeb\xb3\x62\x5a\xb7\xd6\xc9\x04\xbb\x66\x25\x75\xae\x68\xc5\x15\xc5\xbd\x71\x0a\x6c\x58\x18\x75\xec\x2b\xd4\x94\xbf\x39\x85\x06\x3f\x71\xd3\x35\x93\x0d\x36\xc7\x1b\x9d\x27\xd7\x5b\xa1\x0e\xb9\xb1\x87\xf0\xfd\x8d\x40\x36\x2d\xed\x9b\x59\x79\xbf\x6f\x08\x40\x69\xd5\x06\xb6\x7b\x63\x61\x8d\xc4\x0b\x6c\x58\x6f\xe6\xe0\xd1\xf8\xd8\x93\xe3\xc5\xe5\xfe\x58\xb3\xa1\xb8\x26\x5e\xd6\x32\x87\x37\xc9\xb7\xd4\x0c\x27\x0f\xb3\x11\x32\x92\x35\xe8\x96\x6c\xe2\xd2\x8a\xd8\x26\xe0\xdf\x3e\x83\xaf\xac\xa2\xe3\xf0\xeb\x3d\xff\x77\xaf\x5f\x5f\x8e\x20\x6d\xdd\x1c\xfe\x7f\x71\x71\xf1\x84\x20\x51\x24\xc8\xda\x3f\x0e\x2b\xb6\x9d\xc3\xb3\x21\x13\x5f\xe3\xd9\xdb\xf3\xd1\x61\x60\x3f\x3f\x3f\x3f\x60\xfb\xc2\xa5\x7d\xe5\xfa\x59\x4b\x87\x61\xcb\x42\xc1\x8a\x81\x2b\x4c\xe1\xd9\x17\xcc\x5d\x7d\x56\xcc\x66\x27\x3d\xbe\x2d\xee\x6a\xf6\xb0\x60\x4d\xb0\x46\x0f\x19\x42\xed\x68\x91\x47\x61\x2c\xfc\x3c\x4d\x97\x2c\x75\x57\x26\x95\x6d\xd2\x75\x6d\xbd\x35\x8b\xc0\x9a\x2e\x6d\x3c\x31\x63\xc5\x5e\xe2\xb6\x2b\x35\xfb\x3a\x2a\x96\x64\xc8\xa1\x90\x82\x72\x03\xce\x96\x56\x7c\x96\x62\x01\xd6\x40\xd6\x97\xb8\xd2\xe8\x7d\x1e\x29\x14\x8a\x8a\xed\x16\x92\x77\x28\x04\x0f\x0f\x59\x1a\x8e\x8b\x04\xee\x6a\x72\x04\xec\x01\xb5\xb7\x80\xd0\x60\x55\x87\x76\x3b\x42\x15\x14\x05\x2b\x72\x9e\xad\x01\x5c\x21\xf7\x1a\x03\x94\x43\xf2\x3c\xdc\x35\xf9\xcb\x5b\x13\x15\xc7\x76\xc8\x25\xc9\xd2\x76\x2c\x42\xa7\x87\x4f\x80\xed\x16\x1c\x9a\x25\xc1\x2b\x3e\x85\x57\x2c\x30\xcf\x21\xb9\x15\x6a\x3c\x3c\x3c\x0c\xa0\x61\xcd\x86\x41\xda\x6e\x03\x36\xf9\x19\x9b\x70\x8d\xa8\x08\x21\x46\x90\x66\x60\xf5\x14\x33\x3b\x39\xc9\xbc\x38\x6b\x96\xc5\x98\xf8\x04\x13\x9e\x21\x47\xde\x93\xea\xd1\x2f\x9d\x84\x48\x41\x1a\x3d\x0f\x64\xbe\x41\xad\x87\xfa\xee\xa4\x35\xfa\xfe\xda\x9b\x8f\xbc\x7b\x87\x22\x2b\x5d\x18\xd0\x93\x93\x4c\xf1\x6a\x6c\xcf\x4e\xcb\x7d\xa6\x00\xbd\x4e\x42\x9f\x1d\x97\x9d\x84\xda\x87\x37\x03\xd9\xf8\xe3\xa6\x56\xb6\x33\x32\xc6\xbc\x0a\xd6\xa1\xb3\xf0\x48\x9c\x8e\x2a\xeb\x94\x07\x34\x2a\x74\xfb\x88\xc8\xf3\x67\x3a\xe4\xce\x9f\x1f\xdf\x7b\x47\x58\x1d\x36\x03\xc9\x7f\x33\x4c\xfc\x3a\x73\xf0\x4c\xc6\xeb\x03\xdc\x4a\x3f\x07\x95\xa3\x5e\xc0\xcf\xeb\x35\xf0\xfd\x84\x5e\xde\x5b\xc5\x0b\x3e\xca\x27\x24\xf1\xb5\xd4\x04\x0a\x05\x81\xc5\x93\x5e\x7c\xd3\x53\x6a\xf4\x02\x5d\xab\x7a\xde\x17\x47\x61\xa0\xbe\xef\x81\x93\x99\x08\xcd\x49\x15\xaf\xc2\xd8\x0f\xca\x4a\x35\x1f\x29\x97\x4c\x48\x65\x7c\x0d\xf6\xa0\x20\xef\xb0\xee\xde\x8a\x2c\xdd\xfd\xc3\xff\x0d\x00\x00\xff\xff\x84\x52\xf9\xfe\xd4\x07\x00\x00")

func templatesHtmlInventoryHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlInventoryHtml,
		"templates/html/inventory.html",
	)
}

func templatesHtmlInventoryHtml() (*asset, error) {
	bytes, err := templatesHtmlInventoryHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/inventory.html", size: 2004, mode: os.FileMode(436), modTime: time.Unix(1533648851, 0)}
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
	"templates/html/inventory.html": templatesHtmlInventoryHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"inventory.html": &bintree{templatesHtmlInventoryHtml, map[string]*bintree{}},
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

