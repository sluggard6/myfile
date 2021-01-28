package model

//Policy 存储策略
type Policy struct {
	Model
	Type string
	size int64
	Path string
	Sha  string
}

//File 文件
type File struct {
	Model
	Name     string `json:"name"`
	Ext      string `json:"ext"`
	FolderID uint   `json:"-"`
	Size     uint   `json:"size"`
	PolicyID uint   `json:"-"`
	Policy   Policy `gorm:"foreignKey:PolicyID" json:"-"`
}

func (file *File) GetFilesByFolderId() (files *[]File, err error) {
	files = &[]File{}
	err = db.Where("folder_id=?", file.FolderID).Find(files).Error
	return
}
