package middleware

import (
	"net/http"

	"github.com/fauzanebd/gj-restful-api/exception"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "someapikey" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		exception.ErrorHandler(
			writer,
			request,
			exception.AuthorizationError{
				AuthorizationDetails: "you provide wrong api key",
			},
		)
	}
}
