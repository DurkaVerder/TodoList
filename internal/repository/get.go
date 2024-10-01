package repository

import "TodoList/internal/model"

func (r *TaskRepository) GetAllTasks(userId string) ([]model.Task, error) {
	req := "SELECT * FROM task WHERE creator_id = $1"
	rows, err := r.DB.Query(req, userId)
	if err != nil {
		return nil, err
	}
	tasks := []model.Task{}

	for rows.Next() {
		t := model.Task{}
		if err = rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.DueDate, &t.CreatorId); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepository) GetTask(taskId string) (model.Task, error) {
	req := "SELECT * FROM task WHERE id = $1"
	row := r.DB.QueryRow(req, taskId)
	task := model.Task{}
	if err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.DueDate, &task.CreatorId); err != nil {
		return task, err
	}
	return task, nil
}