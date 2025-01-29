package cmd

import (
	"go-cli-todos-app/storage"
	"testing"
)

// Test AddTask function
func TestAddTask(t *testing.T) {
	Tasks = []storage.Task{} // I am resetting tasks before testing

	AddTask("Test Task 1", "High")
	if len(Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(Tasks))
	}

	if Tasks[0].Name != "Test Task 1" {
		t.Errorf("Expected task name to be 'Test Task 1', got %s", Tasks[0].Name)
	}

	if Tasks[0].Priority != "High" {
		t.Errorf("Expected priority 'High', got %s", Tasks[0].Priority)
	}
}

// Testing CompleteTask function
func TestCompleteTask(t *testing.T) {
	Tasks = []storage.Task{
		{ID: 1, Name: "Task 1", Priority: "Medium", Done: false},
	}

	success := CompleteTask(1)
	if !success {
		t.Errorf("Expected task to be marked complete, but it was not")
	}

	if !Tasks[0].Done {
		t.Errorf("Task was not marked as complete")
	}
}

// Testing DeleteTask function
func TestDeleteTask(t *testing.T) {
	Tasks = []storage.Task{
		{ID: 1, Name: "Task 1", Priority: "Low"},
	}

	success := DeleteTask(1)
	if !success {
		t.Errorf("Expected task to be deleted, but it was not")
	}

	if len(Tasks) != 0 {
		t.Errorf("Expected 0 tasks after deletion, got %d", len(Tasks))
	}
}

// Testing ListTasks function
func TestListTasks(t *testing.T) {
	Tasks = []storage.Task{
		{ID: 1, Name: "Sample Task", Priority: "High", Done: false},
	}

	ListTasks() 
}

// Test adding an empty task
func TestAddEmptyTask(t *testing.T) {
	Tasks = []storage.Task{} // Reset task list

	AddTask("", "Low")
	if len(Tasks) != 1 {
		t.Errorf("Task with empty name should still be added (optional validation)")
	}
}

// Test completing a task that doesn't exist
func TestCompleteInvalidTask(t *testing.T) {
	Tasks = []storage.Task{} // Reset task list

	success := CompleteTask(999) // Non-existent task
	if success {
		t.Errorf("Expected false when completing non-existent task, but got true")
	}
}

// Test deleting a task that doesn't exist
func TestDeleteInvalidTask(t *testing.T) {
	Tasks = []storage.Task{} // Reset task list

	success := DeleteTask(999) // Non-existent task
	if success {
		t.Errorf("Expected false when deleting non-existent task, but got true")
	}
}

