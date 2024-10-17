package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
)


func (app *application) ListCoffees(w http.ResponseWriter, r *http.Request)  {
	
	coffees, err := app.DB.AllCoffees()

	out, err := json.Marshal(coffees)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}


func (app *application) CoffeeInfo(w http.ResponseWriter, r *http.Request)  {
	
	coffee, err := app.DB.Coffee(chi.URLParam(r, "id"))

	out, err := json.Marshal(coffee)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}



func (app *application) Home(w http.ResponseWriter, r *http.Request)  {
	var payload = struct {
		Status string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status: "active",
		Message: "Go to coffees roasted",
		Version: "1.0.0",

	}

	out, err := json.Marshal(payload)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
