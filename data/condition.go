package data

type ConditionType int

const (
	NoActions ConditionType = iota
	DoDamage
	NoSpells
	NoOffensiveAction
)

type Condition struct {
	Id          int
	Name        string
	Type        ConditionType
	MinDuration int
	MaxDuration int
	DamageType  DamageType
	ActiveQuote string
	EndQuote    string
	FailQuote   string
}

type ActiveCondition struct {
	Duration  int
	Condition Condition
}

var Conditions []Condition = []Condition{
	{},
	{
		Id:          1,
		Name:        "Asleep",
		Type:        NoActions,
		MinDuration: 2,
		MaxDuration: 4,
		ActiveQuote: "%s is Asleep!",
		EndQuote:    "%s woke up.",
	},
	{
		Id:          2,
		Name:        "Bleeding (Slashing)",
		Type:        DoDamage,
		MinDuration: 3,
		MaxDuration: 6,
		DamageType:  Slashing,
		ActiveQuote: "%s is Bleeding! It takes %d Slashing damage.",
		EndQuote:    "%s is no longer Bleeding.",
		FailQuote:   "%s is Bleeding but it is immune to Slashing damage!",
	},
	{
		Id:          3,
		Name:        "Bleeding (Piercing)",
		Type:        DoDamage,
		MinDuration: 3,
		MaxDuration: 6,
		DamageType:  Piercing,
		ActiveQuote: "%s is Bleeding! It takes %d Piercing damage.",
		EndQuote:    "%s is no longer Bleeding.",
		FailQuote:   "%s is Bleeding but it is immune to Piercing damage!",
	},
	{
		Id:          4,
		Name:        "Burned",
		Type:        DoDamage,
		MinDuration: 3,
		MaxDuration: 6,
		DamageType:  Fire,
		ActiveQuote: "%s is Burned! It takes %d Fire damage.",
		EndQuote:    "%s is no longer Burned.",
		FailQuote:   "%s is Burned but it is immune to Fire damage!",
	},
	{
		Id:          5,
		Name:        "Charmed",
		Type:        NoOffensiveAction,
		MinDuration: 2,
		MaxDuration: 4,
		ActiveQuote: "%s is Charmed!",
		EndQuote:    "%s is no longer Charmed.",
	},
	{
		Id:          6,
		Name:        "Confused",
		Type:        DoDamage,
		MinDuration: 3,
		MaxDuration: 6,
		DamageType:  Psychic,
		ActiveQuote: "%s is Confused! It takes %d Psychic damage.",
		EndQuote:    "%s is no longer Confused.",
		FailQuote:   "%s is Confused but it is immune to Psychic damage!",
	},
	{
		Id:          7,
		Name:        "Cursed",
		Type:        DoDamage,
		MinDuration: 3,
		MaxDuration: 6,
		DamageType:  Necrotic,
		ActiveQuote: "%s is Cursed! It takes %d Necrotic damage.",
		EndQuote:    "%s is no longer Cursed.",
		FailQuote:   "%s is Cursed but it is immune to Necrotic damage!",
	},
	{
		Id:          8,
		Name:        "Doomed",
		Type:        DoDamage,
		MinDuration: 3,
		MaxDuration: 6,
		DamageType:  Radiant,
		ActiveQuote: "%s is Doomed! It takes %d Radiant damage.",
		EndQuote:    "%s is no longer Doomed.",
		FailQuote:   "%s is Doomed but it is immune to Radiant damage!",
	},
	{
		Id:          9,
		Name:        "Electrified",
		Type:        DoDamage,
		MinDuration: 3,
		MaxDuration: 6,
		DamageType:  Thunder,
		ActiveQuote: "%s is Electrified! It takes %d Thunder damage.",
		EndQuote:    "%s is no longer Electrified.",
		FailQuote:   "%s is Electrified but it is immune to Thunder damage!",
	},
	{
		Id:          10,
		Name:        "Entangled",
		Type:        NoActions,
		MinDuration: 2,
		MaxDuration: 4,
		ActiveQuote: "%s is Entangled!",
		EndQuote:    "%s is no longer Entangled.",
	},
	{
		Id:          11,
		Name:        "Frozen",
		Type:        DoDamage,
		MinDuration: 3,
		MaxDuration: 6,
		DamageType:  Cold,
		ActiveQuote: "%s is Frozen! It takes %d Cold damage.",
		EndQuote:    "%s is no longer Frozen.",
		FailQuote:   "%s is Frozen but it is immune to Cold damage!",
	},
	{
		Id:          12,
		Name:        "Paralyzed",
		Type:        NoActions,
		MinDuration: 2,
		MaxDuration: 4,
		ActiveQuote: "%s is Paralyzed!",
		EndQuote:    "%s is no longer Paralyzed.",
	},
	{
		Id:          13,
		Name:        "Petrified",
		Type:        NoActions,
		MinDuration: 2,
		MaxDuration: 4,
		ActiveQuote: "%s is Petrified!",
		EndQuote:    "%s is no longer Petrified.",
	},
	{
		Id:          14,
		Name:        "Poisoned",
		Type:        DoDamage,
		MinDuration: 3,
		MaxDuration: 6,
		DamageType:  Earth,
		ActiveQuote: "%s is Poisoned! It takes %d Earth damage.",
		EndQuote:    "%s is no longer Poisoned.",
		FailQuote:   "%s is Poisoned but it is immune to Earth damage!",
	},
	{
		Id:          15,
		Name:        "Scared",
		Type:        NoActions,
		MinDuration: 2,
		MaxDuration: 4,
		ActiveQuote: "%s is Scared!",
		EndQuote:    "%s is no longer Scared.",
	},
	{
		Id:          16,
		Name:        "Silenced",
		Type:        NoSpells,
		MinDuration: 2,
		MaxDuration: 4,
		ActiveQuote: "%s is Silenced!",
		EndQuote:    "%s is no longer Silenced.",
	},
}
