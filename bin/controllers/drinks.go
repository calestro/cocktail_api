package controllers

import (
	"cocktail/bin/models"
	"cocktail/bin/utils/errs"
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

func Post_Drinks(w http.ResponseWriter, r *http.Request) {
	var model models.DrinkModel

	err := render.Decode(r, &model)

	if err != nil {
		w.WriteHeader(500)

	} else if model.Name == "" || model.Recipe == "" {
		w.WriteHeader(400)
		errs.PostError("Campos Inválidos", r, w)
	} else {
		db().Create(&model)
		w.WriteHeader(201)
	}

}
