package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mostafa20220/go-api/internal/middlewares"
	// "github.com/mostafa20220/go-api/internal/middlewares"
)

func Handler(r* chi.Mux){
	
	// Global Middlewares
	middleware.StripSlashes(r)

	r.Route("/account", func(r chi.Router) {

		// use auth middleware
		r.Use(middlewares.Authorization)

		r.Get("/coins", GetCoinBalance)

	})

}

