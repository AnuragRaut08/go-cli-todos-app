package storage

import (
	"encoding/json"
	"errors"
	"os"
)

// Define Task struct in storage package
type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Priority string `json:"priority"`
	Done     bool   `json:"done"`
}

// File to store tasks
const fileName = "tasks.json"

// SaveTasks saves tasks to a file
func SaveTasks(tasks []Task) error {
	file, err := os.Create(fileName)
	if err != nil {
		return errors.New("failed to create file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		return errors.New("failed to encode tasks")
	}
	return nil
}

// LoadTasks loads tasks from a file
func LoadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // Return empty list if file doesn't exist
		}
		return nil, errors.New("failed to open file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, errors.New("failed to decode tasks")
	}

	return tasks, nil
}
