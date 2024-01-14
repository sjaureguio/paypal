package router

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/sjaureguio/paypal/domain/invoice"
	"github.com/sjaureguio/paypal/domain/order"
	"github.com/sjaureguio/paypal/domain/paypal"
	"github.com/sjaureguio/paypal/domain/subscription"

	storageInvoice "github.com/sjaureguio/paypal/storage/postgres/invoice"
	storageOrder "github.com/sjaureguio/paypal/storage/postgres/order"
	storageSubs "github.com/sjaureguio/paypal/storage/postgres/subscription"

	handler "github.com/sjaureguio/paypal/handlers/paypal"
)

func PayPal(e *echo.Echo, db *sql.DB) {
	useCaseOrder := buildOrder(db)
	useCaseSubs := buildSubs(db)
	useCaseInvoice := buildInvoice(db)

	useCasePayPal := paypal.New(useCaseOrder, useCaseSubs, useCaseInvoice)

	handler.NewRouter(e, useCasePayPal)
}

func buildOrder(db *sql.DB) paypal.Order {
	store := storageOrder.New(db)
	return order.New(store)
}

func buildSubs(db *sql.DB) paypal.Subscription {
	store := storageSubs.New(db)
	return subscription.New(store)
}

func buildInvoice(db *sql.DB) paypal.Invoice {
	store := storageInvoice.New(db)
	return invoice.New(store)
}
