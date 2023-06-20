package router

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/subscription"
	handler "github.com/sjaureguio/paypal/handlers/subscription"
	storage "github.com/sjaureguio/paypal/storage/postgres/subscription"
)

func Subscription(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := subscription.New(store)
	handler.NewRouter(e, useCase)
}
