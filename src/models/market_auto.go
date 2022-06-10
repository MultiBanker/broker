package models

import "time"

type MarketAuto struct {
	ID       string `bson:"_id"`
	AutoSKU  string `bson:"auto_sku"`
	MarketID string `bson:"market_id"`
	CityID   string `bson:"city_id"`
	VIN      string `bson:"vin"`
	Status   string `bson:"status"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
