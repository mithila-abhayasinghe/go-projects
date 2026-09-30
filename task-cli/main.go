package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Get this warnings out of here
// insert robert pattinson odyssey meme
var (
	_ = json.Marshal
	_ = os.Stdout
	_ = flag.Bool
	_ = fmt.Print
)

type TaskStatus int
type TaskBucket []Task

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

func newTask(id int, desc string) Task {
	newTask := Task{
		Id:          id,
		Description: desc,
		Status:      StateTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
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

	desc := truncateTaskDesc(t.Description, 20)

	// %-4d  -> left-align ID in 4 spaces
	// %-11s -> left-align status string in 11 spaces (length of "IN-PROGRESS")
	// %-20s -> left-align description in 20 spaces
	return fmt.Sprintf("[#%-3d]  [%-11s]  %-20s  (%s)",
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

	// test case: update 1 "some task"

	// under assumption trimspace was called before on line

	var argBuffer []string

	cmdIdx := 0

	// fmt.Printf("[debug] %s %d\n", line, len(line))

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

			// fmt.Printf("desc of the task %s\n", unquoted)
			argBuffer = append(argBuffer, unquoted)
			break

		}
		if c == ' ' {
			if i > cmdIdx {
				argBuffer = append(argBuffer, line[cmdIdx:i])
			}
			cmdIdx = i + 1
			continue
		}
	}

	if cmdIdx < len(line) {
		remaining := line[cmdIdx:]
		if len(remaining) > 0 {
			argBuffer = append(argBuffer, remaining)
		}
	}

	// fmt.Println("These are final commands on the command buffer")
	// for i, val := range argBuffer {
	// 	fmt.Printf("%d %s %d\n", i, val, len(val))
	// }
	// fmt.Println(len(argBuffer))

	// sampleCommand := `update 1 "some task"`

	// for i := 0; i < len(sampleCommand); i++ {
	// 	fmt.Printf("%c\n", sampleCommand[i])
	// }

	// s, err := strconv.QuotedPrefix(sampleCommand)
	// fmt.Printf("%q, %v\n", s, err)

	return argBuffer
}

func addTask(tasks *TaskBucket, desc string) *TaskBucket {

	// Don't rely on lenght of the slice for id's
	nextId := func(tasks TaskBucket) int {
		maxId := 0
		for _, t := range tasks {
			if t.Id > maxId {
				maxId = t.Id
			}
		}
		return maxId + 1

	}(*tasks)

	createdNewTask := newTask(nextId, desc)

	*tasks = append(*tasks, createdNewTask)

	return tasks

}

func printTaskList(tasks TaskBucket, state *TaskStatus) {
	if state != nil {
		for _, task := range tasks {
			if *state == task.Status {
				fmt.Println(task)
			}
		}

	} else {
		for _, task := range tasks {
			fmt.Println(task)
		}
	}
}

func changeTaskStatus(tasks *TaskBucket, id *int, state *TaskStatus) (*Task, bool) {

	if id == nil {
		fmt.Println("Id is missing")
		panic("id is missing")
	}

	if state == nil {
		fmt.Println("Status is missing")
		panic("missing state")
	}

	if *id <= 0 {
		fmt.Println("invalid id")
	}

	for i, t := range *tasks {
		if t.Id == *id {
			(*tasks)[i].Status = *state
			return &(*tasks)[i], true
		}
	}
	return nil, false

}

func updateTask(tasks *TaskBucket, id *string, desc *string) (*Task, bool) {

	idx, _ := strconv.Atoi(*id)

	if desc == nil {
		fmt.Println("No input description")
		panic("no description input")
	}

	if idx <= 0 {
		fmt.Println("invalid id")
	}

	for i, t := range *tasks {
		if t.Id == idx {
			(*tasks)[i].Description = *desc
			return &(*tasks)[i], true
		}
	}
	return nil, false

}

func deleteTask(tasks TaskBucket, idStrPtr *string) (TaskBucket, Task, bool) {

	id, _ := strconv.Atoi(*idStrPtr)
	idx := 0

	for i, t := range tasks {
		if t.Id == id {
			idx = i
		}
	}

	deletedTask := tasks[idx]
	tasks = slices.Delete(tasks, idx, idx+1)
	return tasks, deletedTask, true

}

