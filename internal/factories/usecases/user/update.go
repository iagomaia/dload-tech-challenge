package userservicefactories

import (
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	userrepositoriesfactories "github.com/iagomaia/dload-tech-challenge/internal/factories/repositories/user"
	userservices "github.com/iagomaia/dload-tech-challenge/internal/services/usecases/user"
)

var updateUserUseCase usecases.IUpdateUser

func GetUpdateUserUseCase() usecases.IUpdateUser {
	if updateUserUseCase == nil {
		updateUserUseCase = &userservices.UpdateUserUseCase{
			UpdateUserRepository: userrepositoriesfactories.GetUpdateUserRepository(),
		}
	}

	return updateUserUseCase
}
