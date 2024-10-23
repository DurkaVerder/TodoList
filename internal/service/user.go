package service

import (
	"TodoList/internal/model"
)

type UserService interface {
	Login(data model.EnterDataUser) (bool, error)
	Register(data model.EnterDataUser) error
	GetUser() error
	UpdateUser() error
	GetIdUser(data model.EnterDataUser) (int, error)
}

func (s *ServiceManager) Login(data model.EnterDataUser) (bool, error) {
	user, err := s.repo.GetUser(data)
	if err != nil {
		return false, err
	}
	if !(user.Login == data.Login && user.Password == data.Password) {
		return false, nil
	}

	// create JWT and save he in cookie
	return true, nil
}

func (s *ServiceManager) Register(data model.EnterDataUser) error {
	if err := s.repo.AddUser(data); err != nil {
		return err
	}
	// create JWT and save he in cookie
	return nil
}

func (s *ServiceManager) GetUser() error {
	return nil
}

func (s *ServiceManager) UpdateUser() error {
	return nil
}

func (s *ServiceManager) GetIdUser(data model.EnterDataUser) (int, error) {
	userId, err := s.repo.GetIdUser(data)
	if err != nil {
		return -1, err
	}

	return userId, nil
}
