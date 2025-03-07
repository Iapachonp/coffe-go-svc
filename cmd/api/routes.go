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
	mux.Delete("/coffees/{id}", app.Deletecoffee)
	mux.Get("/varietals", app.ListVarietals)
	mux.Post("/varietals", app.CreateVarietal)
	mux.Get("/varietals/{id}", app.VarietalInfo)
	mux.Get("/farmers", app.ListFarmers)
	mux.Post("/farmers", app.CreateFarmer)
	mux.Get("/farmers/{id}", app.FarmerInfo)
	mux.Get("/origins", app.ListOrigins)
	mux.Post("/origins", app.CreateOrigin)
	mux.Get("/origins/{id}", app.OriginInfo)
	mux.Get("/processes", app.ListProcesses)
	mux.Post("/processes", app.CreateProcess)
	mux.Get("/processes/{id}", app.ProcessInfo)
	mux.Get("/prices", app.ListPrices)
	// mux.Post("/prices", app.CreatePrice)
	// mux.Get("/prices/{id}", app.PriceInfo)
	return mux
}
