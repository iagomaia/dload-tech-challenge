package clientsfactories

import "github.com/iagomaia/dload-tech-challenge/internal/infra/repositories"

func GetMongoClient() *repositories.MongoClient {
	return new(repositories.MongoClient)
}
