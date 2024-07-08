package user

import (
	"context"
)

var r Repository

type Repository interface {
	LockUser(ctx context.Context, uid uint32) (*UserData, error)
	UnlockUser(ctx context.Context, uid uint32) (*UserData, error)
}

func RegisterRepository(repository Repository) {
	r = repository
}
