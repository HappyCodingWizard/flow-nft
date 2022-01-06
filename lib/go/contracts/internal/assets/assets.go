// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../contracts/ExampleNFT.cdc (6.872kB)
// ../../../contracts/MetadataViews.cdc (2.228kB)
// ../../../contracts/NonFungibleToken.cdc (4.832kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
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

var _examplenftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x5f\x6f\xdb\xba\x15\x7f\xf7\xa7\x38\x37\x0f\x9d\x8d\x25\xd6\x06\x0c\x7b\x30\xe2\xa6\x45\x73\x8d\xe5\xe1\x06\x45\xeb\x6d\x0f\x45\xb0\x4b\x8b\xc7\x31\x17\x89\x14\x48\xca\x9e\x96\xeb\xef\x3e\x1c\x52\x7f\x28\x89\x4a\xd2\x15\xb8\x42\xd1\x5a\xe2\x39\x87\xe7\xff\xf9\x91\x4d\x12\xd8\x1e\x84\x01\x61\x80\x49\xc0\xff\xb0\xbc\xc8\x10\x04\xfd\x9d\xa3\xb4\xcc\x0a\x25\x41\xed\x81\xc1\x26\x53\x27\xb8\x57\xf2\x6a\x53\xca\x47\xb1\xcb\x10\xb6\xea\x09\xe5\x2c\x49\xe0\xce\x12\xbf\x54\x16\x0a\xa6\x2d\x91\xdb\x03\x82\xda\xef\x45\x2a\x58\x06\xc6\x32\xc9\x99\xe6\xb0\x2b\x2d\x08\x0b\xcc\x98\x32\x47\x0e\x56\xc1\x0e\x89\xff\x88\xba\x02\x23\x72\x91\x31\x4d\x5f\x0f\xea\x04\x39\x93\x15\xdc\x6f\xb6\x06\x4e\xaa\xcc\x78\xa7\x92\x93\x9d\x2a\x8d\xb0\x2f\x65\x4a\xfa\xb1\x4c\xd8\x6a\x39\x9b\x89\xbc\x50\xda\x92\x8e\x8d\x8a\x4e\x43\xd8\x6b\x95\xc3\xc5\x32\x19\x2e\x2c\x53\x9e\x5e\x34\x5c\xbf\xa0\x65\x9c\x59\xf6\x0f\x81\x27\xd3\xb2\xf4\xbe\x7a\xfa\x59\x51\xee\x20\x55\xd2\x6a\x96\x5a\xf8\xd9\x7b\xec\x7e\xb3\x5d\x8d\x37\x7e\x9e\xcd\x00\x00\x88\xe1\xe8\x2c\xb3\x2c\xfb\x5a\x16\x45\x56\xad\xe0\xef\x77\xd2\xfe\xf5\x2f\x1d\x01\x1e\xc9\xb6\x4f\xb5\xdc\x3b\x29\xac\x60\x99\xf8\x2f\xf2\xf9\x62\x40\xf3\x4f\x61\x0f\x5c\xb3\xd3\x5c\xf0\x46\xcc\xa5\x53\x78\x05\x1f\x39\xd7\x68\xcc\xcd\x90\xe5\x16\x0b\x65\x84\xed\x71\x58\x15\xd2\xb7\x0c\x19\x92\x16\x59\x86\xce\xb5\x5f\xad\xd2\xec\x11\x3f\x33\x7b\x58\x41\xf0\x32\x41\xfe\xb9\xdc\x65\x22\xf5\xd4\xdd\xef\x1e\xf1\x2f\x42\x5a\xd4\x93\x72\x5b\x5a\x8d\x46\x95\x3a\x45\x88\xba\x76\x79\x77\xbf\xd9\x5e\xf6\x83\xb6\xfc\x82\x46\x65\x47\xd4\xf0\xec\xa4\x84\xbb\x76\x86\xcf\x46\x6b\x92\xe5\x48\x4a\x68\x21\x1f\x47\x8b\x1c\x4d\xaa\x45\x41\xc6\x4d\xd2\xd8\x43\x99\xef\x24\x13\x59\x4b\xd1\x92\x08\x29\xec\xbc\x7d\x73\x5f\xba\x10\xf4\xbe\x87\x5a\xf4\x57\x22\x2a\xf4\x09\x46\xfb\x77\xcb\x8b\xc0\x17\xf4\x18\xcc\xf6\x4b\xc1\x61\x0d\x82\x8f\x17\x48\x07\x58\x3b\x55\xc6\x8b\x81\x1a\xb0\x0e\x95\x1a\x93\xb6\x0a\xc1\xba\x53\xae\x25\x3b\xbb\x5f\x3d\x2f\xee\x4b\x09\x8f\x68\x5d\x18\xe7\x8b\x15\x7c\xdb\x56\x05\x3e\x0c\x74\xd7\x68\x4b\x2d\xe1\x5b\xef\x23\x3d\x44\x7c\xdd\x4f\x85\x5b\x61\x8a\x8c\x55\xef\xe7\x8b\xcb\xb7\x90\xff\x6d\xbb\xfd\xbc\x6d\x14\x7d\x5f\x57\x5c\xf3\x3c\x04\x9a\x8f\xb4\xd6\x3e\xe7\x48\xcc\xfc\x5f\x70\x14\x78\x5a\xb9\x0d\x16\x2b\xf8\x28\xab\xaf\x56\x97\xa9\xbd\x19\x06\xe1\x24\x6c\x7a\x70\xc4\x83\x15\x7a\x52\x66\xf0\x65\x93\x56\x23\x9e\xc0\x3d\x51\xa6\x79\x94\x03\xda\xb4\x6b\xa3\x3f\xf6\x56\xf3\xf4\xb2\x70\x98\x10\x71\xb6\xc5\x9b\x6d\x1b\xfa\xff\x3b\x2c\xec\xb1\x4e\xdb\x59\x6a\xb1\x1a\x24\xe7\xb4\xad\xb9\xc8\xd1\x56\x05\xae\xe0\x42\xe4\xec\x11\x93\x7f\x17\xf8\x78\xf1\x06\x1b\x83\x04\x09\x14\x96\xa3\xe4\x3f\x47\xfa\x9c\xeb\x8b\x7b\x96\x62\x30\x53\x86\x9d\x75\xd0\xd9\x28\x01\x79\xdd\xda\x2d\xb5\xc5\x15\x7c\x18\x75\xca\xfb\xcd\x76\x11\x2b\xb6\xbb\x5b\x5f\x6a\xbe\x19\x3d\x8c\x48\x76\x4a\x6b\x75\xba\xdf\x6c\x83\xa9\xb1\x58\xc1\xbb\xd8\x06\x13\xcc\x9d\x21\x03\x19\xdd\x02\x71\x0f\xeb\xa3\x50\xc6\x46\x0a\x63\xae\xd1\x94\x99\x85\xf5\x9a\x3c\xba\x80\xdf\x7e\x6b\x3e\xdd\xb8\x96\x46\x3d\x6d\x22\x73\x2e\x3e\x31\x49\xf0\xc4\xab\x15\x38\x18\x34\xee\x51\xa3\x4c\x71\xe5\x70\xc5\xdd\x6d\x83\x5e\x7c\xec\x90\x77\x14\x84\x71\x84\x4c\x95\xd6\x98\xda\x8b\x41\xdc\x5f\x8f\x6f\x17\xcb\xd5\x0b\x11\xbe\x1c\x8f\xba\xcf\x5a\x1d\x05\x47\x1d\x59\xfa\x82\x29\x8a\x63\x74\x69\x2c\x38\x3e\x2c\x3b\xba\xc0\xe5\x49\x02\x5c\x78\x6c\xa5\x2b\xf2\x08\xb9\x2a\x55\x72\xaf\x74\x2e\xe4\x23\xb8\x64\x33\x21\x39\x11\x10\x86\xec\xec\xa5\x02\x82\x93\xb0\x07\x02\x96\xbf\xfa\xd8\xff\x4a\x0e\xde\x0b\xcc\x78\x2f\x63\x08\x1c\xa9\x93\x44\x4e\x78\x6f\x05\x1f\x9e\x3d\x75\x64\xec\xdf\x6f\xb6\xe7\xfe\x74\x85\x79\x74\xc6\xb5\xe2\xe0\xfa\x0a\x9e\xcf\xb1\x16\x9e\x24\x4e\x3d\x82\x52\xa0\x31\x57\x47\x74\x18\x98\x2c\x71\xf0\xcf\xe3\xcc\xd6\x3b\x4c\x72\xf0\x44\xc2\x12\x48\x75\xcb\x2c\xcb\x50\x8f\xb2\xbf\x11\x3b\x6f\x7e\xdc\xdd\x06\xd9\x1f\x2d\xd1\x81\x0d\x0e\x56\x38\x0c\x79\x7d\x35\x30\x68\xe9\x75\x9d\x3f\x61\xb5\x82\x6e\x83\x05\xdc\xdc\x40\xc1\xa4\x48\xe7\x17\xb9\x30\x86\xc2\x74\xbf\xd9\x5e\x2c\xfa\x2d\x09\x73\x31\x40\x90\x6e\x9b\xa5\xe0\x0d\x86\x6c\x77\xd3\x37\x4b\xe6\xf1\xe1\x22\xda\xd6\xae\xaf\x1c\xeb\x84\x6b\xeb\xbe\x04\x96\x3d\x91\x5f\x9d\x5b\xc9\x85\x8c\xf3\x9e\x07\x5b\x07\x9b\x20\xe5\x42\x41\x2d\x53\x5d\x9f\x35\xa3\xe0\xc0\xb4\x66\xd5\xff\xd7\x10\x5f\x72\xb7\xff\xc1\xcc\x4f\xf0\xa1\xdf\xa7\x66\x23\x9e\xae\xab\x11\xd2\xa9\x1d\xd9\x27\x23\x0b\x38\x77\x2a\x4b\x3c\xd5\xc2\x6b\x1b\x82\x1a\x3b\x1d\x44\x7a\x68\xd3\xd0\x1d\x9f\x32\x0e\x4a\xe2\x68\x4f\x95\xf1\x6d\x3c\x33\xbe\x09\xfe\xd0\x1a\x10\x09\x7b\x78\x0a\xa0\x78\xd3\x09\xe0\xf5\x68\x73\x34\x56\xab\xaa\xdd\x77\x22\xde\x7e\xa2\xd4\xb9\xe1\x0a\xc9\x85\xa7\x69\xa7\xb4\x66\x0f\xcc\x02\xd3\x34\xea\x06\xb1\x7f\xc3\x7c\x8a\x83\xc1\x41\x69\x3c\x61\x65\x26\xf4\x6b\xc7\x19\xc9\xf6\x8d\xaa\xe9\xeb\x56\x35\x75\x3f\xad\x58\x92\x80\x51\xde\x82\xae\xf0\x21\x65\x84\x00\x19\x07\x61\x0d\xe4\x75\x7f\x75\x19\x4b\x04\xcd\xd7\x83\xe2\xe6\x87\xc6\x6b\xdc\xf6\x77\x91\xe8\x33\xf3\xca\x80\x3e\xcf\xc6\xc8\xfb\x87\x86\xb5\xd8\xc7\xb2\xf0\x27\x37\xa3\x23\x43\x3c\x49\xe0\x93\x46\x66\xd1\x65\x48\x69\x0f\x4a\xd3\x21\x77\x10\x8d\x2c\x53\x27\xe0\xea\x24\x53\x66\x6c\x78\xe4\x0a\x0b\x41\xe3\x1e\xd6\x53\x5e\x20\xd1\xaf\xb8\x62\xe0\x4e\x12\x47\x45\x3f\xb0\xf7\xbb\xf1\xdd\x84\x7b\x69\xea\x36\x33\x77\xe0\xe0\x8f\xb2\xfa\x52\x4f\xcd\xe7\xf8\x90\x3e\x47\xfa\x95\xdc\xdb\x1f\x36\x9f\xe4\x60\x07\x86\xd6\x4e\xe8\x6b\x4e\xa8\xad\x0e\xf8\x28\xeb\xde\x60\x44\xcc\x49\x75\x77\x19\x0d\xf1\xa6\xeb\xf4\xcd\x8b\xc3\xac\x24\x21\x5f\x13\x3c\x6e\x6e\x84\xea\x56\x23\x2b\x25\xd1\x15\xa9\x2b\x47\xab\x20\xad\x73\xcf\xf5\x62\xcc\x0b\x5b\x0d\x8b\xbd\x89\x9a\xa7\xfc\x99\x48\x3a\x88\x34\x8f\x8e\xef\x28\x84\x6a\x87\x64\xb3\x67\x28\x65\xa0\xfd\x97\x16\x33\x79\xb5\x81\xf1\x5c\x48\x50\x1a\x8c\xa2\xfe\x41\xb3\xbc\xb9\x1e\xf3\xb7\x61\xea\x24\xeb\xeb\xb3\x5a\x04\xdb\x65\xae\x74\x72\x21\xad\x33\xae\x75\x57\x92\x44\xef\x54\xfc\x3d\x4c\x73\x45\x55\x4b\x21\x6e\x0a\x28\xfd\x6b\x6a\x2f\xd1\xbb\x87\x71\xee\xf5\xee\x76\x38\x9c\x9b\x49\x4f\x7f\x64\x8d\x9d\x53\x51\x08\x24\x19\x01\x80\x2a\x1d\x26\xb1\x07\x14\x3a\xfc\xdc\x56\xfe\xa8\x70\x6a\x6d\xe6\x83\xec\xab\x65\xaf\xe0\xdd\xf3\xab\xa8\xf7\xfc\x7b\x5d\xb2\x0c\x67\x7e\x2f\xd1\x86\x15\x44\x78\x57\xa2\x1b\x46\x5d\x7e\x8c\x2c\x85\xfa\xb6\x28\x28\xc5\xe0\x06\x71\x7c\x7a\xf5\xc6\xc5\x4f\xf1\x3d\xf3\x5e\x3c\xb8\x07\x76\x4e\x9c\x94\x17\x23\x63\x5f\xc8\x80\x3f\x18\x60\x69\xaa\x4a\x69\x7b\xf1\x1f\x07\x1d\xc2\xd8\x2e\x07\x20\xee\xfa\xca\xbb\x6b\xb0\x75\xdc\x33\xb0\x9e\x5a\xf8\x63\xdd\x74\xe7\x7f\x5e\xc4\x3b\x89\xbb\xae\x5b\xf4\x0f\x42\xdd\x55\xac\xb3\xcc\xc9\x03\xe3\x04\xb6\x64\xfe\x62\xa1\xa7\xc2\x9f\x7a\x85\xf5\x15\x3d\x6e\xa0\xe8\x70\x28\x98\x3d\x98\x3e\x73\xf4\xca\x15\xd6\x90\x18\xff\x9a\x60\xe4\xc8\x38\x25\xa2\xbb\x7a\x25\x09\xbe\x37\xbe\x41\xc0\xe8\x6a\x36\xbe\xbf\x27\xeb\x99\xd7\x8c\xf4\xa0\xc9\x75\xcd\x86\x7a\x84\x61\x47\xac\x51\x7f\x2d\xb0\x65\xa7\x19\x14\xb4\x83\x17\x1a\x66\xab\x68\x9d\x51\x4b\x92\x3a\xbf\xbe\xea\xb8\x03\x4c\x1b\x75\xe8\xa2\xa7\x75\x5b\xa3\xf5\xf4\x48\x59\xc1\x76\x22\x13\xb6\x82\xbd\xd2\x53\x48\xb0\xa7\x41\x26\xe4\xd3\x75\x38\x2c\xbb\x6d\x5f\x6f\x4e\x97\x61\x9e\x4e\x5f\x09\x9c\xdf\xcf\xc7\xe7\xdb\x58\xb0\x07\x0d\x8b\xe9\x47\xb4\x2f\x79\x63\x16\xa9\xe8\x30\x98\xf5\x88\xf8\x9e\x40\xe6\x9e\xa5\xd7\xd5\xbc\x98\x57\x62\xe8\x19\x83\xf8\x8d\x92\x31\x50\xd2\x9d\x66\xa6\xff\xab\xe4\x3c\x3b\xcf\xfe\x17\x00\x00\xff\xff\x90\x4c\x20\x21\xd8\x1a\x00\x00"

