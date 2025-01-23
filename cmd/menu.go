package cmd

import (
	"fmt"
	"go-cli-todos-app/models"
	"go-cli-todos-app/storage"

)

var Tasks []models.Task // Exported variable
var taskID int

// Initialize the task list by loading tasks from storage
func Init() {
	loadedTasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error initializing tasks:", err)
		return
	}
	Tasks = loadedTasks

	// Set taskID to avoid duplicate IDs
	for _, task := range Tasks {
		if task.ID > taskID {
			taskID = task.ID
		}
	}
}

// Add a new task
func AddTask(title string, priority string) {
	if priority != "High" && priority != "Medium" && priority != "Low" {
		fmt.Println("Invalid priority! Use High, Medium, or Low.")
		return
	}

	taskID++
	task := models.Task{
		ID:        taskID,
		Title:     title,
		Completed: false,
		Priority:  priority,
	}
	Tasks = append(Tasks, task)
	fmt.Println("Task added successfully!")
}

// List all tasks
func ListTasks() {
	if len(Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("\nList of Tasks:")
	for _, task := range Tasks {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d | Title: %s | Status: %s | Priority: %s\n",
			task.ID, task.Title, status, task.Priority)
	}
}

// Complete a task
func CompleteTask(id int) {
	for i := range Tasks {
		if Tasks[i].ID == id {
			Tasks[i].Completed = true
			fmt.Println("Task marked as completed!")
			return
		}
	}
	fmt.Println("Task not found!")
}

// Delete a task
func DeleteTask(id int) {
	for i, task := range Tasks {
		if task.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			fmt.Printf("Task with ID %d deleted successfully!\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found!\n", id)
}
