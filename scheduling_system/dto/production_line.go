package dto

import "gorm.io/gorm"

type ProductionLine struct {
	gorm.Model
	ID     string
	Name   string
	Status ProductionLineStatus
}

type ProductionLineStatus string

const (
	Staring  ProductionLineStatus = "Staring"
	Running                       = "Running"
	Stopping                      = "Stopping"
	Stopped                       = "Stopped"
	Standby                       = "Standby"
)
