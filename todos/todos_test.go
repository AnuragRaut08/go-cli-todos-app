package todo

import (
	"testing"
)

// Test adding a task
func TestAddTask(t *testing.T) {
	todoList := NewTodoList()
	todoList.AddTask("Complete Go CLI app")

	if len(todoList.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(todoList.Tasks))
	}
	if todoList.Tasks[0].Name != "Complete Go CLI app" {
		t.Errorf("Expected task name 'Complete Go CLI app', got '%s'", todoList.Tasks[0].Name)
	}
}

// Test deleting a task
func TestDeleteTask(t *testing.T) {
	todoList := NewTodoList()
	todoList.AddTask("Task to be deleted")
	todoList.DeleteTask(1)

	if len(todoList.Tasks) != 0 {
		t.Errorf("Expected 0 tasks after deletion, got %d", len(todoList.Tasks))
	}
}

// Test listing tasks
func TestListTasks(t *testing.T) {
	todoList := NewTodoList()
	todoList.AddTask("Task 1")
	todoList.AddTask("Task 2")

	tasks := todoList.ListTasks()
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}
