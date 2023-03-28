package main

import (
	"strings"

	"github.com/urfave/cli/v2"
)

func NewAdd() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"touch", "new", "a"},
		Action:  addAction,
	}

}

func addAction(ctx *cli.Context) error {
	tasks, err := Tasks()
	if err != nil {
		return err
	}
	taskName := strings.Join(ctx.Args().Slice(), " ")
	tasks = append(tasks, Task{
		Name: taskName,
		Done: UNDONE_STATUS,
	})
	tasks.Show()
	return Flush(tasks)
}
