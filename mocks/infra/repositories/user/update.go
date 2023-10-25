package userrepositorymocks

import (
	"context"

	models "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
	serviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
	mocks "github.com/iagomaia/dload-tech-challenge/mocks/domain/models"
)

var (
	_ serviceprotocols.IUpdateUserRepository = (*UpdateUserRepositoryMock)(nil)
)

type UpdateUserRepositoryMock struct {
	UpdateMethodCalledTimes int
	UpdateMethodReturn      *models.User
	UpdateMethodError       error
}

func (r *UpdateUserRepositoryMock) Init() {
	r.UpdateMethodCalledTimes = 0
	r.UpdateMethodError = nil
	r.UpdateMethodReturn = mocks.GetUserModelMock()
}

func (r *UpdateUserRepositoryMock) WithCtx(ctx context.Context) serviceprotocols.IUpdateUserRepository {
	return r
}

func (r *UpdateUserRepositoryMock) Update(dto *serviceprotocols.UpdateUserDto) (*models.User, error) {
	r.UpdateMethodCalledTimes++
	return r.UpdateMethodReturn, r.UpdateMethodError
}
