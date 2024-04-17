package color_converter_test

import (
	color_converter "chroma-api/utils"
	"testing"
)

func TestRGBToHEX(t *testing.T) {
	color, colorError := color_converter.RGBToHEX("rgb(23,43,12)")
	if colorError != nil {
		t.Fatalf("Error in the color converter: %s", colorError.Error())
	}
	println(color)
}

func TestHEXToRGB(t *testing.T) {
	color, colorError := color_converter.HEXToRGB("#FFFFFF")
	if colorError != nil {
		t.Fatalf("Error in the color converter: %s", colorError.Error())
	}
	println(color)
}
