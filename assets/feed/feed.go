// Code generated by go-bindata.
// sources:
// templates/feed/atom_1.0.xml
// templates/feed/rss_2.0.xml
// DO NOT EDIT!

package feed

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

var _templatesFeedAtom_10Xml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\x4d\x6f\xdb\x38\x10\x3d\x6f\x7e\xc5\x80\xeb\xc3\x2e\x76\x23\xca\x6e\x9d\x22\x86\x1c\xa0\x49\x50\x20\x40\xda\x43\x3e\x2e\xbd\x04\xb4\x39\x8e\x88\xf2\x43\x20\x47\xb0\x1d\x57\xff\xbd\x20\xa5\xd8\x4a\xec\x00\x41\x9b\x53\xcc\x79\xef\xcd\xd7\x1b\x6d\x36\x20\x71\xa1\x2c\x02\x5b\x20\xca\x07\x41\xce\x3c\x0c\x73\x06\x4d\x73\x54\xc4\x17\x58\x19\x6d\xc3\x94\x95\x44\xd5\x84\xf3\xe5\x72\x99\x2d\x3f\x64\xce\x3f\xf2\x51\x9e\x8f\xf9\x67\x72\x86\xb5\x98\x49\x79\x08\x35\x3c\x3d\x3d\xe5\xab\x92\x8c\x66\x67\x47\x00\x85\x92\x67\x9b\x0d\x64\x97\x2a\x90\x57\xb3\x9a\x94\xb3\x37\xce\xd1\xfd\xcd\x35\x34\x4d\x8c\xdc\xad\x2b\x84\xa6\x29\xb8\x92\x89\x40\x8a\x34\xee\x71\xbe\x09\x13\x51\xd0\x63\x80\xec\xc5\x43\xc1\x5b\x62\x94\xd0\xca\xfe\x80\xd2\xe3\x62\xca\xde\x95\x9b\x33\xf0\xa8\xa7\x4c\x68\x42\x6f\x05\x21\x83\x24\xb6\x4f\xef\xca\x60\x40\xeb\x0a\xa7\x8c\x70\x45\x3c\xf5\xca\x7f\x2f\x71\x1c\x7f\xb6\x32\xba\x2b\x20\xa0\x5e\xbc\x37\xb7\xa8\x2a\xad\xe6\x22\x46\x92\xcc\x7f\xab\x6d\x19\x75\x25\x05\x61\x3b\xf8\xf3\x5a\x69\x79\x29\x08\xb3\x2f\xce\x1b\x41\xc0\x46\x79\x7e\x72\x9c\x0f\x8f\xf3\xd1\xdd\x70\x3c\xc9\x3f\x4e\xf2\xf1\xf7\xfc\xd3\x24\x4f\x26\x28\xf8\x33\x39\x0a\x89\x9a\x4a\xe7\xe3\xbf\x00\x85\x15\xe6\xcd\xbd\x14\x3c\x45\x23\x87\xef\x48\x45\xa8\x67\xa9\x99\xde\xb8\xd8\x9e\xc4\xb9\xae\xfd\x2c\x69\x3c\xc3\xdb\x84\x10\x97\xed\x85\x7d\x44\x18\xa8\xff\x61\xa0\x08\x26\x53\xc8\xae\x08\x4d\x88\xeb\x4f\x7f\x1d\xb2\x40\x4b\x7e\xdd\x15\xda\x5a\x6e\xf0\xf6\xf8\x07\xdb\xf9\xc7\x1f\x8a\xb2\xd8\xc5\x85\x33\x95\xc7\x10\x50\x42\xd3\xfc\xdd\x05\x6e\x4b\x31\x1a\x9f\xbc\x08\x25\xa3\xfe\xd5\x66\x4a\xfb\x7e\x6d\x85\x9d\x01\xfe\xa8\x84\x3d\x4f\xf2\xae\xbd\xed\x81\x1c\xa4\xf5\xee\x00\xa0\xa8\xea\x99\x56\xa1\x6c\xbd\x10\xf1\xd7\x22\xd0\x57\x27\xd5\x42\x75\xe8\x1d\xa2\x65\xf4\xbc\xf3\x8c\xbf\x4f\x4f\xaf\xcd\x01\x50\xcc\x9d\x25\xb4\xf4\x62\xbb\x77\xa5\x0a\x2f\x2e\x13\x22\x4a\x28\x1b\xa0\x93\xbc\x70\xb5\x25\xf8\x09\x65\x6d\x84\x55\x4f\xf8\x30\x77\xc6\x88\xee\xb8\x07\x07\xaf\xde\xe3\xdc\x79\x19\x40\x58\x09\x6a\x2b\x74\xab\x9e\xfa\xbd\xf7\x14\x67\x6b\xc2\xe4\x91\xf9\x2e\x1c\xb9\x3d\xe2\x1e\xbc\x69\xa0\xb6\x3b\x7c\x06\x57\x04\x4b\x11\x60\xee\x31\xb6\x0c\x87\x07\x98\x64\xff\xa1\x12\x41\x0a\x12\xa0\x28\x9e\xf0\xbf\x89\xa8\x45\x20\xe8\x06\x06\xce\xc2\xa1\x89\x66\x05\xef\x86\xd8\xde\xce\xd6\xc6\x9b\x0d\xa0\x95\xe9\xab\xcc\xe3\x67\xf9\x6c\xf7\xf0\x2b\x00\x00\xff\xff\xef\x17\x43\xa7\xc3\x05\x00\x00")

