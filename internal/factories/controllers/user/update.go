package usercontrollersfactories

import (
	userservicefactories "github.com/iagomaia/dload-tech-challenge/internal/factories/usecases/user"
	usercontrollers "github.com/iagomaia/dload-tech-challenge/internal/presentation/controllers/user"
	presentation "github.com/iagomaia/dload-tech-challenge/internal/presentation/protocols"
)

var updateUserController presentation.IHandler

func GetUpdateUserController() presentation.IHandler {
	if updateUserController == nil {
		updateUserController = &usercontrollers.UpdateUserController{
			UseCase: userservicefactories.GetUpdateUserUseCase(),
		}
	}

	return updateUserController
}
