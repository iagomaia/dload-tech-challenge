package userservicesprotocols

import (
	"context"

	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
)

type IGetUserByEmailRepository interface {
	Get(email string) (*usermodels.User, error)
	WithCtx(ctx context.Context) IGetUserByEmailRepository
	Init()
}
