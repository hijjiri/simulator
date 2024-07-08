package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/hijjiri/simulator/core/go/grpc-server/user"
)

type InMemoryRepository struct {
	mu    sync.Mutex
	users map[uint32]*user.UserData
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users: make(map[uint32]*user.UserData),
	}
}

func (repo *InMemoryRepository) LockUser(ctx context.Context, uid uint32) (*user.UserData, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	userData, exists := repo.users[uid]
	if !exists {
		userData = &user.UserData{Uid: uid}
		repo.users[uid] = userData
	}

	if userData.IsLocked {
		return nil, errors.New("user is already locked")
	}

	// ロック状態を示すためにis_lockedをtrueに設定
	userData.IsLocked = true
	return userData, nil
}

func (repo *InMemoryRepository) UnlockUser(ctx context.Context, uid uint32) (*user.UserData, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	userData, exists := repo.users[uid]
	if !exists {
		return nil, errors.New("user not found")
	}

	// ロック解除のためにis_lockedをfalseに設定
	userData.IsLocked = false
	return userData, nil
}
