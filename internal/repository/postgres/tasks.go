package postgres

import "TodoList/internal/model"

func (repo PostgresRepo) AddTask(model.User) error {
	return nil
}

func (repo PostgresRepo) AllTasks() ([]model.Task, error) {
	return []model.Task{}, nil
}

func (repo PostgresRepo) GetTask() (model.Task, error) {
	return model.Task{}, nil
}

func (repo PostgresRepo) ChangeTask(id int, newData model.Task) error {
	return nil
}

func (repo PostgresRepo) DeleteTask(id int) error {
	return nil
}
