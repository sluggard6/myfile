package dao

import (
	"gorm.io/gorm"
)

type UserDao interface {
}

type userDao struct {
	db gorm.DB
}

func NewUserDao() UserDao {
	return &userDao{sqlDb}
}
