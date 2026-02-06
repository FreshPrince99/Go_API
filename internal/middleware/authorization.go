package middleware

import (
	"errors"
	"net/http"

	"github.com/FreshPrince99/Go_API/api"
	"github.com/FreshPrince99/Go_API/internal/tools"
	log "github.com/sirupsen/logrus"

)

var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	// The handler function that will be returned by the middleware, it will be called for every request that goes through this middleware
	// The parameters w and r represent the response writer and the request respectively, they are used to write the response and read the request data
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.URL.Query().Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return 
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r) // if the user is authorized, we call the next handler in the chain, which is the actual handler for the endpoint

	})
}

