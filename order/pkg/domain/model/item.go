package model

import (
	"github.com/google/uuid"
)

type Item struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Price    float64   `json:"price"`
	Quantity int       `json:"quantity"`
}
