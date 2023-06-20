package invoice

import (
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/invoice"
)

const (
	path     = "/v1/invoices"
	pathByID = "/:email"
)

func NewRouter(e *echo.Echo, useCase invoice.UseCase) {
	handler := New(useCase)

	g := e.Group(path)
	g.GET(pathByID, handler.FindByEmail)
}
