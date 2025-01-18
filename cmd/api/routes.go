package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler  {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)
	mux.Get("/", app.Home)
	mux.Get("/coffees", app.ListCoffees)
	mux.Post("/coffees", app.CreateCoffee)
	mux.Get("/coffees/{id}", app.CoffeeInfo)
	mux.Get("/varietals", app.ListVarietals)
	mux.Post("/varietals", app.CreateVarietal)
	mux.Get("/varietals/{id}", app.VarietalInfo)
	mux.Get("/farmers", app.ListFarmers)
	mux.Post("/farmers", app.CreateFarmer)
	mux.Get("/farmers/{id}", app.FarmerInfo)
	return mux
}
