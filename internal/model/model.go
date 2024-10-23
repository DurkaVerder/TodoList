package model

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	DueDate     time.Time `json:"due_date"`
	CreatorId   int       `json:"creator_id"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type EnterDataUser struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
