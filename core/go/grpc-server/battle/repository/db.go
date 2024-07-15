package repository

import (
	"fmt"
	"sync"

	"github.com/hijjiri/simulator/core/go/grpc-server/battle"
)

// インメモリキャッシュのためのマップとミューテックス
var (
	cachedSkills = make(map[uint32]*battle.Skill)
	mu           sync.Mutex
)

// DatastoreRepository構造体の定義
type DatastoreRepository struct{}

// GetSkill関数の実装
func (r *DatastoreRepository) GetSkill(skillId uint32) (*battle.Skill, error) {
	mu.Lock()         // ミューテックスをロック
	defer mu.Unlock() // 関数終了時にミューテックスをアンロック

	skill, ok := cachedSkills[skillId]
	if !ok {
		return nil, fmt.Errorf("skill not found: %v", skillId)
	}
	return skill, nil
}

// NewDatastoreRepository関数の実装
func NewDatastoreRepository() *DatastoreRepository {
	return &DatastoreRepository{}
}

// InitializeSkills関数の実装
func InitializeSkills(skills map[uint32]*battle.Skill) {
	mu.Lock()         // ミューテックスをロック
	defer mu.Unlock() // 関数終了時にミューテックスをアンロック
	for id, skill := range skills {
		cachedSkills[id] = skill
	}
}
