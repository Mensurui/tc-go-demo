package customers

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type CustomerRepository struct {
	conn *pgx.Conn
}

func NewRepository(ctx context.Context, connStr string) (*CustomerRepository, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the database")
		return nil, err
	}

	return &CustomerRepository{
		conn: conn,
	}, nil
}

func (r CustomerRepository) CreateCustomer(ctx context.Context, customer Customer) (Customer, error) {
	query := `
	INSERT INTO customers (name, email)
	VALUES($1, $2)
	RETURNING id
	`
	err := r.conn.QueryRow(ctx, query, customer.Name, customer.Email).Scan(&customer.Id)
	return customer, err
}
