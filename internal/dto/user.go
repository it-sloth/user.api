package dto

import "time"

type UserResponse struct {
	Nickname    string    `json:"nickname"`
	Role        string    `json:"role"`
	Email       string    `json:"email"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	ActivatedAt time.Time `json:"activated_at"`
}

type UserCreateRequest struct {
	Login    string `json:"login"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserCreateResponse struct {
	Guid string `json:"guid"`
}
