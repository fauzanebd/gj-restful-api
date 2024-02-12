package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/fauzanebd/gj-restful-api/model/domain"
)

type CategoryRepoImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepoImpl{}
}

func (repo *CategoryRepoImpl) FindAll(ctx context.Context, tx *sql.Tx) (*[]domain.Category, error) {
	getAllQuery := "SELECT id, name, createdAt, updatedAt FROM Category;"
	row, err := tx.QueryContext(ctx, getAllQuery)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var listOfCategories []domain.Category
	var tempCategory = domain.Category{}
	for row.Next() {
		err := row.Scan(
			&tempCategory.Id,
			&tempCategory.Name,
			&tempCategory.CreatedAt,
			&tempCategory.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		listOfCategories = append(listOfCategories, tempCategory)
	}
	return &listOfCategories, nil
}
func (repo *CategoryRepoImpl) Insert(ctx context.Context, tx *sql.Tx, category domain.Category) (*domain.Category, error) {
	postQuery := "INSERT INTO Category (name, createdAt, updatedAt) VALUES (?, ?, ?);"
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	res, err := tx.ExecContext(
		ctx,
		postQuery,
		category.Name,
		category.CreatedAt.Format("2006-01-02"),
		category.UpdatedAt.Format("2006-01-02"),
	)
	if err != nil {
		return nil, err
	}
	lastCategoryId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	category.Id = int(lastCategoryId)
	return &category, nil
}
func (repo *CategoryRepoImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (*domain.Category, error) {
	findQuery := "SELECT id, name, createdAt, updatedAt FROM Category WHERE id = ?;"
	row, err := tx.QueryContext(ctx, findQuery, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var returnedCategory = domain.Category{}
	if row.Next() {
		err := row.Scan(
			&returnedCategory.Id,
			&returnedCategory.Name,
			&returnedCategory.CreatedAt,
			&returnedCategory.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &returnedCategory, nil
	} else {
		return nil, fmt.Errorf("no category with Id: %d", id)
	}

}
func (repo *CategoryRepoImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) (*domain.Category, error) {
	updateQuery := `
		UPDATE Category 
		SET name = ?, updatedAt = ?
		WHERE id = ?;
	`
	category.UpdatedAt = time.Now()
	_, err := tx.ExecContext(ctx, updateQuery, category.Name, category.UpdatedAt.Format("2006-01-02"), category.Id)
	if err != nil {
		return nil, err
	}
	return &category, nil
}
func (repo *CategoryRepoImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	deleteQuery := "DELETE FROM Category WHERE id = ?;"
	_, err := tx.ExecContext(ctx, deleteQuery, id)
	if err != nil {
		return err
	}
	return nil
}
