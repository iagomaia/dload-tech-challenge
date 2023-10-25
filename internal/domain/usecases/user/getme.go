package userusecases

import (
	"context"

	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
)

type IGetMe interface {
	Get() (*usermodels.User, error)
	WithCtx(ctx context.Context) IGetMe
}
