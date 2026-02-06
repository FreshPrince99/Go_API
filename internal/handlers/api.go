package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/FreshPrince99/Go_API/internal/middleware"
)

func Handler(r *chi.Mux) { // we write the function in capital letter because we want to tell the compiler that it can used outside the package
	// Global middleware
	r.Use(chimiddle.StripSlashes) // remove trailing slashes from the request URL

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance) // GET /account/coins

	})
}