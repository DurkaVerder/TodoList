package repository

import "TodoList/internal/model"

type Repository interface {
	UserRepository
	TaskRepository
}

type UserRepository interface {
	GetByUserData(data model.EnterDataUser) (model.User, error)
	GetByUserId(userId int) (model.User, error)
	AddUser(data model.EnterDataUser) error
	UpdateUser(userId int, newData model.EnterDataUser) error
	DeleteUser(userId int) error
}

type TaskRepository interface {
	AllTasks(userId int) ([]model.Task, error)
	GetTask(taskId int) (model.Task, error)
	AddTask(task model.Task) error
	UpdateTask(taskId int, newData model.Task) error
	DeleteTask(taskId int) error
}
