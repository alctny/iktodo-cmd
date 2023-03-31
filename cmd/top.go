package cmd

import (
	"fmt"
	"strconv"

	"github.com/alctny/todo/dao"
	"github.com/urfave/cli/v2"
)

var topCmd = &cli.Command{
	Action: topAction,
	Name:   "top",
}

func topAction(ctx *cli.Context) error {
	if ctx.NArg() < 1 {
		return fmt.Errorf("top command must set task id, but you dont")
	}
	id, err := strconv.ParseUint(ctx.Args().Get(0), 10, 64)
	if err != nil {
		return err
	}
	ts, err := dao.TodoList()
	if err != nil {
		return err
	}
	nt := len(ts)
	if int(id) >= nt {
		return nil
	}
	ts = append(ts[:id], ts[id+1:]...)
	return dao.FlushAll(ts)
}
