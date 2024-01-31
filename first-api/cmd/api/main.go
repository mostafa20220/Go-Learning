package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mostafa20220/go-api/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){

	log.SetReportCaller(true)

	// use chi
	r := chi.NewRouter()

	handlers.Handler(r)

	fmt.Println("Starting GO Server API...")

	err := http.ListenAndServe("localhost:8080",r)
	
	if err!=nil{
		log.Error(err)
	}

}