package internal

import (
	"fmt"
	"time"
)

type Expense struct {
	Id          int
	Description string
	Amount      float64
	Date        time.Time
}

func NewExpense(id int, desc string, amount float64) Expense {

	return Expense{
		Id:          id,
		Description: desc,
		Amount:      amount,
		Date:        time.Now(),
	}

}

func (e Expense) String() string {

	return fmt.Sprintf("# %-4d %-12s %-14s Rs.%-g",
		e.Id,
		e.Date.Format("2006-01-02"),
		e.Description,
		e.Amount,
	)

}
