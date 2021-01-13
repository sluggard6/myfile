package model

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Name     string
	Email    *string
	Password string
}
