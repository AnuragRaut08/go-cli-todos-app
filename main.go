package main

import (
	"bufio"
	"fmt"
	"go-cli-todos-app/cmd"     // Import your cmd package
	"go-cli-todos-app/storage" // Import your storage package
	"os"
	"strconv"
	"strings"
)

func main() {
	// Initialize tasks from storage
	err := cmd.Init()
	if err != nil {
		fmt.Println("Error initializing tasks:", err)
		return
	}

	// Ensure tasks are saved before exiting
	defer func() {
		if err := storage.SaveTasks(cmd.Tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nTODO CLI Menu:")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter task description: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)

			if task == "" {
				fmt.Println("Task cannot be empty. Please enter a valid task.")
				continue
			}

			fmt.Print("Enter priority (High, Medium, Low): ")
			priority, _ := reader.ReadString('\n')
			priority = strings.TrimSpace(priority)

			if priority != "High" && priority != "Medium" && priority != "Low" {
				fmt.Println("Invalid priority. Please enter High, Medium, or Low.")
				continue
			}

			cmd.AddTask(task, priority)
			fmt.Println("Task added successfully!")

		case "2":
			cmd.ListTasks()

		case "3":
			fmt.Print("Enter task ID to mark as complete: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil || id <= 0 {
				fmt.Println("Invalid task ID. Please enter a valid number.")
				continue
			}

			if cmd.CompleteTask(id) {
				fmt.Println("Task marked as complete!")
			} else {
				fmt.Println("Task ID not found.")
			}

		case "4":
			fmt.Print("Enter task ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil || id <= 0 {
				fmt.Println("Invalid task ID. Please enter a valid number.")
				continue
			}

			if cmd.DeleteTask(id) {
				fmt.Println("Task deleted successfully!")
			} else {
				fmt.Println("Task ID not found.")
			}

		case "5":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
		}
	}
}
