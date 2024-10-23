package service

import (
	"TodoList/internal/model"
)

type UserService interface {
	Login(data model.EnterDataUser) error
	Register(data model.EnterDataUser) error
	GetUser() error
	UpdateUser() error
	GetIdUser(data model.EnterDataUser) (int, error)
}

func (s *ServiceManager) Login(data model.EnterDataUser) error {

	return nil
}

func (s *ServiceManager) Register(data model.EnterDataUser) error {
	if err := s.repo.AddUser(data); err != nil {
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

func (s *ServiceManager) GetIdUser(data model.EnterDataUser) (int, error) {
	userId, err := s.repo.GetIdUser(data)
	if err != nil {
		return -1, err
	}

	return userId, nil
}
