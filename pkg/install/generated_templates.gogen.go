// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 836,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4b\xaf\xd3\x30\x10\x85\xf7\xfe\x15\x47\xba\x8b\x0b\xe8\x26\xa8\x3b\x94\x5d\xdb\x05\x0b\x10\x8b\xf0\xd8\x20\x16\x63\x7b\x42\x4d\x5d\x3b\xf2\x23\x3c\xa2\xfc\x77\x94\xa4\x95\x9a\xb6\x20\x55\xba\x3b\x7b\x7c\xc6\x73\xe6\xe8\x2b\x8a\x42\x3c\xe0\xd3\x8e\x11\x39\x74\x46\x31\x48\x29\x9f\x5d\x7a\x82\xb2\x39\x26\x0e\x08\xde\x72\x7c\x02\x39\xbd\x28\x41\x1a\xa7\x8d\xfb\x0e\x0a\x2c\x1e\xe0\x9d\xfd\x0d\xc7\xac\x59\xa3\xf1\x01\xef\xb2\xe4\xe0\x38\x71\xc4\x4f\x93\x76\x53\x4b\x21\x29\xb2\x1e\x27\x70\x8c\x50\xde\xa5\xe0\x2d\x5e\xd4\x9b\xf5\xf6\x65\x29\xa8\x35\x5f\x38\x44\xe3\x5d\x85\x6e\x25\xf6\xc6\xe9\x0a\x1f\x67\x57\xeb\xd9\x94\x38\x70\x22\x4d\x89\x2a\x01\x58\x92\x6c\xe3\x78\x02\x1c\x1d\xb8\x42\x63\xf3\x2f\x71\x7e\xe9\x7b\x98\x06\xe5\x07\x3a\x70\x6c\x49\x31\x86\xe1\xf8\x3e\x5d\x2b\xf4\xfd\xf2\xb5\xef\xc1\x4e\x0f\x83\x18\x73\x39\x37\x14\x24\xa9\x92\x72\xda\xf9\x60\xfe\x50\x32\xde\x95\xfb\x37\xb1\x34\xfe\x75\xb7\x92\x9c\xe8\xe4\x77\x3b\x27\x54\x7b\xcb\xf7\x9a\x15\x21\x5b\x9e\x24\x05\xa8\x35\x6f\x83\xcf\x6d\xac\xf0\xf5\xf1\xd5\xe3\xb7\xa9\x2f\x70\xf4\x39\x28\x5e\x14\x3b\x0e\xf2\xac\x50\xc0\x79\x57\x1f\x85\x9f\xeb\xf7\xff\xd6\x3e\xc3\x86\x9b\x99\x80\xfb\x17\xf5\x96\x6b\x6e\x46\xd1\x69\xd1\xff\xcc\x17\xc0\x75\xb6\x8b\xff\x62\x96\x3f\x58\xa5\x63\x76\x37\xc1\xb9\xb2\x73\x89\xc1\x25\x27\xb7\xc8\xb0\x71\x3c\x69\x6e\x28\xdb\x34\xa3\x32\x12\xf5\x37\x00\x00\xff\xff\xfd\x7f\x67\x6a\x44\x03\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 6872,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\xdd\x6f\x1b\x37\x12\x7f\xf7\x5f\x31\x50\x0e\x48\x0c\x48\x2b\xbb\x6e\x8b\xc3\xf6\x5c\x5c\x9a\x0f\xd7\x97\xc6\x31\xe2\xf8\x0e\x79\xaa\x29\xee\x48\x4b\x88\x4b\xee\x71\xb8\x52\x05\xa3\xff\xfb\x61\xc8\xfd\xe0\xca\xb2\x5d\xe4\xed\xfa\xd0\xd8\xdc\xe1\x7c\x7f\xfc\x38\x9e\xcd\x66\x47\xa2\x56\xff\x46\x47\xca\x9a\x1c\x44\x5d\xd3\x7c\x73\x7a\xb4\x56\xa6\xc8\xe1\x2d\xd6\xda\xee\x2a\x34\xfe\xa8\x42\x2f\x0a\xe1\x45\x7e\x04\x60\x44\x85\x39\x2c\x75\xf3\xc7\xfd\x3d\xa8\x25\x64\x57\xa2\x42\xaa\x85\x44\xf8\xf3\xcf\xf6\x7b\xf8\x35\x87\xfb\xfb\xf1\xd7\xfb\x7b\x40\x53\x30\x19\xd5\x28\x99\x99\xc3\x5a\x2b\x29\x28\x87\xd3\x23\x00\x42\x8d\xd2\x5b\xc7\x5f\x00\x2a\xe1\x65\xf9\x9b\x58\xa0\xa6\x78\x90\xca\x66\x6a\xef\x84\xc7\xd5\x2e\x7e\xf4\xbb\x1a\x73\xf8\x8c\xd2\xa1\xf0\x78\x04\xe0\xb1\xaa\xb5\xf0\xd8\x32\x4b\x2c\xe0\xff\x84\x31\xd6\x0b\xaf\xac\xe9\x99\x03\xd4\xce\x56\xe8\x4b\x6c\x28\x53\x76\x5e\x5b\xe7\x73\x98\x9c\x9d\x9c\x9d\x4e\xe0\x05\x78\xd4\x3a\xa1\x00\x6f\x81\xa4\x13\x35\xc2\xbc\x42\xef\x94\x24\x36\xae\xb6\xca\xf8\x97\x04\x7c\x39\x6b\x19\xeb\x91\x0d\x7b\x56\x00\x74\xbe\x08\x3f\xa3\xdb\x28\x89\xaf\xa5\xb4\x8d\xf1\x57\x63\x42\x80\x8d\xd5\x4d\x85\x3d\xab\x59\xcb\x6a\xa5\xfc\x6c\x8d\xbb\x5e\x00\xb1\x17\xfc\x20\xb0\x3b\x19\xf8\xcd\xf8\x4a\x11\x02\x9c\x50\x15\xb8\x14\x8d\xf6\x1f\x6d\x81\x39\x9c\x7c\x7f\x72\x02\x2f\x60\x5b\xa2\x81\x8a\xb5\xc1\x02\x1c\x8a\x62\x66\x8d\xde\x4d\x61\x8b\xb0\xb5\xe6\xa5\x87\x05\x82\x58\x68\x64\x7f\xc8\xb2\xb2\xc5\x51\xcb\xf0\x05\x7c\x29\x15\x81\x22\x10\xe0\xab\x7a\x49\xd0\x10\x16\xb0\xb4\x0e\x56\x68\xd0\x09\xaf\xcc\x0a\x6e\x6e\x7e\x85\x35\xee\x28\x83\x4b\x03\x1f\xfe\x4e\xf0\xf3\x39\x9c\x66\xa7\x27\xd3\x9e\x4b\x27\x3b\x9a\x40\x20\x1c\xa6\x7a\x90\x65\x55\x0c\x62\x01\x02\x08\x6b\xc1\x49\xd1\x3a\x0a\xb6\xd8\xb3\x91\xc2\xc0\xd6\x29\xcf\x8a\x66\x87\xfd\xb7\x42\xd3\x3b\x03\xab\xda\xef\xde\x2a\x97\x3a\xb1\xc2\x42\x35\x55\x0e\x1f\xb1\xb2\x6e\x97\xda\x89\xb0\xb4\x5a\xdb\x2d\x5b\xd4\x8a\x56\x14\x4c\x6d\x88\xcf\x04\xc8\x86\xbc\xad\x14\x7b\x60\x6d\xec\xd6\xfc\x5e\x5a\xf2\xd4\xb3\x58\x2a\x8d\x53\xd8\x96\x4a\x96\xb0\xb3\x0d\x6c\x95\xd6\xd1\x28\x6f\xa1\xb0\x5c\x67\x7c\xcc\x97\xf8\x07\x07\x76\x6b\x58\xed\x9e\x81\xc3\xda\x82\x13\xbe\x44\x07\xbe\x14\xa6\x15\xbc\x52\xbe\x6c\x16\x60\xf9\x10\x41\xab\x35\x66\xf0\xd5\x36\x2f\xb5\x06\xa1\xc9\x76\x22\xc6\xce\x06\xe5\x41\x19\x6f\xc3\x1d\x69\x8d\x17\xca\xa0\x9b\xc2\x02\xb5\xdd\x66\x70\x83\x83\x57\x4b\xef\x6b\xca\xe7\xf3\xc2\x4a\xca\x38\xb1\x64\xc1\xa5\x83\x66\xce\xa5\x47\x7e\xbe\x6a\x54\x81\x34\x6f\x08\x67\xb5\x53\x1b\xe1\x31\xa4\x1e\x1b\x92\x95\xbe\xd2\x3d\xa7\x2e\x16\x44\xe5\x4c\x5a\xb3\x54\xab\xfe\x13\x40\x3c\xf8\x28\xea\x3c\x39\x4c\x0b\x69\x96\x5c\xfb\xd6\xb8\x64\xeb\x66\x81\xf3\xc8\x64\x48\xbf\x67\x63\xb2\x55\x54\xf2\x49\x29\x36\x08\x02\x0a\xb5\x5c\xa2\xe3\xa6\xd9\x71\x68\xab\x6a\x68\x8c\x21\x04\x91\x5d\x1a\x04\x6e\x2e\x1b\x55\x60\xe7\xf6\xa5\x5a\x55\xa2\x1e\x14\x51\xbe\x04\x61\x00\x8d\x77\xbb\x60\xc3\x5d\x24\xba\x9b\x82\x30\x05\x34\x46\xda\x8a\xbb\x75\xb8\x1f\xad\xfd\x18\xc2\x29\x4c\xd1\x73\x41\xb3\x09\x1c\x14\x52\x1b\xcf\x07\x11\x60\x37\x7c\x43\x04\x92\x6b\xcf\x46\x20\x74\x02\x6f\x41\x55\xdc\x27\xe1\xe2\xfa\x22\x34\x01\x78\xc5\x66\x91\x5a\x19\x65\x06\xe1\x6c\xdc\x06\x9d\x5a\x2a\x19\x1a\x36\xd4\x8d\xab\x2d\x21\x1d\xff\x05\x47\xf6\x5c\x62\xfb\x88\x5e\x64\x07\xb1\xbc\xbf\xe0\x38\x10\x6e\x35\x94\xe9\x23\x1e\x5b\xd5\x2b\xee\x1f\x94\xb8\x66\xdc\x82\x5f\x3c\xd2\x84\x1f\xde\x3b\xd0\x84\x3b\x77\xf6\x95\xf8\xa0\xff\x27\x13\xa2\xf5\xba\xc3\xd0\x27\x8d\x85\x49\x1e\x2b\x71\x02\xaa\x12\x2b\x8c\xd9\xcf\x17\x32\x78\xaf\x4c\x11\x6c\xae\xb8\xad\x38\x94\x43\xd6\xc6\x96\xa2\x51\x10\x72\xf3\x08\x57\x39\x08\x8c\x13\x40\xf8\xbe\xee\xcb\x66\x91\x15\x56\xae\xd1\x65\xd2\x56\x73\x37\x8f\x3d\x20\xfc\x33\xf7\xa2\x77\x5d\x17\x47\x9e\xf7\x8c\x05\x58\xaa\x17\x2b\x60\x4d\xb3\x9e\x26\x88\xc9\xa1\x65\xa8\x6c\xca\x2d\x3f\xcd\x4e\x7f\xc8\x4e\xc6\xb4\xd7\x8d\xd6\xd7\x56\x2b\xb9\xcb\xe1\x72\x79\x65\xfd\xb5\x43\x4a\xad\x70\x48\xb6\x71\x12\x29\xed\xe3\x0e\xff\xdb\x20\xf9\xd1\x19\x80\xac\x9b\x1c\x7e\x38\xa9\x46\x87\x55\x68\xf5\x39\xfc\xf8\xfd\x47\x35\xc0\x04\xeb\xd2\xcb\xb3\x21\x32\xd7\x01\x32\x9c\x9d\x9c\xf1\xe4\x54\x66\x69\x5d\x15\x52\x56\xe8\x9e\x5a\xab\x0d\x1a\x24\xba\x76\x76\x81\xa9\x06\xec\xd2\x8b\xf1\xd4\x8e\xa2\x22\xc3\xf1\xb1\xf0\x65\x0e\x73\x51\xab\xe8\xe9\xcd\x8f\x73\x55\xa0\xf1\xca\xef\xb2\xba\x59\x24\xb4\xca\x28\xaf\x84\x7e\x8b\x5a\xec\x6e\xb8\x3e\x0b\xca\xe1\x87\x84\xc0\xab\x0a\x6d\xe3\x0f\x7c\xe3\x21\xab\xfe\x3f\x54\x4d\x8a\x76\x14\x98\xc3\xf0\x08\xe2\x98\xbb\x8e\x9a\xa1\x97\x41\xb3\x62\x4e\x54\x32\xce\xb3\x11\x79\x82\xb6\x6d\xbf\x59\x71\xc8\x40\x99\x98\x73\x2f\x29\xde\x21\x2a\xe7\xa3\x36\xd9\xf9\xec\x93\xd1\xbb\x1c\xbc\x6b\x90\xb9\x31\x06\x0a\x1d\x6a\xd1\x36\x76\x2e\xa9\x1a\xdd\xd2\x3a\x89\xcc\x34\x82\x1e\xc6\x3c\x8f\x29\x9e\xe2\x92\xb1\xee\x1b\xe1\x5a\xdd\x23\xd9\xb7\xa9\x9f\xd4\xe8\xa5\x91\xba\x09\x9d\x93\xa1\x5b\x1c\x70\x5d\x57\x8d\xd8\xe0\x19\x28\xd3\x81\x99\x9f\xf8\xea\x1e\xcc\xe8\xbb\x2b\x14\x28\xb5\x70\x0c\xd9\x16\x76\x93\x34\x80\x27\x60\x40\x6c\x8f\xa9\xf1\xce\x5a\x3f\xcf\x88\xca\x47\x0d\x10\x66\x24\x75\x32\x8c\xa8\x49\x94\x3c\xed\x48\x12\x0e\x68\x36\xca\x59\x13\x06\x42\x9c\xb5\x93\x0f\xb7\xbf\xbc\x7b\xf3\xe9\xea\xfd\xe5\xc5\x24\x8e\x80\x29\xfb\xc3\x6e\xd0\xb9\xf1\xbc\x4e\xd8\x84\x11\xb7\xd8\xc5\x69\xea\xf5\x21\x1b\x1f\x0c\xda\x87\x36\x0e\xc9\xc9\xc4\x8f\x1a\xca\x33\x8f\x1f\x1e\x9d\x34\x6e\xd1\x09\x14\x69\xb5\x0b\x31\x49\x58\xec\x03\x9a\x34\xe8\x01\xcd\x74\xd0\x5b\x18\x10\xda\xa3\x33\x0c\xad\x1f\x68\xbc\x74\xb6\xe2\xb4\xe8\x10\xcb\x14\x04\x71\xba\xb5\x53\x95\xdd\xa0\xad\x5c\xd3\xc3\x60\xa3\xd9\xe4\x07\xfc\x32\xb8\x7b\xe4\x97\x8d\xd0\x0d\x3e\xf0\xc9\x73\x49\xbc\x9f\x03\xdd\xcc\x7d\x22\x03\x78\xe4\x8f\x47\xfd\x13\xc3\xfe\x91\xbc\x64\xaa\x88\x6e\x46\x74\xe3\xfe\xf0\x5c\xe5\x6d\x05\x83\x12\x0b\xd4\xd4\xb5\xde\xc1\xaf\x5f\xbe\x5c\xc3\x42\x90\x92\x20\x1a\x5f\x82\x74\x18\x3a\xa9\xd0\x71\xaa\x0f\xef\x01\x66\xb8\x51\x22\x18\x7e\x77\x71\xf9\xe5\xf7\xd7\xb7\x5f\x7e\xbd\xbd\x79\xf7\xf9\x2e\x98\xdb\x1f\x7d\x78\xf7\xf5\x6e\x94\xf0\x1b\xe1\x14\xbf\xe6\xa8\x03\xc8\x09\xc3\x08\x5f\xf6\xe2\xf7\xde\xd9\x6a\x1c\xc3\x48\xf6\x19\x97\xf9\xc8\xf2\x11\x56\xe4\xc6\xc6\x26\x0c\x0e\x60\x9f\xe7\x23\x7f\x44\x17\xc4\x37\x2a\x16\x3c\x89\xa5\x90\x25\x16\x9c\x5a\x69\x6e\xf7\xb0\x9a\x3d\xc5\xdc\xa7\x09\x17\xeb\x5a\xdc\x9c\x5c\x68\xdf\xd8\xe1\xe2\x34\x08\xe1\xb7\x61\xeb\x63\x5f\x22\xa5\xb9\x30\xa0\x57\xbf\xb5\xac\x65\xc3\x7e\x0a\x15\x17\x16\x02\x21\x11\xa1\xb4\xdb\xf0\xfe\xb5\xc6\xa0\x0c\x21\x53\x7e\x9c\x3b\xb3\x59\x6f\x40\x78\xfc\xb0\xf0\xf3\xfe\x28\x6b\x41\x5f\x46\x1b\x99\x49\xdd\x90\x47\x97\x71\x03\xd7\xa9\x4b\x6e\x29\xf6\x9a\xc1\x15\x6f\x22\xe9\xe5\xf5\xc8\x28\x6e\x3b\x84\x3e\xbc\xaf\xc7\x99\x3d\xe8\xd0\xd1\x73\x76\x79\xc7\x94\xe1\xc5\x9b\x8c\xa0\x54\xe3\x96\xfa\xfc\x68\x84\x32\x15\x41\xd5\x50\xd8\x00\x04\xef\x29\x2c\x62\x39\x2d\xc2\x60\x0b\x18\x2f\x3c\xfc\x5f\x75\xaf\xe9\xe3\x54\x97\xae\xb9\xc4\x32\xe4\x04\x4e\xde\xff\x23\x45\x78\x18\xc4\x01\x37\x2b\x94\x3b\x7f\x30\xf6\x52\xb5\x3e\x27\x08\x73\x08\xde\xed\xe7\xdf\xe2\x82\x42\x98\x55\xfc\x76\xa1\x7c\x78\x34\x93\xf2\xd6\xed\xfa\x76\xfd\x9e\x91\x71\xc2\xee\xa9\x9a\xe3\xb4\x49\x6c\x6f\x4b\xe6\x60\x39\xa5\xb5\xd0\x61\xe7\xbf\xbd\x4a\x2b\xf3\x38\x1f\x7e\xff\xf0\xee\xeb\xf1\x3f\xe3\xd3\x3d\xc0\xea\x86\xd0\xcd\x07\x65\xb3\xb4\xd0\xd9\x3f\x5c\x4e\x8d\xd3\xe7\xf7\xf7\x90\x5d\x28\xcf\xc6\x86\x55\xdc\x98\x62\xe1\x84\x91\x65\x47\xf4\x4b\xf8\x2d\x2e\xe5\xd4\x32\x1c\x71\xff\xa2\x43\x37\x19\xc3\xf1\xbd\x9b\x90\x29\xf4\x2f\xab\x4c\x72\x61\x32\x9d\xb4\xbb\x3d\x4d\x98\x5e\x7f\xba\xa9\x39\xe4\xc4\x93\xf1\xd5\x55\x09\xa3\x96\x8c\xc9\xb9\x86\x48\x15\xe8\x62\x38\xf6\x5e\x36\x61\x27\x61\x09\xa1\x31\x05\xba\xbd\x18\x3b\xd4\xc2\xab\x0d\x06\xc8\x49\x5d\x06\xae\x46\x71\xde\xab\xc9\xde\x38\x6a\x16\x85\x72\xa7\xd3\xf8\xef\x77\xfd\xa2\x72\x70\x4e\x58\x44\x1e\x72\x4e\xd8\xee\x75\x5e\xed\xa8\x0e\x30\xb8\x25\x74\x87\xee\x73\x70\xfb\xc8\x31\x0d\x1c\xbe\xff\xae\x12\xea\xa0\x02\xc8\x1f\x3a\x0e\x1d\xd5\xb0\x6a\x3d\x18\x0e\xe4\x56\xb2\xb5\xec\x50\x34\x61\x7d\xc7\x7e\xe2\x89\xad\xfc\xde\x03\x3c\xf5\x55\x3b\xfb\xda\xc9\x76\xfe\xc4\xa8\xeb\x6e\xb4\xbc\xf8\xd6\xf9\x3f\xd6\xb8\x03\x55\xfc\xdc\x93\x3d\x01\x67\x12\xad\x98\x85\xf0\x8d\xc3\xd1\x16\xe0\x80\xac\xf0\x79\x37\xeb\xe9\x69\xd4\xae\xba\x6e\x0d\xca\x43\x29\x28\x8c\x62\x6b\xf4\x0e\x84\x94\x48\xb1\xa3\x97\x18\x17\x69\xaf\xba\x9d\xcd\xdd\x52\x68\xc2\xbb\xe3\x03\xd2\xba\xfb\x63\x07\x93\x77\x8d\xf4\x51\xd0\x36\xbc\xc3\x19\x9b\x35\x1e\x68\x67\x24\x2c\xac\x5d\xaf\x11\x6b\x4e\xd7\x5e\xc6\x64\xa5\xfc\x64\x0a\x15\x0a\xf6\x14\x77\x22\x10\xe1\x71\xdc\x66\x70\x53\x93\x77\x28\xaa\x3e\x95\xf7\xb5\x61\xd6\x33\xf2\xc2\xe3\x39\x77\x86\x47\x03\x6e\xf0\x0f\xdf\x45\x3d\x19\x55\xc2\xc0\xa4\x93\x31\xe9\x06\x49\xc2\xe4\x15\x66\xab\x6c\x0a\xff\x41\x86\x84\x6f\xb4\x6d\x8a\xe3\x2c\x6c\x76\xbc\x5d\xf3\xc3\x82\xa0\x16\xce\x2b\xd9\x68\xe1\x3a\x2f\xb6\x5c\xf6\x67\x60\x2b\xf5\x7c\x4b\xdc\x00\x25\xf3\xca\xb6\xcc\x37\xdb\x5a\xb7\xa6\xfe\x95\xb8\x77\x2d\x08\x3a\x17\x0b\x79\xfa\xdd\xd9\xc3\xff\xa7\x06\xdf\xa0\xdb\x1c\x58\xc8\x33\x1e\x1e\x00\x00\xa7\xea\x4f\xe9\x24\x12\x6b\xee\xe2\x31\x56\x84\x3e\xd9\xf2\xbf\x4c\xfe\x50\x90\x6c\xfc\xd9\xc4\xb0\xb9\x0a\x98\x34\x1b\x95\xa4\x56\xe4\xd1\xcc\x5a\x15\xce\xf3\xb3\x93\xb3\xd3\xa3\xb6\x8c\x5f\x17\x85\x8a\xfb\x00\x9e\x33\xaf\x19\x67\x8e\xfa\xe5\xf0\x7d\x80\x1a\xf7\xf7\xe0\xc2\xd4\x7a\xe6\xf6\x2c\xfc\xb9\x65\x54\xfa\xc3\x4f\x9d\x80\x4f\x75\xcb\xfe\xed\xd5\x4d\x87\x11\x68\xda\x62\xf7\xc6\xb5\x88\x01\x4c\x61\x3d\x81\x0d\xc4\x50\x89\x5d\xd8\xa3\xe8\xcd\xb0\x4d\x33\xa4\xad\x5d\x37\x35\x28\xa2\x06\x09\xac\x01\xb2\x15\xc2\x87\x66\x81\xce\xa0\x47\x62\xee\x4d\x4d\xc3\xb2\xac\x30\xd4\xad\x6a\x26\x57\xd6\xe0\x24\xfd\xf2\x26\x28\x90\xae\xcb\xa2\x70\x1a\x6f\xd0\x3a\x0c\x1e\xf4\x1b\x7d\xe9\x9f\x07\x93\xd3\xc9\xd1\xff\x02\x00\x00\xff\xff\x68\x3f\x72\x25\xd8\x1a\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 874,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\xcd\x6e\x9c\x40\x10\x84\xef\x3c\x45\x49\x7b\x0d\x1b\x61\x69\x2f\xdc\xa2\x38\x89\x2c\x25\xd6\x5e\x9c\x7b\x7b\x68\xf0\x28\xf3\x97\xe9\x66\xb3\x04\xf9\xdd\xa3\xd9\x5f\x36\xf6\x9c\x80\xaa\xaf\xa7\xa6\x80\xba\xae\xab\x15\x3c\x7b\x43\xe6\x85\x3b\x74\x9c\x5c\x9c\x3c\x07\xc5\x28\xdc\xe1\x79\xc2\x57\x37\xee\xa1\x11\x07\x47\xb5\x82\x89\x41\xc9\x06\xce\xb0\x9e\x06\x86\x67\xa5\x8e\x94\xd6\x15\x25\xfb\x93\xb3\xd8\x18\x5a\x50\x4a\xf2\x71\xd7\x54\xbf\x6c\xe8\x5a\xdc\x5f\xc6\x56\x67\x7b\x5b\x01\x81\x3c\xb7\xd7\xdd\xe7\x19\xb6\xc7\xfa\x91\x3c\x4b\x22\xc3\x78\x7d\x3d\x99\x0e\xb7\x2d\xe6\xf9\x56\x9d\x67\x70\xe8\x8a\x4d\x12\x9b\x32\x31\x73\x72\xd6\x90\xb4\x68\x2a\x40\xd8\xb1\xd1\x98\x8b\x02\x78\x52\xf3\xf2\x9d\x9e\xd9\xc9\xf1\xc1\x9b\x00\x15\xa0\xec\x93\x23\xe5\x13\xb2\x08\x5b\x96\xbb\xa1\xdf\xe3\x81\x73\x94\xb2\x2e\x5d\x5d\x98\xfa\x5d\xa6\xac\x43\x9b\x0b\xa1\x6d\xd6\x9b\x75\xb3\xb9\xd5\xb7\xa3\x73\xdb\xe8\xac\x99\x5a\x3c\xf4\x8f\x51\xb7\x99\xa5\xd4\x7a\x76\x51\x1e\x16\xf9\x6a\xd4\x1e\x9b\xe6\x0e\xc0\x0a\x3f\x68\x6f\xfd\xe8\xcb\x0e\x31\x4f\xe5\x95\x8e\xc2\x1f\x60\x03\x3c\x0f\xf4\x3c\x29\xcb\x12\x7c\xc0\xc6\xe3\x06\x14\xfb\x97\xd1\xc7\x8c\x18\x18\x56\xd9\x2f\xed\x09\x4d\x73\xd7\x34\x58\xe1\x9e\x7b\x1a\x9d\x22\xc5\x7c\xcd\xb5\x2a\x9e\xdd\xee\x78\xf9\x14\x4c\xf4\x87\x8f\x4c\x23\x06\x56\xb8\x38\x08\x62\x0f\x26\xf3\x82\xcc\xbf\x47\x16\x05\x85\x0e\x99\x25\xc5\x20\xbc\xbe\x0c\x2a\x53\x6f\x4e\x78\xec\xd3\x38\xcb\x41\xaf\x07\x58\x74\xbf\x8d\x59\xdb\x63\xba\x8b\x2c\x6c\xc6\x6c\x75\xfa\x1c\x83\xf2\x5e\xdb\x05\x97\xc7\xf0\x49\x9e\x84\xf3\xff\xcc\x49\xfa\x96\xe3\x98\xde\x6a\xe4\x5c\xfc\xb3\xcd\x76\x67\x1d\x0f\xfc\x45\x0c\x39\xd2\xc3\xaf\xd0\x93\x13\xae\xfe\x05\x00\x00\xff\xff\x5d\x9a\x63\xab\x6a\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
