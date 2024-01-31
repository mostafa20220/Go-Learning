package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/mostafa20220/go-api/api"
	"github.com/mostafa20220/go-api/internal/tools"
	"github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	// read request data: params
		// username := r.URL.Query().Get("username")
		
	//? use gorilla/schema package this time
	params := api.CoinsAccountParam{}
	decoder := schema.NewDecoder()

	decoder.Decode(&params, r.URL.Query())

	database, err := tools.NewDatabase()

	if err != nil {
		logrus.Error(err)
		api.InternalError(w)
		return
	}

	coinsDetails := (*database).GetUserCoinsDetails(params.Username)

	if coinsDetails == nil {
		api.InternalError(w)
		return
	}

	res := api.CoinsBalanceResponse{
		Code: 200,
		Account: coinsDetails.Coins,
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		logrus.Error(err)
		api.InternalError(w)
		return
	}

}