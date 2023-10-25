package userservices

import (
	"context"

	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	userservicesprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

var (
	_ usecases.IUpdateUser = (*UpdateUserUseCase)(nil)
)

type UpdateUserUseCase struct {
	UpdateUserRepository userservicesprotocols.IUpdateUserRepository
	ctx                  context.Context
}

func (s *UpdateUserUseCase) WithCtx(ctx context.Context) usecases.IUpdateUser {
	return &UpdateUserUseCase{
		UpdateUserRepository: s.UpdateUserRepository,
		ctx:                  ctx,
	}
}

func (s *UpdateUserUseCase) Update(dto *usecases.UpdateUserDto) (*usermodels.User, error) {
	dataDto := mapUpdateUserDomainDtoToDataDto(dto)
	return s.UpdateUserRepository.WithCtx(s.ctx).Update(dataDto)
}

func mapUpdateUserDomainDtoToDataDto(dto *usecases.UpdateUserDto) *userservicesprotocols.UpdateUserDto {
	return &userservicesprotocols.UpdateUserDto{
		Name: dto.Name,
		ID:   dto.ID,
	}
}
