package battle

import (
	"log"

	au "github.com/hijjiri/simulator/core/go/grpc-server/aura"
	d "github.com/hijjiri/simulator/core/go/grpc-server/deck"
	e "github.com/hijjiri/simulator/core/go/grpc-server/extension"
	h "github.com/hijjiri/simulator/core/go/grpc-server/hero"
	lc "github.com/hijjiri/simulator/core/go/grpc-server/landcryptid"
)

func NewBattleUnitFromDeck(position int32, u *d.DeckUnit, h *h.HeroData, e1, e2 *e.ExtensionData, cryptidInfo *lc.CryptidInfo, auras []*au.AuraEffect, skillOrders []int32, heroActiveIndex int32, contentType string) *Unit {
	log.Println("Creating unit at position", position)
	log.Printf("HeroData: %+v\n", h)
	log.Printf("Extension1: %+v\n", e1)
	log.Printf("Extension2: %+v\n", e2)

	// アクティブインデックスを調整
	log.Println("Adjusting hero active index")
	u.HeroActiveIndex = adjustHeroActiveIndex(u.HeroActiveIndex, h.StockedSkillIds)

	// スキルのセット
	log.Println("Setting active skills")
	actives := setActiveSkills(u, h, e1, e2)

	// オーラの初期化
	log.Println("Initializing aura effects")
	auraEffects, effectsMap, stateEffectRateMap, withinActionRateMap := initAuraEffects(auras, contentType)
	log.Printf("Initialized Aura Effects: auraEffects: %+v, effectsMap: %+v, stateEffectRateMap: %+v, withinActionRateMap: %+v\n", auraEffects, effectsMap, stateEffectRateMap, withinActionRateMap)

	// オーラの有効化
	log.Println("Setting aura enabled state")
	auraEnabled := setAuraEnabledState(auraEffects)

	// パラメータの計算
	log.Println("Calculating base params")
	log.Printf("HeroData.Param: %+v\n", h.Param)
	log.Printf("Extension1.Param: %+v\n", e1.Param)
	log.Printf("Extension2.Param: %+v\n", e2.Param)
	hp, phy, intl, agi := calculateBaseParams(h, e1, e2, effectsMap, cryptidInfo)
	log.Printf("Calculated params - HP: %d, PHY: %d, INT: %d, AGI: %d\n", hp, phy, intl, agi)

	// レートの調整
	log.Println("Calculating recovery rate")
	resultRate1000 := calculateRecoveryRate(cryptidInfo)

	// エクステより蘇生効果の有効化を確認
	log.Println("Checking resurrection enabled")
	resurrectionEnabled := checkResurrectionEnabled(e1, e2)

	log.Println("Checking get skill")
	log.Printf("Hero passive skill ID: %v", h.Passive)
	skill, err := GetSkill(h.Passive)
	if err != nil {
		log.Printf("Error getting skill: %v", err)
	}
	log.Printf("Obtained skill: %v", skill)

	return &Unit{
		Base:    createBaseUnit(position, hp, phy, intl, agi, actives, h, e1, e2, resultRate1000, auraEffects, effectsMap, skill),
		Current: createCurrentUnit(position, hp, phy, intl, agi, effectsMap, auraEnabled, stateEffectRateMap, withinActionRateMap, skill, resurrectionEnabled),
	}
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

func initAuraEffects(auras []*au.AuraEffect, contentType string) ([]*BaseUnit_AuraEffectRate, map[Aura_Effect_Param]int32, map[int32]int32, map[uint32]int32) {
	auraEffects := []*BaseUnit_AuraEffectRate{}
	effectsMap := map[Aura_Effect_Param]int32{}
	stateEffectRateMap := map[int32]int32{}
	withinActionRateMap := map[uint32]int32{}

	for _, effect := range auras {
		master, err := au.GetAuraTypeMaster(effect.AuraType)
		if err != nil {
			log.Printf("Failed to get AuraTypeMaster for AuraType %v: %v", effect.AuraType, err)
			continue
		}
		if master != nil {
			isValid := false

			switch contentType {
			case "Quest":
				isValid = master.IsQuest() && master.IsBeforeBattle()
			case "Battle":
				isValid = master.IsBattle() && master.IsBeforeBattle()
			case "simulator":
				isValid = false
			}

			if isValid {
				skill, err := GetAuraSkill(effect.AuraType)
				if err != nil {
					log.Printf("Failed to get AuraSkill for AuraType %v: %v", effect.AuraType, err)
					continue
				}
				if skill != nil {
					if skill.IsWithinAction {
						withinActionRateMap[skill.AuraId] = effect.SecondaryValue
					}
					effectRates := []int32{effect.EffectValue, effect.SecondaryValue}
					for i, skillEffect := range skill.Effects {
						effectsMap[skillEffect.Param] += effectRates[i]
						stateEffectRateMap[int32(skillEffect.Param)] += effectRates[i]
					}
				}
				auraEffect := &BaseUnit_AuraEffectRate{
					AuraType:       effect.AuraType,
					EffectValue:    effect.EffectValue,
					SecondaryValue: effect.SecondaryValue,
				}
				auraEffects = append(auraEffects, auraEffect)
			}
		}
	}

	log.Printf("Initialized Aura Effects: auraEffects: %v, effectsMap: %v, stateEffectRateMap: %v, withinActionRateMap: %v",
		auraEffects, effectsMap, stateEffectRateMap, withinActionRateMap)

	return auraEffects, effectsMap, stateEffectRateMap, withinActionRateMap
}

func setAuraEnabledState(auraEffects []*BaseUnit_AuraEffectRate) map[uint32]bool {
	auraEnabled := map[uint32]bool{}
	for _, aura := range auraEffects {
		auraEnabled[aura.AuraType] = true
	}
	return auraEnabled
}

func calculateBaseParams(h *h.HeroData, e1 *e.ExtensionData, e2 *e.ExtensionData, effectsMap map[Aura_Effect_Param]int32, cryptidInfo *lc.CryptidInfo) (int32, int32, int32, int32) {
	log.Println("Calculating base params")
	if h == nil {
		log.Fatal("HeroData is nil")
	}
	if h.Param == nil {
		log.Fatal("HeroData.Param is nil")
	}
	if e1 == nil {
		log.Fatal("ExtensionData1 is nil")
	}
	if e1.Param == nil {
		log.Fatal("ExtensionData1.Param is nil")
	}
	if e2 == nil {
		log.Fatal("ExtensionData2 is nil")
	}
	if e2.Param == nil {
		log.Fatal("ExtensionData2.Param is nil")
	}
	if effectsMap == nil {
		log.Fatal("effectsMap is nil")
	}
	if cryptidInfo == nil {
		log.Println("cryptidInfo is nil")
	}

	log.Printf("HeroData.Param: %+v\n", h.Param)
	log.Printf("Extension1.Param: %+v\n", e1.Param)
	log.Printf("Extension2.Param: %+v\n", e2.Param)
	log.Printf("effectsMap: %+v\n", effectsMap)

	hp := h.Param.Hp + e1.Param.Hp + e2.Param.Hp + effectsMap[Aura_Effect_INITIAL_HP]
	phy := h.Param.Phy + e1.Param.Phy + e2.Param.Phy + effectsMap[Aura_Effect_INITIAL_PHY]
	intl := h.Param.Int + e1.Param.Int + e2.Param.Int + effectsMap[Aura_Effect_INITIAL_INT]
	agi := h.Param.Agi + e1.Param.Agi + e2.Param.Agi + effectsMap[Aura_Effect_INITIAL_AGI]

	if cryptidInfo != nil {
		if cryptidInfo.Hp != nil {
			hp = hp * (1000 + cryptidInfo.Hp.BuffRate1000) / 1000
		}
		if cryptidInfo.Phy != nil {
			phy = phy * (1000 + cryptidInfo.Phy.BuffRate1000) / 1000
		}
		if cryptidInfo.Int != nil {
			intl = intl * (1000 + cryptidInfo.Int.BuffRate1000) / 1000
		}
		if cryptidInfo.Agi != nil {
			agi = agi * (1000 + cryptidInfo.Agi.BuffRate1000) / 1000
		}
	}

	return hp, phy, intl, agi
}

func calculateRecoveryRate(cryptidInfo *lc.CryptidInfo) map[int32]int32 {
	resultRate1000 := map[int32]int32{}

	if cryptidInfo != nil {
		if cryptidInfo.Recovery != nil {
			resultRate1000[int32(BaseUnit_RECOVERY)] = 1000 + cryptidInfo.Recovery.BuffRate1000
		} else {
			log.Println("cryptidInfo.Recovery is nil")
		}
	} else {
		log.Println("cryptidInfo is nil")
	}

	return resultRate1000
}

func checkResurrectionEnabled(e1 *e.ExtensionData, e2 *e.ExtensionData) bool {
	if e1 == nil {
		log.Println("e1 is nil")
	} else {
		log.Printf("e1.Aura: %v", e1.Aura)
	}
	if e2 == nil {
		log.Println("e2 is nil")
	} else {
		log.Printf("e2.Aura: %v", e2.Aura)
	}
	return e1.Aura == e.AuraIdResurrection || e2.Aura == e.AuraIdResurrection
}

func createBaseUnit(position, hp, phy, intl, agi int32, actives []uint32, h *h.HeroData, e1, e2 *e.ExtensionData, resultRate1000 map[int32]int32, auraEffects []*BaseUnit_AuraEffectRate, effectsMap map[Aura_Effect_Param]int32, skill *Skill) *BaseUnit {
	return &BaseUnit{
		Position:       position,
		Hp:             hp,
		Phy:            phy,
		Int:            intl,
		Agi:            agi,
		Actives:        actives,
		Passive:        h.Passive,
		Character:      h.HeroId,
		Extension1:     e1.ExtensionId,
		Extension2:     e2.ExtensionId,
		IsEnemy:        false,
		ResultRate1000: resultRate1000,
		Auras:          auraEffects,
		EnchantInfos:   skill.EnchantInfos,
	}
}

func createCurrentUnit(position, hp, phy, intl, agi int32, effectsMap map[Aura_Effect_Param]int32, auraEnabled map[uint32]bool, stateEffectRateMap map[int32]int32, withinActionRateMap map[uint32]int32, skill *Skill, resurrectionEnabled bool) *CurrentUnit {
	return &CurrentUnit{
		Position:              position,
		Hp:                    hp,
		Phy:                   phy,
		Int:                   intl,
		Agi:                   agi,
		Charge:                0,
		BonusCharge:           0,
		ActiveCounts:          0,
		PassiveEnabled:        true,
		State:                 CurrentUnit_NORMAL,
		ResurrectionEnabled:   resurrectionEnabled,
		DecoyEnabled:          false,
		DecoyInAction:         false,
		SkillDeathEnabled:     false,
		AuraDeathEnabled:      false,
		DuplicateDecoyEnabled: false,
		CriticalJudg:          false,
		Shield:                0,
		ActiveAuras:           []uint32{},
		AuraEnabled:           auraEnabled,
		LastTargets:           []*Unit{},
		StateEffectRate:       stateEffectRateMap,
		WithinActionValue:     withinActionRateMap,
		EnchantInfos:          skill.EnchantInfos,
	}
}
