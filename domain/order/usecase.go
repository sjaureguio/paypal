package order

import (
	"github.com/google/uuid"
	"github.com/sjaureguio/paypal/models"
)

type UseCase struct {
	storage Storage
}

func New(s Storage) UseCase {
	return UseCase{storage: s}
}

func (uc UseCase) Create(o *models.Order) error {
	o.ID = uuid.New()
	return uc.storage.Create(o)
}

func (uc UseCase) FindByID(ID uuid.UUID) (models.Order, error) {
	return uc.storage.FindByID(ID)
}
