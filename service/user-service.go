package service

import (
	"rock-paper-scissor/app/users"
)

type UserService interface {
	Save(users.User) users.User
	FindAll() []users.User
}

type userService struct {
	users []users.User
}

func New() UserService {
	return &userService{
		users: []users.User{},
	}
}

func (service *userService) Save(user users.User) users.User {
	service.users = append(service.users, user)
	return user
}

func (service *userService) FindAll() []users.User {
	return service.users
}
