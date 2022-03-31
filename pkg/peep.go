package peep

import "context"

type Beer struct {
	ID       int     `json:"ID"`
	Name     string  `json:"name"`
	Brewery  string  `json:"Brewery"`
	Country  string  `json:"country"`
	Price    float32 `json:"price"`
	Currency string  `json:"currency"`
}

type BeerBox struct {
	TotalPrice float32 `json:"price_total"`
}

type FetchAnswer struct {
	Beers []Beer `json:"beers"`
}

type FetchByIdAnswer struct {
	BeerByID Beer `json:"beer"`
}

type PriceByBox struct {
	PriceBox BeerBox `json:"BeerBox"`
}

type Repository interface {
	SetUpDB(context.Context) error
	SeedDB(context.Context) error
}
