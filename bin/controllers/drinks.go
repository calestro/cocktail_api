package controllers

import (
	"cocktail/bin/models"
	"net/http"

	"github.com/go-chi/render"
)

func Get_Drinks(w http.ResponseWriter, r *http.Request) {
	var search []models.DrinkModel

	db().Find(&search)

	res := map[string]*[]models.DrinkModel{
		"data": &search,
	}

	render.JSON(w, r, res)
}
