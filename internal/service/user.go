package service

import (
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/repository"
)

type User struct {
	userRepository *repository.UserRepository
}

func (u *User) Create(user dto.UserCreateRequest) (string, error) {
	return u.userRepository.Check(user.Login, user.Email)
}

func NewUser(userRepository *repository.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}