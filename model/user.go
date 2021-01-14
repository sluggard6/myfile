package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique_index:idx_only_one"`
	Password string
}

func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	result := DB.Where("username = ?", username).First(user)
	return user, result.Error
}
