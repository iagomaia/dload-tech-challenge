package userservicefactories

import (
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	adaptersfactories "github.com/iagomaia/dload-tech-challenge/internal/factories/adapters"
	userrepositoriesfactories "github.com/iagomaia/dload-tech-challenge/internal/factories/repositories/user"
	userservices "github.com/iagomaia/dload-tech-challenge/internal/services/usecases/user"
)

var userLoginUseCase usecases.ILogin

func GetUserLoginUseCase() usecases.ILogin {
	if userLoginUseCase == nil {
		userLoginUseCase = &userservices.LoginUseCase{
			GetUserByEmailRepository: userrepositoriesfactories.GetGetUserByEmailRepository(),
			JwtAdapter:               adaptersfactories.GetJwtAdapter(),
			HashAdapter:              adaptersfactories.GetbcryptAdapter(),
		}
	}

	return userLoginUseCase
}
