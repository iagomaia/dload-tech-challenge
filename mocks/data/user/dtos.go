package dataDtosMocks

import (
	userserviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

func GetCreateUserDtoMock() *userserviceprotocols.CreateUserDto {
	return &userserviceprotocols.CreateUserDto{
		Name:     "some-name",
		Email:    "some-email",
		Password: "some-password",
	}
}

func GetUpdateUserDtoMock() *userserviceprotocols.UpdateUserDto {
	return &userserviceprotocols.UpdateUserDto{
		Name: "some-name",
		ID:   "some-id",
	}
}
