package dto

import (
	"it-sloth/user.api/internal/entity"
)

type UserResponse struct {
	User    entity.User `json:"user"`
	Version string      `json:"version"`
}
