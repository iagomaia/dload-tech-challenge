package userrepositorymocks

import (
	"context"

	models "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
	serviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
	mocks "github.com/iagomaia/dload-tech-challenge/mocks/domain/models"
)

var (
	_ serviceprotocols.IGetUserByIDRepository = (*GetUserByIDRepositoryMock)(nil)
)

type GetUserByIDRepositoryMock struct {
	GetMethodCalledTimes int
	GetMethodCalledWith  string
	GetMethodError       error
	GetMethodReturn      *models.User
}

func (r *GetUserByIDRepositoryMock) Init() {
	r.GetMethodCalledTimes = 0
	r.GetMethodCalledWith = ""
	r.GetMethodError = nil
	r.GetMethodReturn = mocks.GetUserModelMock()
}

func (r *GetUserByIDRepositoryMock) WithCtx(ctx context.Context) serviceprotocols.IGetUserByIDRepository {
	return r
}

func (r *GetUserByIDRepositoryMock) Get(id string) (*models.User, error) {
	r.GetMethodCalledTimes++
	r.GetMethodCalledWith = id
	return r.GetMethodReturn, r.GetMethodError
}
