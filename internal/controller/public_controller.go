package controller

import (
	"net/http"
)

type PublicController struct {
}

func (c *PublicController) Create(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("OK"))
}

func NewPublicController() *PublicController {
	return &PublicController{}
}
