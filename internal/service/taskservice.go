package service

type TaskService interface {
	CreateTask() error
	GetTasks() error
	UpdateTask() error
	DestroyTask() error
}
