package userservices

import (
	"context"
	"net/http"

	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
	utilsmodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	hashprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/hash"
	userservicesprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

var (
	_ usecases.ISignupUser = (*SignupUserUseCase)(nil)
)

type SignupUserUseCase struct {
	CreateUserRepository userservicesprotocols.ICreateUserRepository
	HashAdapter          hashprotocols.IHash
	ctx                  context.Context
}

func (s *SignupUserUseCase) WithCtx(ctx context.Context) usecases.ISignupUser {
	return &SignupUserUseCase{
		CreateUserRepository: s.CreateUserRepository,
		HashAdapter:          s.HashAdapter,
		ctx:                  ctx,
	}
}

func (s *SignupUserUseCase) Signup(dto *usecases.SignupUserDto) (*usermodels.User, error) {
	hash, err := s.HashAdapter.MakeHash(dto.Password)
	if err != nil {
		cErr := utilsmodels.CustomError{
			Status: http.StatusInternalServerError,
			Message: "Error hashing password",
			OriginalError: err,
		}
		return nil,  cErr
	}
	dto.Password = hash
	dataDto := mapDomainDtoToDataDto(dto)
	return s.CreateUserRepository.WithCtx(s.ctx).Create(dataDto)
}

func mapDomainDtoToDataDto(dto *usecases.SignupUserDto) *userservicesprotocols.CreateUserDto {
	return &userservicesprotocols.CreateUserDto{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
