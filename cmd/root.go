package cmd

import (
	"fmt"
	"os"

	"github.com/alctny/todo/dao"
	"github.com/urfave/cli/v2"
)

func Execute() {
	app := cli.App{
		Name:    "todo",
		Version: "1.0.0",
		Commands: []*cli.Command{
			NewAdd(),
			NewClear(),
			NewDone(),
			NewList(),
			NewModify(),
			NewSort(),
		},
		Action: defAction,
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func defAction(ctx *cli.Context) error {
	tasks, err := dao.TodoList()
	if err != nil {
		return err
	}
	tasks.Show()
	return nil
}
