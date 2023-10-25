package userservicefactories

import (
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	adaptersfactories "github.com/iagomaia/dload-tech-challenge/internal/factories/adapters"
	userrepositoriesfactories "github.com/iagomaia/dload-tech-challenge/internal/factories/repositories/user"
	userservices "github.com/iagomaia/dload-tech-challenge/internal/services/usecases/user"
)

var userSignupUseCase usecases.ISignupUser

func GetUserSignupUseCase() usecases.ISignupUser {
	if userSignupUseCase == nil {
		userSignupUseCase = &userservices.SignupUserUseCase{
			CreateUserRepository: userrepositoriesfactories.GetCreateUserRepository(),
			HashAdapter:          adaptersfactories.GetbcryptAdapter(),
		}
	}

	return userSignupUseCase
}
