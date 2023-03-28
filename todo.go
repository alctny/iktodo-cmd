package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var (
	DONE_STATUS   = true
	UNDONE_STATUS = false

	printLayout = map[bool]string{
		DONE_STATUS:   doneLayout,
		UNDONE_STATUS: undoneLayout,
	}
)

const (
	undoneLayout = "\033[1;95;1m%02d %s\033[0m\n" // fg black bg white
	doneLayout   = "\033[1;92;9m%02d %s\033[0m\n" // fg white bg black
)

type Todo []Task

func (td Todo) Show() {
	for i, t := range td {
		fmt.Printf(printLayout[t.Done], i, t.Name)
	}
}

func Tasks() (Todo, error) {
	buf, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}
	var td Todo
	err = json.Unmarshal(buf, &td)
	return td, err
}

func Flush(ts Todo) error {
	buf, err := json.Marshal(ts)
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, buf, os.ModePerm)
}