func examplenftCdcBytes() ([]byte, error) {
	return bindataRead(
		_examplenftCdc,
		"ExampleNFT.cdc",
	)
}

func examplenftCdc() (*asset, error) {
	bytes, err := examplenftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ExampleNFT.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6, 0xc3, 0x45, 0xad, 0x98, 0xac, 0x3, 0x57, 0x4, 0x73, 0x3f, 0x2a, 0x88, 0xd, 0xbf, 0x22, 0x46, 0x57, 0x89, 0xa1, 0xee, 0x63, 0x35, 0x96, 0xf1, 0x3a, 0xf1, 0x2d, 0x12, 0xc5, 0xa1, 0xa8}}
	return a, nil
}

var _metadataviewsCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x51\x6f\xe3\x36\x0c\x7e\xf7\xaf\xe0\xd3\x90\x14\xad\xdd\xbb\x0d\x05\x16\x20\x18\x8a\xe6\xba\x65\xd8\x0d\x45\x9b\xdb\xcb\x70\x18\x64\x89\x8e\xb5\xc9\x92\x41\xd1\xed\x82\xa2\xff\x7d\xa0\x12\xbb\x4e\xdd\x0e\xbb\x6d\x7e\x08\x62\x93\xfa\xf4\xf1\xe3\x27\xaa\x38\x39\xc9\xb2\x4d\x6d\x23\xe8\xe0\x99\x94\x66\xb0\x4d\xeb\xb0\x41\xcf\x11\xb8\x46\x68\x90\x95\x51\xac\x20\xb2\xf2\x46\x91\x81\x96\x42\x1b\x22\x9a\xcc\x7a\xb8\xfe\x69\x7d\x73\x76\x7e\xf1\xf5\x45\x9e\x65\xb7\x58\x2d\xa0\x66\x6e\xe3\xa2\x28\xb6\x96\xeb\xae\xcc\x75\x68\x8a\xe0\x2b\x17\x1e\x8a\xf4\x53\xba\x50\x16\x8d\x8a\x8c\x54\x54\xce\xb6\xb1\x78\x7f\xfe\xfe\xdd\xf9\xb7\xef\x2e\xce\x7c\xc5\x67\xfd\x66\x79\x63\xb2\xec\x8e\xa9\xd3\x1c\x41\x79\x03\x84\x31\x74\xa4\x31\x82\x56\xfe\x99\x22\x04\x8f\x10\x08\x9a\x40\x98\x0d\x4c\x79\xd7\x62\x3c\x05\xad\x9c\x43\x03\xf7\x16\x1f\x62\x0e\x1f\x94\xae\xd3\xff\x14\x06\xc2\x96\x30\x4a\x95\x99\x02\x63\xab\x0a\x49\xf0\xfe\xb0\xde\x40\xa8\x86\xaa\x4f\x21\x76\xba\x06\x15\x41\x81\x26\x54\x1c\x08\x4a\x1b\xb6\xa4\xda\x7a\x97\x05\x02\x05\x3f\xde\x7c\xf8\x1e\x6c\xa3\xb6\x08\x95\x75\x98\x67\x27\x45\x96\xb5\x5d\xf9\xac\xe8\xc7\x03\xd8\x2f\xc2\x04\x1e\xb3\x0c\x00\xa0\x28\xe0\x12\x6e\x31\x06\x77\x8f\x24\x9a\xde\x5b\x83\x11\x94\xd6\x18\x23\x70\x00\x05\x11\x79\xcc\xe5\x50\xc9\x61\xf5\x33\x48\x4c\x3a\x89\x0c\xbd\x4a\x30\xc3\x7c\x9b\x83\xf2\xf0\xf3\xf5\x66\xfe\x42\x32\x96\x6e\x5b\xcf\x48\x95\xd2\xd8\xc3\x70\xe8\x39\x8c\x28\x48\xff\xd3\xa6\xc0\xb5\x62\xb0\x0c\xb1\x6b\xdb\x40\x7c\xcc\x42\x8a\x1d\xb6\x1e\x90\x9f\x8b\x7b\x4c\x59\x7d\x66\xd5\x79\xd8\x22\x27\x31\x66\xf3\x05\xfc\xba\xd9\xb5\xf8\x79\x92\x42\xfb\xd5\x92\x36\xfb\x2d\xb1\x58\x80\x64\xce\x17\x70\xe9\x77\x7b\x6f\x7c\x97\x56\x3d\xbd\x22\xe8\x55\x70\x0e\x35\xdb\xe0\xc1\x4a\xef\xb6\x14\xba\x56\xc4\x4c\x0e\x38\x60\x93\xe8\x60\xf0\x4f\x28\x77\xb0\x5e\x7d\x49\x49\x23\xf8\x69\x71\x65\x20\x0a\x0f\x42\xbc\x4f\x9f\x59\xb3\x80\x4f\x6b\xcf\x17\xdf\xcc\x17\xf0\xd5\x63\xff\xfd\xe9\x35\x61\xd6\xab\xbd\x2c\xfb\xfc\xcf\x2f\x4a\x5c\xd9\xd8\x3a\xb5\xdb\x57\x55\xaa\x68\xf5\xc1\xd4\xa9\x41\x5e\xbb\x4e\x5c\x24\x8d\xf3\xaa\xc1\x74\x76\x0c\x46\x4d\xb6\x15\xb6\x3d\x4a\xa8\xc4\x1c\xa1\xfc\x1d\x35\xe7\xf0\x31\x44\x3e\xbc\x44\x88\x75\xe8\x9c\x79\xe9\x17\xd9\x63\x22\xd0\xc1\x78\x3d\xa5\x63\x25\x1c\x72\xa2\xb0\x80\x3b\x26\xeb\xb7\x93\xe0\x88\xd6\x90\x33\x24\x59\x6f\x79\x36\xbc\xc9\x33\xc6\x3a\x3d\x8a\xbc\x06\xd4\xc7\xe6\x23\x56\xf2\x44\x74\x55\x9e\x94\x59\x26\xc4\x69\x70\x84\x06\xcb\x89\x76\xfb\x66\x1c\xb7\xe4\x87\xcd\xe6\x66\x53\x77\x4d\xe9\x95\x75\x40\xc8\x1d\x79\xe9\x0e\x0f\xdf\x0e\xd3\x41\xc6\xc5\xa0\xfa\x1b\x62\x1e\x83\x4d\x25\xed\xc8\xbe\xa9\x68\x63\x1b\x94\xe1\xf6\x4f\xe4\x1c\xe1\x1c\xab\x39\x01\xf9\x5b\x29\x3b\xb2\xb0\x14\xb0\x69\xa8\x07\x82\xe5\x80\xf9\xa6\x84\xeb\x9b\xeb\xbb\x2f\x96\xb0\x5f\x1c\x39\x10\x9a\x34\xa3\xfd\x68\x10\x83\xf5\x09\xf7\xe5\xc0\x94\x6f\xfb\xb4\x08\x8a\xe4\x26\x48\xb3\x5f\xa3\x91\x39\xc0\x35\x5a\x4a\xa3\x5b\xbc\x6f\x0d\x7a\xb6\x95\x45\x82\xd9\xd5\x7a\x35\xef\x31\x48\x71\x8d\x24\x87\xce\x83\x5c\x1f\x84\x9a\xe1\xd3\xed\x3a\x87\x4b\xd0\xce\xca\x52\xd5\xb6\xce\x6a\x95\x7c\x24\x03\xb8\x8b\xb8\x3f\x4a\x57\xeb\xd5\x68\xe8\x56\x72\xe1\xc8\x29\x75\x41\x99\x74\x6e\xf7\x15\xdc\x5b\x25\xe5\x24\xb2\x5b\xc5\xf8\xa0\x76\x6f\x59\xe6\x58\xbc\xa9\x65\xb4\x4c\x9f\xff\xc1\x32\x23\x9c\xff\x62\x19\x6d\x0d\x2c\x05\xec\xdf\x59\xe6\x29\xfb\x2b\x00\x00\xff\xff\x30\xea\x8c\xb8\xb4\x08\x00\x00"

