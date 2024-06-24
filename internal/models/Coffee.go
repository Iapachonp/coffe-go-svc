package models

import "time"

type Coffee struct {
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
