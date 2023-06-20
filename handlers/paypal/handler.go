package paypal

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/paypal"
	"log"
	"net/http"
)

type Handler struct {
	useCase paypal.PayPal
}

func New(paypal paypal.PayPal) Handler {
	return Handler{useCase: paypal}
}

func (h Handler) Webhook(c echo.Context) error {
	var body []byte
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != nil {
		resp := map[string]string{
			"error":    "no se pudo leer la informacion",
			"internal": err.Error(),
		}
		log.Print(resp)

		return c.JSON(http.StatusBadRequest, resp)
	}

	go func() {
		err = h.useCase.ProcessRequest(c.Request().Header, body)
		if err != nil {
			log.Print("error procesando el webhook", err)
		}
	}()

	resp := map[string]string{"message": "OK"}
	return c.JSON(http.StatusOK, resp)
}
