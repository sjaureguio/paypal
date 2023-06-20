package order

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/order"
	"github.com/sjaureguio/paypal/models"
	"net/http"
)

type Handler struct {
	useCase order.Order
}

func New(uc order.Order) Handler {
	return Handler{useCase: uc}
}

func (h Handler) Create(c echo.Context) error {
	o := models.Order{}

	err := c.Bind(&o)
	if err != nil {
		msg := map[string]string{
			"error":    "La estructura de la orden no es correcta",
			"internal": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	err = h.useCase.Create(&o)
	if err != nil {
		msg := map[string]string{
			"error":    "No pudimos crear al orden",
			"internal": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, map[string]models.Order{"message": o})
}

func (h Handler) FindByID(c echo.Context) error {
	ID := c.Param("id")
	data, err := h.useCase.FindByID(uuid.MustParse(ID))
	if err != nil {
		msg := map[string]string{
			"error":    "No pudimos consultar la orden",
			"internal": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	resp := map[string]models.Order{
		"data": data,
	}

	return c.JSON(http.StatusOK, resp)
}
