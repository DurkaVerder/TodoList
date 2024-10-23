package service

import (
	"TodoList/internal/jwt"
	"TodoList/internal/model"
)

type UserService interface {
	Login(data model.EnterDataUser) (string, error)
	Register(data model.EnterDataUser) (string, error)
	GetUser() error
	UpdateUser() error
}

func (s *ServiceManager) Login(data model.EnterDataUser) (string, error) {
	user, err := s.repo.GetUser(data)
	if err != nil {
		return "", err
	}
	if !(user.Login == data.Login && user.Password == data.Password) {
		return "", nil
	}

	token, err := jwt.GenerateJWT(user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *ServiceManager) Register(data model.EnterDataUser) (string, error) {
	if err := s.repo.AddUser(data); err != nil {
		return "", err
	}

	return s.Login(data)
}

func (s *ServiceManager) GetUser() error {
	return nil
}

func (s *ServiceManager) UpdateUser() error {
	return nil
}
