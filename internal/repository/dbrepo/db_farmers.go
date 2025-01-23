package dbrepo 

import (
	"coffee-svc/internal/models"
	"context"
	"log"
	"strconv"
)

func (m *PostgresDBRepo) AllFarmers() ([]*models.Farmer, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select *
		from farmers
		order by name
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("error no able to query database for farmers")
		return nil, err
	}

	defer rows.Close()

	var farmers []*models.Farmer

	for rows.Next() {
		var farmer models.Farmer
		err := rows.Scan(
			&farmer.FarmerId,
			&farmer.Name,
			&farmer.Description,
			&farmer.Altitude,
		)
		if err != nil {
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		farmers = append(farmers, &farmer)
	}

	return farmers, nil
}

func (m *PostgresDBRepo) Farmer(id string) ([]*models.Farmer, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select *  
		from farmers 
		where farmerid = $1
		order by name
		limit 1 
	`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		log.Println("error no able to query databaases for farmer")
		return nil, err
	}

	defer rows.Close()

	var farmers []*models.Farmer

	for rows.Next() {
		var farmer models.Farmer
		err := rows.Scan(
			&farmer.FarmerId,
			&farmer.Name,
			&farmer.Description,
			&farmer.Altitude,
		)
		if err != nil {
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		farmers = append(farmers, &farmer)
	}
	return farmers, nil	
}

func (m *PostgresDBRepo) PostFarmer(farmer *models.Farmer) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	id := 0
	query := `
	INSERT INTO public.farmers (Name, Description, Altitude ) VALUES
	($1, $2, $3)
	RETURNING farmerid`
	err := m.DB.QueryRowContext(ctx, query,
		farmer.Name,
		farmer.Description,
		farmer.Altitude).Scan(&id)

	if err != nil {
		log.Println("error no able to query databaases for farmers")
		log.Panic(err)
		return "", err
	}

	farmerid := "New farmer created with id:" + strconv.FormatInt(int64(id), 10)
	return farmerid, nil
}
