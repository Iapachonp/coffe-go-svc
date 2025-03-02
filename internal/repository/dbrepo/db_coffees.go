package dbrepo

import (
	"coffee-svc/internal/models"
	"context"
	"log"
	"strconv"
)

func (m *PostgresDBRepo) AllCoffees() ([]*models.CoffeeData, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
			c.coffeeid as ID, c.name as Name , pr.name as Process, o.state || ' ' || o.citytown as Origin, v.description Description, sca SCA, p.grams as Weight , p.values as Price , coalesce(image, '') as Image, Createdat, Updatedat 
		from 
			coffees c
	        join varietals v on c.varietalid =  v.varietalid
		join prices p on c.priceid =  p.priceid
		join origins o on c.originid =  o.originid
		join Processes pr on c.processid =  pr.processid
		order by 
			name
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("error no able to query databaases for coffees")
		return nil, err
	}

	defer rows.Close()

	var coffees []*models.CoffeeData

	for rows.Next() {
		var coffee models.CoffeeData
		err := rows.Scan(
			&coffee.ID,
			&coffee.Name,
			&coffee.Process,
			&coffee.Origin,
			&coffee.Description,
			&coffee.Sca,
			&coffee.Weight,
			&coffee.Price,
			&coffee.Image,
			&coffee.CreatedAt,
			&coffee.UpdatedAt,
		)
		if err != nil {
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		coffees = append(coffees, &coffee)
	}

	return coffees, nil
}

func (m *PostgresDBRepo) Coffee(id string) ([]*models.CoffeeData, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select 
			c.coffeeid as ID, c.name as Name , pr.name as Process, o.state || ' ' || o.citytown as Origin, v.description Description, sca SCA, p.grams as Weight , p.values as Price , coalesce(image, '') as Image, Createdat, Updatedat 
		from 
			coffees c
	        join varietals v on c.varietalid =  v.varietalid
		join prices p on c.priceid =  p.priceid
		join origins o on c.originid =  o.originid
		join Processes pr on c.processid =  pr.processid
		where c.coffeeid = $1
		order by name
		limit 1 
	`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		log.Println("error no able to query databaases for coffees")
		return nil, err
	}

	defer rows.Close()

	var coffees []*models.CoffeeData

	for rows.Next() {
		var coffee models.CoffeeData
		err := rows.Scan(
			&coffee.ID,
			&coffee.Name,
			&coffee.Process,
			&coffee.Origin,
			&coffee.Description,
			&coffee.Sca,
			&coffee.Weight,
			&coffee.Price,
			&coffee.Image,
			&coffee.CreatedAt,
			&coffee.UpdatedAt,
		)
		if err != nil {
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		coffees = append(coffees, &coffee)
	}

	return coffees, nil
}

func (m *PostgresDBRepo) PostCoffee(coffee *models.Coffee) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	id := 0
	query := `
	INSERT INTO coffees (
		Name,
		varietalId,
		FarmerId,
		OriginId,
		ProcessId,
		SCA,
		Acidity,
		Body,
		Balance,
		Clarity,
		Sweetness,
		PriceId,
		Image,
		Createdat,
		Updatedat
		)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	RETURNING CoffeeId`
	err := m.DB.QueryRowContext(ctx, query,
		coffee.Name,
		coffee.VarietalId,
		coffee.FarmerId,
		coffee.OriginId,
		coffee.ProcessId,
		coffee.Sca,
		coffee.Acidity,
		coffee.Body,
		coffee.Balance,
		coffee.Clarity,
		coffee.Sweetness,
		coffee.PriceId,
		coffee.Image,
		coffee.CreatedAt,
		coffee.UpdatedAt).Scan(&id)

	if err != nil {
		log.Println("error no able to query databaases for coffees")
		log.Panic(err)
		return "", err
	}

	coffeid := "New Coffee created with id:" + strconv.FormatInt(int64(id), 10)
	return coffeid, nil
}


func (m *PostgresDBRepo) DeleteCoffee(id string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		DELETE FROM coffees WHERE coffeeid = $1; 
	`

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		log.Println("error no able to query databaases for coffees")
		return false, err
	}

	rows, err := result.RowsAffected()
	
	if err != nil {
		log.Println(err)
		return false, err
	}
	
	if rows != 1 {
		log.Println("expected to affect 1 row, affected %d", rows)
		return false, err
		}
	return true, nil

}
	
