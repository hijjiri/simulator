package repository

import (
	"context"
	"errors"
	"sync"

	u "github.com/hijjiri/simulator/core/go/grpc-server/user"
)

type InMemoryRepository struct {
	mu    sync.Mutex
	users map[uint32]*u.UserData
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users: make(map[uint32]*u.UserData),
	}
}

func (repo *InMemoryRepository) LockUser(ctx context.Context, uid uint32) (*u.UserData, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	user, exists := repo.users[uid]
	if !exists {
		user = &u.UserData{Uid: uid}
		repo.users[uid] = user
	}

	if user.Uid == 0 {
		return nil, errors.New("user is already locked")
	}

	// Simulate locking the user by setting Uid to 0
	user.Uid = 0
	return user, nil
}

func (repo *InMemoryRepository) UnlockUser(ctx context.Context, uid uint32) (*u.UserData, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	user, exists := repo.users[uid]
	if !exists {
		return nil, errors.New("user not found")
	}

	// Simulate unlocking the user by setting Uid to original value
	user.Uid = uid
	return user, nil
}
