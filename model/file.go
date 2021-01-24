package model

//Policy 存储策略
type Policy struct {
	Model
	Type string
	Name string
	Path string
	Sha  string
}

//File 文件
type File struct {
	Model
	Name     string
	FolderID uint
	Size     uint
	policy   Policy
}

func (file *File) GetFilesByFolderId() (files *[]File, err error) {
	files = &[]File{}
	err = db.Where("folder_id=?", file.FolderID).Find(files).Error
	return
}
