package postgres

import "TodoList/internal/model"

func (repo PostgresRepo) AddTask(task model.Task) error {
	req := "INSERT INTO tasks (title, description, status, creator_id) VALUES ($1, $2, $3, $4)"
	_, err := repo.db.Exec(req, task.Title, task.Description, task.Status, task.CreatorId)
	if err != nil {
		return err
	}

	return nil
}

func (repo PostgresRepo) AllTasks(userId int) ([]model.Task, error) {
	req := "SELECT * FROM tasks WHERE creator_id = $1"
	rows, err := repo.db.Query(req, userId)
	if err != nil {
		return nil, err
	}

	tasks := []model.Task{}
	for rows.Next() {
		task := model.Task{}
		err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.DueDate, &task.CreatorId)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repo PostgresRepo) GetTask(taskId int) (model.Task, error) {
	req := "SELECT * FROM tasks WHERE id = $1"
	row := repo.db.QueryRow(req, taskId)

	task := model.Task{}
	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.DueDate, &task.CreatorId)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (repo PostgresRepo) UpdateTask(taskId int, newData model.Task) error {
	req := "UPDATE tasks SET title = $1, description = $2, status = $3, due_date= $4 WHERE id = $5"
	_, err := repo.db.Exec(req, newData.Title, newData.Description, newData.Status, newData.DueDate, taskId)
	if err != nil {
		return err
	}

	return nil
}

func (repo PostgresRepo) DeleteTask(taskId int) error {
	req := "DELETE FROM tasks WHERE id = $1"
	_, err := repo.db.Exec(req, taskId)
	if err != nil {
		return err
	}

	return nil
}
