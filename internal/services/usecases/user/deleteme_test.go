package userservices_test

import (
	"context"
	"net/http"
	"testing"

	utils "github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	userservices "github.com/iagomaia/dload-tech-challenge/internal/services/usecases/user"
	userrepositorymocks "github.com/iagomaia/dload-tech-challenge/mocks/infra/repositories/user"
)

type DeleteMeSutTypes struct {
	UseCase                      *userservices.DeleteMeUseCase
	SoftDeleteUserRepositoryMock *userrepositorymocks.SoftDeleteUserRepositoryMock
}

func GetDeleteMeSutDependencies() *DeleteMeSutTypes {
	softDeleteUserRepositoryMock := &userrepositorymocks.SoftDeleteUserRepositoryMock{}
	softDeleteUserRepositoryMock.Init()

	useCase := &userservices.DeleteMeUseCase{
		SoftDeleteUserRepository: softDeleteUserRepositoryMock,
	}

	return &DeleteMeSutTypes{
		UseCase:                      useCase,
		SoftDeleteUserRepositoryMock: softDeleteUserRepositoryMock,
	}
}

func Test_DeleteMe(t *testing.T) {
	t.Run("should delete signed in in user", func(t *testing.T) {
		sut := GetDeleteMeSutDependencies()
		ctx := context.WithValue(context.Background(), utils.UserIDContextKey, "logged-in-id")
		err := sut.UseCase.WithCtx(
			ctx,
		).Delete()

		if err != nil {
			t.Errorf("failed: %v", err)
		}
		if sut.SoftDeleteUserRepositoryMock.SoftDeleteMethodCalledTimes != 1 {
			t.Error("Delete user by ID repository was not called")
		}
		if sut.SoftDeleteUserRepositoryMock.SoftDeleteMethodCalledWith != "logged-in-id" {
			t.Error("Delete user by ID repostory was not called with logged in user id value")
		}
	})
	t.Run("should return error if failed to soft delete user", func(t *testing.T) {
		sut := GetDeleteMeSutDependencies()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		sut.SoftDeleteUserRepositoryMock.SoftDeleteMethodError = cErr

		ctx := context.WithValue(context.Background(), utils.UserIDContextKey, "logged-in-id")
		err := sut.UseCase.WithCtx(
			ctx,
		).Delete()

		if err == nil {
			t.Errorf("expected error: %v", cErr)
		}
		if sut.SoftDeleteUserRepositoryMock.SoftDeleteMethodCalledTimes != 1 {
			t.Error("Delete user by ID repository was not called")
		}
		if sut.SoftDeleteUserRepositoryMock.SoftDeleteMethodCalledWith != "logged-in-id" {
			t.Error("Delete user by ID repostory was not called with logged in user id value")
		}
	})
}
