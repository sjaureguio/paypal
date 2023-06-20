package router

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/invoice"
	handler "github.com/sjaureguio/paypal/handlers/invoice"
	storage "github.com/sjaureguio/paypal/storage/postgres/invoice"
)

func Invoice(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := invoice.New(store)
	handler.NewRouter(e, useCase)
}
