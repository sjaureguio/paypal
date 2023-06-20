package order

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sjaureguio/paypal/models"
	"github.com/sjaureguio/paypal/storage/postgres"
)

const (
	queryInsert = `INSERT INTO orders(id, 
    					       customer_email, 
    					       is_product, 
    					       is_subscription, 
    					       product_id, 
    					       type_subs, 
    					       price)
					VALUES ($1, $2, $3, $4, $5, $6, $7) 
				`
	queryByID = `SELECT * FROM orders WHERE id = $1`
)

type Order struct {
	db *sql.DB
}

func New(db *sql.DB) Order {
	return Order{db: db}
}

func (o Order) Create(order *models.Order) error {
	emptyCtx := context.Background()
	stmt, err := o.db.PrepareContext(emptyCtx, queryInsert)
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
		order.ID,
		order.CustomerEmail,
		order.IsProduct,
		order.IsSubscription,
		order.ProductID,
		order.TypeSubs,
		order.Price,
	)

	if err != nil {
		return err
	}

	got, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if got != 1 {
		return fmt.Errorf("expected 1 row, got %d", got)
	}

	return nil
}

func (o Order) FindByID(ID uuid.UUID) (models.Order, error) {
	emptyCtx := context.Background()
	stmt, err := o.db.PrepareContext(emptyCtx, queryByID)
	if err != nil {
		return models.Order{}, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			return
		}
	}(stmt)

	row := stmt.QueryRowContext(emptyCtx, ID)

	return o.scan(row)

}

func (o Order) scan(row postgres.RowScanner) (models.Order, error) {
	productIDNull := sql.NullString{}
	typeSubsNull := sql.NullString{}
	updatedNull := sql.NullTime{}
	order := models.Order{}

	err := row.Scan(
		&order.ID,
		&order.CustomerEmail,
		&order.IsProduct,
		&order.IsSubscription,
		&productIDNull,
		&typeSubsNull,
		&order.Price,
		&order.CreatedAt,
		&updatedNull,
	)

	if err != nil {
		return models.Order{}, err
	}

	order.ProductID = uuid.MustParse(productIDNull.String)
	order.TypeSubs = typeSubsNull.String

	return order, nil
}
