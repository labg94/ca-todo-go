package domain

import "time"

const (
	Created = iota
	WorkingOn
	Completed
)

type Todo struct {
	task         string
	status       int
	creationDate time.Time
	updateTime   time.Time
}

func NewTodo(task string) Todo {
	return Todo{task: task, status: Created}
}

func (todo Todo) GetTask() string {
	return todo.task
}

func (todo Todo) UpdateStatus() Todo {
	if todo.status == Completed {
		return todo
	}
	return Todo{task: todo.task, status: todo.status + 1}
}
