package controller

import (
	"TodoList/internal/service"
)

type Controller interface {
	UserController
	TaskController
}

type ControllerManager struct {
	Service service.Service
}

func NewController(s service.Service) *ControllerManager {
	return &ControllerManager{Service: s}
}
