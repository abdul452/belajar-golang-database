package repository

import (
	"context"
	"database/sql"

	"github.com/abdul452/belajar-golang-database/entity"
	_ "github.com/go-sql-driver/mysql"
)

type customerRepositoryImpl struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepositoryImpl{DB: db}
}

func (repository *customerRepositoryImpl) Insert(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	script := `INSERT INTO customer(id, name) VALUES(?, ?)`
	_, err := repository.DB.ExecContext(ctx, script, customer.Id, customer.Name)
	if err != nil {
		return customer, err
	}
	return customer, nil
}
