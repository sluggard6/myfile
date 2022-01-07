package model

type ShareFile struct {
	Model
	Path   string `json:"path"`
	Code   string `json:"code"`
	FileId uint   `json:"fileId"`
}
