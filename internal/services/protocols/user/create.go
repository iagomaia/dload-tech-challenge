package userservicesprotocols

import (
	"context"

	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
)

type CreateUserDto struct {
	Name     string
	Email    string
	Password string
}

type ICreateUserRepository interface {
	Create(dto *CreateUserDto) (*usermodels.User, error)
	WithCtx(ctx context.Context) ICreateUserRepository
	Init()
}
