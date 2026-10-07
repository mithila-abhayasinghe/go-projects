package internal

import (
	"fmt"
	"strings"
)

type ConversionResult struct {
	FromValue float64
	FromUnit  Unit
	ToValue   float64
	ToUnit    Unit
}

func Convert(val float64, from, to Unit) (ConversionResult, error) {
	if from.Dimension != to.Dimension {
		return ConversionResult{}, fmt.Errorf("conversion failed: dimension mismatch %s vs %s", from.Dimension, to.Dimension)
	}

	conversion := val * from.ToBaseRatio / to.ToBaseRatio
	return ConversionResult{
		FromValue: val,
		FromUnit:  from,
		ToValue:   conversion,
		ToUnit:    to,
	}, nil
}

// ----- Temperature Conversion -----

type TempConversionResult struct {
	FromValue float64
	FromUnit  string
	ToValue   float64
	ToUnit    string
}

func ConvertTemperature(val float64, from, to string) (TempConversionResult, error) {

	fromNormalized := tempStringNormalize(from)
	toNormalized := tempStringNormalize(to)

	var result TempConversionResult
	result.FromValue = val
	result.FromUnit = fromNormalized
	result.ToUnit = toNormalized

	switch fromNormalized {
	case "K":
		switch toNormalized {
		case "C":
			formulaResult := val - 273.15
			result.ToValue = formulaResult
		case "F":
			formulaResult := (val-273.15)*9.0/5.0 + 32
			result.ToValue = formulaResult
		}
	case "C":
		switch toNormalized {
		case "K":
			formulaResult := val + 273.15
			result.ToValue = formulaResult
		case "F":
			formulaResult := val*1.8 + 32
			result.ToValue = formulaResult
		}
	case "F":
		switch toNormalized {
		case "K":
			formulaResult := (val-32)*5.0/9.0 + 273.15
			result.ToValue = formulaResult
		case "C":
			formulaResult := (val - 32) * 5.0 / 9.0
			result.ToValue = formulaResult
		}

	}

	return result, nil
}

func tempStringNormalize(str string) string {
	// Normalize first
	normalized := strings.ToLower(strings.TrimSpace(str))

	switch normalized {
	case "celsius", "c":
		return "C"
	case "fahrenheit", "f":
		return "F"
	case "kelvin", "k":
		return "K"
	default:
		return ""
	}
}
