package models

import (
	"github.com/google/uuid"
	"time"
)

const TABLE = "products"

type Product struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	Image          string    `json:"image"`
	IsSubscription bool      `json:"is_subscription"`
	Months         uint8     `json:"months"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Products is a slice of Order
type Products []Product
