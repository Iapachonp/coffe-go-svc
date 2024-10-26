package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"coffee-svc/internal/models"
)


func (app *application) CreateCoffee(w http.ResponseWriter, r *http.Request)  {
	
	var coffee *models.Coffee	

	err := json.NewDecoder(r.Body).Decode(&coffee)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	// Do something with the Coffee struct...
	fmt.Fprintf(w, "Coffee: %+v", coffee)

	created, err := app.DB.PostCoffee(coffee)

	if err != nil{
		fmt.Println(err)
	}

	success := ""
	
	if created != "" {
		success = "operation was successful " + created 	
	} else {
		success = "operation was not successful"
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(success))

}


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



func (app *application) CreateVarietal(w http.ResponseWriter, r *http.Request)  {
	
	var varietal *models.Varietal	

	err := json.NewDecoder(r.Body).Decode(&varietal)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	// Do something with the Varietal struct...
	fmt.Fprintf(w, "Varietal: %+v", varietal)

	created, err := app.DB.PostVarietal(varietal)

	if err != nil{
		fmt.Println(err)
	}

	success := ""
	
	if created != "" {
		success = "operation was successful " + created 	
	} else {
		success = "operation was not successful"
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(success))

}


func (app *application) ListVarietals(w http.ResponseWriter, r *http.Request)  {
	
	varietals, err := app.DB.AllVarietals()

	out, err := json.Marshal(varietals)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}


func (app *application) VarietalInfo(w http.ResponseWriter, r *http.Request)  {
	
	varietal, err := app.DB.Varietal(chi.URLParam(r, "id"))

	out, err := json.Marshal(varietal)
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
