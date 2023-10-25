package userservices_test

import (
	"context"
	"net/http"
	"testing"

	utils "github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	userservices "github.com/iagomaia/dload-tech-challenge/internal/services/usecases/user"
	domainDtosMocks "github.com/iagomaia/dload-tech-challenge/mocks/domain/dtos/user"
	adaptersmocks "github.com/iagomaia/dload-tech-challenge/mocks/infra/adapters"
	userrepositorymocks "github.com/iagomaia/dload-tech-challenge/mocks/infra/repositories/user"
)

type UserLoginSutTypes struct {
	UseCase                      *userservices.LoginUseCase
	GetUserByEmailRepositoryMock *userrepositorymocks.GetUserByEmailRepositoryMock
	BCryptAdapterMock            *adaptersmocks.BCryptAdapterMock
	JwtAdapterMock               *adaptersmocks.JwtAdapterMock
}

func GetUserLoginSutDependencies() *UserLoginSutTypes {
	getUserByEmailRepositoryMock := &userrepositorymocks.GetUserByEmailRepositoryMock{}
	bCryptAdapterMock := &adaptersmocks.BCryptAdapterMock{}
	jwtAdapterMock := &adaptersmocks.JwtAdapterMock{}
	getUserByEmailRepositoryMock.Init()
	bCryptAdapterMock.Init()
	jwtAdapterMock.Init()

	useCase := &userservices.LoginUseCase{
		GetUserByEmailRepository: getUserByEmailRepositoryMock,
		HashAdapter:              bCryptAdapterMock,
		JwtAdapter:               jwtAdapterMock,
	}

	return &UserLoginSutTypes{
		UseCase:                      useCase,
		GetUserByEmailRepositoryMock: getUserByEmailRepositoryMock,
		BCryptAdapterMock:            bCryptAdapterMock,
		JwtAdapterMock:               jwtAdapterMock,
	}
}

func Test_Login(t *testing.T) {
	t.Run("should return jwt", func(t *testing.T) {
		sut := GetUserLoginSutDependencies()
		_, err := sut.UseCase.WithCtx(
			context.Background(),
		).Login(domainDtosMocks.GetloginDomainDtoMock())

		if err != nil {
			t.Errorf("failed: %v", err)
		}
		if sut.GetUserByEmailRepositoryMock.GetMethodCalledTimes != 1 {
			t.Error("Get user by email repository was not called")
		}
		if sut.BCryptAdapterMock.CompareHashMethodCalledTimes != 1 {
			t.Error("CompareHash was not called")
		}
		if sut.JwtAdapterMock.GenerateMethodCalledTimes != 1 {
			t.Error("Generate JWT was not called")
		}
	})
	t.Run("should return error if failed to get user", func(t *testing.T) {
		sut := GetUserLoginSutDependencies()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		loginErr := utils.CustomError{
			Status:  http.StatusUnauthorized,
			Message: "Email/Password doesn't match",
		}
		sut.GetUserByEmailRepositoryMock.GetMethodError = cErr
		sut.GetUserByEmailRepositoryMock.GetMethodReturn = nil

		jwt, err := sut.UseCase.WithCtx(
			context.Background(),
		).Login(domainDtosMocks.GetloginDomainDtoMock())

		if err == nil {
			t.Errorf("expected error: %v", loginErr)
		}
		if jwt != "" {
			t.Error("should return empty string if an error happens")
		}
		if sut.GetUserByEmailRepositoryMock.GetMethodCalledTimes != 1 {
			t.Error("Get user by email repository was not called")
		}
		if sut.BCryptAdapterMock.CompareHashMethodCalledTimes != 0 {
			t.Error("CompareHash should not have been called if failed to get user")
		}
		if sut.JwtAdapterMock.GenerateMethodCalledTimes != 0 {
			t.Error("Generate JWT should not have been called if failed to get user")
		}
	})
	t.Run("should return error if failed to compare password", func(t *testing.T) {
		sut := GetUserLoginSutDependencies()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		loginErr := utils.CustomError{
			Status:  http.StatusUnauthorized,
			Message: "Email/Password doesn't match",
		}
		sut.BCryptAdapterMock.CompareHashError = cErr

		jwt, err := sut.UseCase.WithCtx(
			context.Background(),
		).Login(domainDtosMocks.GetloginDomainDtoMock())

		if err == nil {
			t.Errorf("expected error: %v", loginErr)
		}
		if jwt != "" {
			t.Error("should return empty string if an error happens")
		}
		if sut.GetUserByEmailRepositoryMock.GetMethodCalledTimes != 1 {
			t.Error("Get user by email repository was not called")
		}
		if sut.BCryptAdapterMock.CompareHashMethodCalledTimes != 1 {
			t.Error("CompareHash wasn't called")
		}
		if sut.JwtAdapterMock.GenerateMethodCalledTimes != 0 {
			t.Error("Generate JWT should not have been called if failed to compare password")
		}
	})
	t.Run("should return error if failed to generate token", func(t *testing.T) {
		sut := GetUserLoginSutDependencies()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		loginErr := utils.CustomError{
			Status:  http.StatusUnauthorized,
			Message: "Email/Password doesn't match",
		}
		sut.JwtAdapterMock.GenerateMethodError = cErr
		sut.JwtAdapterMock.GenerateMethodReturn = ""

		jwt, err := sut.UseCase.WithCtx(
			context.Background(),
		).Login(domainDtosMocks.GetloginDomainDtoMock())

		if err == nil {
			t.Errorf("expected error: %v", loginErr)
		}
		if jwt != "" {
			t.Error("should return empty string if an error happens")
		}
		if sut.GetUserByEmailRepositoryMock.GetMethodCalledTimes != 1 {
			t.Error("Get user by email repository was not called")
		}
		if sut.BCryptAdapterMock.CompareHashMethodCalledTimes != 1 {
			t.Error("CompareHash wasn't called")
		}
		if sut.JwtAdapterMock.GenerateMethodCalledTimes != 1 {
			t.Error("Generate JWT wasn't called")
		}
	})
}
