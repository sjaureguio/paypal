package subscription

import (
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/subscription"
)

const (
	path     = "/v1/subscriptions"
	pathByID = "/:email"
)

func NewRouter(e *echo.Echo, useCase subscription.Subscription) {
	handler := New(useCase)

	g := e.Group(path)
	g.GET(pathByID, handler.FindByEmail)
}
