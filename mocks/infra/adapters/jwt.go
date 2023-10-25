package adaptersmocks

import (
	"context"

	serviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/jwt"
)

var (
	_ serviceprotocols.IJwt = (*JwtAdapterMock)(nil)
)

type JwtAdapterMock struct {
	GenerateMethodCalledTimes int
	GenerateMethodReturn      string
	GenerateMethodError       error
	VerifyMethodCalledTimes   int
	VerifyMethodReturn        map[string]any
	VerifyMethodError         error
}

func (r *JwtAdapterMock) Init() {
	r.GenerateMethodCalledTimes = 0
	r.GenerateMethodError = nil
	r.GenerateMethodReturn = "some-token"
	r.VerifyMethodCalledTimes = 0
	r.VerifyMethodError = nil
	r.VerifyMethodReturn = make(map[string]any)
}

func (r *JwtAdapterMock) WithCtx(ctx context.Context) serviceprotocols.IJwt {
	return r
}

func (r *JwtAdapterMock) Generate(userId string, claims map[string]any) (string, error) {
	r.GenerateMethodCalledTimes++
	return r.GenerateMethodReturn, r.GenerateMethodError
}

func (r *JwtAdapterMock) Verify(userToken string) (map[string]any, error) {
	r.VerifyMethodCalledTimes++
	return r.VerifyMethodReturn, r.VerifyMethodError
}
