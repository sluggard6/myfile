package service

import (
	"io"
	"unsafe"

	"github.com/sluggard/myfile/model"
	"github.com/sluggard/myfile/store"
)

type FileService interface {
	SaveFile(file io.Reader, name string, folder *model.Folder) (*model.File, error)
}

type myFileSer struct {
	store store.Store
}

var myFileService *myFileSer

//NewFileService 根据存储类型创建文件服务
func NewFileService(store store.Store) *FileService {
	if myFileService == nil {
		myFileService = &myFileSer{store}
	}
	// return myFileService
	return (*FileService)(unsafe.Pointer(myFileService))
}

func (s *myFileSer) SaveFile(file io.Reader, name string, folder *model.Folder) (*model.File, error) {

	return nil, nil
}
