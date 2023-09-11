package domain

type Store interface {
	GetAllUsers() ([]User, *AppError)
	CreateUser(input User) (User, *AppError)
}
