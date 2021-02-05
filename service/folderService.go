package service

import (
	"fmt"

	"github.com/sluggard/myfile/model"
)

//FolderService 文件夹服务
type FolderService interface {
	GetChildrenFolder(folderID uint) (*[]model.Folder, error)
	GetChildrenFile(folderID uint) (*[]model.File, error)
	CreateChild(parent *model.Folder, name string) (*model.Folder, error)
	CheckFileName(folder *model.Folder, name string) error
	DeleteByLibrary(libraryID uint) error
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
	// folder := &model.Folder{Model: model.Model{ID: folderID}}
	// return folder.GetChildrenFiles()
	file := &model.File{FolderID: folderID}
	return file.GetFilesByFolderID()
}

func (s *folderSer) CreateChild(parent *model.Folder, name string) (*model.Folder, error) {
	return parent.CreateChild(name)
}

func (s *folderSer) CheckFileName(folder *model.Folder, name string) error {
	files, err := s.GetChildrenFile(folder.ID)
	for _, v := range *files {
		if v.Name == name {
			return fmt.Errorf("file name '%s' is exist", name)
		}
	}
	return err
}

func (s *folderSer) DeleteByLibrary(libraryID uint) error {
	var folders []model.Folder
	model.DB().Where("library_id=?", libraryID).Find(folders)
	folderIDs := make([]uint, len(folders))
	for i, folder := range folders {
		folderIDs[i] = folder.ID
	}
	model.DB().Where("folder_id in ?", folderIDs).Delete(&model.File{})
	model.DB().Where("library_id = ?", libraryID).Delete(&model.Folder{})
	return nil
}
