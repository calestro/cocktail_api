package controllers

import (
	"cocktail/bin/models"
	"cocktail/bin/services/serials"
	"net/http"

	"github.com/go-chi/render"
)

func MakerDrink(w http.ResponseWriter, r *http.Request) {
	var model models.DrinkMaker
	render.DecodeJSON(r.Body, &model)
	println(model.Recipe)
	serials.SerialSender(model.Recipe)
}
