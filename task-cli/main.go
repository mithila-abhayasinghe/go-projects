package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Get this fucking warnings out of here
var (
	_ = json.Marshal
	_ = os.Stdout
	_ = flag.Bool
	_ = fmt.Print
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

func newTask(id *int, desc string, status TaskStatus) Task {
	if *id == 0 {
		*id++
	}

	newTask := Task{
		Id:          *id,
		Description: desc,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	*id++
	return newTask

}

// This approach can be optimized not idiomatic in go but works
func (t Task) statusString() string {

	statusStrings := map[int]string{
		0: "TODO",
		1: "IN-PROGRESS",
		2: "DONE",
	}

	return statusStrings[int(t.Status)]

}

func (t Task) String() string {
	ts := t.UpdatedAt.Format("Jan 02 15:04")
	if t.UpdatedAt.IsZero() {
		ts = t.CreatedAt.Format("Jan 02 15:04")
	}

	desc := truncateTaskDesc(t.Description, 28)

	// %-4d  -> left-align ID in 4 spaces
	// %-11s -> left-align status string in 11 spaces (length of "IN-PROGRESS")
	// %-20s -> left-align description in 20 spaces
	return fmt.Sprintf("[#%-3d]  [%-11s]  %-28s  (%s)",
		t.Id,
		t.statusString(),
		desc,
		ts,
	)
}

func truncateTaskDesc(s string, maxLen int) string {
	if len(s) > maxLen {
		// remove truncate string upto maxLen but minus 3 place taken by dots
		return s[:maxLen-3] + "..."
	}
	return s
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	// Main repl loop
	for {
		fmt.Print("task-cli ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		fmt.Printf("[debug] %s %d\n", line, len(line))
	}

	// Create an empty task bucket
	// var taskBucket []Task = make([]Task, 0)
	// noOfTasks := len(taskBucket)

	// Add a new task to the bucket
	// taskBucket = append(taskBucket, newTask(&noOfTasks, *newTaskPtr, StateTodo))

	// List all the tasks
	// func(listCommand string) {
	// 	switch listCommand {
	//
	// 	case "done":
	//
	// 	case "todo":
	//
	// 	case "in-progress":
	//
	// 	defaultf
	// 		for task := range taskBucket {
	// 			fmt.Println(taskBucket[task])
	// 		}
	// 	}
	//
	// }()

	// flagTask := newTask(&noOfTasks, *newTaskPtr, StateTodo)
	// InProgressTask := newTask(&noOfTasks, *InProgressPtr, StateInProgress)
	// DoneTask := newTask(&noOfTasks, *DoneTaskPtr, StateDone)

	// fmt.Println(flagTask)
	// fmt.Println(InProgressTask)
	// fmt.Println(DoneTask)

	// flagTaskJson, _ := json.Marshal(flagTask)
	// fmt.Println(string(flagTaskJson))

	// task1 := newTask("Be Good Go Programmer")
	// task1json, _ := json.Marshal(task1)
	// fmt.Println(string(task1json))

	// task2 := newTask("Do More Projects")
	// task2json, _ := json.Marshal(task2)
	// fmt.Println(string(task2json))

	// task3 := newTask("Learn Even Further")
	// task3json, _ := json.Marshal(task3)
	// fmt.Println(string(task3json))

	// var task Task

	// filepath := "tasklist.json"
	//
	// file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE, 0777)
	// // file, err := os.Open(filepath)
	// if err != nil {
	// 	fmt.Println("File Reading error", err)
	// 	return
	// }
	//
	// defer file.Close()

	// Write stuff
	// encoder := json.NewEncoder(file)
	// err = encoder.Encode(task2json)

	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&task)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(task)

}
