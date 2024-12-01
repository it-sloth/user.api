package controller

import (
	"encoding/json"
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/service"
	"log"
	"net/http"
)

type PublicController struct {
	userService *service.User
}

func (c *PublicController) Create(rw http.ResponseWriter, r *http.Request) {
	var dto dto.UserCreateRequest
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		log.Fatal(err)
		rw.Write([]byte(err.Error()))
		return
	}

	response, err := c.userService.Create(dto)
	if err != nil {
		log.Fatal(err)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(response))
}

func NewPublicController(userService *service.User) *PublicController {
	return &PublicController{
		userService: userService,
	}
}
