package repository

import "TodoList/internal/model"

// Interface for tasks
type TaskGetter interface {
	AllTasks() ([]model.Task, error)
	GetTask() (model.Task, error)
}

type TaskAdder interface {
	AddTask(user model.Task) error
}

type TaskChanger interface {
	ChangeTask(id int, newData model.Task) error
}

type TaskDeleter interface {
	DeleteTask(id int) error
}

// Interface for users
type UserGetter interface {
	AllUsers() ([]model.User, error)
	GetUser() (model.User, error)
}

type UserAdder interface {
	AddUser(user model.User) error
}

type UserChanger interface {
	ChangeUser(id int, newData model.User) error
}

type UserDeleter interface {
	DeleteUser(id int) error
}
