package color_converter

import (
	"chroma-api/handlers"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RGBToHEX(rgbColor string) (string, error) {
	re := regexp.MustCompile(`\(|\)|rgb|RGB`)
	parsedRGBColor := re.ReplaceAllString(rgbColor, "")
	uniqueColors := strings.SplitN(parsedRGBColor, ",", 3)
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
		The input of HEXColor is like #f2b2aa or #f1431f
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

	return fmt.Sprintf("rbg(%d,%d,%d)", red, green, blue), nil
}
