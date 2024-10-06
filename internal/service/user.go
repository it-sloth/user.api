package service

import (
	"it-sloth/user.api/internal/convertor"
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/internal_error"
	"it-sloth/user.api/internal/repository"
)

type User struct {
	userRepository *repository.User
	dataConvertor  *convertor.UserEntity
}

func (u *User) Get(guid string) (dto.UserResponse, internal_error.InternalErrorInterface) {
	user, err := u.userRepository.GetUser(guid)

	if err != nil {
		return dto.UserResponse{}, internal_error.NewUserNotFoundError("user not found", "3OJY9KOR")
	}

	return u.dataConvertor.EntityToGetRequest(user), nil
}

func NewUserService(userRepository *repository.User, dataConvertor *convertor.UserEntity) *User {
	return &User{
		userRepository: userRepository,
		dataConvertor:  dataConvertor,
	}
}
