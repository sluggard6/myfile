package store

import (
	"encoding/base64"
	"encoding/hex"
	"io"
	"os"

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
	fileName := base64.StdEncoding.EncodeToString(sha)

	return &File{"", fileName, hex.EncodeToString(sha), size}
}

func NewTmpFile() (*os.File, error) {
	name := base64.StdEncoding.EncodeToString([]byte(util.UUID()))
	return os.Create(name)
}
