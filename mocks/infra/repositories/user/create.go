package userrepositorymocks

import (
	"context"

	models "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
	serviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
	mocks "github.com/iagomaia/dload-tech-challenge/mocks/domain/models"
)

var (
	_ serviceprotocols.ICreateUserRepository = (*CreateUserRepositoryMock)(nil)
)

type CreateUserRepositoryMock struct {
	CreateMethodCalledTimes int
	CreateMethodReturn      *models.User
	CreateMethodError       error
}

func (r *CreateUserRepositoryMock) Init() {
	r.CreateMethodCalledTimes = 0
	r.CreateMethodError = nil
	r.CreateMethodReturn = mocks.GetUserModelMock()
}

func (r *CreateUserRepositoryMock) WithCtx(ctx context.Context) serviceprotocols.ICreateUserRepository {
	return r
}

func (r *CreateUserRepositoryMock) Create(dto *serviceprotocols.CreateUserDto) (*models.User, error) {
	r.CreateMethodCalledTimes++
	return r.CreateMethodReturn, r.CreateMethodError
}
