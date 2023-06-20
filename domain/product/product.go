package product

import (
	"github.com/google/uuid"
	"github.com/sjaureguio/paypal/models"
)

// ************************
// ******* POR OUT ********
// ************************

type Storage interface {
	FindAll() (models.Products, error)
	FindByID(ID uuid.UUID) (models.Product, error)
}

// ************************
// ******* POR IN *********
// ************************

type Product interface {
	FindAll() (models.Products, error)
	FindByID(ID uuid.UUID) (models.Product, error)
}
