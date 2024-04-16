package random

import (
	"chroma-api/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)



func Router() http.Handler { 
	r := chi.NewRouter()
	r.Get("/random", handlers.RandomHandler)
	return r
}