package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
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

	newTaskPtr := flag.String("add", "empty task", "Add a new task")

	// TODO update/delete find the function for format flag input input

	inProgressPtr := flag.String("mark-in-progres", 1, "To mark a Task to in progress")

	listPtr := flag.String("list", "all", "List all the taks we have")

	// task1 := newTask("Be Good Go Programmer")
	// task1json, _ := json.Marshal(task1)
	// fmt.Println(string(task1json))

	task2 := newTask("Do More Projects")
	task2json, _ := json.Marshal(task2)
	fmt.Println(string(task2json))

	task3 := newTask("Learn Even Further")
	task3json, _ := json.Marshal(task3)
	fmt.Println(string(task3json))

	// var task Task

	filepath := "tasklist.json"

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE, 0777)
	// file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("File Reading error", err)
		return
	}

	defer file.Close()

	// Write stuff
	encoder := json.NewEncoder(file)
	err = encoder.Encode(task2json)

	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&task)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(task)

}
