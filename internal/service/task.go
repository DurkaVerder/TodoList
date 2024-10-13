package service

type TaskService interface {
	CreateTask() error
	GetTasks() error
	UpdateTask() error
	DestroyTask() error
}

func (s *ServiceManager) CreateTask() error {

}

func (s *ServiceManager) GetTasks() error {

}

func (s *ServiceManager) UpdateTask() error {

}

func (s *ServiceManager) DestroyTask() error {

}
