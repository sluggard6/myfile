package model

import "gorm.io/gorm"

type Library struct {
	gorm.Model
	Name   string
	UserId uint
	Owner  User `gorm:"foreignKey:UserId"`
}

func (l *Library) GetLibraryMine(userId uint) (librarys []Library, err error) {
	err = db.Where("user_id=?", userId).Find(librarys).Error
	return
}
