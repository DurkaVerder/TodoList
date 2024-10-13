package postgres

import "TodoList/internal/model"

func (repo PostgresRepo) AddUser(user model.User) error {
	req := "INSERT INTO users (name, login, password) VALUES ($1, $2, $3)"
	_, err := repo.db.Exec(req, user.Name, user.Login, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (repo PostgresRepo) AllUsers() ([]model.User, error) {
	req := "SELECT * FROM users"
	rows, err := repo.db.Query(req)
	if err != nil {
		return nil, err
	}

	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Login, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo PostgresRepo) GetUser(userId int) (model.User, error) {
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

func (repo PostgresRepo) UpdateUser(userId int, newData model.User) error {
	req := "UPDATE users SET name = $1, login = $2, password = $3 WHERE id = $4"
	_, err := repo.db.Exec(req, newData.Name, newData.Login, newData.Password, userId)
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
