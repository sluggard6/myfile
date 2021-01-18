package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique_index:idx_only_one;commit:'用户名'" validate:"required"`
	Password string `gorm:"not null;commit:'用户密码'" validate:"required"`
	Salt     string `grom:"not null;commit:'用户掩码'" json:"-"`
}

func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	result := DB.Where("username = ?", username).First(user)
	return user, result.Error
}

func GetUserById(id int) (*User, error) {
	user := &User{}
	result := DB.Where("id = ?", id).First(user)
	return user, result.Error
}
