package product

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

func (uc UseCase) FindAll() (models.Products, error) {
	return uc.storage.FindAll()
}

func (uc UseCase) FindByID(ID uuid.UUID) (models.Product, error) {
	return uc.storage.FindByID(ID)
}
