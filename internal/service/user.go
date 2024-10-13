package service

type UserService interface {
	Login() error
	Register() error
	GetUser() error
	UpdateUser() error
}

func (s *ServiceManager) Login() error {
	return nil
}

func (s *ServiceManager) Register() error {
	return nil
}

func (s *ServiceManager) GetUser() error {
	return nil
}

func (s *ServiceManager) UpdateUser() error {
	return nil
}
