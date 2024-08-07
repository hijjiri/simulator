package battle

var r Repository

type Repository interface {
	GetSkill(skillId uint32) (*Skill, error)
	GetAuraSkill(auraType uint32) (*Aura, error)
}

func RegisterRepository(repository Repository) {
	r = repository
}
