package service

import (
	"context"

	"github.com/fauzanebd/gj-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request *web.CategoryCreateRequest) (*web.CategoryResponse, error)
	Update(ctx context.Context, request *web.CategoryUpdateRequest) (*web.CategoryResponse, error)
	Delete(ctx context.Context, categoryId int) error
	FindById(ctx context.Context, categoryId int) (*web.CategoryResponse, error)
	FindAll(ctx context.Context) (*[]web.CategoryResponse, error)
}
