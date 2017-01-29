package d2s

// All attribute ids described.
const (
	strength       = 0
	energy         = 1
	dexterity      = 2
	vitality       = 3
	unusedStats    = 4
	unusedSkills   = 5
	currentHP      = 6
	maxHP          = 7
	currentMana    = 8
	maxMana        = 9
	currentStamina = 10
	maxStamina     = 11
	level          = 12
	experience     = 13
	gold           = 14
	stashedGold    = 15
)

// Attributes Represents a characters allocated attributes.
type attributes struct {
	Strength          uint64 `json:"strength"`
	Energy            uint64 `json:"energy"`
	Dexterity         uint64 `json:"dexterity"`
	Vitality          uint64 `json:"vitality"`
	UnusedStats       uint64 `json:"unused_stats"`
	UnusedSkillPoints uint64 `json:"unused_skill_points"`
	CurrentHP         uint64 `json:"current_hp"`
	MaxHP             uint64 `json:"max_hp"`
	CurrentMana       uint64 `json:"current_mana"`
	MaxMana           uint64 `json:"max_mana"`
	CurrentStamina    uint64 `json:"current_stamina"`
	MaxStamina        uint64 `json:"max_stamina"`
	Level             uint64 `json:"level"`
	Experience        uint64 `json:"experience"`
	Gold              uint64 `json:"gold"`
	StashedGold       uint64 `json:"stashed_gold"`
}

// Holds all the references to bit sizes of all attributes.
var attributeBitMap = map[uint64]uint{
	strength:       10,
	energy:         10,
	dexterity:      10,
	vitality:       10,
	unusedStats:    10,
	unusedSkills:   8,
	currentHP:      21,
	maxHP:          21,
	currentMana:    21,
	maxMana:        21,
	currentStamina: 21,
	maxStamina:     21,
	level:          7,
	experience:     32,
	gold:           25,
	stashedGold:    25,
}
