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

type UserSignupSutTypes struct {
	UseCase                  *userservices.SignupUserUseCase
	CreateUserRepositoryMock *userrepositorymocks.CreateUserRepositoryMock
	BCryptAdapterMock        *adaptersmocks.BCryptAdapterMock
}

func GetUserSignupSutDependencies() *UserSignupSutTypes {
	createUserRepositoryMock := &userrepositorymocks.CreateUserRepositoryMock{}
	bCryptAdapterMock := &adaptersmocks.BCryptAdapterMock{}
	createUserRepositoryMock.Init()
	bCryptAdapterMock.Init()

	useCase := &userservices.SignupUserUseCase{
		CreateUserRepository: createUserRepositoryMock,
		HashAdapter:          bCryptAdapterMock,
	}

	return &UserSignupSutTypes{
		UseCase:                  useCase,
		CreateUserRepositoryMock: createUserRepositoryMock,
		BCryptAdapterMock:        bCryptAdapterMock,
	}
}

func Test_Signup(t *testing.T) {
	t.Run("should create a user calling the repository and hashing password", func(t *testing.T) {
		sut := GetUserSignupSutDependencies()
		_, err := sut.UseCase.WithCtx(
			context.Background(),
		).Signup(domainDtosMocks.GetCreateUserDomainDtoMock())

		if err != nil {
			t.Errorf("failed: %v", err)
		}
		if sut.CreateUserRepositoryMock.CreateMethodCalledTimes != 1 {
			t.Error("Create user repository was not called")
		}
		if sut.BCryptAdapterMock.MakeHashMethodCalledTimes != 1 {
			t.Error("MakeHash was not called")
		}
	})
	t.Run("should return error if failed to create user", func(t *testing.T) {
		sut := GetUserSignupSutDependencies()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		sut.CreateUserRepositoryMock.CreateMethodError = cErr
		sut.CreateUserRepositoryMock.CreateMethodReturn = nil

		user, err := sut.UseCase.WithCtx(
			context.Background(),
		).Signup(domainDtosMocks.GetCreateUserDomainDtoMock())

		if err == nil {
			t.Errorf("expected error: %v", cErr)
		}
		if user != nil {
			t.Error("should return nil user if an error happens")
		}
		if sut.BCryptAdapterMock.MakeHashMethodCalledTimes != 1 {
			t.Error("MakeHash was not called")
		}
		if sut.CreateUserRepositoryMock.CreateMethodCalledTimes != 1 {
			t.Error("create user repository not called")
		}
	})
	t.Run("should return error if failed to hash password", func(t *testing.T) {
		sut := GetUserSignupSutDependencies()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		sut.BCryptAdapterMock.MakeHashError = cErr
		sut.BCryptAdapterMock.MakeHashReturn = ""

		user, err := sut.UseCase.WithCtx(
			context.Background(),
		).Signup(domainDtosMocks.GetCreateUserDomainDtoMock())

		if err == nil {
			t.Errorf("expected error: %v", cErr)
		}
		if user != nil {
			t.Error("should return nil user if an error happens")
		}
		if sut.BCryptAdapterMock.MakeHashMethodCalledTimes != 1 {
			t.Error("MakeHash was not called")
		}
		if sut.CreateUserRepositoryMock.CreateMethodCalledTimes != 0 {
			t.Error("create user must not be called if hash failed")
		}
	})
}