func templatesFeedAtom_10XmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesFeedAtom_10Xml,
		"templates/feed/atom_1.0.xml",
	)
}

func templatesFeedAtom_10Xml() (*asset, error) {
	bytes, err := templatesFeedAtom_10XmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/feed/atom_1.0.xml", size: 1475, mode: os.FileMode(420), modTime: time.Unix(1560185831, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFeedRss_20Xml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x53\xcd\x6e\xdb\x3c\x10\x3c\x7f\x79\x8a\x85\x3e\x1f\x5a\x20\x10\x15\xb7\x09\x60\x43\xf6\x21\x09\x0a\xa4\x88\x7b\x88\x9d\xb3\x41\x8b\x6b\x6b\x51\x89\x34\xc8\x55\x9d\x44\xe5\xbb\x17\xa4\xff\xe4\x3a\x39\xf5\x26\x72\x77\x66\x56\xb3\xc3\xb6\x05\x85\x4b\xd2\x08\xc9\x12\x51\xcd\xad\x73\xf3\x7e\x96\x80\xf7\x17\xb9\x75\x0e\x5e\xea\x4a\xbb\x61\x39\x4a\x4a\xe6\xf5\x50\x88\xcd\x66\x93\x6e\xbe\xa4\xc6\xae\xc4\xd5\x60\x30\x10\x2f\x25\xd7\x55\x02\xbf\xd0\x3a\x32\x7a\x94\xf4\xd3\x2c\x19\x5f\x00\xe4\x45\x29\xb5\xc6\x2a\x7c\x03\xe4\x4c\x5c\xe1\xb8\x6d\x21\xbd\x27\xc7\x96\x16\x0d\x93\xd1\x3f\x64\x8d\xe0\x3d\x84\xfb\xd9\xeb\x3a\x7e\xab\x4e\xdd\xe5\x62\x0b\xdc\x92\x54\xa4\x7f\x9e\x71\x3c\x19\xc3\xcf\x4f\x8f\xe0\x7d\x87\x45\xe4\x22\x36\xef\x70\xd2\xf1\x6d\x43\x95\xba\x97\xbc\x1d\xe2\x70\x4a\xbf\x19\x5b\x4b\x86\x64\x62\xf4\x25\x64\x7d\xf8\x2e\x35\xf4\xb3\xec\x06\xae\xae\x87\xd9\xd7\x61\x76\x0d\x93\xe9\x2c\xd8\x91\x8b\x53\x9a\x2d\xb5\x42\x57\x58\x5a\x87\x41\xce\x26\xbb\xad\x1a\xbb\x88\xc8\x6e\x57\xc4\x41\xf8\x65\x2b\xf5\x0a\xa1\x47\x97\xd0\x23\x86\xe1\x08\xd2\x07\xc6\xda\x05\x13\x00\x60\xd7\x97\x13\x63\xbd\x07\xe5\xab\x86\x54\xd0\xe9\x7d\x6c\x41\xef\xe0\x41\x38\x10\xa7\xc1\xe4\x3b\x53\xaf\x2d\x3a\x87\x0a\xbc\xff\x7f\x57\x98\x96\xb2\x7f\x7d\x73\x52\xca\x45\x54\xd8\xaa\xfd\xb7\x57\xdd\xfb\xfe\x4f\xaa\xdd\x8d\x74\x03\xf1\x41\x73\x67\xef\x7f\xd9\x3c\x2b\xc9\x9d\xa4\x04\x0a\xa3\x59\x92\x76\xb0\xa3\xbb\x33\x8d\x66\xf8\x0d\x65\x53\x4b\x4d\x6f\x38\x2f\x4c\x5d\xcb\x5d\xd0\x7a\xef\x26\xd0\x62\x61\xac\x72\x20\xb5\x02\x3a\x10\x4d\xe9\xad\x3b\x57\x87\x71\xf1\xca\x18\x37\x55\x1c\xcb\x01\xdb\x01\x9e\xb5\x7b\x0f\x8d\x3e\xf6\xa7\xf0\xc0\xb0\x91\x0e\x0a\x8b\x92\xf1\x80\x7d\x94\x8e\x27\x46\xd1\x92\xa2\x15\x91\xf6\x13\x97\x08\x4a\xb2\x04\x62\x87\xd5\xf2\x73\x04\x86\x40\x42\xb3\x56\x11\x6d\x74\x97\xe0\x39\xde\x82\xf7\xe9\x3b\xe9\xcb\xc5\x31\x54\x6d\x0b\xa8\x83\x4e\x78\xb3\xe2\xf0\x68\x73\x61\x9d\x1b\x1f\x8b\x7f\x02\x00\x00\xff\xff\xaf\x8e\x35\x08\x27\x04\x00\x00")

func templatesFeedRss_20XmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesFeedRss_20Xml,
		"templates/feed/rss_2.0.xml",
	)
}

func templatesFeedRss_20Xml() (*asset, error) {
	bytes, err := templatesFeedRss_20XmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/feed/rss_2.0.xml", size: 1063, mode: os.FileMode(420), modTime: time.Unix(1560185831, 0)}
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
	"templates/feed/atom_1.0.xml": templatesFeedAtom_10Xml,
	"templates/feed/rss_2.0.xml": templatesFeedRss_20Xml,
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
		"feed": &bintree{nil, map[string]*bintree{
			"atom_1.0.xml": &bintree{templatesFeedAtom_10Xml, map[string]*bintree{}},
			"rss_2.0.xml": &bintree{templatesFeedRss_20Xml, map[string]*bintree{}},
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

