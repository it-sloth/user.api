package factory

import (
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/entity"
	"time"
)

type EntityFactory struct {
}

func (u *EntityFactory) EntityFromCreateDto(userDto dto.UserCreateRequest) entity.User {
	now := time.Now()
	return entity.User{
		Login:    userDto.Login,
		Nickname: userDto.Nickname,
		Email:    userDto.Email,
		Password: userDto.Password,
		Active:   true,
		Role:     entity.Role{
			Name: entity.UserRole,
		},
		CreatedAt:   now.Format(time.DateOnly),
	}
}

func NewEntityFactory() *EntityFactory {
	return &EntityFactory{}
}