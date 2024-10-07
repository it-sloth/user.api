package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"it-sloth/user.api/config"
	"it-sloth/user.api/internal/convertor"
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/error"
	"it-sloth/user.api/internal/repository"
	"it-sloth/user.api/internal/service"
	"it-sloth/user.api/internal/wrapper"
	"net/http"
)

type User struct {
	userService     *service.User
	responseWrapper *wrapper.ResponseWriter
}

func (h *User) Create(rw http.ResponseWriter, r *http.Request) {
	_, cError := h.userService.Create(dto.UserCreateRequest{})

	if cError != nil && cError.Status() == http.StatusConflict {
		h.responseWrapper.WriteError(rw, cError, cError.Status())
	}

	if cError != nil {
		h.UnknownError(rw, "WU28J4D9")
		return
	}
}

func (h *User) Read(rw http.ResponseWriter, request *http.Request) {
	user, cError := h.userService.Read(mux.Vars(request)["guid"])

	if cError != nil && cError.Status() == http.StatusNotFound {
		h.responseWrapper.WriteError(rw, cError, cError.Status())
		return
	}

	if cError != nil {
		h.UnknownError(rw, "JG83HS1O")
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		h.UnknownError(rw, "4H7FGW53")
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	h.responseWrapper.Write(rw, string(response), http.StatusOK)
}

func (h *User) UnknownError(rw http.ResponseWriter, code string) {
	h.responseWrapper.WriteError(rw, error.NewInternal("unknown error occurred", code), http.StatusInternalServerError)
}

func NewUser() *User {
	return &User{
		userService: service.NewUserService(
			repository.NewUserRepository(config.GetEnv()),
			convertor.NewUserEntityConvertor()),
		responseWrapper: wrapper.NewResponseWriter(),
	}
}
