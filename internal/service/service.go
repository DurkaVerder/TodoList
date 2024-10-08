package service

type Service interface {
	TaskManager
	UserManager
}

type TaskManager interface {
	GetAllTasks() error
	GetTask() error
}

type UserManager interface {
}
