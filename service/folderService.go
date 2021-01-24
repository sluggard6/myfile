package service

import "github.com/sluggard/myfile/model"

type FolderService interface {
	GetChildrenFolder(folderId uint) (*[]model.Folder, error)
	GetChildrenFile(folderId uint) (*[]model.File, error)
}

type folderSer struct {
}

var folderService = &folderSer{}

func NewFolderService() FolderService {
	return folderService
}

func (s *folderSer) GetChildrenFolder(folderId uint) (*[]model.Folder, error) {
	folder := &model.Folder{Model: model.Model{ID: folderId}}
	return folder.GetChildren()
}

func (s *folderSer) GetChildrenFile(folderId uint) (*[]model.File, error) {
	file := &model.File{FolderID: folderId}
	return file.GetFilesByFolderId()
}
