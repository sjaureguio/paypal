package product

import (
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/product"
)

const (
	path     = "/v1/products"
	pathAll  = ""
	pathByID = "/:id"
)

func NewRouter(e *echo.Echo, useCase product.Product) {
	handler := New(useCase)

	g := e.Group(path)
	g.GET(pathAll, handler.FindAll)
	g.GET(pathByID, handler.FindByID)
}
