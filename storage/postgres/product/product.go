package product

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/sjaureguio/paypal/models"
	"github.com/sjaureguio/paypal/storage/postgres"
	"log"
)

type Product struct {
	db *sql.DB
}

func New(db *sql.DB) Product {
	return Product{db: db}
}

func (p Product) FindAll() (models.Products, error) {
	rows, err := p.db.Query(query + queryAll)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("Pool de conexión no cerrada ", err)
		}
	}(rows)

	var resp models.Products
	for rows.Next() {
		row, err := p.scan(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resp, nil
}

func (p Product) FindByID(ID uuid.UUID) (models.Product, error) {
	row := p.db.QueryRow(query+queryByID, ID)

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
