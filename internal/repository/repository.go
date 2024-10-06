package repository

type UserRepo interface {
	UserGetter
	UserAdder
	UserChanger
	UserDeleter
}

type TaskRepo interface {
	TaskGetter
	TaskAdder
	TaskChanger
	TaskDeleter
}
