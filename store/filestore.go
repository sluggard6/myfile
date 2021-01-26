package store

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/sluggard/myfile/util"
)

const (
	defaultRoot = "file-data"
	defaultTmp  = ".tmp"
)

type FileStore struct {
	Root string
	Tmp  string
}

type File struct {
	Path string
	Name string
	Sha  string
	Size int
}

func New(root string) (*FileStore, error) {
	if "" == root {
		root = defaultRoot
	}
	var err error
	if root, err = filepath.Abs(root); err != nil {
		return nil, err
	}
	tmp := root + string(filepath.Separator) + defaultTmp
	return &FileStore{root, tmp}, nil
}

func (fs *FileStore) saveFile(reader io.Reader, name string) (*File, error) {
	tmpFile, err := fs.NewTmpFile()
	if err != nil {
		return nil, err
	}
	sha, err := util.SaveAndSha(reader, tmpFile)
	if err != nil {
		return nil, err
	}
	hexString := hex.EncodeToString(sha)
	var fileName = fs.Root + string(filepath.Separator) + strings.Join(makeFilePath(hexString), string(filepath.Separator))
	logrus.Debugf("store file : %s", fileName)
	dir, name := filepath.Split(fileName)
	if err := os.MkdirAll(dir, 0744); err != nil {
		return nil, err
	}

	os.Rename(tmpFile.Name(), fileName)
	return &File{"", fileName, hex.EncodeToString(sha), size}, nil
}

func (fs *FileStore) NewTmpFile() (*os.File, error) {
	name := util.UUID()
	filePath := fs.Tmp + string(filepath.Separator) + name
	return os.Create(filePath)
}

func makeFilePath(sha string) (path []string) {
	var i, folderLength, folderLevel = 0, 10, 4
	path = make([]string, folderLevel+1)
	for ; i < folderLevel; i++ {
		path[i] = sha[i*folderLength : (i+1)*folderLength]
	}
	path[i] = sha[i*folderLength:]
	fmt.Println(path)

	// path = make([]string, 4)
	// path[0] = sha[:8]
	// path[1] = sha[8:16]
	// path[2] = sha[16:24]
	// path[3] = sha[24:]
	// fmt.Println(path)
	// fmt.Print(sha[0:5])
	return
}