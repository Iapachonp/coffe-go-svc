package repository

import (
	"coffee-svc/internal/models"
	"database/sql"
)

type DatabaseRepo interface{
	Connection() *sql.DB
	AllCoffees() ([]*models.Coffee, error) 
}


