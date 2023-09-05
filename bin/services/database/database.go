package database

import (
	"cocktail/bin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connection() {
	dsn := "host=localhost user=postgres password=25265459 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connect REFUSED")
	}
	db.AutoMigrate(&models.PositionDrinks{}, &models.DrinkModel{})

	Db = db
}
