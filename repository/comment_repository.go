package repository

import (
	"context"

	"github.com/abdul452/belajar-golang-database/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comments) (entity.Comments, error)
	FindById(ctx context.Context, id int) (entity.Comments, error)
	FindAll(ctx context.Context) ([]entity.Comments, error)
}
