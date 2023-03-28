package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var (
	dataFile = ".todo"
)

func init() {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		dataFile = filepath.Join(home, dataFile)
		_, err = os.Stat(dataFile)
		if os.IsNotExist(err) {
			err = os.WriteFile(dataFile, []byte("[]"), os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {
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
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
