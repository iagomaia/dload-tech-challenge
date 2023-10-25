package userusecases

import (
	"context"
)

type LoginDto struct {
	Email    string
	Password string
}

type ILogin interface {
	Login(dto *LoginDto) (string, error)
	WithCtx(ctx context.Context) ILogin
}
