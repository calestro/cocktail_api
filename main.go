package main

import (
	"cocktail/bin/routes"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.RouteHeaders().Handler)

	routes.GetEndpoints(r)

	fmt.Println("start")

	http.ListenAndServe(":8282", r)

}
