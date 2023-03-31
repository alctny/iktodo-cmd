package cmd

import (
	"strings"

	"github.com/alctny/todo/dao"
	"github.com/alctny/todo/task"
	"github.com/urfave/cli/v2"
)

var addCmd = &cli.Command{
	Name:    "add",
	Aliases: []string{"touch", "new", "a"},
	Action:  addAction,
}

func addAction(ctx *cli.Context) error {
	tasks, err := dao.TodoList()
	if err != nil {
		return err
	}

	class := ""
	tag := []string{}
	name := ""
	args := strings.Split(strings.Join(ctx.Args().Slice(), " "), ":")
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

	tasks = append(tasks, task.Task{
		Name:  name,
		Done:  task.UNDONE_STATUS,
		Class: class,
		Tag:   tag,
	})
	tasks.Show()
	return dao.FlushAll(tasks)
}
