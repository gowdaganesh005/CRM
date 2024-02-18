package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	NAME  string `json:"name"`
	EMAIL string `json:"email"`
	PHONE int64  `json:"phone"`
}
