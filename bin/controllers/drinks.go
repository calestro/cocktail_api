package controllers

import (
	"cocktail/bin/models"
	"net/http"

	"github.com/go-chi/render"
)

func Get_Drinks(w http.ResponseWriter, r *http.Request) {
	var search []models.DrinkModel
	db().Find(&search)

	render.JSON(w, r, search)
}
