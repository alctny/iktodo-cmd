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

	prefix = map[bool]rune{
		DONE_STATUS:   'V',
		UNDONE_STATUS: 'X',
	}
)

func (t Task) String() string {
	return fmt.Sprintf("[%c] %s", prefix[t.Done], t.Name)
}

type Todo []Task

func (td Todo) Show() {
	for i, t := range td {
		fmt.Printf("%02d %s\n", i, t)
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
