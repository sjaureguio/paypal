package invoice

import (
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/invoice"
	"github.com/sjaureguio/paypal/models"
	"net/http"
)

type Handler struct {
	useCase invoice.Invoice
}

func New(uc invoice.UseCase) Handler {
	return Handler{useCase: uc}
}

func (h Handler) FindByEmail(c echo.Context) error {
	email := c.Param("email")

	invoices, err := h.useCase.FindByEmail(email)
	if err != nil {
		msg := map[string]string{
			"error":    "No se puedo consultar las facturas",
			"internal": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, msg)
	}

	resp := map[string]models.Invoices{
		"data": invoices,
	}

	return c.JSON(http.StatusOK, resp)
}
