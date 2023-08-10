package main

import (
	"cocktail/bin/routes/api_router"
	"cocktail/bin/services/serial_comunication"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	serial_comunication.Teste()
	r := chi.NewRouter()
	r.Use(middleware.DefaultLogger)

	api_router.GetEndpoints(r)

	fmt.Println("start")

	http.ListenAndServe(":8282", r)

}
