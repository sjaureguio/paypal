package product

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/sjaureguio/paypal/models"
	"github.com/sjaureguio/paypal/storage/postgres"
)

const (
	query     = "SELECT * FROM products"
	queryAll  = " ORDER BY name"
	queryByID = " WHERE id = $1"
)

type Product struct {
	db *sql.DB
}

func New(db *sql.DB) Product {
	return Product{db: db}
}

func (p Product) FindAll() (models.Products, error) {
	emptyCtx := context.Background()
	stmt, err := p.db.PrepareContext(emptyCtx, query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(emptyCtx)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var resp models.Products

	for rows.Next() {
		prod, err := p.scan(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, prod)
	}

	return resp, nil
}

func (p Product) FindByID(ID uuid.UUID) (models.Product, error) {
	emptyCtx := context.Background()
	stmt, err := p.db.PrepareContext(emptyCtx, query+queryByID)
	if err != nil {
		return models.Product{}, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			return
		}
	}(stmt)

	row := stmt.QueryRowContext(emptyCtx, ID)

	return p.scan(row)
}

func (p Product) scan(r postgres.RowScanner) (models.Product, error) {
	updateNull := sql.NullTime{}
	resp := models.Product{}

	err := r.Scan(
		&resp.ID,
		&resp.Name,
		&resp.Description,
		&resp.Image,
		&resp.IsSubscription,
		&resp.Months,
		&resp.Price,
		&resp.CreatedAt,
		&updateNull,
	)

	if err != nil {
		return models.Product{}, err
	}

	resp.UpdatedAt = updateNull.Time

	return resp, nil
}
