package product

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/product"
	"github.com/sjaureguio/paypal/models"
	"net/http"
)

type Handler struct {
	useCase product.Product
}

func New(uc product.Product) Handler {
	return Handler{useCase: uc}
}

func (h Handler) FindAll(c echo.Context) error {
	data, err := h.useCase.FindAll()
	if err != nil {
		msg := map[string]string{
			"error":    "No pudimos consultar la info",
			"internal": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, msg)
	}

	resp := map[string]models.Products{"data": data}

	return c.JSON(http.StatusOK, resp)
}

func (h Handler) FindByID(c echo.Context) error {
	ID := c.Param("id")
	data, err := h.useCase.FindByID(uuid.MustParse(ID))
	if err != nil {
		msg := map[string]string{
			"error":    "No se pudo consultar el producto",
			"internal": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, msg)
	}

	resp := map[string]models.Product{
		"data": data,
	}

	return c.JSON(http.StatusOK, resp)
}
