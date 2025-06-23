package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Task - structure fo storing tasks information
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

func loadTasks(path string, tasks []Task) {
	source, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	fmt.Println("Open successful!")
	defer source.Close()

	dec := json.NewDecoder(source)
	err = dec.Decode(&tasks)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
}
func saveTasks(path string, tasks []Task) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Something went wrng while accessing file:", err)
		return
	}
	defer file.Close()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error while encoding:", err)
		return
	}
	file.Write(data)
}
func inputTask() (int, string, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Task ID:")
	input1, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input1 = strings.TrimSpace(input1)

	id, err := strconv.Atoi(input1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title:")
	title, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	title = strings.TrimSpace(title)

	fmt.Printf("Status:")
	input3, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input3 = strings.TrimSpace(input3)

	isdone := strings.ToLower(input3) == "true" || input3 == "1"

	return id, title, isdone

}
func main() {
	tasks := []Task{}

	fmt.Println("Starting the programm: opening the json-file...")
	loadTasks("./data/input.json", tasks)

	fmt.Println("Current list of tasks: ", len(tasks), "items")
	for _, i := range tasks {
		fmt.Printf("[%d] %s: %v\n", i.ID, i.Title, i.IsDone)
	}
	fmt.Println("Adding task now. Enter the following info - one field at a time:")
	id, title, isdone := inputTask()
	tasks = append(tasks, Task{ID: id, Title: title, IsDone: isdone})
	saveTasks("./data/input.json", tasks)
	fmt.Println("Adding done! Restart the program to see list of tasks.")

}
