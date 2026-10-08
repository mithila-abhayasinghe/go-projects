package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	unit "unit-converter/internal"
)

var (
	_ = fmt.Println
	_ = unit.Convert
)

func main() {

	// Examples
	// inputValue := 1
	// result, _ := unit.Convert(float64(inputValue), unit.Kilogram, unit.Gram)
	// fmt.Println(result.ToValue)
	//
	// // something i could not do before Convert hand temperature
	// inputValue = 100
	// result, _ = unit.Convert(float64(inputValue), unit.Celsius, unit.Kelvin)
	// fmt.Println(result.ToValue)

	http.HandleFunc("/form/", formHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

type PageData struct {
	Success bool
	Result  float64
}

var tmpl = template.Must(template.ParseFiles("views/form.html"))

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Debug] - Method ", r.Method)

	if r.Method == http.MethodGet {
		tmpl.Execute(w, PageData{Success: false})
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if r.PostForm.Get("action") == "reset" {
			tmpl.Execute(w, PageData{Success: false})
			return
		}

		val, _ := strconv.ParseFloat(r.FormValue("fromValue"), 64)
		fromUnit := unit.GetUnitName(r.FormValue("from"))
		toUnit := unit.GetUnitName(r.FormValue("to"))

		conversionEvent := unit.ConversionEvent{
			FromValue: val,
			FromUnit:  fromUnit,
			ToUnit:    toUnit,
		}

		result, _ := unit.ConverstionEventHandler(conversionEvent)
		fmt.Println(result)

		tmpl.Execute(w, PageData{
			Success: true,
			Result:  result.ToValue,
		})
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
