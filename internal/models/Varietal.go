package models

type Varietal struct {
	VarietalId       int    `json:"varietalId"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	BeanSize         string `json:"beanSize"`
	Stature          string `json:"stature"`
	QualityPotential string `json:"qualityPotential"`
	OptimalAltitude  string `json:"optimalAltitude"`
}
