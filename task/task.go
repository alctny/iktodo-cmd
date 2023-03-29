package task

import (
	"fmt"
	"strings"
)

const (
	DONE_STATUS   = true
	UNDONE_STATUS = false

	undoneLayout = "\033[1;95;1m%02d \033[1;94;1m%s \033[1;93;1m%s \033[1;95;1m%s\033[0m\n"
	doneLayout   = "\033[1;92;9m%02d \033[1;94;1m%s \033[1;93;1m%s \033[1;92;9m%s\033[0m\n"
)

var (
	printLayout = map[bool]string{
		DONE_STATUS:   doneLayout,
		UNDONE_STATUS: undoneLayout,
	}
)

type Task struct {
	Done  bool     `json:"done"`
	Name  string   `json:"name"`
	Class string   `json:"class"`
	Tag   []string `json:"tag"`
}

type Tasks []Task

func (td Tasks) Show() {
	for i, t := range td {
		// eg: 00 [bug] <go,rust,shell> 遊歷的目的是收集五聖器
		fmt.Printf(printLayout[t.Done], i, t.Class, strings.Join(t.Tag, ","), t.Name)
	}
}

func (td Tasks) Filter(fs ...func(Task) bool) Tasks {
	result := Tasks{}
	for _, filter := range fs {
		res := Tasks{}
		for _, task := range td {
			if filter(task) {
				res = append(res, task)
			}
		}
		result = append(result, res...)
	}
	return result
}
