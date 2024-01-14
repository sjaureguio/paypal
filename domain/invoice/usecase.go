package invoice

import (
	"github.com/google/uuid"
	"github.com/sjaureguio/paypal/models"
	"time"
)

type UseCase struct {
	storage Storage
}

func (uc UseCase) FindByEmail(email string) (models.Invoices, error) {
	// TODO: implement me
	panic("implement me")
}

func New(s Storage) UseCase {
	return UseCase{storage: s}
}

func (uc UseCase) Create(order *models.Order, subsID uuid.UUID) error {
	i := models.Invoice{}

	if order.IsSubscription {
		i.SubscriptionID = subsID
	}

	i.ID = uuid.New()
	i.CustomerEmail = order.CustomerEmail
	i.InvoiceDate = time.Now()
	i.IsProduct = order.IsProduct
	i.IsSubscription = order.IsSubscription
	i.Price = order.Price
	i.ProductID = order.ProductID

	return uc.storage.Create(&i)
}

func (uc UseCase) FIndByEmail(email string) (models.Invoices, error) {
	return uc.storage.FindByEmail(email)
}
