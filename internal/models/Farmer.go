package models

type Farmer struct {
	FarmerId         int    `json:"farmerid"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Altitude         string `json:"altitude"`
}
