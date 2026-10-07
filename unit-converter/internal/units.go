package internal

type Unit string

const (
	// Lenght Units
	Millimeter Unit = "millimeter"
	Centimeter Unit = "centimeter"
	Meter      Unit = "meter"
	Kilometer  Unit = "kilometer"
	Inch       Unit = "inch"
	Foot       Unit = "foot"
	Yard       Unit = "yard"
	Mile       Unit = "mile"

	// Weight Units
	Milligram Unit = "milligram"
	Gram      Unit = "gram"
	Kilogram  Unit = "kilogram"
	Ounce     Unit = "ounce"
	Pound     Unit = "pound"

	// Temperature
	Kelvin     Unit = "kelvin"
	Celsius    Unit = "celsius"
	Fahrenheit Unit = "fahrenheit"
)

type UnitData struct {
	Name        string
	Symbol      string
	Dimension   string
	ToBaseRatio float64
}

var SupportedUnits = map[Unit]UnitData{

	// Length Units (Base Unit: Meter)
	Millimeter: {Name: "Millimeter", Symbol: "mm", Dimension: "length", ToBaseRatio: 0.001},
	Centimeter: {Name: "Centimeter", Symbol: "cm", Dimension: "length", ToBaseRatio: 0.01},
	Meter:      {Name: "Meter", Symbol: "m", Dimension: "length", ToBaseRatio: 1.0},
	Kilometer:  {Name: "Kilometer", Symbol: "km", Dimension: "length", ToBaseRatio: 1000.0},
	Inch:       {Name: "Inch", Symbol: "in", Dimension: "length", ToBaseRatio: 0.0254},
	Foot:       {Name: "Foot", Symbol: "ft", Dimension: "length", ToBaseRatio: 0.3048},
	Yard:       {Name: "Yard", Symbol: "yd", Dimension: "length", ToBaseRatio: 0.9144},
	Mile:       {Name: "Mile", Symbol: "mi", Dimension: "length", ToBaseRatio: 1609.344},

	// Weight Units (Base Unit: Kilogram)
	Milligram: {Name: "Milligram", Symbol: "mg", Dimension: "weight", ToBaseRatio: 0.000001},
	Gram:      {Name: "Gram", Symbol: "g", Dimension: "weight", ToBaseRatio: 0.001},
	Kilogram:  {Name: "Kilogram", Symbol: "kg", Dimension: "weight", ToBaseRatio: 1.0},
	Ounce:     {Name: "Ounce", Symbol: "oz", Dimension: "weight", ToBaseRatio: 0.028349523125},
	Pound:     {Name: "Pound", Symbol: "lb", Dimension: "weight", ToBaseRatio: 0.45359237},
}

func Get(name Unit) UnitData {
	return SupportedUnits[name]
}

// ----- Temperature Conversion -----

// type TempConversionResult struct {
// 	FromValue float64
// 	FromUnit  Unit
// 	ToValue   float64
// 	ToUnit    Unit
// }

func ConvertTemperature(val float64, from, to Unit) (ConversionResult, error) {

	// fromNormalized := tempStringNormalize(from)
	// toNormalized := tempStringNormalize(to)

	var result ConversionResult
	result.FromValue = val
	result.FromUnit = from
	result.ToUnit = to

	switch from {
	case Kelvin:
		switch to {
		case Celsius:
			formulaResult := val - 273.15
			result.ToValue = formulaResult
		case Fahrenheit:
			formulaResult := (val-273.15)*9.0/5.0 + 32
			result.ToValue = formulaResult
		}
	case Celsius:
		switch to {
		case Kelvin:
			formulaResult := val + 273.15
			result.ToValue = formulaResult
		case Fahrenheit:
			formulaResult := val*1.8 + 32
			result.ToValue = formulaResult
		}
	case Fahrenheit:
		switch to {
		case Kelvin:
			formulaResult := (val-32)*5.0/9.0 + 273.15
			result.ToValue = formulaResult
		case Celsius:
			formulaResult := (val - 32) * 5.0 / 9.0
			result.ToValue = formulaResult
		}

	}

	return result, nil
}

// func tempStringNormalize(str string) string {
// 	// Normalize first
// 	normalized := strings.ToLower(strings.TrimSpace(str))
//
// 	switch normalized {
// 	case "celsius", "c":
// 		return "C"
// 	case "fahrenheit", "f":
// 		return "F"
// 	case "kelvin", "k":
// 		return "K"
// 	default:
// 		return ""
// 	}
// }
