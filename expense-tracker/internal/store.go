package internal

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Store []Expense

var MemStore Store = make(Store, 0)

func AddExpense(desc string, amt float64) (Expense, error) {
	if strings.TrimSpace(desc) == "" {
		return Expense{}, errors.New("description cannot be empty")
	}
	if amt <= 0 {
		return Expense{}, errors.New("amount must be greater than zero")
	}

	exp := NewExpense(nextId(), desc, amt)
	MemStore = append(MemStore, exp)
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
			return MemStore[i], nil
		}
	}
	return Expense{}, errors.New("id does not exist")
}

func DeleteExpense(didx int) (Expense, error) {
	if didx <= 0 || didx > len(MemStore) {
		return Expense{}, errors.New("id must be valid (1 <= id <= len(n))")
	}
	for i := range MemStore {
		if MemStore[i].Id == didx {
			deletedExpense := MemStore[i]
			MemStore = slices.Delete(MemStore, i, i+1)
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

func csvPersist() {
	fmt.Println("CSV Written")
}

func csvLoad() {
	fmt.Println("CSV Loaded")
}
