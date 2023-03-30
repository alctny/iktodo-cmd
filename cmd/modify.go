package cmd

import (
	"strconv"
	"strings"

	"github.com/alctny/todo/dao"
	"github.com/urfave/cli/v2"
)

func newModify() *cli.Command {
	return &cli.Command{
		Name:    "rename",
		Aliases: []string{"mv"},
		Action:  modifyAction,
	}
}

func modifyAction(ctx *cli.Context) error {
	id_i64, err := strconv.ParseInt(ctx.Args().First(), 10, 64)
	if err != nil {
		return err
	}
	id := int(id_i64)
	tasks, err := dao.TodoList()
	if err != nil {
		return err
	}
	if id >= len(tasks) {
		return nil
	}
	newName := strings.Join(ctx.Args().Slice()[1:], " ")
	tasks[id].Name = newName
	tasks.Show()
	return dao.FlushAll(tasks)
}
