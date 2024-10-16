package repository

import "TodoList/internal/model"

type Repository interface {
	UserRepository
	TaskRepository
}

type UserRepository interface {
	GetIdUser(login, password string) (int, error)
	GetUser(userId int) (model.User, error)
	AddUser(user model.User) error
	UpdateUser(userId int, newData model.User) error
	DeleteUser(userId int) error
}

type TaskRepository interface {
	AllTasks(userId int) ([]model.Task, error)
	GetTask(taskId int) (model.Task, error)
	AddTask(task model.Task) error
	UpdateTask(taskId int, newData model.Task) error
	DeleteTask(taskId int) error
}
