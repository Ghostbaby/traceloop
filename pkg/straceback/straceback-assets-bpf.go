// Code generated by go-bindata.
// sources:
// ../dist/straceback-guess-bpf.o
// ../dist/straceback-main-bpf.o
// ../dist/straceback-tailcall-bpf.o
// DO NOT EDIT!

package straceback

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

var _stracebackGuessBpfO = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x59\x4d\x6c\x14\x47\x16\x7e\x3d\xe3\xf1\xd8\xe6\x67\x0c\xac\x17\xe3\x85\x55\xaf\x90\x57\x5e\xb4\x18\x9b\x5f\xaf\xb9\xd8\xd6\x0e\x3f\x8a\x21\x96\xf1\x28\x10\x89\x34\x4d\x4f\xdb\x9e\x30\xf6\xb4\xbb\xdb\xe0\x19\x47\x49\xa4\x28\x52\x22\x0e\x81\x48\x89\x94\x88\x03\x0e\x39\x58\xf9\x91\x9c\x4b\x4c\x4e\xf6\x29\x58\xc9\x21\x56\x94\x1f\x4e\x04\x45\x91\x82\x22\x22\x50\x14\x09\x1f\x22\x3a\xaa\xaa\x57\xd3\xdd\xd5\xdd\xf6\xa0\x44\xe1\x00\x2d\x41\x75\x7d\xf5\x5e\x7d\xaf\xde\xab\x7a\xf5\x7a\xfc\x42\xba\xf7\x40\x4c\x92\x80\x3f\x12\xdc\x03\xb7\xe7\x3e\xc6\x26\xf7\xbd\x0b\xff\x5f\x0f\x12\xcc\x6d\x64\xd8\x2c\x2a\x4d\x36\x5d\x70\x48\x3b\xf7\x0e\xeb\x27\x63\x00\x17\x1c\xc7\x69\x14\x26\x7d\x99\x72\x01\xcc\x55\xb3\x7e\x43\xf5\x27\xb4\x2d\xaa\xd8\x97\x66\xcb\x72\x6b\x89\x5c\xad\xab\xb7\x95\xf4\x6b\x58\xbf\x38\xc6\x5e\x66\x63\xc8\xbf\x6d\x89\xf2\x17\x4f\xe3\x3c\xb1\x8b\x10\x07\x80\x4c\xec\x43\xca\x97\x91\x3e\x80\x35\x12\xe1\x61\x7a\x27\x2f\xbe\x47\x5b\x62\xdf\xb8\xd1\x29\x93\xf7\xe6\x6c\x1d\x4c\x36\x2d\x38\x1c\x37\x75\x35\x4b\xde\x95\x21\xb3\x30\xd9\x34\x5f\xc6\x55\xeb\x0c\x95\xdf\x2e\xdb\xc3\x93\x4d\x57\x19\x5e\x0b\x50\x18\x1c\xb4\xc8\xbb\x6e\xcb\xf6\xe4\xdb\x33\xcc\x1f\x53\xe8\x0f\x09\x60\xc6\x71\x1c\x62\x2f\x51\x9e\x8d\x03\x2c\x26\xd8\xba\xaa\x39\x9f\xaa\xe9\x44\x36\xdb\x29\x37\x73\xbe\xd9\x24\xd3\xb7\x4a\x9f\x39\xdc\xdf\xd9\x3a\x80\x33\x1e\x3b\x5d\x7b\x0c\x9b\xdb\xb3\x1c\xff\x3f\x90\x7f\x58\xe0\xcf\x2a\x76\xce\xe5\x77\xe7\xd7\x55\x53\x21\xef\xda\x70\x2e\x1f\xee\x07\x2d\xcf\x79\xad\xd2\x97\x82\x9d\x8b\x2b\xda\xb3\x15\xed\x21\x86\x78\xed\xc9\x8d\xe8\x9d\x44\x56\x6e\xce\xd6\x85\xf3\x9e\xb5\x5d\xde\x85\x15\x79\x36\x21\x4f\x87\xc0\xa3\x67\x23\xe2\x8f\xf1\x30\x15\xcd\x0c\xe7\xaf\xcc\xdf\x3c\xde\xb7\x3c\xbc\x61\x71\x34\xcc\xc2\x04\xdd\xc3\x9e\xf8\xfb\xf9\x46\x2d\x91\x8f\xef\x8b\xe5\xe2\xdc\x95\x74\x79\x27\x4b\xb7\x98\x7e\xe9\x26\xb6\xd7\xb1\xc5\x38\x71\x3f\x96\xe6\xb1\xbd\x5a\x8e\x67\x8c\x9e\xf3\x19\x3c\x67\xf5\x78\x6e\x6b\xf0\xfc\xa3\x5d\xdb\x98\x5e\x51\x6d\x64\xe3\x9a\xcc\xc6\x3b\x71\x7e\x5c\x6f\x51\x6b\x41\xfc\xba\xcf\x1e\x55\xeb\xa0\xb8\xd6\xc9\xec\xcc\xc4\xf6\x30\xb9\xd2\x34\xca\x5d\xa6\xad\x56\xba\x44\xdb\x14\xa3\xf1\xad\x7f\x1a\xd7\x4f\x2c\x9b\x7b\xc9\xcd\x1f\x55\xd4\x2e\x66\x4f\xf1\x0a\x9b\x2f\xb5\x21\xa8\x7f\xd9\xa3\xcf\xf5\x86\x80\xc5\xf1\x1c\xb6\x8d\x49\x80\x5f\x1c\xc7\x01\x7c\x4e\x1a\x27\xc0\xef\x27\xb6\x9e\x04\xbc\x09\x5e\x7f\x84\xd9\x3b\xb3\x8c\xbd\x01\x5e\xc9\xcf\xbb\xa5\x3e\x85\xeb\x61\x71\x69\x88\xaf\x41\x3e\xb6\xce\xb0\xf5\x2d\x55\xb2\xbe\x00\x4f\x0c\x79\x58\x9e\xcd\xc4\xab\xca\xeb\xa5\xf1\x68\x62\xf8\x2c\xca\xbf\x81\x6d\x51\x6d\x89\xb4\x83\xc7\xa9\x6a\x19\x3b\xf8\xfc\x8d\x31\xc1\x1e\x99\x0d\xa8\x53\x2c\x8e\x9c\x7f\x48\x42\x7d\x6c\x1b\xa4\x9f\x1d\xae\x7f\xdf\x71\xee\x94\xf5\xff\xf5\x53\xe0\xbe\x5a\xaa\xe8\xbe\xfa\x4f\x79\xdd\x75\x40\xce\xef\x57\xe5\xf3\x31\x21\x37\x67\x35\xcc\x77\xf4\x5e\x91\x0d\x93\xbc\x17\x86\x0e\x67\xbd\x79\x45\x1e\xb7\xe9\x3d\x31\x6a\x09\xe7\x3c\x97\xd7\xe8\x9a\xf2\xf9\xfd\xe5\x7b\x45\x02\x18\x2d\x98\x23\x0c\x97\x6d\x7e\xfe\xd8\xcd\x04\xa0\x62\x3c\xc2\xf6\xd3\xbf\xc1\x3d\xf7\xaa\x86\xfe\x1f\x73\xfd\x29\xae\x73\x35\xdd\xaf\xd7\x1c\x6f\xdc\x8a\x57\x2e\x47\x9e\x93\x4b\x15\xc4\xef\xa4\xb1\xbb\xec\x2f\xc9\x77\x2e\x76\x61\xfc\x98\xfd\x4f\x49\x00\x29\x1a\xc7\xe9\xd0\xb8\xcc\x5d\xc1\x7e\x1c\xf7\x0d\xc6\x43\x8c\xd7\x6c\x95\xbb\x1e\xb2\x44\xeb\xd9\x05\x4f\x3c\xec\x02\x84\xdc\x27\xfe\x78\xb8\x7e\x57\xf3\x79\x2a\xaf\xa9\xb6\x5e\xf6\xfb\xbb\xcc\xbe\xe5\xfc\xbe\xc9\xeb\x77\xcc\x33\x2b\xf9\xdd\xdd\x4f\x0b\xde\xfd\x20\x33\xfe\xbc\xef\xde\x2d\xc8\x3a\xbd\x27\x74\x4d\xb6\xbd\xf6\x0e\xaa\xb9\x3c\xc5\xb3\xee\x3e\x09\xb3\xaf\xc9\x63\x5f\x02\x5e\x65\xeb\x99\xba\x84\xe7\xe8\xd6\x8a\xfe\x9f\x09\x39\x27\xa2\xdf\x13\xf0\x3c\x9b\xe7\xef\x50\x8e\x3f\x78\xe2\xef\x9d\xff\xe6\x72\xe7\x0e\xeb\x9f\x86\x24\xe6\x6d\xf4\x63\x46\xfa\x27\xad\xef\x78\x5d\x58\x1c\xc3\xfc\x5a\x1f\xdc\xa7\xde\x7c\x37\x17\x77\xe7\x7f\xb0\xfc\xba\xe4\xcf\xaf\x63\xd1\xf9\xf5\xee\x1f\xca\xaf\x77\x1f\x2c\xbf\x8e\x45\xe7\xd7\xdb\x7f\x4a\x7e\xbd\xfd\x97\xe4\xd7\x06\xf8\x81\xed\x43\x8c\xcf\x9c\xea\xda\x25\xca\xb3\x3c\x75\x83\xca\xbf\xe8\x1f\x26\xae\x27\x4f\xea\x3e\xb6\x08\xa4\x52\x28\x40\xb8\x6a\xf0\x5b\x26\xec\x7b\xe7\x51\x7d\x24\xdc\x1f\x55\x41\x9f\x3e\xd2\x4f\xfc\xb1\x5f\x42\x9f\xc7\x7e\x09\x7f\x78\x7e\x69\x03\x74\x0e\xaf\x43\x6c\xf2\x35\x79\x5a\xd5\xce\x3c\x64\x03\x1f\xd2\x43\x6a\x58\x52\xbf\x92\x1a\x97\xd4\xb7\xa4\xe6\x22\xf5\x16\xa9\x91\x49\x7d\x4c\x6a\xe7\x3a\x8f\x3c\xa9\xbd\x48\xdd\xe5\xca\xd9\x05\xaf\x08\xa9\x75\x48\x9d\x43\x6a\x21\x52\x07\x91\x5a\x89\x4c\xcc\x05\x0a\x83\x83\x96\x6e\xcb\x36\xf9\x76\xde\x2e\xdb\xc3\xa6\xae\x66\x95\x21\xb3\x30\x6e\x74\xd2\x69\xfc\xe3\x06\x8d\x4e\xb6\xd3\x65\xf0\x8f\x6b\x79\x5d\x35\x15\x6d\x38\x97\xcf\x2a\x76\xce\x2b\x07\xa1\x7c\x67\xed\xdc\x88\xee\x93\x0a\xe3\x33\x15\xcd\xe4\xa4\xfe\xf1\x51\xcb\x30\x0b\x13\x45\x36\x74\xb0\xaf\x17\x00\xcb\x1a\x90\x4a\xfd\x50\xf3\xdc\x2a\x89\x7c\xdb\x13\x0c\x3f\x27\xa1\x05\x99\xfe\x06\x00\xff\x45\xec\x54\x6c\xe5\xb8\xdc\xa1\x97\x60\x1c\x16\x6b\xfd\xf8\x8f\x88\xdf\x14\xf0\xef\x10\x6f\x4b\xf8\xf1\x69\xc4\x6b\xaa\xfd\xf8\x5b\x88\xb7\x08\xf8\x3d\xc4\x2f\x08\xf8\x17\x88\x1f\xaa\xf1\xe3\xd7\x10\x6f\x14\xe4\x5f\xe7\x76\x26\xfd\xf8\xb7\x14\x4f\x40\x9f\x70\xc9\xcf\x20\xde\x25\xe0\x53\x88\x9f\x12\xf0\xf3\x88\x2f\x84\x14\x0b\x71\x88\x07\x41\x8a\x27\x02\x58\x4e\x22\xf1\x4a\x06\xf0\x2a\xc4\x85\xe5\x52\xfb\x08\x2e\x0b\xf8\x16\xc4\x5b\x04\xfc\x2c\xe2\xf5\x02\xfe\x0c\xe2\x8d\x02\xfe\x35\xdd\x3f\x6b\x03\xf6\x3c\x4d\x71\xd1\x1a\x80\x3c\xc5\x57\x07\xf0\x8f\x29\x5e\x17\xc0\x0f\x50\x7c\x55\x00\x1f\xa7\x78\xd0\x6f\x49\x8a\x07\xfd\xf6\x39\xc5\x6b\x03\x38\x5f\x3f\x09\xcb\x46\x00\xb8\x5e\xed\xf6\xb7\x00\xc0\x7c\xd2\xed\x93\xef\xb1\x57\x6a\xfc\xe3\x1d\xb5\xfe\x71\xbe\x04\x3e\x9f\x38\x3f\x96\xe7\x65\xfd\x2e\x4f\x9f\x7c\x4f\x6d\x46\xf9\x3a\xfc\xdd\xba\xcd\xd3\xa7\xb5\x67\xab\xad\x4f\xd8\x40\xcf\xbd\x51\xc8\x8d\xda\x8a\x62\x15\x2d\x45\x9f\xc8\xd9\xd0\x6a\xea\x79\x77\x60\x87\xa9\x9e\x23\x63\x24\x9b\x59\x3b\xca\x42\x23\xaa\x61\xed\x20\x69\x4e\x21\x03\x1e\x6c\x68\x5c\xb7\x2c\xc5\xb2\x55\x7b\xdc\x62\x08\x57\x0e\xb2\x8d\xda\xba\xb9\x32\x1d\x95\x12\xf9\x5c\x90\x66\x64\x65\x44\x35\x40\xc9\xe7\x34\x7d\xd4\xd2\xe9\x94\xad\xfa\xb0\x32\x68\xaa\x23\x3a\xb4\x5a\xb6\x69\xab\xa7\xa1\xd5\x2a\x8e\xd0\xd6\x2c\x64\x55\x5b\x85\xbe\xee\xfe\xee\x23\x4a\x5f\xff\x93\x3d\x69\xa5\x7b\x40\x49\x1f\x3f\x3c\xa0\x1c\xe9\x3e\xf6\x04\x64\x8e\xa5\x95\xee\xfe\x83\xca\xe1\xa3\xff\x4f\x1f\x57\xba\x8f\x29\x4c\xb2\x37\x7d\xf4\xe0\xc0\xa1\x0a\x44\xe8\x68\x7f\x7a\x20\x14\x3f\x9a\xe9\xed\x55\x7a\x4e\x0c\xa4\x39\xd8\xdb\xd3\xd3\xae\xfc\x8f\x34\x6d\xbc\x69\xef\x60\xe8\x3e\xec\xee\x63\xdd\xbd\xdc\x74\xb2\xa2\xf6\xd6\xf6\xbd\x6c\x78\xe7\x1e\x6c\x77\x33\xb1\xf6\x5d\xa8\xc6\xdb\x9d\xd8\xb6\xa3\x5c\x5b\x60\xfb\x3e\xf0\x73\x1b\x6b\x2e\xf1\x31\xf0\xef\x30\x8e\x90\x9f\xc4\x74\x25\xe1\x3f\x21\x7d\xe2\xdf\x74\x82\x4f\x95\xd0\x7f\x6d\x05\xfd\x16\x21\x05\x88\x19\xe4\x3c\x40\xc8\x29\x06\x98\xc7\xdf\x07\x4e\x61\x7f\x1d\xae\x93\xeb\xf3\x7c\x26\x47\xf0\x0f\x23\xaf\x78\xd7\x89\xfc\x9b\x23\xf8\x65\xf4\x5f\x9b\x87\x3f\x11\xc2\xff\x1b\xf2\x8b\x57\xea\xc2\x6a\xbf\x7e\x14\x7f\x06\xf5\xc5\x18\x2e\xa1\xbe\xf8\x5d\x2c\xfa\xff\xd3\x08\xfd\xfa\xb5\x95\xe9\x7f\x14\xa1\xdf\x92\xaa\x4c\xbf\x3b\x42\xbf\xab\xbe\x32\xfd\xa1\x08\xfd\xe3\xeb\x2a\xd3\xe7\xbf\x1f\xee\x14\x70\x63\x3d\x6b\xf7\x0b\xfa\xe5\x7a\x09\xdb\x6f\x22\xf8\xdf\xdf\x10\xce\x27\x9e\x9f\xef\x23\xe2\x3f\x83\xfa\xe2\x39\x12\xe3\x7f\x23\x62\xff\xf5\xe1\xfe\xe3\xf7\xfc\x3a\xbc\x2f\xc4\xfd\xf7\x6b\x08\x37\x79\xee\x22\xff\xa2\xe7\x3e\xaa\xf7\xe8\xf3\x7b\xff\xf7\x00\x00\x00\xff\xff\x15\x56\xe7\xfe\xf0\x1d\x00\x00")

