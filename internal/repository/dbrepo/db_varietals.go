package dbrepo 

import (
	"coffee-svc/internal/models"
	"context"
	"log"
	"strconv"
)

func (m *PostgresDBRepo) AllVarietals() ([]*models.Varietal, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select *
		from varietals
		order by name
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("error no able to query database for varietals")
		return nil, err
	}

	defer rows.Close()

	var varietals []*models.Varietal

	for rows.Next() {
		var varietal models.Varietal
		err := rows.Scan(
			&varietal.VarietalId,
			&varietal.Name,
			&varietal.Description,
			&varietal.BeanSize,
			&varietal.Stature,
			&varietal.QualityPotential,
			&varietal.OptimalAltitude,
		)
		if err != nil {
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		varietals = append(varietals, &varietal)
	}

	return varietals, nil
}

func (m *PostgresDBRepo) Varietal(id string) ([]*models.Varietal, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select *  
		from varietals 
		where varietalid = $1
		order by name
		limit 1 
	`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		log.Println("error no able to query databaases for varietal")
		return nil, err
	}

	defer rows.Close()

	var varietals []*models.Varietal

	for rows.Next() {
		var varietal models.Varietal
		err := rows.Scan(
			&varietal.VarietalId,
			&varietal.Name,
			&varietal.Description,
			&varietal.BeanSize,
			&varietal.Stature,
			&varietal.QualityPotential,
			&varietal.OptimalAltitude,
		)
		if err != nil {
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		varietals = append(varietals, &varietal)
	}
	return varietals, nil	
}

func (m *PostgresDBRepo) PostVarietal(varietal *models.Varietal) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	id := 0
	query := `
	INSERT INTO public.varietals (Name, Description, BeanSize, Stature, QualityPotential, OptimalAltitude) VALUES
	($1, $2, $3, $4, $5, $6)
	RETURNING varietalId`
	err := m.DB.QueryRowContext(ctx, query,
		varietal.Name,
		varietal.Description,
		varietal.BeanSize,
		varietal.Stature,
		varietal.QualityPotential,
		varietal.OptimalAltitude).Scan(&id)

	if err != nil {
		log.Println("error no able to query databaases for varietals")
		log.Panic(err)
		return "", err
	}

	varitealid := "New varietal created with id:" + strconv.FormatInt(int64(id), 10)
	return varitealid, nil
}
