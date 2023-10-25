package userrepositoriesfactories

import (
	userrepositories "github.com/iagomaia/dload-tech-challenge/internal/infra/repositories/user"
	userserviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

var getUserByIDRepository userserviceprotocols.IGetUserByIDRepository

func GetGetUserByIDRepository() userserviceprotocols.IGetUserByIDRepository {
	if getUserByIDRepository == nil {
		getUserByIDRepository = &userrepositories.GetUserByIDRepository{}
		getUserByIDRepository.Init()
	}

	return getUserByIDRepository
}
