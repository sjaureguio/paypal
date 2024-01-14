package product

import (
	"github.com/google/uuid"
	"github.com/sjaureguio/paypal/models"
)

// ************************
// ******* PORT OUT *******
// ************************

type Storage interface {
	FindAll() (models.Products, error)
	FindByID(ID uuid.UUID) (models.Product, error)
}

// ************************
// ******* PORT IN ********
// ************************

// Equivalent a service

type Product interface {
	FindAll() (models.Products, error)
	FindByID(ID uuid.UUID) (models.Product, error)
}
