package exception

import (
	"net/http"

	"github.com/fauzanebd/gj-restful-api/helper"
	"github.com/fauzanebd/gj-restful-api/model/web"
	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if errorVal, ok := err.(NotFoundError); ok {
		notFoundError(writer, request, errorVal)
	} else if errorVal, ok := err.(validator.ValidationErrors); ok {
		validationError(writer, request, errorVal)
	} else if errorVal, ok := err.(AuthorizationError); ok {
		authorizationError(writer, request, errorVal)
	} else {
		internalServerError(writer, request, err)
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	helper.WriteToResponseBodyJSON(writer, web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
		Data:   err.(error).Error(),
	})
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	helper.WriteToResponseBodyJSON(writer, web.WebResponse{
		Code:   http.StatusNotFound,
		Status: "NOT FOUND",
		Data:   err.(error).Error(),
	})
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {

	helper.WriteToResponseBodyJSON(writer, web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err.(error).Error(),
	})
}

func authorizationError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	helper.WriteToResponseBodyJSON(writer, web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
		Data:   err.(error).Error(),
	})
}
