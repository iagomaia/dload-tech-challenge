package userrepositoriesfactories

import (
	userrepositories "github.com/iagomaia/dload-tech-challenge/internal/infra/repositories/user"
	userserviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

var updateUserRepository userserviceprotocols.IUpdateUserRepository

func GetUpdateUserRepository() userserviceprotocols.IUpdateUserRepository {
	if updateUserRepository == nil {
		updateUserRepository = &userrepositories.UpdateUserRepository{}
		updateUserRepository.Init()
	}

	return updateUserRepository
}
