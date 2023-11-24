package dto

import "gorm.io/gorm"

type Robot struct {
	gorm.Model
	ID    string
	State string
}