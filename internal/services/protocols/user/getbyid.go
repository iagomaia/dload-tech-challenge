package userservicesprotocols

import (
	"context"

	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
)

type IGetUserByIDRepository interface {
	Get(id string) (*usermodels.User, error)
	WithCtx(ctx context.Context) IGetUserByIDRepository
	Init()
}
