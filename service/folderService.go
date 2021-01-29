package service

import "github.com/sluggard/myfile/model"

//FolderService 文件夹服务
type FolderService interface {
	GetChildrenFolder(folderID uint) (*[]model.Folder, error)
	GetChildrenFile(folderID uint) (*[]model.File, error)
	CreateChild(parent *model.Folder, name string) (*model.Folder, error)
}

type folderSer struct {
}

var folderService = &folderSer{}

//NewFolderService 创建FolderService实例
func NewFolderService() FolderService {
	return folderService
}

func (s *folderSer) GetChildrenFolder(folderID uint) (*[]model.Folder, error) {
	folder := &model.Folder{Model: model.Model{ID: folderID}}
	return folder.GetChildren()
}

func (s *folderSer) GetChildrenFile(folderID uint) (*[]model.File, error) {
	file := &model.File{FolderID: folderID}
	return file.GetFilesByFolderID()
}

func (s *folderSer) CreateChild(parent *model.Folder, name string) (*model.Folder, error) {
	return parent.CreateChild(name)
}
