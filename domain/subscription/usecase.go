package subscription

import (
	"github.com/google/uuid"
	"github.com/sjaureguio/paypal/models"
	"time"
)

const (
	month = 1
	year  = 12
)

type UseCase struct {
	storage Storage
}

func New(s Storage) UseCase {
	return UseCase{storage: s}
}

func (uc UseCase) Create(s *models.Subscription) error {
	months := month

	if s.TypeSubs == models.Annual {
		months = year
	}

	s.ID = uuid.New()
	s.BeginsAt = time.Now()
	s.EndsAt = s.BeginsAt.AddDate(0, months, 0)
	s.Status = models.Active

	return uc.storage.Create(s)
}

func (uc UseCase) FindByEmail(email string) (models.Subscriptions, error) {
	return uc.storage.FindByEmail(email)
}
