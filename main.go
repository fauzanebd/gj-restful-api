package main

import (
	"net/http"

	"github.com/fauzanebd/gj-restful-api/app"
	"github.com/fauzanebd/gj-restful-api/controller"
	"github.com/fauzanebd/gj-restful-api/middleware"
	"github.com/fauzanebd/gj-restful-api/repository"
	"github.com/fauzanebd/gj-restful-api/service"
	"github.com/go-playground/validator"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	categoryRepository := repository.NewCategoryRepository()
	db, err := app.GetDBConnection("parseTime=True")
	if err != nil {
		panic(err)
	}
	validate := validator.New()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	categoryRouter := app.NewCategoryRouter(categoryController)
	authMiddleware := middleware.NewAuthMiddleware(categoryRouter)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
