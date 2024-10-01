package model

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"` // Используй time.Time вместо time.DateTime
	DueDate     time.Time `json:"due_date"`   // Поправлено название поля
	CreatorId   int       `json:"creator_id"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
