package database

import (
	"gorm.io/gorm"
)

type IDatabase interface {
	Connect()
	GetDB() *gorm.DB
}
