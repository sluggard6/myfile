package store

import (
	"encoding/base64"
	"io"

	"github.com/sluggard/myfile/util"
)

type FileStore struct {
	root string
	tmp  string
}

type File struct {
	Path string
	Name string
	Sha  string
	Size int
}

func saveFile(reader io.Reader, name string) *File {
	size, sha, _ := util.ShaReader(reader)
	fileName := base64.StdEncoding.EncodeToString([]byte(sha))

	return &File{"", fileName, sha, size}
}
