package internal

import (
	"fmt"
)

type ConversionEvent struct {
	FromValue float64
	FromUnit  Unit
	ToUnit    Unit
}

type ConversionResult struct {
	FromValue float64
	FromUnit  Unit
	ToValue   float64
	ToUnit    Unit
}

func ConverstionEventHandler(e ConversionEvent) (ConversionResult, error) {
	return Convert(e.FromValue, e.FromUnit, e.ToUnit)
}

func Convert(val float64, from, to Unit) (ConversionResult, error) {

	// Handle Temperature workaround
	if isTemperature(from, to) {
		return ConvertTemperature(val, from, to)
	}
	f := GetUnitData(from)
	t := GetUnitData(to)
	if f.Dimension != t.Dimension {
		return ConversionResult{}, fmt.Errorf("conversion failed: dimension mismatch %s vs %s", f.Dimension, t.Dimension)
	}
	conversion := val * f.ToBaseRatio / t.ToBaseRatio
	return ConversionResult{
		FromValue: val,
		FromUnit:  from,
		ToValue:   conversion,
		ToUnit:    to,
	}, nil
}

func isTemperature(f, t Unit) bool {

	fCheck := (f == Kelvin || f == Celsius || f == Fahrenheit)
	tCheck := (t == Kelvin || t == Celsius || t == Fahrenheit)

	if fCheck && tCheck {
		return true
	}
	return false
}
