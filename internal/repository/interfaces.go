package repository

import "TodoList/internal/model"

// Interface for tasks
type TaskGetter interface {
	AllTasks(userId int) ([]model.Task, error)
	GetTask(taskId int) (model.Task, error)
}

type TaskAdder interface {
	AddTask(task model.Task) error
}

type TaskChanger interface {
	ChangeTask(taskId int, newData model.Task) error
}

type TaskDeleter interface {
	DeleteTask(taskId int) error
}

// Interface for users
type UserGetter interface {
	AllUsers() ([]model.User, error)
	GetUser(userId int) (model.User, error)
}

type UserAdder interface {
	AddUser(user model.User) error
}

type UserChanger interface {
	ChangeUser(userId int, newData model.User) error
}

type UserDeleter interface {
	DeleteUser(userId int) error
}
