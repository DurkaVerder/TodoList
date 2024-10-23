package controller

import (
	"TodoList/internal/cookie"
	"TodoList/internal/service"
)

type Controller interface {
	UserController
	TaskController
}

type ControllerManager struct {
	Service service.Service
	Cookie  cookie.Cookie
}

func NewController(s service.Service, c cookie.Cookie) *ControllerManager {
	return &ControllerManager{Service: s, Cookie: c}
}
