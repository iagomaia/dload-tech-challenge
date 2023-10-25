package mocks

import (
	domain "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
)

func GetloginDomainDtoMock() *domain.LoginDto {
	return &domain.LoginDto{
		Email:    "some-email",
		Password: "some-password",
	}
}
