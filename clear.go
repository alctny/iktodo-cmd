package main

import "github.com/urfave/cli/v2"

func NewClear() *cli.Command {
	return &cli.Command{
		Name:    "clear",
		Aliases: []string{"c", "cls"},
		Action:  clearAction,
	}

}

func clearAction(ctx *cli.Context) error {
	tasks, err := Tasks()
	if err != nil {
		return err
	}
	undone := Todo{}
	for _, t := range tasks {
		if !t.Done {
			undone = append(undone, t)
		}
	}
	undone.Show()
	return Flush(undone)
}
