package dto

type UserCreateRequest struct {
	Login    string `json:"login"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserCreateResponse struct {
	Guid string `json:"guid"`
}