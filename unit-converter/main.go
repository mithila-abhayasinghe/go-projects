package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	unit "unit-converter/internal"
)

func main() {
	http.HandleFunc("/converter/", formHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// UnitCategories maps each category to its supported unit keys
var UnitCategories = map[string][]string{
	"length":      {"mm", "cm", "m", "km", "inch", "ft", "yard", "miles"},
	"weight":      {"mg", "gm", "kg", "ounce", "lbs"},
	"temperature": {"C", "F", "K"},
}

// CategoryMeta holds display metadata for navigation and labels
type CategoryMeta struct {
	ID    string // "length", "weight", "temperature"
	Title string // "Length", "Weight", "Temperature"
}

var AvailableCategories = []CategoryMeta{
	{ID: "length", Title: "Length"},
	{ID: "weight", Title: "Weight"},
	{ID: "temperature", Title: "Temperature"},
}

type PageData struct {
	Success         bool
	Result          float64
	Categories      []CategoryMeta // For rendering the nav bar
	CurrentCategory string         // The active category ID
	AvailableUnits  []string       // Units strictly for the active category
	FromValue       string
	FromUnit        string
	ToUnit          string
}

var tmpl = template.Must(template.ParseFiles("views/converter.html"))

func formHandler(w http.ResponseWriter, r *http.Request) {

	// determine active category
	category := r.URL.Query().Get("category")
	if r.Method == http.MethodGet {
		if cat := r.FormValue("category"); cat != "" {
			category = cat
		}
	}

	if _, ok := UnitCategories[category]; !ok {
		category = "length"
	}

	units := UnitCategories[category]

	if r.Method == http.MethodGet {
		tmpl.Execute(w, PageData{
			Categories:      AvailableCategories,
			CurrentCategory: category,
			AvailableUnits:  units,
		})
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if r.PostForm.Get("action") == "reset" {
			tmpl.Execute(w, PageData{
				Categories:      AvailableCategories,
				CurrentCategory: category,
				AvailableUnits:  units,
			})
			return
		}

		rawVal := r.PostForm.Get("fromValue")
		fromKey := r.PostForm.Get("from")
		toKey := r.PostForm.Get("to")

		val, _ := strconv.ParseFloat(rawVal, 64)
		fromUnit := unit.GetUnitName(fromKey)
		toUnit := unit.GetUnitName(toKey)

		conversionEvent := unit.ConversionEvent{
			FromValue: val,
			FromUnit:  fromUnit,
			ToUnit:    toUnit,
		}

		result, err := unit.ConverstionEventHandler(conversionEvent)

		if err != nil {
			tmpl.Execute(w, PageData{
				Categories:      AvailableCategories,
				CurrentCategory: category,
				AvailableUnits:  units,
				FromValue:       rawVal,
				FromUnit:        fromKey,
				ToUnit:          toKey,
			})
			return
		}

		tmpl.Execute(w, PageData{
			Success:         true,
			Result:          result.ToValue,
			Categories:      AvailableCategories,
			CurrentCategory: category,
			AvailableUnits:  units,
			FromValue:       rawVal,
			FromUnit:        fromKey,
			ToUnit:          toKey,
		})
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
