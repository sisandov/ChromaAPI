package handlers

import (
	"encoding/json"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
    jsonResponse := map[string]string{
        "message": "Hello, world!",
    }

    jsonResponseBytes, err := json.Marshal(jsonResponse)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponseBytes)
}
