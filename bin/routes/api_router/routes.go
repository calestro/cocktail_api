package api_router

import (
	"cocktail/bin/controllers"

	"github.com/go-chi/chi/v5"
)

func GetEndpoints(r *chi.Mux) {

	///////*****/	Ping	/**********///////

	r.Get("/ping", controllers.Get_Ping)

	r.Get("/position", controllers.Get_Position)
	r.Post("/position", controllers.POST_Position)

	///////*****/	Drinks	/**********///////
	r.Get("/drinks", controllers.Get_Drinks)

	///////*****/	Maker	/**********///////
	r.Post("/maker", controllers.MakerDrink)

}
