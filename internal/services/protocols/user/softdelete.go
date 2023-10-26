package userservicesprotocols

import (
	"context"
)

type ISoftDeleteUserRepository interface {
	SoftDelete(id string) error
	WithCtx(ctx context.Context) ISoftDeleteUserRepository
	Init()
}