func continuousPersist(tasks TaskBucket) {
	filepath := "tasklist.json"

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("file reading error ", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(tasks)

	if err != nil {
		log.Fatal("file encoding errr ", err)
	}

	// research write-to-temp then Rename pattern

}

func loadPersistedTasks() TaskBucket {
	var taskBucket TaskBucket
	filepath := "tasklist.json"

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("file reading error ", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&taskBucket)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return TaskBucket{}
		}
		log.Fatal("json decoding err ", err)
	}

	return taskBucket
}

func main() {

	// Create an empty task bucket
	var taskBucket TaskBucket = make(TaskBucket, 0)

	// load once
	taskBucket = loadPersistedTasks()

	reader := bufio.NewReader(os.Stdin)

	// Main repl loop
	for {
		fmt.Print("task-cli ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		// fmt.Println("[debug] before line: ", line)
		args := parseArgs(line)
		// fmt.Println("[debug] arguments : ", args)

		rootCommand, rootOk := func() (string, bool) {
			var rootCommand string
			if len(args) > 0 {
				rootCommand = args[0]
				return rootCommand, true
			}
			return rootCommand, false
		}()

		if !rootOk {
			fmt.Println("Please Enter a Command")
			// print help page maybe
		}

		// subCommand could be an number index or filter for list
		secondPosition, _ := func() (string, bool) {
			var secondPosition string
			if len(args) > 1 {
				secondPosition = args[1]
				return secondPosition, true

			}
			return secondPosition, false
		}()

		thirdPosition, _ := func() (string, bool) {
			var thirdPosition string
			if len(args) > 2 {
				thirdPosition = args[2]
				return thirdPosition, true
			}
			return thirdPosition, false
		}()

		switch rootCommand {

		case "add":
			if secondPosition != "" {
				addTask(&taskBucket, secondPosition)
			} else {
				fmt.Println("No Task Has been given")
			}
			continuousPersist(taskBucket)

		case "update":
			task, ok := updateTask(&taskBucket, &secondPosition, &thirdPosition)
			if !ok {
				fmt.Println("Failed to update, Task Do not exist")
			} else {
				fmt.Println("Task Updated successfully")
				fmt.Println(task)
			}
			continuousPersist(taskBucket)

		case "delete":
			var task Task
			var ok bool
			taskBucket, task, ok = deleteTask(taskBucket, &secondPosition)
			if ok {
				fmt.Println("Task Deleted successfully")
				fmt.Println(task)
			}
			continuousPersist(taskBucket)

		case "mark-todo":
			idx, _ := strconv.Atoi(secondPosition)
			status := StateTodo
			task, ok := changeTaskStatus(&taskBucket, &idx, &status)
			if !ok {
				fmt.Println("Failed to change state, Task Do not exist")
			} else {
				fmt.Println("Status changed successfully for task")
				fmt.Println(task)
			}
			continuousPersist(taskBucket)

		case "mark-in-progress":
			idx, _ := strconv.Atoi(secondPosition)
			status := StateInProgress
			task, ok := changeTaskStatus(&taskBucket, &idx, &status)
			if !ok {
				fmt.Println("Failed to change state, Task Do not exist")
			} else {
				fmt.Println("Status changed successfully for task")
				fmt.Println(task)
			}
			continuousPersist(taskBucket)

		case "mark-done":
			idx, _ := strconv.Atoi(secondPosition)
			status := StateDone
			task, ok := changeTaskStatus(&taskBucket, &idx, &status)
			if !ok {
				fmt.Println("Failed to change state, Task Do not exist")
			} else {
				fmt.Println("Status changed successfully for task")
				fmt.Println(task)
			}
			continuousPersist(taskBucket)

		case "list", "ls":

			switch secondPosition {

			case "done":
				status := StateDone
				printTaskList(taskBucket, &status)
			case "todo":
				status := StateTodo
				printTaskList(taskBucket, &status)
			case "in-progress":
				status := StateInProgress
				printTaskList(taskBucket, &status)
			default:
				printTaskList(taskBucket, nil)
			}

		case "clear":
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()

		case "exit", "e":
			os.Exit(0)
		}

	}

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
