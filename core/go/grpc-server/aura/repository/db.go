package repository

import (
	"fmt"
	"sync"

	"github.com/hijjiri/simulator/core/go/grpc-server/aura"
)

// インメモリキャッシュのためのマップとミューテックス
var (
	cacheAuraTypesMap = make(map[uint32]*aura.AuraTypeMaster)
	mu                sync.Mutex
)

// DatastoreRepository構造体の定義
type DatastoreRepository struct{}

// GetAuraTypeMaster関数の実装
func (r *DatastoreRepository) GetAuraTypeMaster(auraType uint32) (*aura.AuraTypeMaster, error) {
	mu.Lock()         // ミューテックスをロック
	defer mu.Unlock() // 関数終了時にミューテックスをアンロック

	auraTypeMaster, ok := cacheAuraTypesMap[auraType]
	if !ok {
		return nil, fmt.Errorf("aura type not found auraType=%v", auraType)
	}
	return auraTypeMaster, nil
}

// NewDatastoreRepository関数の実装
func NewDatastoreRepository() *DatastoreRepository {
	return &DatastoreRepository{}
}

// InitializeAuraTypes関数の実装
func InitializeAuraTypes(auraTypes map[uint32]*aura.AuraTypeMaster) {
	mu.Lock()         // ミューテックスをロック
	defer mu.Unlock() // 関数終了時にミューテックスをアンロック
	for id, auraType := range auraTypes {
		cacheAuraTypesMap[id] = auraType
	}
}
