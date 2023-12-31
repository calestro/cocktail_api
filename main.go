package main

import (
	"cocktail/bin/routes/api_router"
	"cocktail/bin/services/alexa_maker"
	"cocktail/bin/services/database"
	"cocktail/bin/services/serials"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)
	api_router.GetEndpoints(r)
	fmt.Println("start")

	database.Connection()
	serials.Connection()

	http.ListenAndServe(":8282", r)

	alexa_maker.DevicesCreate()

}
