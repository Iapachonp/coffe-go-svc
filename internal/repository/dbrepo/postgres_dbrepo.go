package dbrepo

import (
	"coffee-svc/internal/models"
	"context"
	"database/sql"
	"log"
	"time"
	"fmt"
)

type PostgresDBRepo struct {

	DB *sql.DB
}

const dbTimeout = time.Second * 3


func (m *PostgresDBRepo) Connection() (*sql.DB) {

	return m.DB


}

func printRows(rows *sql.Rows){

    colTypes, err := rows.ColumnTypes()
    if err != nil {
		log.Println("error no able to print rows for coffees")
	}
    for _,s := range colTypes {
      log.Println("cols type:", s.Name(), "  " ,s.DatabaseTypeName(), " to gp " , s.ScanType());
    }
}

func (m *PostgresDBRepo) AllCoffees() ([]*models.Coffee, error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
			c.coffeeid as ID, c.name as Name , pr.name as Proccess, o.state || ' ' || o.citytown as Origin, v.description Description, sca SCA, p.grams as Weight , p.values as Price , coalesce(image, '') as Image, Createdat, Updatedat 
		from 
			coffees c
	        join varietals v on c.varietalid =  v.varietalid
		join prices p on c.priceid =  p.priceid
		join origin o on c.originid =  o.originid
		join Proccess pr on c.proccessid =  pr.proccessid
		order by 
			name
	`
	
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("error no able to query databaases for coffees")
		return nil, err
	}

	defer rows.Close()

	var coffees []*models.Coffee

	
	for rows.Next() {
		var coffee models.Coffee
		err := rows.Scan(
			&coffee.ID, 
			&coffee.Name,
			&coffee.Proccess,
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


func (m *PostgresDBRepo) Coffee(id string) ([]*models.Coffee, error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := fmt.Sprintf(`
		select 
			c.coffeeid as ID, c.name as Name , pr.name as Proccess, o.state || ' ' || o.citytown as Origin, v.description Description, sca SCA, p.grams as Weight , p.values as Price , coalesce(image, '') as Image, Createdat, Updatedat 
		from 
			coffees c
	        join varietals v on c.varietalid =  v.varietalid
		join prices p on c.priceid =  p.priceid
		join origin o on c.originid =  o.originid
		join Proccess pr on c.proccessid =  pr.proccessid
		where c.coffeeid = %s
		order by name
		limit 1 
	`, id )

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("error no able to query databaases for coffees")
		return nil, err
	}

	defer rows.Close()

	var coffees []*models.Coffee

	
	for rows.Next() {
		var coffee models.Coffee
		err := rows.Scan(
			&coffee.ID, 
			&coffee.Name,
			&coffee.Proccess,
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
