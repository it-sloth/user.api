package dto

import "time"

type UserResponse struct {
	Guid        string    `json:"guid"`
	Nickname    string    `json:"nickname"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	ActivatedAt time.Time `json:"activated_at"`
}

type UserCreateRequest struct {
	Login    string `json:"login"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
