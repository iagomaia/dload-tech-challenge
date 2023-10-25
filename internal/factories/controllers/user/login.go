package usercontrollersfactories

import (
	userservicefactories "github.com/iagomaia/dload-tech-challenge/internal/factories/usecases/user"
	usercontrollers "github.com/iagomaia/dload-tech-challenge/internal/presentation/controllers/user"
	presentation "github.com/iagomaia/dload-tech-challenge/internal/presentation/protocols"
)

var userLoginController presentation.IHandler

func GetUserLoginController() presentation.IHandler {
	if userLoginController == nil {
		userLoginController = &usercontrollers.LoginController{
			UseCase: userservicefactories.GetUserLoginUseCase(),
		}
	}

	return userLoginController
}
