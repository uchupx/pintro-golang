package model

import "time"

type Transaction struct {
	Id        uint64    `json:"id" db:"id"`
	Customer  string    `json:"customer" db:"customer"`
	Quantity  uint64    `json:"quantity" db:"quantity"`
	Price     float64   `json:"price" db:"price"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}
