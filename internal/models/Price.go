package models

import (
	"github.com/jackc/pgtype"
)

type Price struct {
	PriceId int          `json:"priceid"`
	Prices  pgtype.JSONB `json:"prices"`
}

type InnerPrice struct {
	Grams string  `json:"grams"`
	Price float64 `json:"price,string"`
}

type PriceGo struct {
	PriceId int          `json:"priceid"`
	Prices  []InnerPrice `json:"prices"`
}
