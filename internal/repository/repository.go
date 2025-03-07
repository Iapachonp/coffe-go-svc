package repository

import (
	"coffee-svc/internal/models"
	"database/sql"
)

type DatabaseRepo interface{
	Connection() *sql.DB
	AllCoffees() ([]*models.CoffeeData, error) 
	Coffee(id string) ([]*models.CoffeeData, error)
	DeleteCoffee(id string) (bool, error)
	PostCoffee( Coffee *models.Coffee) (string, error) 
	AllVarietals() ([]*models.Varietal, error) 
	Varietal(id string) ([]*models.Varietal, error)
	PostVarietal( Varietal *models.Varietal) (string, error)
	AllFarmers() ([]*models.Farmer, error) 
	Farmer(id string) ([]*models.Farmer, error)
	PostFarmer( Farmer *models.Farmer) (string, error)
	AllOrigins() ([]*models.Origin, error) 
	Origin(id string) ([]*models.Origin, error)
	PostOrigin( Origin *models.Origin) (string, error)
	AllProcesses() ([]*models.Process, error) 
	Process(id string) ([]*models.Process, error)
	PostProcess( Process *models.Process) (string, error)
	AllPrices() ([]models.PriceGo, error) 
	Price(id string) ([]*models.Price, error)
	PostPrice( Price *models.Price) (string, error)
}


