package storage

import (
	"encoding/json"
	"errors"
	"go-cli-todos-app/models"
	"os"
)

const fileName = "tasks.json"

// SaveTasks saves the task list to a JSON file.
func SaveTasks(tasks []models.Task) error {
	file, err := os.Create(fileName)
	if err != nil {
		return errors.New("could not create tasks.json file: " + err.Error())
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		return errors.New("could not encode tasks to JSON: " + err.Error())
	}
	return nil
}

// LoadTasks loads the task list from a JSON file.
func LoadTasks() ([]models.Task, error) {
	file, err := os.Open(fileName)
	if err != nil {
		// Return an empty slice if the file doesn't exist (first run)
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, errors.New("could not open tasks.json file: " + err.Error())
	}
	defer file.Close()

	var tasks []models.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, errors.New("could not decode tasks from JSON: " + err.Error())
	}
	return tasks, nil
}
