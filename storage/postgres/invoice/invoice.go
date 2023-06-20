package invoice

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sjaureguio/paypal/models"
	"github.com/sjaureguio/paypal/storage/postgres"
)

const (
	queryInsert  = "INSERT INTO invoices (id, invoice_date, customer_email, is_product, is_subscription, product_id, subscription_id, price) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	querySelect  = "SELECT * FROM invoices"
	queryByID    = " WHERE id = $1"
	queryByEmail = " WHERE email = $1"
)

type Invoice struct {
	db *sql.DB
}

func New(db *sql.DB) Invoice {
	return Invoice{db: db}
}

func (i Invoice) Create(invoice *models.Invoice) error {
	emptyCtx := context.Background()

	stmt, err := i.db.PrepareContext(emptyCtx, queryInsert)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			return
		}
	}(stmt)

	row, err := stmt.ExecContext(
		emptyCtx,
		invoice.ID,
		invoice.InvoiceDate,
		invoice.CustomerEmail,
		invoice.IsProduct,
		invoice.IsSubscription,
		invoice.ProductID,
		invoice.SubscriptionID,
		invoice.Price,
	)

	if err != nil {
		return err
	}

	got, err := row.RowsAffected()

	if err != nil {
		return err
	}

	if got != 1 {
		return fmt.Errorf("expected 1 row affected, got %d", got)
	}

	return nil
}

func (i Invoice) FindByID(ID uuid.UUID) (models.Invoice, error) {
	emptyCtx := context.Background()

	stmt, err := i.db.PrepareContext(emptyCtx, querySelect+queryByID)
	if err != nil {
		return models.Invoice{}, err
	}

	defer stmt.Close()

	row, err := stmt.QueryContext(emptyCtx, ID)
	if err != nil {
		return models.Invoice{}, err
	}

	return i.scan(row)
}

func (i Invoice) FindByEmail(email string) (models.Invoices, error) {
	emptyCtx := context.Background()

	stmt, err := i.db.PrepareContext(emptyCtx, querySelect+queryByEmail)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(emptyCtx, email)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var resp models.Invoices

	for rows.Next() {
		row, err := i.scan(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, row)
	}

	return resp, nil
}

func (i Invoice) scan(r postgres.RowScanner) (models.Invoice, error) {
	updateAtNull := sql.NullTime{}
	m := models.Invoice{}

	err := r.Scan(
		&m.ID,
		&m.InvoiceDate,
		&m.CustomerEmail,
		&m.IsProduct,
		&m.IsSubscription,
		&m.ProductID,
		&m.SubscriptionID,
		&m.Price,
		&m.CreatedAt,
		&updateAtNull,
	)
	if err != nil {
		return models.Invoice{}, err
	}

	m.UpdatedAt = updateAtNull.Time

	return m, nil
}
