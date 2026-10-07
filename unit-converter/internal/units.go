package internal

type Unit struct {
	Name        string
	Symbol      string
	Dimension   string
	ToBaseRatio float64
}

var SupportedUnits = map[string]Unit{
	// Length Units (Base Unit: Meter)
	"millimeter": {Name: "Millimeter", Symbol: "mm", Dimension: "length", ToBaseRatio: 0.001},
	"centimeter": {Name: "Centimeter", Symbol: "cm", Dimension: "length", ToBaseRatio: 0.01},
	"meter":      {Name: "Meter", Symbol: "m", Dimension: "length", ToBaseRatio: 1.0},
	"kilometer":  {Name: "Kilometer", Symbol: "km", Dimension: "length", ToBaseRatio: 1000.0},
	"inch":       {Name: "Inch", Symbol: "in", Dimension: "length", ToBaseRatio: 0.0254},
	"foot":       {Name: "Foot", Symbol: "ft", Dimension: "length", ToBaseRatio: 0.3048},
	"yard":       {Name: "Yard", Symbol: "yd", Dimension: "length", ToBaseRatio: 0.9144},
	"mile":       {Name: "Mile", Symbol: "mi", Dimension: "length", ToBaseRatio: 1609.344},

	// Weight Units (Base Unit: Kilogram)
	"milligram": {Name: "Milligram", Symbol: "mg", Dimension: "weight", ToBaseRatio: 0.000001},
	"gram":      {Name: "Gram", Symbol: "g", Dimension: "weight", ToBaseRatio: 0.001},
	"kilogram":  {Name: "Kilogram", Symbol: "kg", Dimension: "weight", ToBaseRatio: 1.0},
	"ounce":     {Name: "Ounce", Symbol: "oz", Dimension: "weight", ToBaseRatio: 0.028349523125},
	"pound":     {Name: "Pound", Symbol: "lb", Dimension: "weight", ToBaseRatio: 0.45359237},
}
