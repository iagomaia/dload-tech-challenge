package userrepositorymocks

import (
	"context"

	models "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
	serviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
	mocks "github.com/iagomaia/dload-tech-challenge/mocks/domain/models"
)

var (
	_ serviceprotocols.IGetUserByEmailRepository = (*GetUserByEmailRepositoryMock)(nil)
)

type GetUserByEmailRepositoryMock struct {
	GetMethodCalledTimes int
	GetMethodError       error
	GetMethodReturn      *models.User
}

func (r *GetUserByEmailRepositoryMock) Init() {
	r.GetMethodCalledTimes = 0
	r.GetMethodError = nil
	r.GetMethodReturn = mocks.GetUserModelMock()
}

func (r *GetUserByEmailRepositoryMock) WithCtx(ctx context.Context) serviceprotocols.IGetUserByEmailRepository {
	return r
}

func (r *GetUserByEmailRepositoryMock) Get(email string) (*models.User, error) {
	r.GetMethodCalledTimes++
	return r.GetMethodReturn, r.GetMethodError
}
