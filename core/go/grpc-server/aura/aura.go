package aura

func GetAuraTypeMaster(auraType uint32) (*AuraTypeMaster, error) {
	return r.GetAuraTypeMaster(auraType)
}

func (auraTypeMaster *AuraTypeMaster) IsBattle() bool {
	return auraTypeMaster.Occasions == Occasions_BATTLE
}

func (auraTypeMaster *AuraTypeMaster) IsQuest() bool {
	return auraTypeMaster.Occasions == Occasions_BATTLE || auraTypeMaster.Occasions == Occasions_QUEST
}

func (auraTypeMaster *AuraTypeMaster) IsCpMining() bool {
	return auraTypeMaster.Occasions == Occasions_CP_MINING
}

func (auraTypeMaster *AuraTypeMaster) IsBeforeBattle() bool {
	return auraTypeMaster.Trigger == Trigger_BEFORE
}

func (auraTypeMaster *AuraTypeMaster) IsDuringBattle() bool {
	return auraTypeMaster.Trigger == Trigger_DURING
}
