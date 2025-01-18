package repository

import (
	"coffee-svc/internal/models"
	"database/sql"
)

type DatabaseRepo interface{
	Connection() *sql.DB
	AllCoffees() ([]*models.CoffeeData, error) 
	Coffee(id string) ([]*models.CoffeeData, error)
	PostCoffee( Coffee *models.Coffee) (string, error) 
	AllVarietals() ([]*models.Varietal, error) 
	Varietal(id string) ([]*models.Varietal, error)
	PostVarietal( Varietal *models.Varietal) (string, error)
	AllFarmers() ([]*models.Farmer, error) 
	Farmer(id string) ([]*models.Farmer, error)
	PostFarmer( Farmer *models.Farmer) (string, error)
}


