package models

import "gorm.io/gorm"

type DrinkModel struct {
	gorm.Model
	ID     int
	Name   string `validate:"required"`
	Recipe string `validate:"required" `
}
