package service

import (
	"context"
	"database/sql"

	"github.com/fauzanebd/gj-restful-api/exception"
	"github.com/fauzanebd/gj-restful-api/helper"
	"github.com/fauzanebd/gj-restful-api/model/domain"
	"github.com/fauzanebd/gj-restful-api/model/web"
	"github.com/fauzanebd/gj-restful-api/repository"
	"github.com/go-playground/validator"
)

type CategorySeviceImpl struct {
	CategoryRepo repository.CategoryRepository
	DB           *sql.DB
	Validate     *validator.Validate
}

func NewCategoryService(categoryRepo repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategorySeviceImpl{
		CategoryRepo: categoryRepo,
		DB:           db,
		Validate:     validate,
	}
}

func (service *CategorySeviceImpl) Create(ctx context.Context, request *web.CategoryCreateRequest) (*web.CategoryResponse, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer tx.Commit()
	catRes, err := service.CategoryRepo.Insert(
		ctx,
		tx,
		domain.Category{Name: request.Name},
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	categoryResponse := helper.ToCategoryResponse(*catRes)
	return &categoryResponse, nil

}
func (service *CategorySeviceImpl) Update(ctx context.Context, request *web.CategoryUpdateRequest) (*web.CategoryResponse, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer tx.Commit()

	// check if theres category with given Id
	_, err = service.CategoryRepo.FindById(ctx, tx, request.Id)
	if err != nil {
		tx.Rollback()
		return nil, exception.NotFoundError{NotFoundDetails: err.Error()}
	}

	catRes, err := service.CategoryRepo.Update(
		ctx,
		tx,
		domain.Category{Id: request.Id, Name: request.Name},
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	categoryResponse := helper.ToCategoryResponse(*catRes)
	return &categoryResponse, nil
}
func (service *CategorySeviceImpl) Delete(ctx context.Context, categoryId int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	defer tx.Commit()

	// check if theres category with given Id
	_, err = service.CategoryRepo.FindById(ctx, tx, categoryId)
	if err != nil {
		tx.Rollback()
		return exception.NotFoundError{NotFoundDetails: err.Error()}
	}

	err = service.CategoryRepo.Delete(ctx, tx, categoryId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (service *CategorySeviceImpl) FindById(ctx context.Context, categoryId int) (*web.CategoryResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer tx.Commit()
	catRes, err := service.CategoryRepo.FindById(ctx, tx, categoryId)
	if err != nil {
		tx.Rollback()
		return nil, exception.NotFoundError{NotFoundDetails: err.Error()}
	}
	categoryResponse := helper.ToCategoryResponse(*catRes)
	return &categoryResponse, nil
}
func (service *CategorySeviceImpl) FindAll(ctx context.Context) (*[]web.CategoryResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer tx.Commit()
	catResults, err := service.CategoryRepo.FindAll(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if len((*catResults)) > 0 {
		catResponses := make([]web.CategoryResponse, len(*catResults))
		for idx, val := range *catResults {
			catResponses[idx] = helper.ToCategoryResponse(val)
		}
		return &catResponses, nil
	} else {
		return nil, exception.NotFoundError{NotFoundDetails: "No categories data found"}
	}
}
