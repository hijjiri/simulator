package repository

import (
	"fmt"
	"sync"

	"github.com/hijjiri/simulator/core/go/grpc-server/battle"
)

type InMemoryRepository struct {
	mu    sync.Mutex
	cache map[uint32]*battle.Skill
	auras map[uint32]*battle.Aura
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		cache: make(map[uint32]*battle.Skill),
		auras: make(map[uint32]*battle.Aura),
	}
}

func (r *InMemoryRepository) GetSkill(skillId uint32) (*battle.Skill, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	skill, ok := r.cache[skillId]
	if !ok {
		defaultSkill, exists := r.cache[1]
		if !exists {
			return nil, fmt.Errorf("default skill not found")
		}
		return defaultSkill, fmt.Errorf("skill not found: %v, returning default skill", skillId)
	}
	return skill, nil
}

func (r *InMemoryRepository) GetAuraSkill(auraType uint32) (*battle.Aura, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	aura, ok := r.auras[auraType]
	if !ok {
		return nil, fmt.Errorf("aura skill not found: %v", auraType)
	}
	return aura, nil
}

func (r *InMemoryRepository) InitializeSkills(skills map[uint32]*battle.Skill) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for id, skill := range skills {
		r.cache[id] = skill
	}
}

func (r *InMemoryRepository) InitializeAuras(auras map[uint32]*battle.Aura) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for id, aura := range auras {
		r.auras[id] = aura
	}
}
