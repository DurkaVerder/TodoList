package service

import (
	"TodoList/internal/repository"
)

type Service interface {
	UserService
	TaskService
}

type ServiceManager struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *ServiceManager {
	return &ServiceManager{repo: repo}
}
