package main

import (
	"strconv"

	"github.com/urfave/cli/v2"
)

func NewDone() *cli.Command {
	return &cli.Command{
		Name:   "done",
		Action: done,
	}

}

func done(ctx *cli.Context) error {
	id_i64, err := strconv.ParseInt(ctx.Args().First(), 10, 64)
	if err != nil {
		return err
	}
	id := int(id_i64)
	tasks, err := Tasks()
	if err != nil {
		return err
	}
	if id >= len(tasks) {
		return nil
	}
	tasks[id].Done = !tasks[id].Done
	tasks.Show()
	return Flush(tasks)
}
