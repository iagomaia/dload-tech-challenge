package userservices

import (
	"context"
	"net/http"

	"github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	hashprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/hash"
	jwtprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/jwt"
	userservicesprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

var (
	_ usecases.ILogin = (*LoginUseCase)(nil)
)

type LoginUseCase struct {
	GetUserByEmailRepository userservicesprotocols.IGetUserByEmailRepository
	HashAdapter              hashprotocols.IHash
	JwtAdapter               jwtprotocols.IJwt
	ctx                      context.Context
}

func (s *LoginUseCase) WithCtx(ctx context.Context) usecases.ILogin {
	return &LoginUseCase{
		GetUserByEmailRepository: s.GetUserByEmailRepository,
		HashAdapter:              s.HashAdapter,
		JwtAdapter:               s.JwtAdapter,
		ctx:                      ctx,
	}
}

func (s *LoginUseCase) Login(dto *usecases.LoginDto) (string, error) {
	cErr := utils.CustomError{
		Status:  http.StatusUnauthorized,
		Message: "Email/Password doesn't match",
	}

	user, err := s.GetUserByEmailRepository.WithCtx(s.ctx).Get(dto.Email)
	if err != nil {
		return "", cErr
	}
	err = s.HashAdapter.CompareHash(user.Password, dto.Password)
	if err != nil {
		return "", cErr
	}
	claims := map[string]any{
		"name":  user.Name,
		"email": user.Email,
	}
	jwt, err := s.JwtAdapter.Generate(user.Id, claims)
	if err != nil {
		return "", cErr
	}
	return jwt, nil
}
