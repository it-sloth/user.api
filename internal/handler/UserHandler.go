package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"it-sloth/user.api/config"
	"it-sloth/user.api/internal/convertor"
	"it-sloth/user.api/internal/repository"
	"net/http"
)

type UserHandler struct {
	userRepository *repository.User
	userConvertor  *convertor.UserEntity
}

func (h UserHandler) GetUser(rw http.ResponseWriter, request *http.Request) {
	user, err := h.userRepository.GetUser(mux.Vars(request)["name"])

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	data := h.userConvertor.EntityToGetRequest(user)

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

func NewPublic() *UserHandler {
	return &UserHandler{
		userRepository: repository.NewUserRepository(config.GetEnv()),
		userConvertor:  convertor.NewUserEntityConvertor(),
	}
}
