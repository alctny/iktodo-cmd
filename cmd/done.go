package cmd

import (
	"strconv"

	"github.com/alctny/todo/dao"
	"github.com/urfave/cli/v2"
)

var doneCmd = &cli.Command{
	Name:    "done",
	Aliases: []string{"d", "finish"},
	Action:  doneAction,
}

func doneAction(ctx *cli.Context) error {
	ids := []int{}
	for _, v := range ctx.Args().Slice() {
		id_i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		ids = append(ids, int(id_i64))
	}

	tasks, err := dao.TodoList()
	if err != nil {
		return err
	}

	for _, id := range ids {
		if id >= len(tasks) {
			continue
		}
		tasks[id].Done = !tasks[id].Done
	}
	tasks.Show()
	return dao.FlushAll(tasks)
}
