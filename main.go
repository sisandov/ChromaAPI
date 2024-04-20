package main

import (
	"chroma-api/router"
	"fmt"
	"net/http"
)

func main() {
	r := router.Router()
	serverError := http.ListenAndServe(":3333", r)
	if serverError != nil {
		fmt.Printf("Error on server start: %s", serverError)
	}
}
