package controllers

import (
	"net/http"

	"github.com/go-chi/render"
)

func Get_Ping(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{"message": "Alive"}

	render.JSON(w, r, m)
}
