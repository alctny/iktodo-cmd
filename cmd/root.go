package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() {
	app := cli.App{
		Name:    "todo",
		Version: "1.0.0",
		Action:  defAction,
		Commands: []*cli.Command{
			newAdd(),
			newClear(),
			newDone(),
			newList(),
			newModify(),
			newSort(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func defAction(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return listAction(ctx)
	}
	return addAction(ctx)
}
