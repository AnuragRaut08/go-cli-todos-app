package cmd

import (
	"errors"
	"fmt"
	"go-cli-todos-app/storage" // Import storage package
)

// Global task list
var Tasks []storage.Task // Use Task from storage package

// Init initializes the task list from storage
func Init() error {
	var err error
	Tasks, err = storage.LoadTasks()
	if err != nil {
		return errors.New("failed to load tasks")
	}
	return nil
}

// AddTask adds a new task
func AddTask(name, priority string) {
	newTask := storage.Task{
		ID:       len(Tasks) + 1,
		Name:     name,
		Priority: priority,
		Done:     false,
	}
	Tasks = append(Tasks, newTask)
}

// CompleteTask marks a task as complete
func CompleteTask(id int) bool {
	for i := range Tasks {
		if Tasks[i].ID == id {
			Tasks[i].Done = true
			return true
		}
	}
	return false
}

// DeleteTask removes a task by ID
func DeleteTask(id int) bool {
	for i, task := range Tasks {
		if task.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return true
		}
	}
	return false
}

// ListTasks displays all tasks
func ListTasks() {
	if len(Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("\nTask List:")
	for _, task := range Tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("[%d] %s (Priority: %s) %s\n", task.ID, task.Name, task.Priority, status)
	}
}
