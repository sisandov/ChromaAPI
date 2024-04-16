package main

import (
	"net/http"
	"chroma-api/router"
)

func main() {
	r:= router.Router()
	http.ListenAndServe(":3333", r)
}
