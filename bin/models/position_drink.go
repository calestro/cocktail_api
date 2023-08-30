package models

type PositionDrinks struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `validate:"required", json:"name"`
	Position int    `validate:"required" json:"position"`
	Type     string `json:"type", validate:"required"`
}
