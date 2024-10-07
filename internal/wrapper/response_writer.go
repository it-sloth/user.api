package wrapper

import (
	"encoding/json"
	"it-sloth/user.api/internal/dto"
	"it-sloth/user.api/internal/error"
	"net/http"
)

type ResponseWriter struct {
}

func (w *ResponseWriter) WriteError(rw http.ResponseWriter, error error.InternalError, code int) {
	errorDto, err := json.Marshal(dto.ErrorResponse{
		Code:    error.Code(),
		Message: error.Error(),
	})

	if err != nil {
		panic(err)
	}

	w.Write(rw, string(errorDto), code)
}

func (w *ResponseWriter) Write(rw http.ResponseWriter, body string, httpCode int) {
	rw.Header().Set("Content-Type", "application/json")
	http.Error(rw, body, httpCode)
}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{}
}
