package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler  {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)
	mux.Get("/", app.Home)
	mux.Get("/coffees", app.ListCoffees)
	mux.Get("/coffees/{id}", app.CoffeeInfo)
	return mux
}
