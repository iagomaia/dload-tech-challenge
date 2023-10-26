package userservices

import (
	"context"

	utilsmodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	userservicesprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

var (
	_ usecases.IDeleteMe = (*DeleteMeUseCase)(nil)
)

type DeleteMeUseCase struct {
	SoftDeleteUserRepository userservicesprotocols.ISoftDeleteUserRepository
	ctx                      context.Context
}

func (s *DeleteMeUseCase) WithCtx(ctx context.Context) usecases.IDeleteMe {
	return &DeleteMeUseCase{
		SoftDeleteUserRepository: s.SoftDeleteUserRepository,
		ctx:                      ctx,
	}
}

func (s *DeleteMeUseCase) Delete() error {
	userId := s.ctx.Value(utilsmodels.UserIDContextKey).(string)
	return s.SoftDeleteUserRepository.WithCtx(s.ctx).SoftDelete(userId)
}
