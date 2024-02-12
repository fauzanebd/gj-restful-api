package test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/fauzanebd/gj-restful-api/app"
	"github.com/fauzanebd/gj-restful-api/controller"
	"github.com/fauzanebd/gj-restful-api/middleware"
	"github.com/fauzanebd/gj-restful-api/model/web"
	"github.com/fauzanebd/gj-restful-api/repository"
	"github.com/fauzanebd/gj-restful-api/service"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB(additionalParameters ...string) (*sql.DB, error) {
	additionalParams := strings.Join(additionalParameters, "&")
	var db *sql.DB
	var err error
	if len(additionalParameters) > 0 {
		db, err = sql.Open("mysql", fmt.Sprintf("root:makan@tcp(localhost:3306)/test_categoryapi?%s", additionalParams))
	} else {
		db, err = sql.Open("mysql", "root:makan@tcp(localhost:3306)/test_categoryapi")
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, err
}

func setupRouter() (http.Handler, error) {
	db, err := setupTestDB("parseTime=True")
	if err != nil {
		return nil, err
	}
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewCategoryRouter(categoryController)
	return middleware.NewAuthMiddleware(router), nil
}

func TestCreateCategorySuccess(t *testing.T) {
	router, err := setupRouter()
	if err != nil {
		t.Errorf("Error setting up router : %v", err)
	}
	testData := web.CategoryCreateRequest{
		Name: "Shopping",
	}
	jsonData, err := json.Marshal(testData)
	if err != nil {
		t.Errorf("Error creating json test data : %v", err)
	}

	request := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:3000/api/categories",
		strings.NewReader(string(jsonData)),
	)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "someapikey")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	webResp := &web.WebResponse{}
	bodyRes := recorder.Result().Body
	decoder := json.NewDecoder(bodyRes)
	decoder.Decode(webResp)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Shopping", webResp.Data.(map[string]interface{})["name"].(string))
}

func TestCreateCategoryFailed(t *testing.T) {
	router, err := setupRouter()
	if err != nil {
		t.Errorf("Error setting up router : %v", err)
	}
	testData := web.CategoryCreateRequest{
		Name: "",
	}
	jsonData, err := json.Marshal(testData)
	if err != nil {
		t.Errorf("Error creating json test data : %v", err)
	}

	request := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:3000/api/categories",
		strings.NewReader(string(jsonData)),
	)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "someapikey")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	// expected to return bad request because of the empty name given in body
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestUpdateCategorySuccess(t *testing.T) {}

func TestUpdateCategoryFailed(t *testing.T) {}

func TestGetCategoryByIDSuccess(t *testing.T) {}

func TestGetCategoryByIDFailed(t *testing.T) {}

func TestDeleteCategorySuccess(t *testing.T) {}

func TestDeleteCategoryFailed(t *testing.T) {}

func TestGetListCategoriesSuccess(t *testing.T) {}

func TestUnauthorizedRequest(t *testing.T) {}
