package router

import (
	"chroma-api/handlers"
	"chroma-api/router/random"
	"chroma-api/custommiddleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)



func Router() http.Handler { 
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(custommiddleware.BasicAuth)
	r.Get("/", handlers.RootHandler)

	randomRouter := random.Router()
	r.Mount("/", randomRouter)
	return r
}
