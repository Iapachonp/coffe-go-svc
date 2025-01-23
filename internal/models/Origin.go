package models

type Origin struct {
	OriginId         int    `json:"originid"`
	Country          string `json:"country"`
	State            string `json:"state"`
	CityTown         string `json:"citytown"`
}
