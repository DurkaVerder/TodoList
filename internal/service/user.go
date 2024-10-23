package service

import (
	"TodoList/internal/jwt"
	"TodoList/internal/model"
)

type UserService interface {
	Login(data model.EnterDataUser) (string, error)
	Register(data model.EnterDataUser) (string, error)
	GetUser(userId int) (model.User, error)
	UpdateUser() error
}

func (s *ServiceManager) Login(data model.EnterDataUser) (string, error) {
	user, err := s.repo.GetByUserData(data)
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

func (s *ServiceManager) GetUser(userId int) (model.User, error) {
	user, err := s.repo.GetByUserId(userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *ServiceManager) UpdateUser() error {
	return nil
}
