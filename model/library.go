package model

import "gorm.io/gorm"

type Library struct {
	gorm.Model
	Name   string
	UserId uint
	Owner  User `gorm:"foreignKey:UserId"`
}
