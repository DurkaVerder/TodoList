package postgres

import "TodoList/internal/model"

func (repo PostgresRepo) AddUser(data model.EnterDataUser) error {
	req := "INSERT INTO users (name, login, password) VALUES ($1, $2, $3)"
	_, err := repo.db.Exec(req, data.Name, data.Login, data.Password)
	if err != nil {
		return err
	}

	return nil
}

func (repo PostgresRepo) GetByUserId(userId int) (model.User, error) {
	req := "SELECT * FROM users WHERE id = $1"
	row := repo.db.QueryRow(req, userId)

	user := model.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Login, &user.Password)
	if err != nil {
		return model.User{}, err
	}

	err = row.Err()
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo PostgresRepo) GetByUserData(data model.EnterDataUser) (model.User, error) {
	req := "SELECT * FROM users WHERE login = $1 AND password = $2"
	row := repo.db.QueryRow(req, data.Login, data.Password)

	user := model.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Login, &user.Password)
	if err != nil {
		return model.User{}, err
	}

	err = row.Err()
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (repo PostgresRepo) UpdateUserName(userId int, name string) error {
	req := "UPDATE users SET name = $1 WHERE id = $2"
	_, err := repo.db.Exec(req, name, userId)
	if err != nil {
		return err
	}

	return nil
}

func (repo PostgresRepo) UpdateUserPassword(userId int, password string) error {
	req := "UPDATE users SET password = $1 WHERE id = $2"
	_, err := repo.db.Exec(req, password, userId)
	if err != nil {
		return err
	}

	return nil
}

func (repo PostgresRepo) DeleteUser(userId int) error {
	req := "DELETE FROM users WHERE id = $1"
	_, err := repo.db.Exec(req, userId)
	if err != nil {
		return err
	}

	return nil
}
