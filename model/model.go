package model

import (
	"gorm.io/gorm"
)
type Vehicle struct {
	gorm.Model
	VehicleId 			string
	VehicleName 		string
	VehicleNumber 		string
	VehicleVinNumber 	string
	VehicleSerialNumber string
}