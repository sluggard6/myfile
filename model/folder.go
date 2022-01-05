package model

import (
	"github.com/sirupsen/logrus"
	"github.com/sluggard/myfile/common"
)

//Folder 文件夹
type Folder struct {
	Model
	Name      string `json:"name"`
	ParentID  uint
	LibraryID uint
}

//PolicyType 策略类型
type PathType string

const (
	//MyFile 自带的文件存储方案
	FolderType  PathType = "folder"
	LibraryType PathType = "library"
)

type PathItem struct {
	Name string   `json:"name"`
	ID   uint     `json:"id"`
	Type PathType `json:"type"`
}

func (f *Folder) IsRoot() bool {
	return f.ParentID == 0
}

func (f *Folder) GetChildren() (subFolders *[]Folder, err error) {
	subFolders = &[]Folder{}
	err = db.Where("parent_id=?", f.ID).Find(subFolders).Error
	return
}

//GetChildrenFiles 查询目录下的所有文件
// func (f *Folder) GetChildrenFiles() (files *[]File, err error) {
// 	err = db.Where("folder_id=?", f.ID).Find(files).Error
// 	return
// }

func (f *Folder) GetParent() (parent *Folder, err error) {
	if f.IsRoot() {
		err = common.CommonError{Message: "root folder has not Parent"}
		return
	}
	parent = &Folder{}
	GetById(parent, f.ParentID)
	return
}

func (f *Folder) GetPath() (path []PathItem, err error) {
	var appendFolder = func(folder *Folder) {
		path = append(path, PathItem{ID: folder.ID, Name: folder.Name, Type: FolderType})
	}
	tp := f
	for !tp.IsRoot() {
		defer appendFolder(tp)
		tp, err = tp.GetParent()
		if err != nil {
			logrus.Debug(err)
		}
	}
	path = append(path, PathItem{ID: tp.ID, Name: tp.Name, Type: FolderType})
	return
}

func (f *Folder) CreateChild(name string) (child *Folder, err error) {
	child = &Folder{Name: name, ParentID: f.ID, LibraryID: f.LibraryID}
	err = db.Create(child).Error
	return
}
