package model

import (
	"gorm.io/gorm"
)

type LibraryRole uint

const (
	Owner LibraryRole = 0
	Read  LibraryRole = 1
	Write LibraryRole = 2
)

type UserLibraryRole struct {
	gorm.Model
	UserId    uint
	LibraryId uint
	// User      User    `gorm:"foreignKey:ID;references:UserId"`
	Library Library `gorm:"foreignKey:ID;references:LibraryId"`
	Role    LibraryRole
}
