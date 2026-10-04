package internal

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var fileName string = "expenses.csv"

type Store []Expense

var MemStore Store = csvLoad()

func AddExpense(desc string, amt float64) (Expense, error) {
	if strings.TrimSpace(desc) == "" {
		return Expense{}, errors.New("description cannot be empty")
	}
	if amt <= 0 {
		return Expense{}, errors.New("amount must be greater than zero")
	}

	exp := NewExpense(nextId(), desc, amt)
	MemStore = append(MemStore, exp)
	saveCSV()
	return exp, nil
}

func UpdateExpense(idx int, desc string, amt float64) (Expense, error) {
	if strings.TrimSpace(desc) == "" {
		return Expense{}, errors.New("description cannot be empty")
	}
	if amt <= 0 {
		return Expense{}, errors.New("amount must be greater than zero")
	}
	for i := range MemStore {
		if MemStore[i].Id == idx {
			MemStore[i].Description = desc
			MemStore[i].Amount = amt
			saveCSV()
			return MemStore[i], nil
		}
	}
	return Expense{}, errors.New("id does not exist")
}

func DeleteExpense(didx int) (Expense, error) {
	if didx <= 0 {
		return Expense{}, errors.New("id must be valid (1 <= id <= len(n))")
	}
	for i := range MemStore {
		if MemStore[i].Id == didx {
			deletedExpense := MemStore[i]
			MemStore = slices.Delete(MemStore, i, i+1)
			saveCSV()
			return deletedExpense, nil

		}
	}
	return Expense{}, errors.New("id does not exist")

}

func ListExpense() {
	for _, exp := range MemStore {
		fmt.Println(exp)
	}
}

func SummaryExpense(month int) error {
	// this does the inverse if the input fall out of the range
	if month < 1 || month > 12 {
		return fmt.Errorf("invalid month %d: must be between 1 and 12", month)
	}
	for i := range MemStore {
		if MemStore[i].Date.Month() == time.Month(month) {
			fmt.Println(MemStore[i])
		}
	}
	return nil

}

func nextId() int {
	currentId := 0
	for _, exp := range MemStore {
		if exp.Id > currentId {
			currentId = exp.Id
		}
	}
	// remember to return + 1
	return currentId + 1
}

func csvConverter() [][]string {
	csvRecords := make([][]string, 0)
	// Add header row first
	csvRecords = append(csvRecords, []string{"Id", "Description", "Amount", "Date"})

	for _, r := range MemStore {
		// // id;description;record;date
		// var builder strings.Builder
		//
		// idString := strconv.Itoa(r.Id)
		// builder.WriteString(idString)
		// builder.WriteString(";")
		//
		// builder.WriteString(r.Description)
		// builder.WriteString(";")
		//
		// // 'g': compact format (switches between scientific and standard decimal notation)
		// // -1 : shortest representation that still parses back to the exact same float64
		// // 64 : bit size (float64)
		// amountStr := strconv.FormatFloat(r.Amount, 'g', -1, 64)
		// builder.WriteString(amountStr)
		// builder.WriteString(";")
		//
		// timeStr := r.Date.Format(time.RFC3339)
		// builder.WriteString(timeStr)
		//
		// rowString := builder.String()
		// row := strings.Split(rowString, ";")

		// learn this one
		row := []string{
			strconv.Itoa(r.Id),
			r.Description,
			strconv.FormatFloat(r.Amount, 'g', -1, 64),
			r.Date.Format(time.RFC3339),
		}

		csvRecords = append(csvRecords, row)

	}
	return csvRecords
}

func csvParser() [][]string {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return make([][]string, 0)
		}
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	csvRecords := make([][]string, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		csvRecords = append(csvRecords, record)
	}
	return csvRecords
}

func saveCSV() {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(csvConverter())
	if err != nil {
		panic(err)
	}
}

func csvLoad() Store {
	newMemStore := make(Store, 0)
	records := csvParser()

	// If empty or only contains the header, return empty store
	if len(records) <= 1 {
		return newMemStore
	}

	// Start from index 1 to skip the header
	for _, record := range records[1:] {

		// Parse data
		idInt, _ := strconv.Atoi(record[0])
		amountFloat64, _ := strconv.ParseFloat(record[2], 64)
		dateTimeTime, _ := time.Parse(time.RFC3339, record[3])

		exp := Expense{
			Id:          idInt,
			Description: record[1],
			Amount:      amountFloat64,
			Date:        dateTimeTime,
		}

		newMemStore = append(newMemStore, exp)
	}
	return newMemStore
}
