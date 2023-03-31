package cmd

import (
	"strings"

	"github.com/alctny/todo/dao"
	"github.com/alctny/todo/task"
	"github.com/urfave/cli/v2"
)

const (
	listClassFlag  = "class"
	listTagFlag    = "tag"
	listDoneFlag   = "done"
	listUndoneFlag = "undone"
)

func newList() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls", "show"},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: listClassFlag},
			&cli.StringFlag{Name: listTagFlag},
			&cli.BoolFlag{Name: listDoneFlag},
			&cli.BoolFlag{Name: listUndoneFlag},
		},
		Action: listAction,
	}
}

func listAction(ctx *cli.Context) error {
	tasks, err := dao.TodoList()
	if err != nil {
		return err
	}

	classParameter := strings.TrimSpace(ctx.String(listClassFlag))
	tagParameter := strings.TrimSpace(ctx.String(listTagFlag))

	flagFilter := map[string]func(t task.Task) bool{
		listClassFlag: func(t task.Task) bool {
			return t.Class == classParameter
		},

		listTagFlag: func(t task.Task) bool {
			for _, tag := range t.Tag {
				if tag == tagParameter {
					return true
				}
			}
			return false
		},

		listDoneFlag: func(t task.Task) bool {
			return t.Done == task.DONE_STATUS
		},

		listUndoneFlag: func(t task.Task) bool {
			return t.Done == task.UNDONE_STATUS
		},
	}

	flags := ctx.FlagNames()
	filters := []func(t task.Task) bool{}
	for _, v := range flags {
		filters = append(filters, flagFilter[v])
	}
	tasks = tasks.Filter(filters...)

	tasks.Show()
	return nil
}