func metadataviewsCdcBytes() ([]byte, error) {
	return bindataRead(
		_metadataviewsCdc,
		"MetadataViews.cdc",
	)
}

func metadataviewsCdc() (*asset, error) {
	bytes, err := metadataviewsCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "MetadataViews.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xca, 0x82, 0x90, 0x83, 0xe, 0xe3, 0x36, 0x70, 0x31, 0x5d, 0x1a, 0x23, 0x7c, 0x31, 0x8e, 0xb0, 0xc3, 0x5d, 0xf2, 0x2b, 0xc8, 0x61, 0x58, 0xbd, 0x11, 0x78, 0x1c, 0x45, 0xd0, 0xa2, 0x32, 0xe4}}
	return a, nil
}

var _nonfungibletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x41\x8f\xdb\xba\x11\xbe\xeb\x57\xcc\xcb\x03\x9a\xdd\xc0\x6b\xf7\x50\xf4\x60\x20\x68\xda\xb7\x6f\x01\x5f\xb6\x0f\x5b\x17\x3d\x04\x01\x4c\x8b\x23\x9b\x08\x45\x2a\x24\x65\xc7\x0d\xf6\xbf\x17\x33\x24\x25\xca\xf6\x26\x9b\x5b\x73\x89\x57\x22\xbf\x99\xf9\xe6\x9b\x8f\xd4\xe2\xdd\xbb\xaa\xfa\xf5\x57\x58\xef\x11\x1e\xb4\x3d\xc2\xa3\x35\x77\x0f\xbd\xd9\xa9\xad\x46\x58\xdb\xcf\x68\xc0\x07\x61\xa4\x70\x92\x17\x6e\x1e\xad\xc9\xef\xf9\xf5\x06\x6a\x6b\x82\x13\x75\x00\x65\x02\xba\x46\xd4\x58\x55\x84\x37\xfc\x09\x61\x2f\x02\x08\xad\xc1\x58\x73\xd7\x64\xf4\xc0\xe8\x79\xb7\x87\xda\xf6\x5a\xd2\xdf\x8d\x75\x2d\x04\x3b\xaf\x56\x0d\x08\xe8\x3d\x3a\x38\x0a\x13\x3c\x04\x0b\x12\x3b\x6d\x4f\x20\xc0\xe0\x11\x4c\x13\x86\xfd\x33\x08\x7b\x54\x6e\xcc\xe6\xc8\x70\x06\x51\x56\xc1\x82\x6a\x3b\x8d\x2d\x9a\x40\xcb\xe0\xbc\x88\x31\xd7\x39\xe7\x7e\x89\xb3\x17\x07\xca\x18\x1a\xab\x89\x26\x2a\x86\x80\x5c\xaf\xd1\x83\x30\x12\x8c\x68\x95\xd9\x55\x5c\x6a\x98\x54\xef\x3b\xac\x55\xa3\xd0\xcf\x13\x83\x0f\xeb\x0d\x38\xf4\xb6\x77\x99\xaa\xda\x3a\x1c\x1e\x41\x38\x75\x89\x33\x87\x9d\x43\x8f\x54\xbb\x30\xf0\xf8\xb0\x06\x65\x18\xdd\xb7\xc2\x8d\xb5\x27\xe0\xdf\xac\xd6\x58\x07\x65\xcd\x06\x9e\x26\xf8\x23\x34\xa1\xfa\x60\x1d\x65\xcd\xd4\xbe\xf5\x8c\x5b\x0f\x7b\xe7\xd5\x8a\x5a\x59\xeb\x5e\xf2\xa2\x06\x8f\xd0\xf4\x86\xdf\x71\x0b\x04\x33\x40\x59\xd8\xa3\x41\x47\x8f\x50\x78\xa5\x4f\x55\x6b\x0f\xa9\xad\x9e\x12\x25\x5a\x6c\x1f\xc0\x36\xbc\xba\x0c\xc1\xf9\xfe\xe1\xec\x41\x49\x74\x1b\x5e\xb9\x79\xc2\x1a\xd5\x81\xfe\x1c\xd2\x1d\x48\xf4\x5c\x87\x2f\x9f\x80\xc4\x5a\x0b\x87\x45\x72\x47\x15\xf6\xe0\x6d\x8b\xd0\x39\x64\xd0\xce\x7a\xa6\x49\x2a\x5e\x51\x25\x56\xbf\xf4\xca\x21\x27\x35\x72\x56\x74\xb7\x46\x17\x84\x32\xa9\xa7\x0c\xb4\xc5\xbd\x38\x28\xeb\x86\x69\xf0\x51\x29\x27\xa0\x14\x3c\x76\xc2\x89\x80\xb0\xc5\x5a\xf4\x94\x66\x80\x9d\x3a\xa0\xe7\x18\xac\x60\xfa\x21\xb6\x4a\xab\x70\xa2\x48\x7e\x4f\xfb\x04\x38\x6c\xd0\xa1\xa9\x91\x44\x1a\x15\x5c\xa6\x44\xe9\x5a\xa3\x4f\x80\x5f\x3b\xeb\x13\x5e\xa3\x50\xcb\xa8\xba\xb1\x76\x65\xc0\x1a\x04\xeb\xa0\xb5\x0e\xab\xc4\xf9\x48\xd7\x1c\x56\x34\x83\xde\xa6\xc4\x28\x29\x7f\x9e\x55\x2b\x3e\x23\xd4\xbd\x0f\xb6\x1d\x9a\x90\x48\x9b\x0c\xd0\xb4\x11\x34\x96\x16\x0e\xc2\x29\xdb\x13\xa4\x32\xbb\xd4\x0b\x82\x8f\x7a\x98\x57\xd5\x3f\x4e\xd0\x7b\xe2\x73\x40\xe6\x12\x46\xa0\x59\x4a\xca\x36\x2c\xc9\xa9\xc6\x3d\xd4\xc2\x80\x47\x23\x2b\xda\xe5\xa2\x58\xb2\xda\x3a\x44\x77\x17\xec\x1d\xfd\x3f\xe3\xd8\x24\x3c\x6a\x99\xd9\x51\x7e\x1c\x84\xa7\x99\xd2\x12\x50\x23\xa1\x6a\xd0\x28\x77\xe8\xaa\x8b\x71\x5a\x5b\x0e\x95\xa7\x8e\x54\x6f\x6c\xd8\xa3\xe3\x14\x67\x83\x2d\xb1\x37\x78\xe2\xe6\xc4\xd0\xd2\x89\x38\x1a\x8f\x0f\xeb\xaa\x71\xb6\xbd\xe8\x29\xfb\x94\x81\x3a\x3b\x88\xc4\xce\x7a\x15\x86\x4e\x82\x35\x93\x58\x6f\x7d\x35\xd5\x68\x6d\xa9\x13\x21\xca\x37\x38\x61\x7c\x83\x6e\x5e\x55\xef\x16\x55\xb5\x58\xb0\x93\xb7\x24\xde\x38\xd5\xe7\xd6\x3c\x87\x7f\x32\x74\xf9\x96\x9a\xa5\x35\x6d\x56\x6d\x67\x5d\x88\x6d\x29\xfa\xad\x7c\xe1\xed\x8b\x45\xd5\xf5\xdb\x2b\xd0\x97\xae\xfa\xad\xaa\x00\x00\x52\x56\xc1\x06\xa1\xc1\xf4\xed\x16\x1d\x7b\x42\x6c\x1d\x2b\x55\xf9\xe8\x7a\xca\x00\x7e\x55\x3e\xf0\x44\xd0\x5e\x0a\x75\x10\x2e\x6e\xfe\x57\xdf\x75\xfa\xb4\x84\x7f\xaf\x4c\xf8\xeb\x5f\x06\xf0\xdf\x0f\x31\x4d\x11\x00\x5b\x15\x02\x4a\x38\x12\xc7\xa9\x0f\x45\xaa\x54\x87\x0a\x4a\x68\xf5\x5f\x94\x69\xfb\x10\x06\x19\xe6\xb7\xb4\x78\x35\x2e\xbc\xb9\xbd\x16\x4a\xf9\x69\x34\x91\x0e\x34\xe5\x07\x25\x98\x59\xde\xa7\x8c\x54\xb5\x08\xac\xc6\xc1\x38\x2f\x7c\x31\x01\x07\x38\x8a\x02\x04\x48\x47\xf3\x32\xdb\xc5\x02\x56\x17\x7b\x95\x07\x63\x43\xf4\x5d\x10\x75\x6d\x7b\x13\xde\x7a\x36\x7b\xb1\xc3\x19\x6c\x08\x66\xc3\xad\x86\x2d\xc2\xc6\x28\xbd\x99\x5f\xe7\xe0\x3f\x29\xf4\x8d\x92\x99\xec\x19\x67\xb1\x84\xbf\x4b\xe9\xd0\xfb\xbf\x5d\xa5\xe4\x25\x3e\x92\xc6\x51\xf2\x20\x4d\x0e\x82\xb3\xaa\x42\x66\x2a\x59\xdd\x6b\x88\x2a\xd1\x5f\x28\xe8\x3e\x2e\x99\xd4\x13\xec\xb5\x6a\x56\xd3\x4b\x4b\x92\x90\x1f\xce\xff\xf1\x7a\x72\x1e\xe9\xf2\xd0\x82\x15\xa9\xef\x1b\xaf\x28\xe6\xa0\x37\xea\x4b\x8f\xb0\xba\x4f\xa4\x89\x7a\xcf\x32\xdd\x0b\x3f\x2c\x25\x40\x8d\x01\xc6\x84\xf9\xd5\xf3\x90\xe7\x53\x3c\xc3\xda\x81\x7b\xf2\x93\x94\x1c\xa9\xec\x9a\x81\x52\x0d\x79\x3f\x5f\xa5\x1a\x65\xe2\x19\x94\x32\x27\x53\x42\x19\x1d\x8f\x30\x13\x1e\x3b\x3c\xd5\x72\x59\xeb\xe3\xc3\x7a\x79\x5e\xe6\x0f\x73\x2f\x38\xb6\xd0\xa2\x54\x74\x72\x66\xb9\x7b\xc8\xb6\x59\x98\xe6\x2b\xb8\xce\x97\x89\x29\xdf\x83\x27\x3b\xa4\xcb\xc9\x70\x8d\x1a\x62\x14\x9a\x22\xd7\x8b\x8b\x54\x80\x78\x1a\x47\x46\xdc\xa4\xb4\xa6\x37\x03\xec\x4d\xfe\xb1\xba\xcf\xb5\xde\x2e\xe1\xc3\x94\x0f\xde\x48\xf7\x90\xe9\x23\xfa\xe7\xd0\xf7\x3a\xcc\x95\x84\xf7\xef\xa1\xc4\x7a\x43\x42\x59\xdd\x67\xe5\x8f\x5e\x10\x67\xaa\xed\x7d\xa0\x21\xe6\xab\xa0\x68\x11\x44\x1c\x17\xba\xd9\xa0\xa7\x51\x58\xdd\xbf\x99\x44\x7b\xae\xa6\xbf\x7e\xd0\x8d\x34\x53\x3e\xf3\xf0\x53\xad\xc8\x17\xb9\xec\xff\x29\x50\x3e\xe9\x82\xf8\x3c\x36\x42\xf0\x2f\xe1\x76\x3d\x4b\x99\x7a\x20\xa4\x2c\x5b\x70\x16\xba\x08\x5f\x76\x24\x81\xdf\x30\x3f\xb1\x05\xb7\x2f\x17\xca\x03\x33\xb8\x64\x3a\xc6\x6b\xdb\xb6\x7c\xd7\xca\x1b\xba\x7e\xab\x95\xdf\x43\x63\xdd\xf0\x71\x31\xc9\xe5\x85\xfa\xc7\x8c\xff\x20\x84\xfa\x6c\x36\xbe\x9b\x6e\xb9\x68\x87\x61\x75\xef\x6f\x6e\x97\xf0\x31\x6a\xeb\xd3\xc5\x92\xad\x75\xce\x1e\x1f\x1f\xd6\x85\xb5\xdd\x2e\xe1\x4f\x79\x58\xaf\x1b\x46\x2a\x28\x0d\x80\xa9\x1d\x5d\x27\x26\x9f\x1f\x85\x4d\x6c\x31\xdf\xb4\x65\xfe\xfa\x18\xee\x06\xe4\x34\xd9\x5f\x5e\x14\xc6\x48\xc7\x72\x98\xd2\xd9\x20\x92\xd9\x35\xba\x4a\xd9\xdc\x2b\x7e\x27\x1c\xdf\x50\xf7\x56\xcb\xd1\x95\x53\x3e\x57\x24\x92\xef\x0d\x74\x80\x48\x5a\xbb\x84\x0f\xdf\x22\x3f\x4b\xda\xfb\x5c\xfd\x5f\xd8\xc4\xf7\x06\x24\xce\xc7\xe5\x40\x8c\xb9\x78\x90\x03\x39\x25\xd0\xb0\x29\x44\x17\x49\x1b\x95\x04\xe1\x9c\x38\xbd\x4e\x8d\x25\x60\x54\x22\x38\x0c\xbd\x33\x69\x62\x9d\x38\x65\x7b\xa2\x77\x71\xa6\x1c\xe6\x9e\xd4\xd7\x7b\xf2\x82\xae\xcb\x60\x4f\x39\x4a\x52\x37\xca\xf1\x2b\x29\xde\xc4\xcb\x2f\xe1\x2b\x71\x16\x0b\xf0\x76\x3c\xbf\x63\x73\xf8\xf3\xc1\xa1\x90\x20\x45\x10\x4c\x11\xdf\xc1\x5b\x0c\x7b\x2b\xd3\xa9\xa3\xc2\xcf\x4c\xd8\xb9\xc7\x3b\xbc\x62\xf1\x1e\x75\x33\x1f\x54\xf8\x51\xc9\x4f\xf0\xcb\x7b\x30\x4a\x2f\xe1\x0d\x61\x48\x8b\xf1\xe2\xc6\xf7\xde\xcb\xaa\x7e\x79\xad\x8f\xd7\x0e\x45\xc0\xdf\xdb\x2e\x9c\x8a\x0f\x86\xf8\x94\x5b\x86\xf4\xea\xd2\xc9\x21\x7e\x4e\x45\xce\xcf\x25\x5d\x12\x79\x62\x0a\xed\x91\xe9\xf7\x55\x49\xd2\xd5\xd8\xd4\xe0\x0f\x45\x2a\x85\x0b\x5e\x9e\x86\xe9\x24\xcc\xd2\x98\x6b\x34\xbb\xb0\xa7\x63\xf1\xcf\xe9\x34\x8c\x31\x64\x39\x8a\xf9\x18\xe4\xca\x0a\xa2\x32\x35\xcf\xd5\xff\x02\x00\x00\xff\xff\x33\x4d\x81\x27\xe0\x12\x00\x00"

func nonfungibletokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_nonfungibletokenCdc,
		"NonFungibleToken.cdc",
	)
}

func nonfungibletokenCdc() (*asset, error) {
	bytes, err := nonfungibletokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NonFungibleToken.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdb, 0x61, 0xca, 0x9d, 0xaa, 0x66, 0x36, 0xdf, 0xbc, 0x51, 0xdb, 0x7b, 0x51, 0xd8, 0x3d, 0x6f, 0x4e, 0x9c, 0x8e, 0x50, 0x28, 0x7c, 0x18, 0x1d, 0x2, 0xb2, 0xc2, 0x2b, 0x26, 0xa1, 0xfe, 0x2d}}
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
	"ExampleNFT.cdc":       examplenftCdc,
	"MetadataViews.cdc":    metadataviewsCdc,
	"NonFungibleToken.cdc": nonfungibletokenCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
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
	"ExampleNFT.cdc": {examplenftCdc, map[string]*bintree{}},
	"MetadataViews.cdc": {metadataviewsCdc, map[string]*bintree{}},
	"NonFungibleToken.cdc": {nonfungibletokenCdc, map[string]*bintree{}},
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
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
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
