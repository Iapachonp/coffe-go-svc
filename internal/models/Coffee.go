package models

import "time"

type Coffee struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	VarietalId  int	      `json:"varietalId,string"`	
	FarmerId    int       `json:"farmerId,string"`
	OriginId    int       `json:"originId,string"`
	ProcessId   int       `json:"processId,string"`
	Description string    `json:"description"`
	PriceId     int       `json:"priceId,string"`
	Sca         float64   `json:"sca,string"`
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
	Process     string    `json:"process"`
	Origin      string    `json:"origin"`
	Description string    `json:"description"`
	Price       float64   `json:"priceid"`
	Sca         float64   `json:"sca"`
	Weight      int	      `json:"weight"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
