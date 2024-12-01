package service

import (
	"errors"
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/factory"
	"it-sloth/user.api/internal/repository"
)

type User struct {
	userRepository *repository.UserRepository
	userFactory    *factory.EntityFactory
}

func (u *User) Create(userDto dto.UserCreateRequest) (string, error) {
	guid, err := u.userRepository.Check(userDto.Login, userDto.Email)
	if err != nil {
		return "", errors.New("internal error")
	}

	if guid != "" {
		return "", errors.New("user already exists")
	}

	user := u.userFactory.EntityFromCreateDto(userDto)
	guid, err = u.userRepository.Create(user)
	if err != nil {
		return "", errors.New("internal error")
	}

	return guid, nil
}

func NewUser(userRepository *repository.UserRepository, userFactory *factory.EntityFactory) *User {
	return &User{
		userRepository: userRepository,
		userFactory:    userFactory,
	}
}