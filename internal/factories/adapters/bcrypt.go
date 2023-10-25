package adaptersfactories

import (
	adapters "github.com/iagomaia/dload-tech-challenge/internal/infra/adapters"
	hashprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/hash"
)

var bcryptAdapter hashprotocols.IHash

func GetbcryptAdapter() hashprotocols.IHash {
	if bcryptAdapter == nil {
		bcryptAdapter = &adapters.BCryptAdapter{}
	}

	return bcryptAdapter
}
