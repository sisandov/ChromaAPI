package router

import (
	"chroma-api/handlers"
	"chroma-api/router/random"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)



func NewRouter() http.Handler { 
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.RootHandler)

	randomRouter := random.Router()
	r.Mount("/", randomRouter)
	return r
}
