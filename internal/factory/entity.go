package factory

import (
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/entity"
	"time"
)

type EntityFactory struct {
}

func (u *EntityFactory) EntityFromCreateDto(userDto dto.UserCreateRequest) entity.User {
	return entity.User{
		Login:    userDto.Login,
		Nickname: userDto.Nickname,
		Email:    userDto.Email,
		Password: userDto.Password,
		Active:   true,
		Role:     entity.Role{
			Name: entity.UserRole,
		},
		CreatedAt:   time.Now().Format("YYYY-MM-DD"),
		UpdatedAt:   time.Now().Format("YYYY-MM-DD"),
	}
}

func NewEntityFactory() *EntityFactory {
	return &EntityFactory{}
}