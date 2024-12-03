package wrapper

import (
	"encoding/json"
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/error"
	"net/http"
)

type ResponseWriter struct {
}

func (w *ResponseWriter) WriteError(rw http.ResponseWriter, error error.InternalError) {
	errorDto, err := json.Marshal(dto.ErrorResponse{
		Code:    error.Code(),
		Message: error.Error(),
	})

	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	http.Error(rw, string(errorDto), error.Status())
}

func (w *ResponseWriter) WriteSuccess(rw http.ResponseWriter, body any, httpCode int) {
	dto, err := json.Marshal(body)

	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	http.Error(rw, string(dto), httpCode)
}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{}
}