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

// Show show all task
func (ts Tasks) Show() {
	for i, t := range ts {
		// eg: 00 [bug] <go,rust,shell> 遊歷的目的是收集五聖器
		fmt.Printf(printLayout[t.Done], i, t.Class, strings.Join(t.Tag, ","), t.Name)
	}
}

// List just print, no color, no strikethrough
func (ts Tasks) List() {
	for i, t := range ts {
		fmt.Printf("%02d %s %s %s\n", i, t.Class, strings.Join(t.Tag, ","), t.Name)
	}
}

// Show show task which task satisfy fs
func (ts Tasks) ShowSome(fs ...func(t Task) bool) {

}

// Filter
func (ts Tasks) Filter(fs ...func(Task) bool) Tasks {
	result := Tasks{}
	for _, t := range ts {
		for _, f := range fs {
			if f != nil && !f(t) {
				goto next
			}
		}
		result = append(result, t)
	next:
	}
	return result
}