func stracebackGuessBpfOBytes() ([]byte, error) {
	return bindataRead(
		_stracebackGuessBpfO,
		"straceback-guess-bpf.o",
	)
}

func stracebackGuessBpfO() (*asset, error) {
	bytes, err := stracebackGuessBpfOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "straceback-guess-bpf.o", size: 7664, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stracebackMainBpfO = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x96\x3b\x6f\x13\x41\x10\xc7\xff\x67\xc7\xc4\x79\x80\x83\x44\x90\x25\x52\xa4\xa4\x40\x8e\x49\x45\x99\xb3\xb8\x84\x08\x27\x58\x8e\x83\x02\xcd\x6a\x63\x0e\xb0\xf0\x4b\x77\x07\x24\x3c\x04\x12\x8a\x84\xe8\x28\xa8\x41\x88\x8f\x80\x04\x54\xa1\x00\x89\x06\x44\x81\x50\x44\x45\x41\x49\x41\x81\x44\x0a\x84\xd1\xed\xcd\x61\x67\xee\x0e\x5b\x34\x2e\xc8\x48\xc9\x66\x7e\x3b\xb3\xb3\x37\x3b\x3b\xd9\x5b\x46\x7e\x36\xa6\x69\xf0\x45\xc3\x0f\xb4\xb5\xb6\x6c\x0d\xb5\xff\x9e\xa1\xdf\xfb\xa1\x61\xf3\xa0\xc7\x36\x00\x14\x00\x5c\x1f\xde\x6e\xb9\xfa\xe6\x63\x8f\x0f\xc6\x80\xed\x56\xab\x95\x66\x8b\x6e\xa8\x58\xc0\x38\x12\x4a\x97\x71\x8f\x6f\x4a\x6f\x4c\xc7\x82\xf6\xa3\x00\x9e\x91\xfe\x80\xc6\x7e\xc7\xbf\xbd\xd3\xec\x5b\xcb\x93\xd4\x2f\x1a\x09\xa4\x52\x64\xe0\xc6\x4c\x02\x18\x08\xfa\xfe\xd7\x12\xa7\x9c\xec\xe6\x65\xa7\xec\xe6\x25\x5c\xfc\x7b\x94\x05\x25\x07\x80\x7b\x65\x6d\xc7\x92\x65\x73\x55\x96\x2f\xf5\x79\x83\x7d\x92\xb9\x42\x1e\x6e\xef\x19\x23\x5d\xbb\x56\x44\xf2\xc6\x88\x76\x08\x80\xcb\xd2\xc4\xfd\x5a\x3a\x00\xe0\x08\x63\x7f\x93\x7b\xaa\x8f\xc6\xd1\x64\xfc\x8e\xe2\x89\x00\xf7\xac\xe3\xa1\x6b\xc5\xa9\xf7\x76\x8a\xd4\xdc\x7d\x0e\x06\xf8\x4f\x78\x3c\xc9\xf8\x31\xb2\x9f\x64\x7c\x9c\xf8\x61\xc6\xeb\xc4\xc7\x18\x3f\x4d\x3c\xcd\xf8\x5b\x15\x77\x34\xb0\x9f\x57\x8a\xef\x0d\xf0\x17\x8a\xf3\x5d\x02\x67\x15\x1f\x09\xf0\x27\x8a\x0f\x05\xf8\xac\xe2\xc3\x01\x5e\x56\x3c\x98\xcf\x41\xc5\x83\xf9\xf4\xbf\xdf\x3d\x1e\x77\xb7\x2b\x1d\x7a\xaa\xcb\xbc\x9b\xa3\x09\xd2\x87\xa9\x17\x65\x3b\x74\x75\xed\x32\x8e\xb9\xe6\x40\xdd\xba\x66\xa3\x52\x77\x84\xb0\xd7\x6d\x61\xae\x55\x1c\x64\x2c\xb3\xda\x9e\x98\xb2\xe4\x55\x77\xae\x2c\xab\x55\x7b\xea\x8f\x51\x4d\x36\xed\x29\x47\x56\xaa\xc2\x9d\xe8\x60\xbe\x69\x70\xed\xba\x63\x5a\xdd\x17\x57\x56\x7c\xf5\x36\x2c\x5f\xb0\x1a\x97\x9b\xa2\x26\x9b\x10\x57\x4c\xcb\xae\x34\xea\x10\xd5\x4a\xd9\xac\xdb\xa6\x5a\x3c\x63\x5e\x14\xe7\x2d\x59\x33\x91\xb1\x1d\xcb\x91\xab\xc8\xd8\xeb\x35\x35\x5a\x8d\x73\xd2\x91\x28\xe8\x45\x7d\x41\x14\x8a\xa7\x72\x86\xd0\x4b\xc2\x58\x99\x2f\x89\x05\x7d\xe9\x24\x96\x97\x0c\xa1\x17\xe7\xc4\xfc\xe2\x71\x63\x45\xe8\x4b\xc2\xb3\xcc\x1b\x8b\x73\xa5\x13\x3d\x98\xa8\xd9\xa2\x51\x0a\xe5\x8b\xcb\xf9\xbc\xc8\x9d\x29\x19\x3e\xcc\xe7\x72\x47\xc5\xb4\x3b\x64\xc5\x74\xa0\x00\xfe\x51\xbe\xd0\x79\x73\x99\xa1\xd2\xbe\xcf\xde\x53\xfc\xcd\xa8\xd1\xcf\x1e\xee\x1f\x11\x6f\x80\xe9\x37\xbb\xf8\xf3\x5e\xc5\x6f\xdc\x3a\x10\x72\xab\x80\xcf\xd4\x5a\xfc\x7e\x31\x46\xdf\x99\xec\xd0\xfd\xf9\xb0\xf8\x2f\x7b\x8c\x3f\x11\x11\xdf\x37\xec\x8c\x9f\x08\x89\xff\x9d\xe2\xb3\x67\x28\x66\x28\xd1\xd9\x2e\xf1\x9f\x92\x3f\x3f\xc3\x26\xf9\xf3\xf7\x30\xcf\xff\xc3\x08\xff\xbb\xb1\xde\xfc\xf5\x08\xff\x47\xf1\xde\xfc\x97\x23\xfc\x9f\x0f\xf4\xe6\xff\x2e\xc2\xff\x7d\x22\xdc\x9e\xd7\xef\xeb\x08\xff\x0f\x11\xfe\x5c\xff\x14\x71\x7e\x5b\xe4\xcf\xef\x01\x3f\xbf\x8f\x11\xf5\x33\x19\x52\x3f\xfb\x42\xea\xe7\x6b\x48\x6c\x57\xd2\x54\xd0\x6f\xe8\x83\x35\x56\x7f\xfe\xff\xbf\xdf\x01\x00\x00\xff\xff\x6e\x9b\x5a\x47\x18\x0e\x00\x00")

func stracebackMainBpfOBytes() ([]byte, error) {
	return bindataRead(
		_stracebackMainBpfO,
		"straceback-main-bpf.o",
	)
}

func stracebackMainBpfO() (*asset, error) {
	bytes, err := stracebackMainBpfOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "straceback-main-bpf.o", size: 3608, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stracebackTailcallBpfO = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5c\x7d\x6c\x9b\xc7\x79\x3f\x52\xa2\x49\xd9\x8e\x45\x7f\xc8\xa6\xe9\xd8\xa6\xed\xb8\xe3\xe4\xd9\xe1\xcb\x6f\x13\xf9\x90\x9c\x2a\x51\x63\xc7\xd1\xfc\x31\x73\x5d\x5d\x4a\x71\x98\xc8\x4b\xec\x68\xa2\x56\x5b\xd6\x86\x08\x68\xbb\xaa\x6e\x37\x28\x4d\x56\xa8\x6d\xba\x38\x69\x3a\x68\xd8\x96\xa9\x43\x06\xb9\x45\x37\x69\x6b\x50\x38\x68\x10\x68\x59\xd0\x09\x68\xb3\xea\x8f\x66\x15\xda\xb5\x15\xb0\xb5\xd1\x80\xc2\x1c\xde\xbb\xdf\xf3\x92\xef\x73\x24\xa5\x37\xd6\xfc\x55\x13\xb0\xef\xbd\xdf\x3d\x5f\xf7\xdc\xdd\xef\x8e\x7c\xcf\x7e\xba\xed\xc0\xfd\x6e\x97\x4b\xd0\xc7\x25\xde\x13\xa5\x5a\xe9\x33\xd2\x5d\x7a\x6e\xc1\xdf\x7e\xe1\x12\x13\xeb\x15\xf6\x49\x21\x84\x4f\x08\x31\xe1\x2b\xd5\x57\x99\xf5\x06\x55\x1f\xf8\xe2\x7c\x91\x70\x8f\x10\xa2\xbf\x4b\x09\x0e\x04\xe7\x24\x3e\xb0\x7c\x58\x96\xe3\x5e\xc8\x9f\x1b\x91\xf5\x27\xce\x5d\x50\xe5\xf9\x97\x64\x79\x22\xf8\xb2\x92\xff\xe2\xa8\x2a\xcf\xcd\xa2\x9c\x41\x39\x8d\x72\x0a\xe5\x25\x94\x93\x28\x2f\xa2\x1c\x43\x39\x84\x72\x10\xe5\x59\x94\x3d\x28\xbb\x51\x76\xca\x92\xfa\x37\xb0\x3c\xab\xea\x2f\xa9\xba\xd7\x2d\xc4\x5c\xb1\x58\x0c\xb0\xe4\x7d\x52\xe6\xb4\x94\x87\xa6\x86\xef\xd4\x49\xbd\x17\xa1\xe7\x12\x62\xac\x58\x2c\x8e\xbb\x85\xf0\x43\xde\x2f\xf3\xe3\x97\xed\x27\xbf\xa2\xfa\x3f\x10\x44\xfc\x41\x15\x47\xff\xb0\xd2\xef\xef\x0a\x20\x8f\xe8\x67\x50\xc5\xdb\xff\x9c\x0f\xed\x21\xb4\x23\x1f\x41\xd5\xaf\xfe\xcf\xf9\xd1\x1e\x46\x3b\xf2\x16\x54\xfd\x1f\x77\xa9\xb8\xc7\x3d\xaa\x3c\xde\xa3\x3a\x66\xd6\xcd\x4f\x63\x02\xfe\x9f\x55\xfe\xfb\x3b\x23\xc8\x0b\xc6\x61\xb9\xca\xe7\x38\xfa\x7f\xbc\x0f\xfa\xc8\x4b\x63\x07\xf4\x9f\x53\xf1\xf5\xe7\xd3\x4a\xff\xc3\x18\xcf\x0f\xab\x71\xb1\xfc\xf7\x32\xff\x28\x2d\xfb\x85\xca\xf6\xcd\xbc\xda\xfc\xbb\xd1\xee\xa6\xf8\x55\xff\x2d\x7d\x8c\x5f\xe3\x36\x55\x3e\x11\x54\xf9\x9f\x78\x19\xe3\x55\x2f\xc4\x70\xb1\x58\x9c\xe8\x52\xf5\x80\xdb\x3e\xde\x81\x3a\x21\x8a\xc5\x62\x91\xea\x66\xbc\x9d\x18\xd7\x8d\x6c\xbe\xcc\x9b\x76\xbe\x82\x7a\x9d\x10\xd9\x0a\xf3\x67\xbc\x5e\x95\xa6\xbe\x9b\xcd\x9b\xa9\xcb\xc5\xe2\x40\x70\xf4\xb2\xec\xc7\x33\x0a\x6f\x72\xbd\x60\xc5\x45\x71\x14\x8b\xc5\xc6\xdd\xdb\x9e\xb7\xf0\xcb\x65\xf8\xe6\x6d\x5f\xb0\xe5\x65\xa0\xb9\x43\xe5\xbf\xb9\x1d\x65\x0b\xca\x34\xca\x08\xca\x30\xca\x10\xca\x00\x4a\x3f\x4a\x1f\x4a\x81\x72\xfe\xb2\x2a\xe7\x50\xce\xa2\x9c\x41\x39\x8d\x72\x0a\xe5\x25\x55\x9e\x1f\x93\xe5\x78\x9d\xea\xff\x40\xe6\x22\xda\x27\x65\x59\x40\xbc\xe5\x79\x97\xfd\x32\xd6\x5b\xfd\x9d\x2b\xcf\xc3\x9a\x55\x12\x9f\xd8\xa0\xfa\x7b\xcc\x2b\x44\x23\xc6\x7d\x50\x08\xb1\xd3\xbb\x5c\xf2\x13\x1f\xa7\x6a\xeb\x9a\xe6\x4b\x93\x90\xcb\x5a\x3c\xee\x15\xf2\xa9\xb1\x07\xf3\xcb\xe5\xb3\xf2\x6b\xda\xdf\xbd\xa6\x34\xcf\xcc\xfa\xc4\x5a\x25\xe7\x11\x6a\x00\x0a\xcd\x0f\x5b\xeb\xef\x69\xd9\x5f\x95\x87\x63\x2e\x21\x13\x99\x77\xad\x92\xf6\xf3\xae\x9d\xd2\x7f\x93\xeb\x77\x30\xee\x2e\x59\xf7\x88\x87\x94\xdf\x13\x6a\x7d\x9b\x7c\x6a\xe2\xfd\x2f\xaa\x79\x42\xeb\xc1\x8c\xdf\x9c\x5a\x8f\x0b\x21\xcc\x15\x78\x06\x65\x93\x38\xa8\xe2\x38\xa7\xf2\xea\x11\x07\x64\x3d\xef\x5a\x21\xf3\xd2\xe4\xda\x26\xf5\x4c\x7f\x1e\xd9\x7e\xaf\xcd\x5f\xb9\x1f\x4f\x4d\x3f\xad\xe8\xf7\x3d\x88\x7f\x8b\x58\x26\xcb\x9d\xc2\x0b\xfb\x3e\xd9\x1e\xab\x6a\xdf\x57\xd3\x7e\x02\xf6\xa3\xb0\xbf\x5d\x66\xd8\xb4\x5b\x27\xf1\x70\x55\xbb\x75\x35\xed\x36\xc3\xee\x6f\x54\xd5\xaf\xaf\xa9\xbf\x1d\xfa\x5b\xab\xea\x2f\xab\xa9\x1f\x84\xfe\x86\xaa\xfa\xde\x9a\xfa\x6b\xa0\xdf\x58\x55\xdf\x5d\x53\x7f\x39\xf4\xbd\x55\xf5\x07\x6b\xea\xbb\xad\xf9\x6f\xce\xc3\x42\xb0\x43\xe3\xd7\xb1\xcb\xce\xf8\x75\xb4\x8c\x5f\xfb\x9f\xf1\x61\xbc\x6f\x5e\x1e\x74\x49\x5e\x98\xac\xc8\x8b\xbf\xbe\x7c\x78\x07\xf8\xf0\x28\xc6\x9f\xf8\xf0\x00\xe6\xa9\x3a\xa7\x38\xe7\x43\xc5\xa7\x25\x3e\xdc\x2f\xeb\x25\x3e\x0c\x31\x3e\xbc\xc7\xe6\x6f\xf1\x7c\xd8\x82\x7e\xdf\x8d\xf8\x37\x83\x0f\xef\x90\x2b\xed\xa8\x2b\x26\xf9\xae\x92\xdd\xa5\xe1\x41\xdd\xae\x33\x1e\xd4\xf5\x9d\xf1\xa0\xae\xef\x8c\x07\x75\x7d\x67\x3c\xa8\xeb\x3b\xe3\x41\x5d\xff\xda\xf2\xa0\x1f\xe3\xfd\x17\x96\xbe\x9d\x07\xbf\x64\xe1\x76\x1e\x1c\xb1\xd6\xa7\xb8\xce\x79\xd0\x5d\xc6\x83\xb7\xf8\x6f\x21\xfe\x53\xdf\xb3\xae\x1e\xff\x85\xb4\xf5\xb0\x34\xfc\xa7\xdb\x5d\x1a\xfe\xd3\xed\x3a\xe3\x3f\x5d\xdf\x19\xff\xe9\xfa\xce\xf8\x4f\xd7\x77\xc6\x7f\xba\xbe\x33\xfe\xd3\xf5\xaf\x2d\xff\x05\x30\xde\x37\xef\x39\xb0\xee\xd6\x39\xf0\x7d\xf0\xa0\x5a\xef\x57\x8f\x07\xc3\xda\xba\x58\x1a\x1e\xd4\xed\x2e\x0d\x0f\xea\x76\x9d\xf1\xa0\xae\xef\x8c\x07\x75\x7d\x67\x3c\xa8\xeb\x3b\xe3\x41\x5d\xdf\x19\x0f\xea\xfa\xd7\x96\x07\x43\x18\xef\x9b\x97\x07\xeb\x6f\xf1\xe0\xfb\xe0\x41\xf5\x7b\xfd\xd5\xe3\xc1\x88\xb6\x2e\x96\x86\x07\x75\xbb\x4b\xc3\x83\xba\x5d\x67\x3c\xa8\xeb\x3b\xe3\x41\x5d\xdf\x19\x0f\xea\xfa\xce\x78\x50\xd7\x77\xc6\x83\xba\xfe\xb5\xe5\xc1\x30\xc6\xfb\xf7\x2c\x7d\x3b\x0f\x66\x2d\xdc\xce\x83\x47\xad\xf5\x29\xae\x73\x1e\xf4\xbc\x2f\x1e\xdc\x69\xf5\xfb\x7a\xe6\xc1\x81\x20\xe7\xb9\x3b\xc0\x73\x19\xf0\x5c\x13\xc6\x97\x78\x6e\x25\xe6\xa1\x7a\xaf\xe8\x9c\xe7\x6e\x93\x7a\x25\x9e\x5b\xa1\xea\xe0\x5f\xe2\x5d\xb2\x7f\xbd\xcc\x73\x7a\xff\xf9\xe7\x28\xab\xf3\xb2\x7a\xcf\x5a\x29\xfe\xda\xbc\xfc\x63\xe4\xe3\x47\x45\x51\x91\x97\xdf\x29\xfa\xaa\xd8\xad\xcd\xcb\x3f\x80\xdd\xef\xc3\x2e\xe7\xe5\xb7\xaa\xc6\x5b\x9b\x97\xdf\x86\xdd\x7f\xad\xaa\x5f\x9b\x97\xdf\x80\xfe\xeb\x55\xf5\x6b\xf3\xf2\x6b\xd0\xff\xe7\xaa\xfa\xb5\x79\xf9\x9b\xd0\xff\x7a\x55\xfd\xda\xbc\xfc\x2a\xf4\xbf\xa6\xe6\x19\xbf\xaf\xe1\x2d\xd5\x57\x95\xdf\x6f\x38\xbf\xc8\xfb\x1a\x58\x47\x03\xcd\x23\x56\xdd\xfc\x3c\xd1\x8c\x7b\x1b\xe7\xd8\xbd\x8d\xf3\xb8\xb7\xd1\x8c\xf7\xfc\xcd\xb8\x2f\xd0\x8c\xfb\x07\xcd\xb8\xa7\xd0\x8c\xfb\x0c\xcd\xb8\xf7\xd0\x8c\x7b\x1b\xcd\xea\xde\x06\xdd\x8f\xa0\x7b\x11\x8b\xbe\x87\x81\xfe\x36\x79\x55\x22\xf8\x3d\x8c\x81\xa0\xe2\xd9\xc5\xf2\x1a\xdd\xeb\x38\xda\xb0\x0a\xe3\x32\x6d\x8d\x4b\xf9\x7d\x8e\xa5\xba\x3f\xc0\xd7\x77\x79\xfc\x3e\x19\xbf\xda\x2f\x26\x9e\x29\xc3\x65\x9e\xd4\xbe\xd1\xdf\x0b\x3f\xf6\xfe\xf8\x77\x6f\xfb\xaa\xf2\x57\x87\xf1\xcf\xe0\x5e\x4c\x06\xf7\x62\x32\xb8\x17\x93\xc1\xbd\x98\x0c\xee\xc5\x64\x3a\x51\x66\x51\x62\x7f\xcc\x60\x7f\xcc\x60\x7f\xcc\x60\x7f\xcc\x60\x7f\xcc\x60\x7f\xcc\x60\x7f\xcc\x60\x7f\xcc\x60\x7f\xcc\x60\x7f\xcc\xa8\xfd\x91\xee\x37\x0c\x04\xd5\xbe\x68\xf2\xa7\x5b\xcd\x47\xfc\xfe\xab\xf6\xcf\x02\xe2\x0e\xd4\xdb\xf7\xb9\x09\xdc\xe7\xc8\xa9\xd7\xce\x62\x73\xab\xcf\xca\x77\xf9\x7e\x7f\xdc\xa0\x7b\x33\x78\x7f\x8e\xf1\x39\xb5\x4d\x9d\x8b\x68\x1f\xf3\xd0\xfe\x90\xf9\x94\xed\x3c\xee\x11\xab\x2c\xbb\xb6\xfd\x74\xad\xd7\xd2\x37\xf7\x87\x63\x6e\xb5\x9f\xee\x74\xd7\xcb\xf5\xf5\x38\x78\xac\xff\x45\x35\x4e\x74\xff\xa4\x7f\xb5\xb0\xe2\xa8\xf8\xfd\x60\x1d\xf2\xf2\xa0\xca\xd3\x84\x3a\xc6\x49\xfb\x72\x9f\x74\x37\xaa\x7d\xd2\xbd\x49\xed\x93\xee\xe3\xb2\xbd\xc9\x4d\xfb\xa4\x3a\xbf\x93\xdf\x7e\x7c\xdf\x18\xf7\xd1\x7e\x39\xb3\xc8\xfd\xf2\x88\xca\xc7\xf9\x21\xf0\xcd\x21\x59\xcf\xbb\x03\x6a\xff\x71\xef\x50\xfb\x8f\x9b\xf6\x9f\xbb\x2a\xfa\x2d\xf7\x57\x7b\x1f\xba\x1f\xe3\x70\x1f\xfa\x13\x52\xfb\x85\x9b\xf6\x8b\xc8\x82\xf6\x6b\xef\x1b\x19\xd8\x4f\xc1\x3e\xf6\x39\xb7\x4b\xf2\xb5\x47\xdc\xb1\xa0\xfd\xda\xbc\xbe\x07\xf6\x77\x2d\x68\xa7\xf6\xfe\xb4\x13\x76\xb6\x2f\x68\xa7\xf6\x3e\x71\x3b\xec\x6c\x5c\xd0\x4e\xed\xfd\x6e\x2d\xec\xe0\x5e\xc0\x4b\xb0\xb3\x43\xd9\x39\xea\xf2\xaa\xef\x6d\x65\xf6\x6a\x9f\x0b\x1a\x60\x6f\x99\x16\x87\xb3\x73\xd6\x90\x76\xce\x9a\xbd\x92\xef\x13\xbd\xe0\x0f\xc6\x33\xbb\xf7\xb9\x10\xef\x97\x95\xde\x0d\xc4\xa7\x72\x1f\x0f\xce\x57\xe4\x57\xe2\x55\xce\xa3\x66\xff\x6d\xdf\x97\x5a\xbd\x5a\xfe\xca\x79\x95\xf8\xb3\x10\xe4\xbc\xa9\x88\x8b\x78\xf7\x94\xb1\xd2\x26\xef\xc1\xf7\xcb\xab\xc6\xab\x6b\x69\xdc\xaa\xf1\x2a\x7e\x67\x71\x6f\x04\xaf\x2a\xfe\x2b\xf1\xaa\xe2\xa5\xfe\x67\xb1\x8f\x38\xe6\x53\xf5\x3b\x4d\x89\x4f\x1f\x94\xf5\xbc\x7b\x3d\xf8\x74\x2b\xe3\xd3\x84\xcd\xdf\xe2\x79\xf4\x5e\xe4\xf7\x2e\xc4\x7f\x3b\xe3\xd1\xe6\xaa\x76\x6b\xf3\x67\x1c\x76\x0d\xd8\xdd\x28\x57\xf0\x51\x77\x48\xf2\x62\x25\x7b\xb5\xf9\x92\xfd\xbe\x52\x41\xdf\xd1\xef\x2b\x15\xf4\x6b\xf3\x23\xfb\x7d\xa5\x82\x7e\x6d\x5e\x64\xbf\xaf\x3c\x4b\xef\xe7\x75\x3e\xac\xcd\x6b\xd5\xf9\xb0\x36\x8f\xfe\x7f\xf3\xa1\xea\x4f\xc0\x67\x53\xf3\xef\x0e\xab\xdf\x97\x6f\x24\x1e\x74\x97\xf1\x20\xe7\xbf\x85\xce\x95\x01\x0f\xe3\xc3\xdf\x56\xfc\x74\xbc\x75\xb9\x95\x7f\x51\xc6\x7f\xfc\x77\xa8\xeb\x96\x0f\xc1\x7f\xf4\x7b\x95\xce\x87\x2b\xc1\x87\x1b\xc0\x87\xea\xfc\x57\xe2\xc3\x56\xcc\xfb\x80\x36\x6f\x6b\xf3\xe0\x7e\xf4\xb7\x5d\x96\x0b\xf3\x9f\x6e\x7f\x69\xf8\x4f\xb7\x7b\x65\xfc\xa7\xdb\x73\xc6\x7f\xba\xbe\x33\xfe\xd3\xf5\x9d\xf1\x9f\xae\xef\x8c\xff\xe8\xfd\xfc\xcd\xc2\x7f\xaa\x3f\xbb\xc3\xea\x3d\xdb\x8d\xc4\x77\x75\x8b\x38\xf7\xdd\xe2\xbd\x2b\x3d\x07\xaa\xf7\xb0\x57\xef\x1c\x18\xd2\xd6\xc7\xd2\xf0\xa0\x6e\xf7\xca\x78\x50\xb7\xe7\x8c\x07\x75\x7d\x67\x3c\xa8\xeb\x3b\xe3\x41\x5d\xdf\x19\x0f\xd2\xfb\xf9\x9b\x85\x07\x55\x7f\xf4\x73\xe0\x8d\xf7\x7d\xb8\xfe\x16\x2f\x5e\x85\xf3\x60\x58\x9b\xbf\x4b\x7b\x1e\xd4\xed\x2f\x0d\x0f\xea\x76\xaf\x8c\x07\x75\x7b\xce\x78\x50\xd7\x77\xc6\x83\xba\xbe\x33\x1e\xd4\xf5\x9d\xf1\x20\xbd\x9f\xbf\x59\x78\x50\xf5\xe7\x46\x3c\x0f\x7a\x6e\xf1\xde\x55\x38\x0f\xf2\xfb\x57\x0e\xcf\x83\xe7\x9c\x9e\x07\x23\xda\xfa\x58\x1a\x1e\xd4\xed\x5e\x19\x0f\xea\xf6\x9c\xf1\xa0\xae\xef\x8c\x07\x75\x7d\x67\x3c\xa8\xeb\x3b\xe3\x41\xa5\x7f\xb3\xf0\xe0\x62\xde\xa3\xab\xf9\x14\xf8\x95\x59\x1f\xb4\x37\x9b\x4b\xda\xfc\x34\x5e\x46\x09\xa0\xb1\x11\x02\xf5\x65\x7f\xd4\x5f\xb7\x3e\x02\x77\x13\xcc\x71\x97\xb3\x09\x79\x31\x87\xae\xd0\xd7\xdb\x75\x22\xff\x48\xd7\x89\x27\xae\x71\x80\xd7\xe8\x43\x79\x49\xeb\x73\xed\xd7\xfa\xf3\x40\xc7\x01\xb9\x5f\xfb\x51\x77\x9d\x3b\x24\x7c\x7f\xb4\xc2\xb5\x09\xf7\x57\x02\xc0\xc7\xb0\x4f\xae\x13\x42\xfc\x16\xb0\x0b\xa1\x85\xed\x7f\x4b\x9e\x7f\xea\x44\xba\xd1\x8e\x6f\x76\x2b\xbc\x83\xe1\x77\xfb\x14\x7e\x91\xf1\x4d\xbc\x41\xe1\x9d\x4c\xbe\x13\xf2\x3d\x0c\x6f\xf3\x28\x7c\xcc\x6f\xc7\xbf\x86\x78\xe6\x99\xfc\x2e\xd8\xcf\x32\xf9\xed\xb0\x33\xb4\xda\x8e\xff\x29\xec\xb4\x33\x7c\xce\xa5\xf0\x00\xb3\xf3\x9d\x65\x0a\x9f\x66\xfd\x0a\x22\x7e\x1f\xb3\xf3\x02\xe4\x2f\x31\x3b\x3f\xab\x07\xce\xe4\x7f\x1f\xf1\x8c\xb2\x7e\xfd\x3b\xe2\xf1\xad\xb1\xe3\x3e\xf4\xb7\x85\xe1\x3f\xf7\x2a\xbc\x83\xe1\x9f\x03\x3e\xc2\xf0\x5f\x22\xfe\x96\x75\x76\xfc\x0c\xe4\xb3\x0c\xff\x63\xf4\xab\x9b\xe1\x3f\x47\x9e\xcf\x32\xfc\xcb\x98\x27\x17\x9b\xec\xf8\x0f\xe1\x57\x30\xfc\x51\xd8\xef\x64\xf8\x39\xd8\x19\x5e\x6f\xc7\x27\x61\xa7\x83\xe1\xad\x88\x3f\xc4\xec\xfc\x15\xe2\xf4\x33\xf9\xcb\xc8\xff\x14\x93\x7f\x18\x7e\xa7\x98\xfc\x57\xe1\x77\x8c\xf5\x37\x0c\xbf\xfe\x0d\x76\xfc\x41\xf4\xab\x9d\xe1\x9f\x42\x3c\x59\x86\xbf\x84\x79\x72\x81\xe1\xad\x64\x67\xa3\x1d\x1f\x86\x7c\x27\xc3\xbf\x8f\x7e\xf5\x30\xfc\x57\xe8\xd7\x20\xc3\x7b\xd1\xaf\xc9\xa0\x1d\xff\x19\xfc\xfa\x18\x7e\x06\x7e\xbb\x19\xbe\x03\x76\xc2\x1e\x3b\x7e\x10\xf8\xc8\x26\x16\x27\xec\x67\x19\x1e\x43\x7e\xc2\xcc\xfe\x0b\xe8\x57\x80\xc9\xbb\xd1\xaf\x69\x26\xbf\x0b\x7e\xa7\x99\xfc\xd7\xe1\xf7\x22\xcb\x83\x1f\x7e\x03\xb7\xdb\xf1\xe3\xe8\x6f\x07\xc3\x07\x10\x4f\x27\xc3\x5f\x47\x3c\xa3\x0c\x1f\x85\xdf\x16\x96\x9f\x23\x64\x7f\x8b\x1d\x1f\x87\x9d\x6e\x86\xef\x80\x9d\x0e\x66\x67\x90\xd6\x2f\xe3\xf8\x06\xc8\x5f\xd8\x6a\xc7\xef\x45\xfc\x82\xc9\x17\x69\x5e\x31\xfb\xa7\x61\x3f\xb2\xcd\x8e\xbf\x83\xbc\xcd\x31\x3b\x19\xd8\x99\x64\x7e\x3f\x83\x7e\x8d\x31\xf9\xb7\x68\x5f\x60\xf8\x61\xf8\xed\x61\x7e\xbf\x01\xbf\x17\x58\x7e\x1a\xe1\x77\x8e\xe1\x1f\xa0\xfe\xb2\x78\xba\x11\x4f\x9a\xe1\xcf\x82\x87\xdb\x19\xfe\x0a\xe2\x1c\x66\xf8\x2c\xe4\xd5\x38\x97\x3e\x19\xe4\x7f\x9e\x9d\xa7\x05\xec\xcc\xb2\x3c\x6f\xa2\xf5\xce\xec\x74\x20\xfe\x11\x26\xff\x6d\xf8\x9d\xf4\xda\xf1\x1f\x21\x6f\x93\xcc\xce\x9f\xc0\xbe\x60\xf2\x6f\x20\x6f\x3e\x26\x9f\x82\xdf\x16\x26\xff\xb7\xf0\x3b\xc8\xf0\x4f\xc3\x4e\x9a\xed\x9b\x6f\x22\x9e\x6e\x66\xff\x0f\x29\x1e\x1f\x8b\x07\xe3\x9b\x66\xf8\x37\x61\xbf\x9d\xe1\xbf\x09\x7c\x98\xe1\x7f\x4f\x76\x96\xdb\x71\x2f\xad\x3b\x86\x6f\xa1\x75\xcd\xf0\x4e\xcc\x93\x1e\x86\xff\x10\xf8\x59\xd6\xdf\xbd\xe8\xef\xd8\x0a\x3b\x3e\x82\x78\xe6\x99\x9d\xf5\xf0\x9b\x65\xf2\x1b\x60\x67\x68\xa5\x1d\x7f\x1a\x76\xda\x19\xfe\x2e\xf1\x24\xb3\x93\x44\x9c\x3e\x26\xdf\x87\x71\x1c\x61\xf1\x7f\x16\xf3\xf3\x12\xb3\x73\x99\xce\x39\xcc\xce\x63\x88\x67\x94\xf5\xeb\x6d\xc4\xe3\xbb\xcd\x8e\xff\x04\xf1\xb4\x30\xfc\x03\x74\xce\x64\xf8\xc7\x29\x4e\x86\x0f\xc8\x38\x3d\x22\xbd\xca\x8e\xbf\xea\x55\x78\x0b\xc3\x07\xeb\x15\x1e\x62\xe7\xae\xef\xb9\x15\x3e\xc9\xe4\x8f\xc1\xfe\x59\x26\x7f\x97\x47\xe1\xed\xec\xbc\x37\x5a\xa7\xf0\x19\x66\x67\x05\xec\x5f\x62\x76\xfe\x06\x71\x5e\xac\xb3\xe3\x71\xea\x17\x93\xff\x97\x65\x0a\xf7\x31\xbf\x1b\x11\xcf\x20\xf3\xfb\x09\xc4\x33\xc4\xe4\x83\x0d\x0a\x9f\x62\xf8\x6a\xf8\x15\xec\xbc\xfa\x25\xc4\xe9\x67\xf8\xe7\x11\x4f\x96\xe1\x5d\xf0\x3b\xca\xce\x9f\x6f\xba\x14\x7e\x89\xe1\xbb\x60\x67\x8a\xe5\xe1\xd3\xf0\x1b\x58\x6b\xc7\xd7\x40\x3e\xc2\xf0\xb7\x31\xbe\x69\x86\xff\x03\xfc\x86\xd9\xf9\x6d\x17\xf2\xd6\xc1\xfc\x7e\x17\x7e\x2f\x31\x3b\xff\x49\x79\x66\x76\x9e\xc3\xf8\x76\xb0\xf3\xe4\x63\x88\x67\x9a\xd9\xff\x1e\xf2\x3c\xcb\xec\x7f\x04\x7e\xa7\x98\xfd\x6f\xc1\x2f\x3f\xb7\x07\x60\xdf\xcf\xfc\xfe\x01\xe2\x39\xcb\xec\x8f\xc3\xef\x30\x93\x3f\x8c\x7c\x4e\x33\xfc\x02\xfc\xfa\xd8\x79\xf8\xaf\xe1\x37\xc0\xf0\x5f\x60\xdc\x3b\x19\xfe\x3c\xfc\x4e\xb2\xf3\xed\x56\xf4\x77\x86\xe1\x5f\x80\xfd\x48\xc0\x8e\xdf\x06\xfb\xed\x0c\xbf\x87\xf2\xcf\xf0\x65\xb0\x1f\x61\xe7\xbd\xd3\xe8\xd7\x14\x93\xff\x6f\xd8\x19\x62\xf2\xef\x21\x3f\x9d\x6c\x1f\xdf\x81\x75\x34\xcd\x7e\xd3\x3a\x89\xfe\x66\xd9\xb9\xf4\x5d\xd8\x99\x63\x7e\x0f\x22\x9e\x69\xe6\xf7\x1b\xd4\x5f\x86\xaf\x43\x9c\x01\x66\xbf\x8d\xc6\x97\xd9\xff\x37\xf8\x1d\x61\xf2\xa7\x91\xe7\x19\x86\x7f\x1e\x7e\xfd\xec\xfc\xfc\x16\xfc\x86\x18\xfe\xbf\x58\x5f\xdd\x0c\x7f\x19\xf9\xef\x61\xf9\x19\x43\x3c\x17\xd9\x39\xd9\x4b\x79\x60\xf8\x24\xfc\x86\x37\xdb\xf1\xd5\xc8\x7f\x0b\xc3\x7d\xc8\x43\x3b\xc3\xdf\x25\x7e\x60\xe7\xc3\x2c\xfa\x3b\xcd\xe4\x8b\x94\x4f\x26\x2f\x68\x3e\xb0\x73\xe0\x7e\xd8\x9f\x67\x76\x5e\x41\xfc\x33\xcc\xce\x7b\x14\x0f\x3b\xd7\xfd\x14\x7e\xf9\xf7\x82\x3c\xf2\x19\x62\x7e\xdf\x46\xde\x86\x98\x5f\x03\xf6\xf9\xf9\xff\x65\x8c\xd7\x2c\xc3\x5f\x87\xdf\x00\x3b\x87\xff\x23\xfc\x86\x19\xfe\x20\x8d\x2f\xc3\x0b\xc8\x67\x80\x8d\xfb\x0c\xfc\xce\xb0\xfe\xc6\x90\xcf\x49\xb6\xbe\xe6\x61\x7f\x8e\xd9\xf9\x33\x9a\xcf\xcc\xce\x3b\xe8\xef\x30\x93\x6f\x47\x3c\x61\x76\xfe\xfc\x27\xc4\xc3\xbf\x4f\x1d\x23\x7e\x60\xf1\xfc\x80\xfa\xcb\xec\xcc\x43\x9e\x9f\xb7\x5f\x43\x3c\x73\x0c\xbf\x0f\x38\x3f\x87\xdf\x49\xeb\x9d\xe1\x7f\x89\xf9\xd3\xc3\xf0\x6f\x23\x9e\x0b\xec\xdc\x7b\x8a\xf2\xc9\xf0\x6d\xb4\x5f\x34\xd8\xf1\x8f\xd3\xfa\x62\xf8\xf3\xc8\x4f\x84\xe1\x1f\x81\xfd\x10\x3b\xef\xbd\x0a\xfb\x93\x4c\x3e\x4b\xfb\x11\x93\x37\x10\x7f\x3b\x3b\x67\x0e\x23\x9f\x33\xcc\xce\x4f\x91\x9f\x4b\xcc\xce\x03\xb0\xcf\xcf\xf9\xcf\x61\x3e\xfb\x98\xfd\x46\x1a\x47\x66\xff\x0c\xad\x23\x26\xef\x81\xdf\x29\x86\xef\x81\x5f\xc1\xce\xc3\x9f\x41\xde\xfc\x0c\x3f\x45\x3c\xc9\xce\x03\x9f\xa0\xfd\x82\xc9\x7f\x14\xf1\x8c\xb1\x73\xef\x1b\x14\x0f\x3f\x0f\xc3\x7e\x88\x9d\x03\x85\x3c\x3d\xd7\xe9\xa0\xc4\x3d\x1a\xb6\xdf\x25\x84\x9f\x4f\x4e\xf3\xfc\x23\x14\xce\xa6\x95\xd8\x0e\x79\xfe\x13\x73\x11\xf2\x61\x86\x67\x21\xcf\x8e\x9f\xa2\x05\x38\xdb\xbe\xe4\xbf\xa7\xf4\x8b\x15\x5a\x3c\xaf\x48\x7c\xa5\x86\x1f\x97\x38\x8f\x52\x0d\x93\x5f\x2c\xd7\xf0\x93\x12\x6f\xd0\xf0\x5e\x89\xeb\x79\xdb\x22\x71\x3d\x6f\x7e\xf0\x85\x59\xfc\x52\xe0\xa5\x19\xea\xff\x63\x7d\x8f\x57\xf5\x5f\x08\x21\x66\xeb\xed\xf5\x4e\x9f\x5d\x7e\xbe\xc1\xde\xde\x7d\x1b\x6b\x5f\xc5\xda\xd7\xd8\xdb\xc5\x3a\x7b\x7b\xcf\x06\x7b\xbb\x6f\xa3\xbd\xfd\xec\xed\xf6\x76\xff\x16\x16\xdf\x56\x7b\xfb\xac\xb0\xf7\x37\xed\xb2\xcb\x0f\xbb\xec\xf2\x01\xaf\xbd\x1e\x5e\xc9\xea\xab\xed\xf5\xc8\x7a\x7b\x3d\xbd\xc9\x5e\x6f\x09\xd9\xeb\x83\x21\xbb\x7f\x3a\x46\x98\x23\xfe\x63\x7a\xcf\x86\xfa\x4f\xcc\x87\x3d\x7d\xf9\xb3\x7d\xe2\x54\x57\x4f\xe1\xce\x9e\xde\xa7\x1e\xc9\xe7\xba\xfa\x72\x85\xfe\x42\x2e\x7f\xf6\x64\x9f\x90\x2f\xe1\x7a\x9e\x3a\x79\xba\x2f\x57\x02\xf7\xf4\xe6\x9f\x2c\x35\xdc\xd9\xdb\x75\xc6\x6c\x3b\xd1\xf5\xe4\x93\x85\x3b\x2d\x21\x69\x30\xff\xb1\xfc\xe9\xbe\x82\x7a\x26\x11\xdd\xe6\xe9\xbe\x7c\xef\xc2\x46\xa5\x54\xee\x63\xf9\xde\xc2\xc9\xa7\x4e\x8b\xdc\x93\x27\x4f\xe4\x4f\x17\xf2\x52\x6f\x4f\xbe\x3b\xf7\x58\x6f\xd7\xa9\xbc\xd8\x53\xe8\xeb\xed\xeb\x7a\x44\xec\x29\xf4\x9f\x92\x65\xef\x53\x8f\x76\xf5\x75\x89\x8e\xd6\x43\xad\x0f\xe5\x3a\x0e\x3d\xbc\xaf\x2d\xd7\x7a\x24\xd7\x96\xfd\xd0\x91\xdc\x43\xad\x87\xf7\x8b\xa3\x87\xdb\x72\xad\x87\x1e\xc8\x7d\xe8\xe0\x07\xdb\xb2\xb9\xd6\xc3\x39\x25\x79\xa0\xed\xe0\x03\x47\xda\x17\x21\x22\x5b\x0f\xb5\x1d\xa9\x88\x1f\x3c\x7a\xe0\x40\x6e\xdf\xef\x1e\x69\x23\xf0\xc0\xbe\x7d\x46\x6e\xaf\x59\x44\x54\x61\xe4\xf6\x52\x15\xf5\x34\xea\xd1\x24\x80\x04\xca\x68\x02\x2d\x71\x00\x06\x3d\xc4\xd0\x10\xa3\x06\x02\x0c\x42\xa2\x64\x94\x1a\xa2\xd4\x60\x50\x8b\x41\xca\xf4\x10\x8d\x90\x30\x3d\xec\x4d\x23\xc4\xb4\xaa\xa7\x51\x37\xe8\x21\x85\x06\x23\x05\x20\x09\x80\xca\x68\x92\x54\x08\x31\x08\x49\x90\x48\x82\x44\x08\x31\x08\x89\xa3\x8c\xc6\x49\x36\x46\x48\x94\x90\x28\x21\x06\x99\x89\xa0\x29\x85\x4e\xa4\x80\xef\x05\x60\x10\x92\xa6\x96\x34\x5a\x52\x28\x93\x29\x8a\x1e\x12\x71\x00\x54\x46\xe3\x68\x88\x01\x88\xa1\x1e\xa5\x07\x83\x5a\x0c\xab\x89\x74\xa3\x24\x42\x88\x41\x48\xd4\x20\x19\x83\x64\x08\x31\x08\x89\x46\x08\x89\x00\xd9\x9b\x44\x47\x93\xd4\x3f\x3c\xa4\x93\x18\x25\x94\x46\x8a\x24\xe8\x21\x99\xa4\x8e\x92\x08\x3d\x24\xd0\x92\x48\xd2\x64\x24\x51\x42\xe2\xd4\x12\xa7\x96\x18\x3d\x44\x2d\x6b\xd4\x2f\x6a\x89\x50\x4b\x84\xec\x53\xcc\x09\x1a\x0b\x42\xd2\x84\xa4\xf0\x90\x44\x4b\x02\xf5\x38\xea\x71\x12\xa4\x87\x58\x82\x06\x83\x6c\x59\x48\x82\x3a\x41\xd3\x88\x00\x42\x0c\x42\x0c\x0b\x21\x11\x83\x00\x42\x0c\x42\xa2\x11\x0a\x09\x23\x12\xc7\x88\xc4\x69\xea\xe1\x21\x8d\x06\x2a\x8d\x34\x1a\x52\x00\x52\xa4\x41\x80\xf5\x90\x8c\xd3\x50\x91\x08\x21\x89\x38\x75\x8e\x44\x08\x31\x08\x31\x08\x89\x93\x6c\x9c\x64\x09\x31\x08\x89\x59\x2d\xd4\x80\xac\xa2\x99\x06\x8c\x1e\xd2\x31\x0a\x19\x0f\x49\x6a\x49\x02\x48\x00\xa0\xd2\xa0\x87\x38\x04\x62\x64\x22\x46\x12\xf4\x10\x8d\xd1\x40\x11\x40\x88\x11\xb5\x64\x69\x92\x91\x88\x85\x44\x08\x89\x10\x82\x71\x89\xd2\xb8\xe0\x21\x8d\x86\x34\x35\x10\x60\x10\x92\x02\x40\xa5\x91\x42\x43\x34\x19\xa5\x81\x21\x65\x7a\x88\x26\xa8\x89\x1e\x0c\x0b\x89\x93\x70\x9c\x9a\x2c\x24\x46\xea\x04\x50\x19\xb1\x5c\x1a\x48\xac\x81\xc4\x1a\x34\xde\x78\x88\xa3\x81\x4a\x23\x8e\x86\x18\x80\x98\x41\x4b\x84\x54\xa9\xc5\x7a\x88\x52\x19\x25\x11\x42\x0c\x52\x36\x2c\x51\x12\x31\x2c\x11\x92\x89\x90\x4c\x84\x64\xe8\x61\x6f\x04\xa3\x11\xc1\x20\x44\x28\xf9\x11\x4a\x3e\x1e\xa2\x29\x6a\x4a\x51\x13\x21\xd1\x24\xc9\x24\x49\x86\x10\x83\x90\x68\x82\x90\x04\x21\x71\x42\xe2\x84\xc4\x48\x9d\x1e\xa2\x24\x62\x3d\x18\x11\xed\xf0\xfb\x3e\x3e\x6f\x96\x8e\xc8\xb6\x4f\x36\xa7\xca\xbb\xd9\x91\x9c\xbf\x46\x73\xe1\x0f\xfb\x7a\x8d\xff\x95\x4e\xff\xf0\xab\x67\x9f\x5d\x40\x7f\x8c\xbd\x77\xe4\xdf\x28\x86\x84\xa8\xf0\xad\x41\x88\x51\xf5\xdf\x77\x88\x39\xd4\x1b\xd1\x4f\xd2\xa7\xef\x3d\xa9\x2a\xfe\xe7\xe0\x97\xdf\xd9\xe1\xfe\x63\x55\xfc\x8f\x7c\x14\xfa\x65\xfe\x3d\x15\xfc\xff\x07\xfc\xb3\xd7\x3f\x62\x78\x9f\x2a\xf9\x20\x73\xff\x59\xe8\xf3\x31\x9c\x84\x3e\xbf\x47\xc8\xf3\xff\x68\x15\xfd\xe9\xfb\x16\xa7\xef\xad\xa2\x3f\xf7\xc1\xc5\xe9\x5f\xac\xa2\xef\xbb\xbf\xb2\x3c\x9f\x7f\x7f\x57\x45\x7f\x65\x15\x7d\x5e\x7f\xad\x4a\xfe\xfd\xd0\xe7\xf3\x98\xe7\x7f\xb2\xca\xf8\xa7\xb1\x7e\x68\xfa\x98\xe3\x7f\x5b\x85\xf1\xff\x6e\x05\xdf\xe6\xa7\x03\xfe\xb3\x65\xdf\x17\xff\xab\x4c\x9f\xbe\x9f\xff\x5f\x00\x00\x00\xff\xff\x6a\x12\x59\xb7\x98\x6c\x00\x00")

func stracebackTailcallBpfOBytes() ([]byte, error) {
	return bindataRead(
		_stracebackTailcallBpfO,
		"straceback-tailcall-bpf.o",
	)
}

func stracebackTailcallBpfO() (*asset, error) {
	bytes, err := stracebackTailcallBpfOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "straceback-tailcall-bpf.o", size: 27800, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"straceback-guess-bpf.o": stracebackGuessBpfO,
	"straceback-main-bpf.o": stracebackMainBpfO,
	"straceback-tailcall-bpf.o": stracebackTailcallBpfO,
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
	"straceback-guess-bpf.o": &bintree{stracebackGuessBpfO, map[string]*bintree{}},
	"straceback-main-bpf.o": &bintree{stracebackMainBpfO, map[string]*bintree{}},
	"straceback-tailcall-bpf.o": &bintree{stracebackTailcallBpfO, map[string]*bintree{}},
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

