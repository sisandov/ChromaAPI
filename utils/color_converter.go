package color_converter

import (
	"chroma-api/handlers"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func ParseRGB(rgbRaw string) []string {
	re := regexp.MustCompile(`\(|\)|rgb|RGB|\s`)
	parsedRGBColor := re.ReplaceAllString(rgbRaw, "")
	return strings.SplitN(parsedRGBColor, ",", 3)
}

func ParseHSL(hslRaw string) []string {
	re := regexp.MustCompile(`\(|\)|hsl|HSL|\s`)
	parsedHSLColor := re.ReplaceAllString(hslRaw, "")
	return strings.SplitN(parsedHSLColor, ",", 3)
}

func RGBToHEX(rgbColor string) (string, error) {
	uniqueColors := ParseRGB(rgbColor)
	if len(uniqueColors) != 3 {
		return "", errors.New("color have a bad format")
	}

	redInt, errRed := strconv.ParseInt(uniqueColors[0], 10, 64)
	greenInt, errGreen := strconv.ParseInt(uniqueColors[1], 10, 64)
	blueInt, errBlue := strconv.ParseInt(uniqueColors[2], 10, 64)

	if errRed != nil || errGreen != nil || errBlue != nil {
		return "", errors.New("there was an error during the parsing of the numbers")
	}

	red := handlers.NumberToHex(redInt)
	green := handlers.NumberToHex(greenInt)
	blue := handlers.NumberToHex(blueInt)

	return fmt.Sprintf("#%s%s%s", red, green, blue), nil
}

func HEXToRGB(HEXColor string) (string, error) {
	/*
		The input of HEXColor is like #f2b2aa or #f1431f, akso can be in uppercase
	*/

	parseHEX := strings.ToUpper(strings.Replace(strings.Trim(HEXColor, " "), "#", "", -1))

	redHex := parseHEX[0:2]
	greenHex := parseHEX[2:4]
	blueHex := parseHEX[4:6]
	// Add check that every string is of length 2
	red, errRed := strconv.ParseInt(redHex, 16, 64)
	green, errGreen := strconv.ParseInt(greenHex, 16, 64)
	blue, errBlue := strconv.ParseInt(blueHex, 16, 64)

	if errRed != nil || errGreen != nil || errBlue != nil {
		return "", errors.New("error parsing the colors")
	}

	return fmt.Sprintf("rgb(%d,%d,%d)", red, green, blue), nil
}

func RGBToHSL(rgbColor string) (string, error) {
	// Input can be rgb(123, 32, 321)
	uniqueColors := ParseRGB(rgbColor)
	if len(uniqueColors) != 3 {
		return "", fmt.Errorf("error parsing colors %v", uniqueColors)
	}

	redInt, errRed := strconv.ParseInt(uniqueColors[0], 10, 64)
	greenInt, errGreen := strconv.ParseInt(uniqueColors[1], 10, 64)
	blueInt, errBlue := strconv.ParseInt(uniqueColors[2], 10, 64)

	if errRed != nil || errGreen != nil || errBlue != nil {
		return "", fmt.Errorf("error parsing numbers red: %s, green: %s, blue: %s", errRed, errBlue, errGreen)
	}

	normalRed := float64(redInt) / 255
	normalGreen := float64(greenInt) / 255
	normalBlue := float64(blueInt) / 255

	minColor := min(normalRed, normalGreen, normalBlue)
	maxColor := max(normalRed, normalGreen, normalBlue)
	differenceColor := maxColor - minColor

	// Determine the lightness
	lightness := (maxColor + minColor) / 2

	// Saturation
	var saturation float64
	if lightness == 0 {
		saturation = 0
	} else {
		saturation = (differenceColor) / (1 - math.Abs(2*lightness-1))
	}

	// Wheel tree
	var hue float64 = 0

	if saturation == 0 {
		hue = 0
	} else {
		switch int64(max(redInt, greenInt, blueInt)) {
		case redInt:
			hue = 60*(normalGreen-normalBlue)/differenceColor + 0
		case greenInt:
			hue = 60*(normalBlue-normalRed)/differenceColor + 120
		default:
			hue = 60*(normalBlue-normalGreen)/differenceColor + 240
		}
	}

	return fmt.Sprintf("hsl(%.0f,%.2f,%.2f)", hue, saturation, lightness), nil
}

func HEXToHSL(HEXColor string) (string, error) {
	rgbColor, errTransform := HEXToRGB(HEXColor)

	if errTransform != nil {
		return "", fmt.Errorf("cannot transform hex to rgb first")
	}

	hslColor, errHSL := RGBToHSL(rgbColor)
	if errHSL != nil {
		return "", fmt.Errorf("problem generating the hsl color from rgb: %s", errHSL.Error())
	}

	return hslColor, nil
}

func HSLToRGB(HSLColor string) (string, error) {
	parsedHSL := ParseHSL(HSLColor)
	hue, hueError := strconv.ParseFloat(parsedHSL[0], 64)
	saturation, saturationError := strconv.ParseFloat(parsedHSL[1], 64)
	lightness, lightnessError := strconv.ParseFloat(parsedHSL[2], 64)

	if hueError != nil || saturationError != nil || lightnessError != nil {
		return "", fmt.Errorf("error in values of hsl %s %s %s", hueError, saturationError, lightnessError)
	}
	var red, green, blue float64

	hueNormalized := hue / 360

	if saturation == 0 {
		red = lightness
		green = lightness
		blue = lightness
	} else {
		var q float64
		if lightness < 0.5 {
			q = lightness * (1 + saturation)
		} else {
			q = lightness + saturation - lightness*saturation
		}
		p := 2*lightness - q
		red = ParseHSLSingleColor(p, q, hueNormalized+float64(1)/3)
		green = ParseHSLSingleColor(p, q, float64(hueNormalized))
		blue = ParseHSLSingleColor(p, q, hueNormalized-float64(1)/3)
	}

	return fmt.Sprintf("rgb(%.0f,%.0f,%.0f)", red*255, green*255, blue*255), nil
}

func ParseHSLSingleColor(p, q, t float64) float64 {
	switch {
	case t < 0:
		t += 1
	case t > 1:
		t -= 1
	}

	switch {
	case t < float64(1)/6:
		return p + (q-p)*6*t
	case t < float64(1)/2:
		return q
	case t < float64(2)/3:
		return p + (q-p)*(float64(2)/3-t)*6
	}
	return p
}

func HSLToHEX(HSLColor string) (string, error) {
	rgbColor, rgbColorError := HSLToRGB(HSLColor)
	if rgbColorError != nil {
		return "", fmt.Errorf("error parsing the HSL Color to HSL: %s", rgbColorError)
	}

	println("HELLO")
	println(rgbColor)

	return RGBToHEX(rgbColor)
}
