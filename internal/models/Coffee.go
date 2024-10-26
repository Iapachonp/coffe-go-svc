package models

import "time"

type Coffee struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	VarietalId  int	      `json:"varietalId"`	
	FarmerId    int       `json:"farmerId"`
	OriginId    int       `json:"originId"`
	ProcessId   int       `json:"processId"`
	Description string    `json:"description"`
	PriceId     int       `json:"price"`
	Sca         float64   `json:"sca"`
	Acidity     string    `json:"acidity"`
	Body        string    `json:"body`
	Balance     string    `json:"balance`
	Clarity     string    `json:"clarity`
	Sweetness   string    `json:"sweetness` 
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type CoffeeData struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Proccess    string    `json:"proccess"`
	Origin      string    `json:"origin"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Sca         float64   `json:"sca"`
	Weight      int	      `json:"weight"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
