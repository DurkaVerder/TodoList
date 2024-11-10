package service

import (
	"TodoList/internal/model"
	"time"
)

type TaskService interface {
	CreateTask(task model.EnterDataTask, userId int) error
	GetTasks(userId int) ([]model.Task, error)
	UpdateTask(data model.EnterDataTask, taskId int) error
	DestroyTask(taskId int) error
}

func (s *ServiceManager) CreateTask(data model.EnterDataTask, userId int) error {
	task := model.Task{Title: data.Title, Description: data.Description, Status: data.Status, CreatorId: userId}
	if err := s.repo.AddTask(task); err != nil {
		return err
	}
	return nil
}

func (s *ServiceManager) GetTasks(userId int) ([]model.Task, error) {
	tasks, err := s.repo.AllTasksByUser(userId)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *ServiceManager) UpdateTask(data model.EnterDataTask, taskId int) error {
	task, err := s.repo.GetTask(taskId)
	if err != nil {
		return err
	}
	task = s.compareAndChangeData(data, task)
	if err := s.repo.UpdateTask(taskId, task); err != nil {
		return err
	}
	return nil
}

func (s *ServiceManager) DestroyTask(taskId int) error {
	if err := s.repo.DeleteTask(taskId); err != nil {
		return err
	}
	return nil
}

func (s *ServiceManager) compareAndChangeData(data model.EnterDataTask, task model.Task) model.Task {
	if task.Title != data.Title {
		task.Title = data.Title
	}
	if task.Status != data.Status {
		task.Status = data.Status
	}
	if task.Description != data.Description {
		task.Description = data.Description
	}
	time := time.Now()
	task.DueDate = &time
	return task
}
