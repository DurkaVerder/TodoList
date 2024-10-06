package postgres

import "TodoList/internal/model"

func (repo PostgresRepo) AddUser(model.User) error {
	return nil
}

func (repo PostgresRepo) AllUsers() ([]model.User, error) {
	return []model.User{}, nil
}

func (repo PostgresRepo) GetUser() (model.User, error) {
	return model.User{}, nil
}

func (repo PostgresRepo) ChangeUser(id int, newData model.User) error {
	return nil
}

func (repo PostgresRepo) DeleteUser(id int) error {
	return nil
}
