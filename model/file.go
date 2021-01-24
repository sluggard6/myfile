package model

type Policy struct {
	Model
	Type string
	Name string
	Path string
	Sha  string
}

type File struct {
	Model
	Name     string
	FolderId uint
	policy   Policy
}

func (file *File) GetFilesByFolderId() (files *[]File, err error) {
	err = db.Where("folder_id=?", file.FolderId).Find(files).Error
	return
}
