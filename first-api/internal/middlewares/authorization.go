package middlewares

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/mostafa20220/go-api/api"
	"github.com/mostafa20220/go-api/internal/tools"
)

var unAuthorizedError = errors.New("You are not Authorized to do this action")

func Authorization(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userName := r.URL.Query().Get("username")
		token := r.Header.Get("authorization") 


		if userName== "" || token == "" {
			log.Error(unAuthorizedError)
			api.ResponseError(w, unAuthorizedError)
			return
		}

		// compare with the db
		database, err := tools.NewDatabase()
		if err != nil {
			log.Error(err)
			api.ResponseError(w,err)
			return
		}

		userLoginDetails := (*database).GetUserLoginDetails(userName)

		if userLoginDetails == nil || userLoginDetails.Token != token {
			log.Error(unAuthorizedError)
			api.ResponseError(w,unAuthorizedError)
			return 
		}

		next.ServeHTTP(w,r)

	})
}