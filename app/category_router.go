package app

import (
	"github.com/fauzanebd/gj-restful-api/controller"
	"github.com/fauzanebd/gj-restful-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewCategoryRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// Error handler
	router.PanicHandler = exception.ErrorHandler

	return router
}
