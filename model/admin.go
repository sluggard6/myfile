package model

import "time"

type Admin struct {
	Id        uint
	Name      string
	Email     *string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
