package color_converter_test

import (
	color_converter "chroma-api/utils"
	"fmt"
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

func TestRGBToHSL(t *testing.T) {
	color, colorError := color_converter.RGBToHSL("rgb(123, 123, 123)")
	if colorError != nil {
		t.Fatalf("Error in the RGBToHSL converter: %s", colorError.Error())
	}
	println(color)
}

func TestHEXToHSL(t *testing.T) {
	colorHSL, colorError := color_converter.HEXToHSL("#AAAAAA")
	if colorError != nil {
		t.Fatalf("Error in the RGBToHSL converter: %s", colorError.Error())
	}
	println(colorHSL)
}

func TestParseRGB(t *testing.T) {
	rgbSlice := color_converter.ParseRGB("rgb(1,2,3)")
	if len(rgbSlice) != 3 {
		t.Fatalf("Error on slice %v", rgbSlice)
	}

	println(fmt.Sprintf("%v", rgbSlice))
}

func TestParseHSLToRGB(t *testing.T) {
	hslExample, hslError := color_converter.HSLToRGB("hsl(204,0.5,0.3)")
	if hslError != nil {
		t.Fatalf("Error on test %s", hslError.Error())
	}
	println(hslExample)
}

func TestParseHSLToHEX(t *testing.T) {
	hexColor, hexColorError := color_converter.HSLToHEX("hsl(204,0.5,0.3)")
	if hexColorError != nil {
		t.Fatalf("Error on test %s", hexColorError.Error())
	}
	println(hexColor)
}
