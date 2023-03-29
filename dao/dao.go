package dao

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/alctny/todo/task"
)

var dataFile = ".todo"

func init() {
	if err := createDataFile(); err != nil {
		panic(err)
	}
}

func createDataFile() error {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		dataFile = filepath.Join(home, dataFile)
		_, err = os.Stat(dataFile)
		if os.IsNotExist(err) {
			err = os.WriteFile(dataFile, []byte("[]"), os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func TodoList() (task.Tasks, error) {
	buf, err := os.ReadFile(dataFile)
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
	return os.WriteFile(dataFile, buf, os.ModePerm)
}
