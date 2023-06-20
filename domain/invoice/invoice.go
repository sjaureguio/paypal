package invoice

import (
	"github.com/sjaureguio/paypal/models"
)

// ************************
// ******* POR OUT ********
// ************************

type Storage interface {
	Create(i *models.Invoice) error
	FindByEmail(email string) (models.Invoices, error)
}

// ************************
// ******* POR IN *********
// ************************

type Invoice interface {
	FindByEmail(email string) (models.Invoices, error)
}
