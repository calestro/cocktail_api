package controllers

import (
	"cocktail/bin/models"
	"net/http"

	"github.com/go-chi/render"
)

func MakerDrink(w http.ResponseWriter, r *http.Request) {
	var model models.DrinkMaker
	render.DecodeJSON(r.Body, &model)
	SerialSender(model.Recipe)
}
