package models

type Client struct {
	UserID    string
	VIN       string
	ChosenSKU string
	Status    string
}

type DealerCar struct {
	ID       string
	MarketID string
	SKU      string
	VINs     []string
	CityID   string
}
