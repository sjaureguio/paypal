package subscription

import (
	"github.com/sjaureguio/paypal/models"
)

// ************************
// ******* POR OUT ********
// ************************

type Storage interface {
	Create(s *models.Subscription) error
	FindByEmail(email string) (models.Subscriptions, error)
}

// ************************
// ******* POR IN *********
// ************************

type Subscription interface {
	FindByEmail(email string) (models.Subscriptions, error)
}
