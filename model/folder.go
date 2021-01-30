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
	err = db.Where("id=?", f.ParentID).Find(parent).Error
	return
}

func (f *Folder) GetPath() (path []Folder, err error) {
	var appendFolder = func(folder *Folder) {
		path = append(path, *folder)
		logrus.Debug("path:%d:{id:%s,parent:%s}\n", len(path), folder.ID, folder.ParentID)
		return
	}
	tp := f
	for !tp.IsRoot() {
		defer appendFolder(tp)
		tp, _ = tp.GetParent()
	}
	return
}

func (f *Folder) CreateChild(name string) (child *Folder, err error) {
	child = &Folder{Name: name, ParentID: f.ID, LibraryID: f.LibraryID}
	err = db.Create(child).Error
	return
}
