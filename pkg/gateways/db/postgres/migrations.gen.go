package postgres

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __000001_initialize_schema_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x72\x75\xf7\xf4\xb3\xe6\xe2\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xc8\x4c\xa9\x88\x4f\xcd\x2b\x29\xca\x4c\x2d\x8e\x4f\x4c\x4e\xce\x2f\xcd\x2b\x89\x4f\xce\x49\x2c\x2e\xb6\x26\xa4\x2a\xbd\x28\xbf\xb4\x80\xa0\xaa\xe2\xd2\x24\xe2\x14\x66\xa6\xc0\x1c\x16\xe2\xe8\xe4\xe3\xaa\x00\x95\x87\xea\x0b\x89\x0c\x70\x55\x50\xc8\x2f\x48\x2d\x4a\x2c\xc9\xcc\xcf\x8b\x2f\xa9\x2c\x48\x45\x91\x42\x73\x3b\x97\xb3\xbf\xaf\xaf\x67\x88\x35\x17\x20\x00\x00\xff\xff\x23\x79\x57\x99\xf3\x00\x00\x00")

func _000001_initialize_schema_down_sql() ([]byte, error) {
	return bindata_read(
		__000001_initialize_schema_down_sql,
		"000001_initialize_schema.down.sql",
	)
}

var __000001_initialize_schema_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xc1\x8e\x82\x30\x10\x86\xcf\xf2\x14\x73\x53\x13\xdf\xc0\x13\x62\x63\x9a\x48\x31\x50\x12\x77\x2f\x4d\xa5\xdd\x4d\xb3\x58\xd8\xb6\x6c\xe4\xed\x37\x54\x21\x60\xd8\xac\x9c\x9a\xcc\xc7\x37\x33\xff\xec\xd0\x01\x93\x6d\x10\x44\x29\x0a\x29\x02\xfa\x76\x42\xc0\x8b\xa2\x6a\xb4\x63\x45\xc9\xad\x85\x30\x03\x44\xf2\x18\x56\xc1\x62\x59\x2a\x7e\x51\xa5\x72\xed\x72\x13\x2c\x96\xdc\x5a\xe9\xac\x7f\x2a\x5d\x54\x57\xe9\x9f\xf2\x56\x4b\x6d\x1f\xef\xef\xa6\x83\x83\xf5\x53\x87\xaa\x96\x86\x3b\x55\x69\xe6\xda\x5a\x4e\x5a\x08\x79\x51\xce\xff\x5c\x18\x29\x94\x9b\xfe\x1c\xee\x8e\x08\xa4\x76\x46\x49\xdb\xe1\x4a\xc0\xf4\xcb\x73\xbc\x87\xda\xa8\x2b\x37\x2d\x7c\xc9\x76\x13\x2c\xa6\xeb\x00\x40\x18\x45\x49\x4e\x28\x8b\x8e\x61\x96\x81\xae\x1c\xe8\xa6\x2c\x47\xe4\xa7\xa9\x9a\xba\x23\x29\x3a\xd3\x39\xc0\x36\x97\x3b\xf3\x17\x30\xcc\xf5\x0c\x0c\x9b\xf7\x03\x27\x27\x94\x86\x14\x27\x84\xf9\x64\xc6\xae\x6b\xa7\x1a\xed\x86\xc9\x44\xf5\x23\x8d\x1d\x89\x00\x76\xf8\xf0\x84\x38\xc3\xb5\xe5\x85\x4f\xda\x8f\xe4\xe3\x19\x01\x85\x91\xdc\x49\xc1\x78\xdf\x87\xe2\x18\x65\x34\x8c\x4f\xf4\x7d\xe0\x40\xc8\x0f\xde\x94\x0e\xa2\x3c\x4d\x11\xa1\x6c\x80\xc6\xb7\xc1\x64\x8f\xce\xa0\xc4\x8d\x3d\xee\xc3\xa6\xc1\x27\xa4\x3f\xdc\x6a\x52\x58\x6f\xff\x37\xdc\xc3\x9e\x31\xf8\xc2\x2b\x86\xe1\x62\x33\x92\xbe\xf6\x8a\x47\x89\x39\x83\x12\x3e\x88\x24\x8e\x31\xdd\x06\xbf\x01\x00\x00\xff\xff\xd4\x30\xa6\x73\x54\x03\x00\x00")

func _000001_initialize_schema_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_initialize_schema_up_sql,
		"000001_initialize_schema.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"000001_initialize_schema.down.sql": _000001_initialize_schema_down_sql,
	"000001_initialize_schema.up.sql": _000001_initialize_schema_up_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_initialize_schema.down.sql": &_bintree_t{_000001_initialize_schema_down_sql, map[string]*_bintree_t{
	}},
	"000001_initialize_schema.up.sql": &_bintree_t{_000001_initialize_schema_up_sql, map[string]*_bintree_t{
	}},
}}
