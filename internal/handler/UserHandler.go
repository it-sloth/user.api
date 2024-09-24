package handler

import (
	"encoding/json"
	"it-sloth/user.api/internal/dto"
	"net/http"
	"time"
)

type UserHandler struct {
}

func (h UserHandler) GetUser(rw http.ResponseWriter, request *http.Request) {
	data := dto.UserResponse{
		Users:   h.makeUserList(),
		Version: "0.0.0.3",
	}

	response, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(response)
	if err != nil {
		return
	}
}

func (h UserHandler) makeUserList() []dto.User {
	users := []dto.User{
		{Nickname: "Admin", MemberSince: time.Now()},
		{Nickname: "Bob", MemberSince: time.Now()},
		{Nickname: "Alice", MemberSince: time.Now()},
		{Nickname: "Daniel", MemberSince: time.Now()},
	}

	return users
}

func NewPublic() *UserHandler {
	return &UserHandler{}
}
