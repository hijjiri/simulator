package aura

var r Repository

type Repository interface {
	GetAuraTypeMaster(auraType uint32) (*AuraTypeMaster, error)
}

func RegisterRepository(repository Repository) {
	r = repository
}
