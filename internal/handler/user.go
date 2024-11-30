package handler

import (
	"encoding/json"
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/error"
	"it-sloth/user.api/internal/service"
	"it-sloth/user.api/internal/wrapper"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	userService     *service.User
	responseWrapper *wrapper.ResponseWriter
}

func (h *User) Create(rw http.ResponseWriter, r *http.Request) {
	_, cError := h.userService.Create(dto.UserCreateRequest{})

	if cError != nil && cError.Status() == http.StatusConflict {
		h.responseWrapper.WriteError(rw, cError)
	}

	if cError != nil {
		h.UnknownError(rw, "WU28J4D9")
		return
	}
}

func (h *User) Read(rw http.ResponseWriter, request *http.Request) {
	user, cError := h.userService.Read(mux.Vars(request)["guid"])

	if cError != nil && cError.Status() == http.StatusNotFound {
		h.responseWrapper.WriteError(rw, cError)
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

	h.responseWrapper.Write(rw, string(response), http.StatusOK)
}

func (h *User) UnknownError(rw http.ResponseWriter, code string) {
	h.responseWrapper.WriteError(rw, error.Make("unknown error occurred", code, http.StatusInternalServerError))
}

// NewUser initializes a new User handler with its dependencies.
func NewUser(userService *service.User, responseWrapper *wrapper.ResponseWriter) *User {
	return &User{
		userService: userService,
		responseWrapper: responseWrapper,
	}
}
