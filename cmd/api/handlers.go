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
	fmt.Println(err)
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

func (app *application) Deletecoffee(w http.ResponseWriter, r *http.Request)  {
	
	coffee, err := app.DB.DeleteCoffee(chi.URLParam(r, "id"))

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

// ------------ farmers --------------- 

func (app *application) CreateFarmer(w http.ResponseWriter, r *http.Request)  {
	
	var farmer *models.Farmer	

	err := json.NewDecoder(r.Body).Decode(&farmer)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	// Do something with the Farmer struct...
	fmt.Fprintf(w, "Farmer: %+v", farmer)

	created, err := app.DB.PostFarmer(farmer)

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


func (app *application) ListFarmers(w http.ResponseWriter, r *http.Request)  {
	
	farmers, err := app.DB.AllFarmers()

	out, err := json.Marshal(farmers)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}


func (app *application) FarmerInfo(w http.ResponseWriter, r *http.Request)  {
	
	farmer, err := app.DB.Farmer(chi.URLParam(r, "id"))

	out, err := json.Marshal(farmer)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}


// ------------ origins --------------- 

func (app *application) CreateOrigin(w http.ResponseWriter, r *http.Request)  {
	
	var origin *models.Origin	

	err := json.NewDecoder(r.Body).Decode(&origin)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	// Do something with the Origin struct...
	fmt.Fprintf(w, "Origin: %+v", origin)

	created, err := app.DB.PostOrigin(origin)

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


func (app *application) ListOrigins(w http.ResponseWriter, r *http.Request)  {
	
	farmers, err := app.DB.AllOrigins()

	out, err := json.Marshal(farmers)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}


func (app *application) OriginInfo(w http.ResponseWriter, r *http.Request)  {
	
	farmer, err := app.DB.Origin(chi.URLParam(r, "id"))

	out, err := json.Marshal(farmer)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

// ------------ processes --------------- 

func (app *application) CreateProcess(w http.ResponseWriter, r *http.Request)  {
	
	var process *models.Process	

	err := json.NewDecoder(r.Body).Decode(&process)

	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	// Do something with the Process struct...
	fmt.Fprintf(w, "Process: %+v", process)

	created, err := app.DB.PostProcess(process)

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


func (app *application) ListProcesses(w http.ResponseWriter, r *http.Request)  {
	
	farmers, err := app.DB.AllProcesses()

	out, err := json.Marshal(farmers)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}


func (app *application) ProcessInfo(w http.ResponseWriter, r *http.Request)  {
	
	farmer, err := app.DB.Process(chi.URLParam(r, "id"))

	out, err := json.Marshal(farmer)
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


// ------------ Prices --------------- 

// func (app *application) CreatePrice(w http.ResponseWriter, r *http.Request)  {
// 	
// 	var price *models.Price	
//
// 	err := json.NewDecoder(r.Body).Decode(&price)
//
// 	if err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         fmt.Println(err)
// 	return
// 	}
//
// 	// Do something with the Price struct...
// 	fmt.Fprintf(w, "Price: %+v", price)
//
// 	created, err := app.DB.PostPrice(price)
//
// 	if err != nil{
// 		fmt.Println(err)
// 	}
//
// 	success := ""
// 	
// 	if created != "" {
// 		success = "operation was successful " + created 	
// 	} else {
// 		success = "operation was not successful"
// 	}
//
// 	w.Header().Set("Content-Type","app-application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(success))
//
// }


func (app *application) ListPrices(w http.ResponseWriter, r *http.Request)  {
	
	prices, err := app.DB.AllPrices()

	out, err := json.Marshal(prices)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","app-application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}


// func (app *application) PriceInfo(w http.ResponseWriter, r *http.Request)  {
// 	
// 	farmer, err := app.DB.Price(chi.URLParam(r, "id"))
//
// 	out, err := json.Marshal(farmer)
// 	if err != nil{
// 		fmt.Println(err)
// 	}
//
// 	w.Header().Set("Content-Type","app-application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(out)
//
// }
