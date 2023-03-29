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

	class := ""
	tag := []string{}
	name := ""
	args := strings.Split(ctx.Args().First(), ":")
	// class:tag1,tag2...:task
	switch len(args) {
	case 2:
		class, name = args[0], args[1]
	case 1:
		name = args[0]
	default:
		class = args[0]
		tag = strings.Split(args[1], ",")
		name = strings.Join(args[2:], " ")
	}

	tasks = append(tasks, Task{
		Name:  name,
		Done:  UNDONE_STATUS,
		Class: class,
		Tag:   tag,
	})
	tasks.Show()
	return Flush(tasks)
}
