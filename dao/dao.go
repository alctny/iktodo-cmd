package dao

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alctny/todo/task"
)

const DATAFILE = ".iktodo"

var datafile = DATAFILE

func init() {
	if err := Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Init() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	datafile = filepath.Join(home, DATAFILE)

	if _, err := os.Stat(datafile); err != nil {
		if err = os.WriteFile(datafile, []byte("[]"), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func TodoList() (task.Tasks, error) {
	buf, err := os.ReadFile(datafile)
	if err != nil {
		return nil, err
	}
	var td task.Tasks
	err = json.Unmarshal(buf, &td)
	return td, err
}

func FlushAll(ts task.Tasks) error {
	buf, err := json.Marshal(ts)
	if err != nil {
		return err
	}
	return os.WriteFile(datafile, buf, os.ModePerm)
}
