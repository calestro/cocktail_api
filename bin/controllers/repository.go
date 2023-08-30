package controllers

import (
	"cocktail/bin/services/database"

	"gorm.io/gorm"
)

func db() *gorm.DB {
	return database.Db
}
