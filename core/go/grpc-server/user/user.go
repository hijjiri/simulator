package user

import (
	"context"
)

func LockUser(ctx context.Context, uid uint32) (*UserData, error) {
	return r.LockUser(ctx, uid)
}

func UnlockUser(ctx context.Context, uid uint32) (*UserData, error) {
	return r.UnlockUser(ctx, uid)
}
