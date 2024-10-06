package service

import (
	"TodoList/internal/model"
	"TodoList/internal/repository"

	"github.com/labstack/echo"
)

type TaskService struct {
	TaskRepository *repository.TaskRepository
}

// logic layer

func (s *TaskService) GetAllTask(ctx echo.Context) ([]model.Task, error) {
	tasks, err := s.TaskRepository.GetAllTasks(ctx.Param("userid"))
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTask(ctx echo.Context) (model.Task, error) {
	return s.TaskRepository.GetTask(ctx.Param("taskid"))
}

func (s *TaskService) AddTask(ctx echo.Context) error {
	addedTask := model.Task{} // Read from context
	s.TaskRepository.AddTask(addedTask)
	return nil
}

func (s *TaskService) DeleteTask(ctx echo.Context) error {
	return nil
}

func (s *TaskService) ChangeTask(ctx echo.Context) error {
	return nil
}
