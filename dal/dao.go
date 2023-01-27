package dao

import "gorm.io/gorm"

type User struct {
	userId   string `json:"userId"`
	email    string `json:"email"`
	password string `json:"password"`
	userName string `json:"userName"`
	gorm.Model
}
