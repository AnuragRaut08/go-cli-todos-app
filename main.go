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
	// Initialize tasks
	cmd.Init()
	defer func() {
		err := storage.SaveTasks(cmd.Tasks)
		if err != nil {
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
			fmt.Print("Enter task: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)

			fmt.Print("Enter priority (High, Medium, Low): ")
			priority, _ := reader.ReadString('\n')
			priority = strings.TrimSpace(priority)

			cmd.AddTask(task, priority)

		case "2":
			cmd.ListTasks()

		case "3":
			fmt.Print("Enter task ID to mark complete: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid task ID. Please enter a number.")
				continue
			}
			cmd.CompleteTask(id)

		case "4":
			fmt.Print("Enter task ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid task ID. Please enter a number.")
				continue
			}
			cmd.DeleteTask(id)

		case "5":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
