package usercontrollersfactories

import (
	userservicefactories "github.com/iagomaia/dload-tech-challenge/internal/factories/usecases/user"
	usercontrollers "github.com/iagomaia/dload-tech-challenge/internal/presentation/controllers/user"
	presentation "github.com/iagomaia/dload-tech-challenge/internal/presentation/protocols"
)

var meController presentation.IHandler

func GetMeController() presentation.IHandler {
	if meController == nil {
		meController = &usercontrollers.MeController{
			UseCase: userservicefactories.GetGetMeUseCase(),
		}
	}

	return meController
}
