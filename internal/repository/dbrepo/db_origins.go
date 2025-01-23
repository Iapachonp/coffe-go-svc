package dbrepo 

import (
	"coffee-svc/internal/models"
	"context"
	"log"
	"strconv"
)

func (m *PostgresDBRepo) AllOrigins() ([]*models.Origin, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select *
		from origins
		order by country 
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("error no able to query database for origins")
		return nil, err
	}

	defer rows.Close()

	var origins []*models.Origin

	for rows.Next() {
		var origin models.Origin
		err := rows.Scan(
			&origin.OriginId,
			&origin.Country,
			&origin.State,
			&origin.CityTown,
		)
		if err != nil {
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		origins = append(origins, &origin)
	}

	return origins, nil
}

func (m *PostgresDBRepo) Origin(id string) ([]*models.Origin, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select *  
		from origins 
		where originid = $1
		order by country 
		limit 1 
	`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		log.Println("error no able to query databaases for origin")
		return nil, err
	}

	defer rows.Close()

	var origins []*models.Origin

	for rows.Next() {
		var origin models.Origin
		err := rows.Scan(
			&origin.OriginId,
			&origin.Country,
			&origin.State,
			&origin.CityTown,
		)
		if err != nil {
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		origins = append(origins, &origin)
	}
	return origins, nil	
}

func (m *PostgresDBRepo) PostOrigin(origin *models.Origin) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	id := 0
	query := `
	INSERT INTO public.origins (Country, State, CityTown ) VALUES
	($1, $2, $3)
	RETURNING originid`
	err := m.DB.QueryRowContext(ctx, query,
		origin.Country,
		origin.State,
		origin.CityTown).Scan(&id)

	if err != nil {
		log.Println("error no able to query databaases for origins")
		log.Panic(err)
		return "", err
	}

	originid := "New origin created with id:" + strconv.FormatInt(int64(id), 10)
	return originid, nil
}
