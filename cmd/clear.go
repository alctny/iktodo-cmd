package cmd

import (
	"github.com/alctny/todo/dao"
	"github.com/alctny/todo/task"
	"github.com/urfave/cli/v2"
)

func NewClear() *cli.Command {
	return &cli.Command{
		Name:    "clear",
		Aliases: []string{"c", "cls"},
		Action:  clearAction,
	}

}

func clearAction(ctx *cli.Context) error {
	tasks, err := dao.TodoList()
	if err != nil {
		return err
	}
	undone := task.Tasks{}
	for _, t := range tasks {
		if !t.Done {
			undone = append(undone, t)
		}
	}
	undone.Show()
	return dao.FlushAll(undone)
}
