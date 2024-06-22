package main

import (
	"coffee-svc/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)


func (app *application) ListCoffees(w http.ResponseWriter, r *http.Request)  {
	

	colombia := models.Coffee{
		ID: 0,
		Name: "Variedad Colombia",
		Proccess: "Lavado",
		Sca: 8.5,
		Description: "set the des here",
		Weight: "500 gr",
		Price: 38000,
		Origin: "Huila",

	}

	tabi := models.Coffee{
		ID: 1,
		Name: "Variedad Tabi",
		Proccess: "Natural",
		Sca: 8.7,
		Description: "set the des here",
		Weight: "500 gr",
		Price: 48000,
		Origin: "Huila",

	}

	bourbonRojo := models.Coffee{
		ID: 2,
 		Name: "Variedad Bourbon Rojo",
		Proccess: "Lavado",
		Sca: 8.7,
		Description: "set the des here",
		Weight: "500 gr",
		Price: 50000,
		Origin: "Huila",

	}

	coffees := []models.Coffee {
		colombia, tabi, bourbonRojo,
	}

	out, err := json.Marshal(coffees)
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
		Message: "Go to coffees rosted",
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
