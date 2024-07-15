package battle

func GetSkill(id uint32) (*Skill, error) {
	return r.GetSkill(id)
}

func GetAuraSkill(auraType uint32) (*Aura, error) {
	return r.GetAuraSkill(auraType)
}
