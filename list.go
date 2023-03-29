package main

import (
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	listClassFlag = "class"
	listTagFlag   = "tag"
)

func NewList() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls", "show"},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: listClassFlag},
			&cli.StringFlag{Name: listTagFlag},
		},
		Action: listAction,
	}
}

func listAction(ctx *cli.Context) error {
	tasks, err := Tasks()
	if err != nil {
		return err
	}

	if ctx.IsSet(listClassFlag) {
		class := strings.TrimSpace(ctx.String(listClassFlag))
		tasks = tasks.Filter(func(t Task) bool {
			return t.Class == class
		})
	}

	if ctx.IsSet(listTagFlag) {
		tagArg := strings.TrimSpace(ctx.String(listTagFlag))
		tasks = tasks.Filter(func(t Task) bool {
			for _, tag := range t.Tag {
				if tag == tagArg {
					return true
				}
			}
			return false
		})
	}

	tasks.Show()
	return nil
}
