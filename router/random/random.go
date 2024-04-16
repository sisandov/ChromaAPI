package random

import (
	"chroma-api/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/random/hex", handlers.RandomHEXHandler)
	r.Get("/random/rgb", handlers.RandomRGBHandler)
	return r
}
