package handlers_test

import (
	"chroma-api/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHexGenerator(t *testing.T) {

	hex := handlers.GenerateRandomHex()
	if len(hex) != 2 {
		t.Errorf("Length of generated hex should be 2 but is %s", hex)
	}
}

func TestRGBNumber(t *testing.T) {

	rgbNumber := handlers.GenerateRandomRGBNumber()
	if rgbNumber < 0 && rgbNumber > 255 {
		t.Errorf("The number should be between 0 and 255 but is: %d", rgbNumber)
	}
}

func TestHEXHandler(t *testing.T) {
	rr := httptest.NewRecorder()

	// Call the handler function directly
	handlers.RandomHEXHandler(rr, nil)

	// Check the status code of the response
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	// Parse the JSON response body
	var jsonResponse map[string]string

	err := json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
	if err != nil {
		t.Errorf("error parsing JSON response: %v", err)
	}

	hexCodeLength := 7
	if len(jsonResponse["color"]) != hexCodeLength {
		t.Errorf("HEX color code should have length 7 but has length %v", len(jsonResponse["color"]))
	}
}
