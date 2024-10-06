package postgres

import "TodoList/internal/model"

type UserAdder interface {
	AddUser(user model.User) error
}

type UserGetter interface {
	AllUsers() ([]model.User, error)
	GetUser() (model.User, error)
}

