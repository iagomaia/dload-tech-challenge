package userservicesprotocols

import (
	"context"

	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
)

type UpdateUserDto struct {
	ID   string
	Name string
}

type IUpdateUserRepository interface {
	Update(dto *UpdateUserDto) (*usermodels.User, error)
	WithCtx(ctx context.Context) IUpdateUserRepository
	Init()
}
