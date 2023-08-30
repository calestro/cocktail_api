package models

type DrinkModel struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `validate:"required"`
	Recipe string `json:"recipe", validate:"required"`
}
