package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Done  bool     `json:"done"`
	Name  string   `json:"name"`
	Class string   `json:"class"`
	Tag   []string `json:"tag"`
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
	undoneLayout = "\033[1;95;1m%02d [%s] <%s> %s\033[0m\n" // fg black bg white
	doneLayout   = "\033[1;92;9m%02d [%s] <%s> %s\033[0m\n" // fg white bg black
)

type Todo []Task

func (td Todo) Show() {
	for i, t := range td {
		// 00 [bug] <go,rust,shell> 遊歷的目的是收集五聖器
		fmt.Printf(printLayout[t.Done], i, t.Class, strings.Join(t.Tag, ","), t.Name)
	}
}

func (td Todo) Filter(fs ...func(Task) bool) Todo {
	result := Todo{}
	for _, filter := range fs {
		res := Todo{}
		for _, task := range td {
			if filter(task) {
				res = append(res, task)
			}
		}
		result = append(result, res...)
	}
	return result
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
