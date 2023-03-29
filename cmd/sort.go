package cmd

import (
	"github.com/alctny/todo/dao"
	"github.com/alctny/todo/task"
	"github.com/urfave/cli/v2"
)

func NewSort() *cli.Command {
	return &cli.Command{
		Name:   "sort",
		Action: sortAction,
	}

}

func sortAction(ctx *cli.Context) error {
	tasks, err := dao.TodoList()
	if err != nil {
		return err
	}
	undone := task.Tasks{}
	done := task.Tasks{}
	for _, tsk := range tasks {
		switch tsk.Done {
		case task.DONE_STATUS:
			done = append(done, tsk)
		case task.UNDONE_STATUS:
			undone = append(undone, tsk)
		}
	}
	undone = append(undone, done...)
	undone.Show()
	return dao.FlushAll(undone)
}
