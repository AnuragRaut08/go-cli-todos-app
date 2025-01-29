package todo

type Task struct {
	ID   int
	Name string
	Done bool
}

type TodoList struct {
	Tasks []Task
}

// Constructor for TodoList
func NewTodoList() *TodoList {
	return &TodoList{}
}

// AddTask adds a new task
func (t *TodoList) AddTask(name string) {
	newTask := Task{ID: len(t.Tasks) + 1, Name: name, Done: false}
	t.Tasks = append(t.Tasks, newTask)
}

// DeleteTask removes a task by ID
func (t *TodoList) DeleteTask(id int) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			break
		}
	}
}

// ListTasks returns all tasks
func (t *TodoList) ListTasks() []Task {
	return t.Tasks
}
