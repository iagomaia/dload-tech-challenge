package userservices

import (
	"context"

	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
	utilsmodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	userservicesprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

var (
	_ usecases.IGetMe = (*GetMeUseCase)(nil)
)

type GetMeUseCase struct {
	GetUserByIDRepository userservicesprotocols.IGetUserByIDRepository
	ctx                   context.Context
}

func (s *GetMeUseCase) WithCtx(ctx context.Context) usecases.IGetMe {
	return &GetMeUseCase{
		GetUserByIDRepository: s.GetUserByIDRepository,
		ctx:                   ctx,
	}
}

func (s *GetMeUseCase) Get() (*usermodels.User, error) {
	userId := s.ctx.Value(utilsmodels.UserIDContextKey).(string)
	return s.GetUserByIDRepository.WithCtx(s.ctx).Get(userId)
}
