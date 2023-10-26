package userservicefactories

import (
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	userrepositoriesfactories "github.com/iagomaia/dload-tech-challenge/internal/factories/repositories/user"
	userservices "github.com/iagomaia/dload-tech-challenge/internal/services/usecases/user"
)

var deleteMeUseCase usecases.IDeleteMe

func GetDeleteMeUseCase() usecases.IDeleteMe {
	if deleteMeUseCase == nil {
		deleteMeUseCase = &userservices.DeleteMeUseCase{
			SoftDeleteUserRepository: userrepositoriesfactories.GetSoftDeleteUserRepository(),
		}
	}

	return deleteMeUseCase
}
