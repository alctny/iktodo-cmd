package main

import "github.com/urfave/cli/v2"

func NewSort() *cli.Command {
	return &cli.Command{
		Name:   "sort",
		Action: sort,
	}

}

func sort(ctx *cli.Context) error {
	tasks, err := Tasks()
	if err != nil {
		return err
	}
	undone := Todo{}
	done := Todo{}
	for _, task := range tasks {
		switch task.Done {
		case DONE_STATUS:
			done = append(done, task)
		case UNDONE_STATUS:
			undone = append(undone, task)
		}
	}
	undone = append(undone, done...)
	return Flush(undone)
}
