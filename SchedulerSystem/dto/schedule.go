package dto

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	ID               string
	Name             string
	Description      string
	ProductionLineId string
}
