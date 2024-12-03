package factory

import (
	"encoding/json"
	"io"
	"it-sloth/user.api/internal/dto"
)

type DtoFactory struct{
}

func (f *DtoFactory) UserCreateDto(rc io.ReadCloser) (dto.UserCreateRequest, error) {
	var dto dto.UserCreateRequest
	err := json.NewDecoder(rc).Decode(&dto)

	if err != nil {
		return dto, err
	}

	return dto, nil
}

func (f *DtoFactory) UserCreateResponseDto(guid string) (dto.UserCreateResponse) {
	return dto.UserCreateResponse{
		Guid: guid,
	}
}

func NewDtoFactory() *DtoFactory {
	return &DtoFactory{}
}