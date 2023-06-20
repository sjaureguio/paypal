package models

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	ID            uuid.UUID `json:"id"`
	CustomerEmail string    `json:"customer_email"`
	Status        string    `json:"status"`
	TypeSubs      string    `json:"type_subs"`
	BeginsAt      time.Time `json:"begins_at"`
	EndsAt        time.Time `json:"ends_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Subscriptions is a slice of Subscription
type Subscriptions []Subscription

const (
	Monthly = "monthly"
	Annual  = "Annual"
)

const (
	Active     = "ACTIVE"
	Terminated = "TERMINATED"
)
