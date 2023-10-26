package userrepositorymocks

import (
	"context"

	serviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

var (
	_ serviceprotocols.ISoftDeleteUserRepository = (*SoftDeleteUserRepositoryMock)(nil)
)

type SoftDeleteUserRepositoryMock struct {
	SoftDeleteMethodCalledTimes int
	SoftDeleteMethodCalledWith  string
	SoftDeleteMethodError       error
}

func (r *SoftDeleteUserRepositoryMock) Init() {
	r.SoftDeleteMethodCalledTimes = 0
	r.SoftDeleteMethodCalledWith = ""
	r.SoftDeleteMethodError = nil
}

func (r *SoftDeleteUserRepositoryMock) WithCtx(ctx context.Context) serviceprotocols.ISoftDeleteUserRepository {
	return r
}

func (r *SoftDeleteUserRepositoryMock) SoftDelete(id string) error {
	r.SoftDeleteMethodCalledTimes++
	r.SoftDeleteMethodCalledWith = id
	return r.SoftDeleteMethodError
}
