package repository

import (
	"context"

	"github.com/abdul452/belajar-golang-database/entity"
)

type CustomerRepository interface {
	Insert(ctx context.Context, customer entity.Customer) (entity.Customer, error)
	// FindById(ctx context.Context, id string) (entity.Customer, error)
	// FindAll(ctx context.Context) ([]entity.Customer, error)
}
