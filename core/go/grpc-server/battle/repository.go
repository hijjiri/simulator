package battle

var r Repository

type Repository interface {
	GetSKill(skillId uint32) (*Skill, error)
}

func RegisterRepository(repository Repository) {
	r = repository
}
