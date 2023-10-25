package usercontrollersfactories

import (
	userservicefactories "github.com/iagomaia/dload-tech-challenge/internal/factories/usecases/user"
	usercontrollers "github.com/iagomaia/dload-tech-challenge/internal/presentation/controllers/user"
	presentation "github.com/iagomaia/dload-tech-challenge/internal/presentation/protocols"
)

var userSignupController presentation.IHandler

func GetUserSignupController() presentation.IHandler {
	if userSignupController == nil {
		userSignupController = &usercontrollers.CreateUserController{
			UseCase: userservicefactories.GetUserSignupUseCase(),
		}
	}

	return userSignupController
}
