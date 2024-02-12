package controller

import (
	"net/http"
	"strconv"

	"github.com/fauzanebd/gj-restful-api/helper"
	"github.com/fauzanebd/gj-restful-api/model/web"
	"github.com/fauzanebd/gj-restful-api/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}

func (c CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	categoryRequest := &web.CategoryCreateRequest{}
	err := helper.ReadFromRequestBodyJSON(request, categoryRequest)
	if err != nil {
		panic(err)
	}

	catRes, err := c.CategoryService.Create(request.Context(), categoryRequest)
	if err != nil {
		panic(err)
	}

	err = helper.WriteToResponseBodyJSON(writer, web.WebResponse{
		Code:   200,
		Status: "Category created successfully",
		Data:   catRes,
	})
	if err != nil {
		panic(err)
	}
}
func (c CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	if err != nil {
		panic(err)
	}

	categoryRequest := &web.CategoryUpdateRequest{}
	err = helper.ReadFromRequestBodyJSON(request, categoryRequest)
	if err != nil {
		panic(err)
	}
	categoryRequest.Id = categoryId

	catRes, err := c.CategoryService.Update(request.Context(), categoryRequest)
	if err != nil {
		panic(err)
	}

	err = helper.WriteToResponseBodyJSON(writer, web.WebResponse{
		Code:   200,
		Status: "Category updated successfully",
		Data:   catRes,
	})
	if err != nil {
		panic(err)
	}

}
func (c CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	if err != nil {
		panic(err)
	}
	err = c.CategoryService.Delete(request.Context(), categoryId)
	if err != nil {
		panic(err)
	}

	err = helper.WriteToResponseBodyJSON(writer, web.WebResponse{
		Code:   200,
		Status: "Category deleted successfully",
	})
	if err != nil {
		panic(err)
	}
}
func (c CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	if err != nil {
		panic(err)
	}
	catRes, err := c.CategoryService.FindById(request.Context(), categoryId)
	if err != nil {
		panic(err)
	}

	err = helper.WriteToResponseBodyJSON(writer, web.WebResponse{
		Code:   200,
		Status: "Category found",
		Data:   catRes,
	})
	if err != nil {
		panic(err)
	}
}
func (c CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	catResults, err := c.CategoryService.FindAll(request.Context())
	if err != nil {
		panic(err)
	}

	err = helper.WriteToResponseBodyJSON(writer, web.WebResponse{
		Code:   200,
		Status: "Categories found",
		Data:   catResults,
	})
	if err != nil {
		panic(err)
	}
}
