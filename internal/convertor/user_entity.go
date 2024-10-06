package convertor

import (
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/entity"
)

type UserEntity struct{}

func (t *UserEntity) EntityToGetRequest(user entity.User) dto.UserResponse {
	return dto.UserResponse{
		Guid:        user.Id,
		Nickname:    user.Nick,
		Role:        user.Role.Name,
		CreatedAt:   user.CreatedAt,
		ActivatedAt: user.ActivatedAt,
	}
}

func NewUserEntityConvertor() *UserEntity {
	return &UserEntity{}
}
