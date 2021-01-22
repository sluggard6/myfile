package model

type Folder struct {
	Model
	Name      string
	ParentId  uint
	LibraryId uint
}
