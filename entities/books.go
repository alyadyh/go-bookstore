package entities

import "time"

type Books struct {
	Id              uint
	Title           string
	Author          Authors
	Publisher       Publishers
	Genre           Genres
	PublicationDate string
	ISBN            string
	Price           float32
	StockQty        int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
