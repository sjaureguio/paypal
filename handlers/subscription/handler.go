package subscription

import (
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/subscription"
	"github.com/sjaureguio/paypal/models"
	"net/http"
)

type Handler struct {
	useCase subscription.Subscription
}

func New(uc subscription.Subscription) Handler {
	return Handler{useCase: uc}
}

func (h Handler) FindByEmail(c echo.Context) error {
	email := c.Param("email")

	subscriptions, err := h.useCase.FindByEmail(email)
	if err != nil {
		msg := map[string]string{
			"error":    "No se pudo consultar las suscripciones",
			"internal": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, msg)
	}

	resp := map[string]models.Subscriptions{
		"data": subscriptions,
	}

	return c.JSON(http.StatusOK, resp)
}
