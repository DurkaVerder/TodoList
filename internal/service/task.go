package service

type TaskService interface {
	CreateTask() error
	GetTasks() error
	UpdateTask() error
	DestroyTask() error
}

func (s *ServiceManager) CreateTask() error {
	return nil
}

func (s *ServiceManager) GetTasks() error {
	return nil
}

func (s *ServiceManager) UpdateTask() error {
	return nil
}

func (s *ServiceManager) DestroyTask() error {
	return nil
}
