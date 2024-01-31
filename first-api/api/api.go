package api

import (
	"encoding/json"
	"net/http"
)

type CoinsAccountParam struct{
	Username string
}

type CoinsBalanceResponse struct{
	Code int
	Account int64
}

type Error struct{
	Code int
	Message string
}


func writeError(w http.ResponseWriter ,message string, code int){
	res := Error {
		Code : code,
		Message: message,
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(res)
}

func ResponseError(w http.ResponseWriter ,err error){
	writeError(w, err.Error(), http.StatusBadRequest)
}

func InternalError(w http.ResponseWriter){
	writeError(w, "An Unexpected Internal Error Occurred", http.StatusInternalServerError)
}