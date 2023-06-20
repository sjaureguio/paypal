package order

import (
	"github.com/google/uuid"
	"github.com/sjaureguio/paypal/models"
)

// ************************
// ******* POR OUT ********
// ************************

type Storage interface {
	Create(o *models.Order) error
	FindByID(ID uuid.UUID) (models.Order, error)
}

// ************************
// ******* POR IN *********
// ************************

type Order interface {
	Create(o *models.Order) error
	FindByID(ID uuid.UUID) (models.Order, error)
}
