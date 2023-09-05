package controllers

import (
	"cocktail/bin/services/database"
	"cocktail/bin/services/serials"

	"gorm.io/gorm"
)

func db() *gorm.DB {
	return database.Db
}

func SerialSender(recipe string) {
	serials.SerialConnect.Write([]byte(recipe))
}
