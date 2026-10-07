package main

import (
	"fmt"
	unit "unit-converter/internal"
)

func main() {

	fmt.Println("Unit Converter")

	inputValue := 1000
	fromUnit := unit.Gram
	toUnit := unit.Kilogram
	result, _ := unit.Convert(float64(inputValue), fromUnit, toUnit)
	fmt.Println(result.ToValue)

	inputValue = 1
	result, _ = unit.Convert(float64(inputValue), unit.Kilogram, unit.Gram)
	fmt.Println(result.ToValue)

	// something i could not do before Convert hand temperature
	inputValue = 100
	result, _ = unit.Convert(float64(inputValue), unit.Celsius, unit.Kelvin)
	fmt.Println(result.ToValue)

	// inputValue = 1
	// result, _ = unit.Convert(float64(inputValue), unit.Get(unit.Kilogram), unit.Get(unit.Gram))
	// fmt.Println(result.ToValue)

	// inputValue := 1000
	// fromUnit := unit.Get(unit.Gram)
	// toUnit := unit.Get(unit.Kilogram)
	// result, _ := unit.Convert(float64(inputValue), , )
	// fmt.Println(result.ToValue)

	// inputValue = 1
	// fromUnit = internal.SupportedUnits["kilogram"]
	// toUnit = internal.SupportedUnits["pound"]
	// result, _ = internal.Convert(float64(inputValue), fromUnit, toUnit)
	// fmt.Println(result.ToValue)

	// inputValueT := 100
	// fromUnitT := "C"
	// toUnitT := "F"
	// resultT, _ := internal.ConvertTemperature(float64(inputValueT), fromUnitT, toUnitT)
	// fmt.Println(resultT)

}
