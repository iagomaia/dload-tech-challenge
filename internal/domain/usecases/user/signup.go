package userusecases

import (
	"context"

	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
)

type SignupUserDto struct {
	Name     string
	Email    string
	Password string
}

type ISignupUser interface {
	Signup(dto *SignupUserDto) (*usermodels.User, error)
	WithCtx(ctx context.Context) ISignupUser
}
