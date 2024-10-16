package service

import (
	"TodoList/internal/model"
)

type UserService interface {
	Login(user model.User) error
	Register(newUser model.User) error
	GetUser() error
	UpdateUser() error
	GetIdUser(user model.User) (int, error)
}

func (s *ServiceManager) Login(user model.User) error {
	return nil
}

func (s *ServiceManager) Register(newUser model.User) error {
	if err := s.repo.AddUser(newUser); err != nil {
		return err
	}

	return nil
}

func (s *ServiceManager) GetUser() error {
	return nil
}

func (s *ServiceManager) UpdateUser() error {
	return nil
}

func (s *ServiceManager) GetIdUser(user model.User) (int, error) {
	userId, err := s.repo.GetIdUser(user.Login, user.Password)
	if err != nil {
		return -1, err
	}

	return userId, nil
}
