package service

import (
	"it-sloth/user.api/internal/convertor"
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/error"
	"it-sloth/user.api/internal/repository"
	"net/http"
)

type User struct {
	userRepository *repository.User
	dataConvertor  *convertor.UserEntity
}

func (u *User) Create(request dto.UserCreateRequest) (dto.UserCreateResponse, error.InternalError) {
	guid, _ := u.userRepository.CheckUser(request.Login, request.Email)
	if guid != "" {
		return dto.UserCreateResponse{}, error.Make("user already exists", "8D2H37G6", http.StatusConflict)
	}

	return dto.UserCreateResponse{Guid: "non"}, nil
}

func (u *User) Read(guid string) (dto.UserResponse, error.InternalError) {
	user, err := u.userRepository.GetUser(guid)

	if err != nil {
		return dto.UserResponse{}, error.Make("user not found", "3OJY9KOR", http.StatusNotFound)
	}

	return u.dataConvertor.EntityToGetRequest(user), nil
}

func NewUserService(userRepository *repository.User, dataConvertor *convertor.UserEntity) *User {
	return &User{
		userRepository: userRepository,
		dataConvertor:  dataConvertor,
	}
}
