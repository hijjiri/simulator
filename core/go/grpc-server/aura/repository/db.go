package repository

import (
	"fmt"
	"sync"

	"github.com/hijjiri/simulator/core/go/grpc-server/aura"
)

type InMemoryRepository struct {
	mu    sync.Mutex
	cache map[uint32]*aura.AuraTypeMaster
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		cache: make(map[uint32]*aura.AuraTypeMaster),
	}
}

func (r *InMemoryRepository) GetAuraTypeMaster(auraType uint32) (*aura.AuraTypeMaster, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	auraTypeMaster, ok := r.cache[auraType]
	if !ok {
		return nil, fmt.Errorf("aura type not found auraType=%v", auraType)
	}
	return auraTypeMaster, nil
}

func (r *InMemoryRepository) InitializeAuraTypes(auraTypes map[uint32]*aura.AuraTypeMaster) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for id, auraType := range auraTypes {
		r.cache[id] = auraType
	}
}
