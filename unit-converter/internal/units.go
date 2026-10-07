package internal

type UnitName string

const (
	// Lenght Units
	Millimeter UnitName = "millimeter"
	Centimeter UnitName = "centimeter"
	Meter      UnitName = "meter"
	Kilometer  UnitName = "kilometer"
	Inch       UnitName = "inch"
	Foot       UnitName = "foot"
	Yard       UnitName = "yard"
	Mile       UnitName = "mile"

	// Weight Units
	Milligram UnitName = "milligram"
	Gram      UnitName = "gram"
	Kilogram  UnitName = "kilogram"
	Ounce     UnitName = "ounce"
	Pound     UnitName = "pound"
)

type Unit struct {
	Name        string
	Symbol      string
	Dimension   string
	ToBaseRatio float64
}

var SupportedUnits = map[UnitName]Unit{

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
