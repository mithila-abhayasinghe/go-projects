package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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

func parseArgs(line string) []string {

	// update 1 "some task"

	// under assumption trimspace was called before on line

	var argBuffer []string

	cmdIdx := 0

	fmt.Printf("[debug] %s %d\n", line, len(line))

	// meh good enough eh i expect nothing after string with quotes
	for i, c := range line {
		// fmt.Printf("%c\n", character)
		if c == '"' || c == '\'' || c == '`' {
			desc, err := strconv.QuotedPrefix(line[i:])
			if err != nil {
				log.Fatal(err)
			}

			unquoted, err := strconv.Unquote(desc)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("desc of the task %s\n", unquoted)
			argBuffer = append(argBuffer, unquoted)
			break

		}
		if c == ' ' {
			argBuffer = append(argBuffer, line[cmdIdx:i])
			cmdIdx = i + 1
			continue
		}
	}

	fmt.Println("These are final commands on the command buffer")
	for i, val := range argBuffer {
		fmt.Printf("%d %s %d\n", i, val, len(val))
	}
	fmt.Println(len(argBuffer))

	// sampleCommand := `update 1 "some task"`

	// for i := 0; i < len(sampleCommand); i++ {
	// 	fmt.Printf("%c\n", sampleCommand[i])
	// }

	// s, err := strconv.QuotedPrefix(sampleCommand)
	// fmt.Printf("%q, %v\n", s, err)

	return args
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
		_ = parseArgs(line)
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
