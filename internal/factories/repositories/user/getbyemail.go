package userrepositoriesfactories

import (
	userrepositories "github.com/iagomaia/dload-tech-challenge/internal/infra/repositories/user"
	userserviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

var getUserByEmailRepository userserviceprotocols.IGetUserByEmailRepository

func GetGetUserByEmailRepository() userserviceprotocols.IGetUserByEmailRepository {
	if getUserByEmailRepository == nil {
		getUserByEmailRepository = &userrepositories.GetUserByEmailRepository{}
		getUserByEmailRepository.Init()
	}

	return getUserByEmailRepository
}
