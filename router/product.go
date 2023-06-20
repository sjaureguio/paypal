package router

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/product"
	handler "github.com/sjaureguio/paypal/handlers/product"
	storage "github.com/sjaureguio/paypal/storage/postgres/product"
)

func Product(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := product.New(store)
	handler.NewRouter(e, useCase)
}
