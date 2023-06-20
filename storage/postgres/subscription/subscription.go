package subscription

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sjaureguio/paypal/models"
	"github.com/sjaureguio/paypal/storage/postgres"
)

const (
	queryInsert        = "INSERT INTO subscriptions (id, customer_email, status, types_subs, begins_at, ends_At) VALUES $1, $2, $3, $4, $5, $6"
	querySelect        = "SELECT * FROM subscriptions"
	querySelectByEmail = " WHERE email = $1"
)

type Subscription struct {
	db *sql.DB
}

func New(db *sql.DB) Subscription {
	return Subscription{db: db}
}

func (s Subscription) Create(subs *models.Subscription) error {
	emptyCtx := context.Background()

	stmt, err := s.db.PrepareContext(emptyCtx, queryInsert)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row, err := stmt.ExecContext(
		emptyCtx,
		subs.ID,
		subs.CustomerEmail,
		subs.Status,
		subs.TypeSubs,
		subs.BeginsAt,
		subs.EndsAt,
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

func (s Subscription) FindByEmail(email string) (models.Subscriptions, error) {
	emptyCtx := context.Background()

	stmt, err := s.db.PrepareContext(emptyCtx, querySelectByEmail)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(emptyCtx, email)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var resp models.Subscriptions
	for rows.Next() {
		row, err := s.scan(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, row)
	}

	return resp, nil
}

func (s Subscription) scan(r postgres.RowScanner) (models.Subscription, error) {
	updateAtNull := sql.NullTime{}
	m := models.Subscription{}

	err := r.Scan(
		&m.ID,
		&m.CustomerEmail,
		&m.Status,
		&m.TypeSubs,
		&m.BeginsAt,
		&m.EndsAt,
		&m.CreatedAt,
		&updateAtNull,
	)
	if err != nil {
		return models.Subscription{}, err
	}

	m.UpdatedAt = updateAtNull.Time

	return m, nil
}
