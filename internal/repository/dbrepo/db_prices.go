package dbrepo 

import (
	"coffee-svc/internal/models"
	"context"
	"log"
	"strconv"
	"encoding/json"
	
)

func (m *PostgresDBRepo) AllPrices() ([]models.PriceGo, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select priceid, to_jsonb(prices) from prices group by priceid ; `

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("error no able to query database for prices")
		return nil, err
	}

	defer rows.Close()

	var prices []*models.Price

	for rows.Next() {
		var price models.Price
		err := rows.Scan(
			&price.PriceId,
			&price.Prices,
		)
		if err != nil {
			log.Println("heyyyyyyyyyyyyyyyyyyyy")
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		prices = append(prices, &price)
	}

	var pricesGo []models.PriceGo
	
	for price := range prices {

		var innerPrice []models.InnerPrice 
		err = json.Unmarshal(prices[price].Prices.Bytes, &innerPrice)
		if err != nil {
			log.Println("Error unmarshaling data:", err)
			return nil, err
		}
		priceid := prices[price].PriceId
		pricego := models.PriceGo{ priceid, innerPrice}
		pricesGo = append(pricesGo , pricego )
	}
	return pricesGo, nil
}

func (m *PostgresDBRepo) Price(id string) ([]*models.Price, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select *  
		from prices 
		where priceid = $1
		order by country 
		limit 1 
	`

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		log.Println("error no able to query databaases for price")
		return nil, err
	}

	defer rows.Close()

	var prices []*models.Price

	for rows.Next() {
		var price models.Price
		err := rows.Scan(
			&price.PriceId,
			&price.Prices,
		)
		if err != nil {
			printRows(rows)
			log.Panic(err)
			return nil, err
		}

		prices = append(prices, &price)
	}
	return prices, nil	
}

func (m *PostgresDBRepo) PostPrice(price *models.Price) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	id := 0
	query := `
	INSERT INTO public.prices (Country, State, CityTown ) VALUES
	($1, $2, $3)
	RETURNING priceid`
	err := m.DB.QueryRowContext(ctx, query,
		price.Prices).Scan(&id)

	if err != nil {
		log.Println("error no able to query databaases for prices")
		log.Panic(err)
		return "", err
	}

	priceid := "New price created with id:" + strconv.FormatInt(int64(id), 10)
	return priceid, nil
}
