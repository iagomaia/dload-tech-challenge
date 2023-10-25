package adaptersfactories

import (
	adapters "github.com/iagomaia/dload-tech-challenge/internal/infra/adapters"
	jwtprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/jwt"
)

var jwtAdapter jwtprotocols.IJwt

func GetJwtAdapter() jwtprotocols.IJwt {
	if jwtAdapter == nil {
		jwtAdapter = &adapters.JwtAdapter{}
	}

	return jwtAdapter
}
