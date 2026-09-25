package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"time"
)

type TaskStatus int

const (
	StateTodo TaskStatus = iota
	StateInProgress
	StateDone
)

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func newTask(desc string) *Task {
	newTask := Task{
		Id:          rand.IntN(100),
		Description: desc,
		Status:      StateTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return &newTask
}

func main() {

	// task1 := newTask("Be Good Go Programmer")
	// task1json, _ := json.Marshal(task1)
	// fmt.Println(string(task1json))

}
