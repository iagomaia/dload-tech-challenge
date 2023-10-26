package usercontrollersfactories

import (
	userservicefactories "github.com/iagomaia/dload-tech-challenge/internal/factories/usecases/user"
	usercontrollers "github.com/iagomaia/dload-tech-challenge/internal/presentation/controllers/user"
	presentation "github.com/iagomaia/dload-tech-challenge/internal/presentation/protocols"
)

var deleteMeController presentation.IHandler

func GetDeleteMeController() presentation.IHandler {
	if deleteMeController == nil {
		deleteMeController = &usercontrollers.DeleteMeController{
			UseCase: userservicefactories.GetDeleteMeUseCase(),
		}
	}

	return deleteMeController
}
