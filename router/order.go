package router

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/order"
	handler "github.com/sjaureguio/paypal/handlers/order"
	storage "github.com/sjaureguio/paypal/storage/postgres/order"
)

func Order(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := order.New(store)
	handler.NewRouter(e, useCase)
}
