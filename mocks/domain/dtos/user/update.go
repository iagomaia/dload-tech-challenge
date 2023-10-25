package mocks

import (
	domain "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
)

func GetUpdateUserDomainDtoMock() *domain.UpdateUserDto {
	return &domain.UpdateUserDto{
		Name: "some-name",
		ID:   "some-id",
	}
}
