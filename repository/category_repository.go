package repository

import (
	"context"
	"database/sql"

	"github.com/fauzanebd/gj-restful-api/model/domain"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) (*[]domain.Category, error)
	Insert(ctx context.Context, tx *sql.Tx, category domain.Category) (*domain.Category, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (*domain.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) (*domain.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
}
