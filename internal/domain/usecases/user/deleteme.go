package userusecases

import (
	"context"
)

type IDeleteMe interface {
	Delete() error
	WithCtx(ctx context.Context) IDeleteMe
}
