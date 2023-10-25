package userservices_test

import (
	"context"
	"net/http"
	"testing"

	utils "github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	userservices "github.com/iagomaia/dload-tech-challenge/internal/services/usecases/user"
	userrepositorymocks "github.com/iagomaia/dload-tech-challenge/mocks/infra/repositories/user"
)

type GetMeSutTypes struct {
	UseCase                   *userservices.GetMeUseCase
	GetUserByIDRepositoryMock *userrepositorymocks.GetUserByIDRepositoryMock
}

func GetMeSutDependencies() *GetMeSutTypes {
	getUserByIDRepositoryMock := &userrepositorymocks.GetUserByIDRepositoryMock{}
	getUserByIDRepositoryMock.Init()

	useCase := &userservices.GetMeUseCase{
		GetUserByIDRepository: getUserByIDRepositoryMock,
	}

	return &GetMeSutTypes{
		UseCase:                   useCase,
		GetUserByIDRepositoryMock: getUserByIDRepositoryMock,
	}
}

func Test_GetMe(t *testing.T) {
	t.Run("should get siggned in user", func(t *testing.T) {
		sut := GetMeSutDependencies()
		ctx := context.WithValue(context.Background(), utils.UserIDContextKey, "logged-in-id")
		_, err := sut.UseCase.WithCtx(
			ctx,
		).Get()

		if err != nil {
			t.Errorf("failed: %v", err)
		}
		if sut.GetUserByIDRepositoryMock.GetMethodCalledTimes != 1 {
			t.Error("Get user by ID repository was not called")
		}
		if sut.GetUserByIDRepositoryMock.GetMethodCalledWith != "logged-in-id" {
			t.Error("Get user by ID repostory was not called with logged in user id value")
		}
	})
	t.Run("should return error if failed to get user", func(t *testing.T) {
		sut := GetMeSutDependencies()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		sut.GetUserByIDRepositoryMock.GetMethodError = cErr
		sut.GetUserByIDRepositoryMock.GetMethodReturn = nil

		ctx := context.WithValue(context.Background(), utils.UserIDContextKey, "logged-in-id")
		user, err := sut.UseCase.WithCtx(
			ctx,
		).Get()

		if err == nil {
			t.Errorf("expected error: %v", cErr)
		}
		if user != nil {
			t.Error("should return nil user if an error happens")
		}
		if sut.GetUserByIDRepositoryMock.GetMethodCalledTimes != 1 {
			t.Error("Get user by ID repository was not called")
		}
		if sut.GetUserByIDRepositoryMock.GetMethodCalledWith != "logged-in-id" {
			t.Error("Get user by ID repostory was not called with logged in user id value")
		}
	})
}
