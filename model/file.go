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
