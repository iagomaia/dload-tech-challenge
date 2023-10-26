package userrepositoriesfactories

import (
	userrepositories "github.com/iagomaia/dload-tech-challenge/internal/infra/repositories/user"
	userserviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

var softDeleteUserRepository userserviceprotocols.ISoftDeleteUserRepository

func GetSoftDeleteUserRepository() userserviceprotocols.ISoftDeleteUserRepository {
	if softDeleteUserRepository == nil {
		softDeleteUserRepository = &userrepositories.SoftDeleteUserRepository{}
		softDeleteUserRepository.Init()
	}

	return softDeleteUserRepository
}
