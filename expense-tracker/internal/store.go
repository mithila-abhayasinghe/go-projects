package expense

import "fmt"

type Store []Expense

var MemStore Store = make(Store, 0)

func AddExpense() {
	fmt.Println("Add Called")
}

func UpdateExpense() {
	fmt.Println("Update Called")
}

func DeleteExpense() {
	fmt.Println("Delete Called")
}

func ListExpense() {
	fmt.Println("List Called")
}

func SummaryExpense() {
	fmt.Println("Summary Called")
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
