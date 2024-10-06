package repository

import "TodoList/internal/model"

func (r *TaskRepository) AddTask(task model.Task) error {
	req := "INSERT INTO task (title, description, status, createdAt, duedate, creatorid) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.DB.Exec(req, task.Title, task.Description, task.Status, task.CreatedAt, task.DueDate, task.CreatorId)
	if err != nil {
		return err
	}
	return nil
}
