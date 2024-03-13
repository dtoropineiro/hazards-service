package handlers

import (
	"strconv"
	"strings"
)

type MagnitudeHandler struct{}

func GetAnnotationSize(magnitude string) float32 {
	magnitudeFloat := extractFloatFromString(magnitude)
	magnitudeIntegerPart := int(magnitudeFloat)

	if magnitudeIntegerPart < 2 {
		return 10.0
	} else {
		return 10.0 + float32(magnitudeIntegerPart)*5.0
	}
}

func extractFloatFromString(stringWithUnit string) float32 {
	numberOnly := strings.TrimFunc(stringWithUnit, func(r rune) bool {
		return !((r >= '0' && r <= '9') || r == '.')
	})

	parsedFloat, _ := strconv.ParseFloat(numberOnly, 32)
	return float32(parsedFloat)
}
