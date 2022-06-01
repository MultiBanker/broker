package models

import "time"

type MarketAuto struct {
	ID       string
	AutoSKU  string
	MarketID string
	CityID   string
	VIN      string

	CreatedAt time.Time
	UpdatedAt time.Time
}
