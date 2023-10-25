package mocks

import (
	domain "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
)

func GetCreateUserDomainDtoMock() *domain.SignupUserDto {
	return &domain.SignupUserDto{
		Name:     "some-name",
		Email:    "some-email",
		Password: "some-password",
	}
}
