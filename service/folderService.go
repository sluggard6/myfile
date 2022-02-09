package service

import (
	"fmt"

	"github.com/sluggard/myfile/model"
)

//FolderService 文件夹服务
type FolderService interface {
	// GetTree(folderId uint) (*[]model.Folder, error)
	GetChildrenFolder(folderID uint) (*[]model.Folder, error)
	GetChildrenFile(folderID uint) (*[]model.File, error)
	CreateChild(parent *model.Folder, name string) (*model.Folder, error)
	CheckFileName(parentID uint, name string) error
	CheckFolderName(parentID uint, name string) error
	DeleteByLibrary(libraryID uint) error
	UpdateFolder(folder *model.Folder) error
}

type folderSer struct {
}

var folderService = &folderSer{}

//NewFolderService 创建FolderService实例
func NewFolderService() FolderService {
	return folderService
}

// func (s *folderSer) GetTree(folderID uint) (*[]model.Folder, error) {
// folder := &model.Folder{Model: model.Model{ID: folderID}}
// for !folder.IsRoot()
// return
// }

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

func (s *folderSer) CheckFileName(parentID uint, name string) error {
	files, err := s.GetChildrenFile(parentID)
	for _, v := range *files {
		if v.Name == name {
			return fmt.Errorf("file name '%s' is exist", name)
		}
	}
	return err
}

func (s *folderSer) CheckFolderName(parentID uint, name string) error {
	folders, err := s.GetChildrenFolder(parentID)
	for _, v := range *folders {
		if v.Name == name {
			return fmt.Errorf("file name '%s' is exist", name)
		}
	}
	return err
}

func (s *folderSer) DeleteByLibrary(libraryID uint) error {
	var folders *[]model.Folder = &[]model.Folder{}
	model.DB().Where("library_id=?", libraryID).Find(folders)
	folderIDs := make([]uint, len(*folders))
	for i, folder := range *folders {
		folderIDs[i] = folder.ID
	}
	model.DB().Where("folder_id in ?", folderIDs).Delete(&model.File{})
	model.DB().Where("library_id = ?", libraryID).Delete(&model.Folder{})
	return nil
}

func (s *folderSer) UpdateFolder(folder *model.Folder) error {
	if err := s.CheckFolderName(folder.ParentID, folder.Name); err != nil {
		return err
	}
	return model.DB().Save(folder).Error
}
