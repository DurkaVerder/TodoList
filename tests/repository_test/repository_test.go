package repository_test

import (
	"TodoList/internal/model"
	"TodoList/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	task := model.Task{Title: "Test title", Description: "Test description", Status: "Test status", CreatorId: 1}
	mockRepo.EXPECT().AddTask(task).Return(nil).Times(1)

	err := mockRepo.AddTask(task)

	assert.NoError(t, err)
}

func TestAllTasksByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)

	userId := 1
	expectedTasks := []model.Task{
		{Id: 1, Title: "Task 1", Description: "Description 1", CreatorId: userId},
		{Id: 2, Title: "Task 2", Description: "Description 2", CreatorId: userId},
	}
	mockRepo.EXPECT().AllTasksByUser(userId).Return(expectedTasks, nil).Times(1)

	tasks, err := mockRepo.AllTasksByUser(userId)
	assert.NoError(t, err)
	assert.Equal(t, expectedTasks, tasks)
}
