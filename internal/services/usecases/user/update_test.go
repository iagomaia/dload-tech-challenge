package userservices_test

import (
	"context"
	"net/http"
	"testing"

	utils "github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	userservices "github.com/iagomaia/dload-tech-challenge/internal/services/usecases/user"
	domainDtosMocks "github.com/iagomaia/dload-tech-challenge/mocks/domain/dtos/user"
	userrepositorymocks "github.com/iagomaia/dload-tech-challenge/mocks/infra/repositories/user"
)

type UpdateUserSutTypes struct {
	UseCase                  *userservices.UpdateUserUseCase
	UpdateUserRepositoryMock *userrepositorymocks.UpdateUserRepositoryMock
}

func GetUpdateUserSutDependencies() *UpdateUserSutTypes {
	updateUserRepositoryMock := &userrepositorymocks.UpdateUserRepositoryMock{}
	updateUserRepositoryMock.Init()

	useCase := &userservices.UpdateUserUseCase{
		UpdateUserRepository: updateUserRepositoryMock,
	}

	return &UpdateUserSutTypes{
		UseCase:                  useCase,
		UpdateUserRepositoryMock: updateUserRepositoryMock,
	}
}

func Test_Update(t *testing.T) {
	t.Run("should update a user calling the repository", func(t *testing.T) {
		sut := GetUpdateUserSutDependencies()
		_, err := sut.UseCase.WithCtx(
			context.Background(),
		).Update(domainDtosMocks.GetUpdateUserDomainDtoMock())

		if err != nil {
			t.Errorf("failed: %v", err)
		}
		if sut.UpdateUserRepositoryMock.UpdateMethodCalledTimes != 1 {
			t.Error("Update user repository was not called")
		}
	})
	t.Run("should return error if failed to update user", func(t *testing.T) {
		sut := GetUpdateUserSutDependencies()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		sut.UpdateUserRepositoryMock.UpdateMethodError = cErr
		sut.UpdateUserRepositoryMock.UpdateMethodReturn = nil

		user, err := sut.UseCase.WithCtx(
			context.Background(),
		).Update(domainDtosMocks.GetUpdateUserDomainDtoMock())

		if err == nil {
			t.Errorf("expected error: %v", cErr)
		}
		if user != nil {
			t.Error("should return nil user if an error happens")
		}
		if sut.UpdateUserRepositoryMock.UpdateMethodCalledTimes != 1 {
			t.Error("update user repository not called")
		}
	})
}
