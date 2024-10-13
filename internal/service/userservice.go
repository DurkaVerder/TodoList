package service

type UserService interface {
	Login() error
	Register() error
	GetUser() error
	UpdateUser() error
}
