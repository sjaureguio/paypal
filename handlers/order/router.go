package order

import (
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/order"
)

const (
	path       = "/v1/orders"
	pathCreate = ""
	pathByID   = "/:id"
)

func NewRouter(e *echo.Echo, useCase order.Order) {
	handler := New(useCase)

	g := e.Group(path)
	g.POST(pathCreate, handler.Create)
	g.GET(pathByID, handler.FindByID)
}
