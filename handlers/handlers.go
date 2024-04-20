package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse := map[string]string{
		"message": "Hello, world!",
	}

	CommonResponse(w, r, jsonResponse)
}

const totalNumbers = 255

func GenerateRandomRGBNumber() int64 {
	return int64(rand.IntN(totalNumbers))
}

func CommonResponse(w http.ResponseWriter, r *http.Request, jsonResponse map[string]string) {
	jsonResponseBytes, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, ResponseError := w.Write(jsonResponseBytes)
	if ResponseError != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GenerateRandomHex() string {
	return NumberToHex(GenerateRandomRGBNumber())
}

func NumberToHex(num int64) string {
	generatedHexNumber := strings.ToUpper(strconv.FormatInt(num, 16))
	return strings.Repeat("0", 2-len(generatedHexNumber)) + generatedHexNumber
}

func RandomHEXHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse := map[string]string{
		"color": fmt.Sprintf("#%s%s%s", GenerateRandomHex(), GenerateRandomHex(), GenerateRandomHex()),
	}
	CommonResponse(w, r, jsonResponse)
}

func RandomRGBHandler(w http.ResponseWriter, r *http.Request) {
	red := fmt.Sprint(GenerateRandomRGBNumber())
	green := fmt.Sprint(GenerateRandomRGBNumber())
	blue := fmt.Sprint(GenerateRandomRGBNumber())

	jsonResponse := map[string]string{
		"color": fmt.Sprintf("(%s, %s, %s)", red, green, blue),
		"red":   red,
		"green": green,
		"blue":  blue,
	}
	CommonResponse(w, r, jsonResponse)
}
