package battle

func GetSkill(id uint32) (*Skill, error) {
	return r.GetSkill(id)
}

func GetAuraSkill(id uint32) (*Aura, error) {
	return r.GetAuraSkill(id)
}
