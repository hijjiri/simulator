package battle

import (
	d "github.com/hijjiri/simulator/core/go/grpc-server/deck"
	e "github.com/hijjiri/simulator/core/go/grpc-server/extension"
	h "github.com/hijjiri/simulator/core/go/grpc-server/hero"
)

func NewBattleUnitFromDeck(position int32, u *d.DeckUnit, h *h.HeroData, e1, e2 *e.ExtensionData, skillOrders []int32, heroActiveIndex int32, contentType string) *Unit {
	// アクティブインデックスを調整
	u.HeroActiveIndex = adjustHeroActiveIndex(u.HeroActiveIndex, h.StockedSkillIds)

	// スキルのセット
	// actives := setActiveSkills(u, h, e1, e2)

	// パラメータの計算
	// hp, phy, intl, agi := calculateBaseParams(h, e1, e2, nil, nil)

	// エクステより蘇生効果の有効化を確認
	// resurrectionEnabled := checkResurrectionEnabled(e1, e2)

	// skill, _ := GetSkill(h.Passive)

	// return &Unit{
	// 	Base:    createBaseUnit(position, hp, phy, intl, agi, actives, h, e1, e2, 0, nil, nil, skill),
	// 	Current: createCurrentUnit(position, hp, phy, intl, agi, nil, false, nil, nil, skill, resurrectionEnabled),
	// }
	return &Unit{}
}

func adjustHeroActiveIndex(heroActiveIndex int32, stockedSkillIds []uint32) int32 {
	if int32(len(stockedSkillIds)) <= heroActiveIndex {
		return 0
	}
	return heroActiveIndex
}

func setActiveSkills(u *d.DeckUnit, h *h.HeroData, e1 *e.ExtensionData, e2 *e.ExtensionData) []uint32 {
	tmpActives := []uint32{h.StockedSkillIds[u.HeroActiveIndex], e1.Active, e2.Active}
	actives := make([]uint32, 3)
	for i, order := range u.SkillOrders {
		actives[i] = tmpActives[order]
	}
	return actives
}

// func calculateBaseParams(h *h.HeroData, e1 *e.ExtensionData, e2 *e.ExtensionData, cryptidInfo *lc.CryptidInfo) (int32, int32, int32, int32) {
// 	hp := h.Param.Hp + e1.Param.Hp + e2.Param.Hp + effectsMap[Aura_Effect_INITIAL_HP]
// 	phy := h.Param.Phy + e1.Param.Phy + e2.Param.Phy + effectsMap[Aura_Effect_INITIAL_PHY]
// 	intl := h.Param.Int + e1.Param.Int + e2.Param.Int + effectsMap[Aura_Effect_INITIAL_INT]
// 	agi := h.Param.Agi + e1.Param.Agi + e2.Param.Agi + effectsMap[Aura_Effect_INITIAL_AGI]

// 	if cryptidInfo != nil {
// 		hp = hp * (1000 + cryptidInfo.Hp.BuffRate1000) / 1000
// 		phy = phy * (1000 + cryptidInfo.Phy.BuffRate1000) / 1000
// 		intl = intl * (1000 + cryptidInfo.Int.BuffRate1000) / 1000
// 		agi = agi * (1000 + cryptidInfo.Agi.BuffRate1000) / 1000
// 	}

// 	return hp, phy, intl, agi
// }

// func calculateRecoveryRate(cryptidInfo *lc.CryptidInfo) map[int32]int32 {
// 	resultRate1000 := map[int32]int32{}

// 	if cryptidInfo != nil {
// 		resultRate1000[int32(BaseUnit_RECOVERY)] = 1000 + cryptidInfo.Recovery.BuffRate1000
// 	}

// 	return resultRate1000
// }

// func checkResurrectionEnabled(e1 *e.ExtensionData, e2 *e.ExtensionData) bool {
// 	return e1.Aura == e.AuraIdResurrection || e2.Aura == e.AuraIdResurrection
// }

// func createBaseUnit(position, hp, phy, intl, agi int32, actives []uint32, h *h.HeroData, e1, e2 *e.ExtensionData, resultRate1000 map[int32]int32, auraEffects []*BaseUnit_AuraEffectRate, effectsMap map[Aura_Effect_Param]int32, skill *Skill) *BaseUnit {
// 	return &BaseUnit{
// 		Position:       position,
// 		Hp:             hp,
// 		Phy:            phy,
// 		Int:            intl,
// 		Agi:            agi,
// 		Actives:        actives,
// 		Passive:        h.Passive,
// 		Character:      h.HeroId,
// 		Extension1:     e1.ExtensionId,
// 		Extension2:     e2.ExtensionId,
// 		IsEnemy:        false,
// 		ResultRate1000: resultRate1000,
// 		Auras:          auraEffects,
// 		EnchantInfos:   skill.EnchantInfos,
// 	}
// }

// func createCurrentUnit(position, hp, phy, intl, agi int32, auraEnabled map[uint32]bool, stateEffectRateMap map[int32]int32, withinActionRateMap map[uint32]int32, skill *Skill, resurrectionEnabled bool) *CurrentUnit {
// 	return &CurrentUnit{
// 		Position:              position,
// 		Hp:                    hp,
// 		Phy:                   phy,
// 		Int:                   intl,
// 		Agi:                   agi,
// 		Charge:                0,
// 		BonusCharge:           0,
// 		ActiveCounts:          0,
// 		PassiveEnabled:        true,
// 		State:                 CurrentUnit_NORMAL,
// 		ResurrectionEnabled:   resurrectionEnabled,
// 		DecoyEnabled:          false,
// 		DecoyInAction:         false,
// 		SkillDeathEnabled:     false,
// 		AuraDeathEnabled:      false,
// 		DuplicateDecoyEnabled: false,
// 		CriticalJudg:          false,
// 		Shield:                0,
// 		ActiveAuras:           []uint32{},
// 		AuraEnabled:           auraEnabled,
// 		LastTargets:           []*Unit{},
// 		StateEffectRate:       stateEffectRateMap,
// 		WithinActionValue:     withinActionRateMap,
// 		EnchantInfos:          skill.EnchantInfos,
// 	}
// }
