package userrepositoriesfactories

import (
	userrepositories "github.com/iagomaia/dload-tech-challenge/internal/infra/repositories/user"
	userserviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
)

var createUserRepository userserviceprotocols.ICreateUserRepository

func GetCreateUserRepository() userserviceprotocols.ICreateUserRepository {
	if createUserRepository == nil {
		createUserRepository = &userrepositories.CreateUserRepository{}
		createUserRepository.Init()
	}

	return createUserRepository
}
