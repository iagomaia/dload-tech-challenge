package adaptersmocks

import (
	"context"

	serviceprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/hash"
)

var (
	_ serviceprotocols.IHash = (*BCryptAdapterMock)(nil)
)

type BCryptAdapterMock struct {
	MakeHashMethodCalledTimes    int
	CompareHashMethodCalledTimes int
	MakeHashReturn               string
	MakeHashError                error
	CompareHashError             error
}

func (r *BCryptAdapterMock) Init() {
	r.MakeHashMethodCalledTimes = 0
	r.CompareHashMethodCalledTimes = 0
	r.MakeHashError = nil
	r.CompareHashError = nil
	r.MakeHashReturn = "hashed-password"
}

func (r *BCryptAdapterMock) WithCtx(ctx context.Context) serviceprotocols.IHash {
	return r
}

func (r *BCryptAdapterMock) MakeHash(input string) (string, error) {
	r.MakeHashMethodCalledTimes++
	return r.MakeHashReturn, r.MakeHashError
}

func (r *BCryptAdapterMock) CompareHash(source, target string) error {
	r.CompareHashMethodCalledTimes++
	return r.CompareHashError
}
