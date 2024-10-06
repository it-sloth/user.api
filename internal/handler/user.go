package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"it-sloth/user.api/config"
	"it-sloth/user.api/internal/convertor"
	"it-sloth/user.api/internal/repository"
	"it-sloth/user.api/internal/service"
	"it-sloth/user.api/internal/wrapper"
	"net/http"
)

type User struct {
	userService     *service.User
	responseWrapper *wrapper.ResponseWriter
}

func (h User) GetUser(rw http.ResponseWriter, request *http.Request) {
	user, notFound := h.userService.Get(mux.Vars(request)["guid"])

	if notFound != nil {
		h.responseWrapper.WriteError(rw, notFound, http.StatusNotFound)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	h.responseWrapper.Write(rw, string(response), 200)
}

func NewUser() *User {
	return &User{
		userService: service.NewUserService(
			repository.NewUserRepository(config.GetEnv()),
			convertor.NewUserEntityConvertor()),
		responseWrapper: wrapper.NewResponseWriter(),
	}
}
