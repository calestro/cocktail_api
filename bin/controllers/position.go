package controllers

import (
	"cocktail/bin/models"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func Get_Position(w http.ResponseWriter, r *http.Request) {

	var search []models.PositionDrinks

	db().Table("position_drinks").Select([]string{"name", "position"}).Find(&search)

	res := map[string]*[]models.PositionDrinks{
		"data": &search,
	}

	render.JSON(w, r, res)
}

func POST_Position(w http.ResponseWriter, r *http.Request) {
	var model *models.PositionDrinks
	render.DecodeJSON(r.Body, &model)
	fmt.Println(model)
	db().Create(&model)
	w.WriteHeader(201)

}
