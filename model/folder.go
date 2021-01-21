package model

import (
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Name      string
	ParentId  uint
	LibraryId uint
}
