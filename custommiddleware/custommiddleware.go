package custommiddleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func BasicAuth(next http.Handler) http.Handler {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not load .env")
	}

	authKey := os.Getenv("AUTH_KEY")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header not present", http.StatusUnauthorized)
			return
		}

		authToken := strings.Split(authHeader, " ")
		if authToken[0] != "Bearer" {
			http.Error(w, "Wrong Authorization header format. The correct format is: Bearer <your token>", http.StatusUnauthorized)
			return
		}
		if authToken[1] != authKey {
			http.Error(w, "Wrong Authorization token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
