package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID             uuid.UUID `json:"id"`
	CustomerEmail  string    `json:"customer_email"`
	IsProduct      bool      `json:"is_product"`
	IsSubscription bool      `json:"is_subscription"`
	ProductID      uuid.UUID `json:"product_id"`
	TypeSubs       string    `json:"type_subs"`
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Orders is a slice of Order
type Orders []Order
