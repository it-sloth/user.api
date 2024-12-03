package controller

import (
	"it-sloth/user.api/internal/factory"
	"it-sloth/user.api/internal/service"
	"it-sloth/user.api/internal/wrapper"
	"net/http"
)

type PublicController struct {
	userService *service.User
	dtoFactory *factory.DtoFactory
	responseWriterWrapper *wrapper.ResponseWriter
}

func (c *PublicController) Create(rw http.ResponseWriter, r *http.Request) {
	dto, err := c.dtoFactory.UserCreateDto(r.Body)
	if err != nil {
		panic(err)
	}

	guid, err := c.userService.Create(dto)
	if err != nil {
		panic(err)
	}

	respDto := c.dtoFactory.UserCreateResponseDto(guid)
	c.responseWriterWrapper.WriteSuccess(rw, respDto, http.StatusCreated)
}

func NewPublicController(userService *service.User, dtoFactory *factory.DtoFactory, responseWriterWrapper *wrapper.ResponseWriter) *PublicController {
	return &PublicController{
		userService: userService,
		dtoFactory: dtoFactory,
		responseWriterWrapper: responseWriterWrapper,
	}
}
