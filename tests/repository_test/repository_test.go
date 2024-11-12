package repository_test

import (
	"TodoList/internal/model"
	"TodoList/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAllTasksByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)

	// check valid userId

	userId := 1
	expectedTasks := []model.Task{
		{Id: 1, Title: "Task 1", Description: "Description 1", CreatorId: userId},
		{Id: 2, Title: "Task 2", Description: "Description 2", CreatorId: userId},
	}
	mockRepo.EXPECT().AllTasksByUser(userId).Return(expectedTasks, nil).Times(1)

	tasks, err := mockRepo.AllTasksByUser(userId)
	assert.NoError(t, err)
	assert.Equal(t, expectedTasks, tasks)

	// check invalid userId

	userId = -1
	mockRepo.EXPECT().AllTasksByUser(userId).Return(nil, errors.New("")).Times(1)
	_, err = mockRepo.AllTasksByUser(userId)
	assert.Error(t, err)
}

func TestGetTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)

	taskId := 1
	validTask := model.Task{Id: taskId, Title: "Task 1", Description: "Description 1", CreatorId: 1}
	mockRepo.EXPECT().GetTask(taskId).Return(validTask, nil).Times(1)

	task, err := mockRepo.GetTask(taskId)
	assert.NoError(t, err)
	assert.Equal(t, validTask, task)
}

func TestAddTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)

	// check valid Task

	validTask := model.Task{Title: "Test title", Description: "Test description", Status: "Test status", CreatorId: 1}
	mockRepo.EXPECT().AddTask(validTask).Return(nil).Times(1)
	err := mockRepo.AddTask(validTask)
	assert.NoError(t, err)

	// check invalid task

	invalidTask := model.Task{}
	mockRepo.EXPECT().AddTask(invalidTask).Return(errors.New("")).Times(1)

	err = mockRepo.AddTask(invalidTask)
	assert.Error(t, err)

}

func TestUpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)

	// check valid data

	task := model.Task{Title: "Test title", Description: "Test description", Status: "Test status", CreatorId: 1}
	taskId := 1
	mockRepo.EXPECT().UpdateTask(taskId, task).Return(nil).Times(1)

	err := mockRepo.UpdateTask(taskId, task)
	assert.NoError(t, err)
}

func TestDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepository(ctrl)

	taskId := 1
	mockRepo.EXPECT().DeleteTask(taskId).Return(nil).Times(1)

	err := mockRepo.DeleteTask(taskId)
	assert.NoError(t, err)

	taskId = -1
	mockRepo.EXPECT().DeleteTask(taskId).Return(errors.New("")).Times(1)

	err = mockRepo.DeleteTask(taskId)
	assert.Error(t, err)
}
