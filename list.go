package main

import "github.com/urfave/cli/v2"

func NewList() *cli.Command {
	return &cli.Command{
		Name:    "ls",
		Aliases: []string{"list", "show"},
		Action:  listAction,
	}
}

func listAction(ctx *cli.Context) error {
	tasks, err := Tasks()
	if err != nil {
		return err
	}
	tasks.Show()
	return nil
}
