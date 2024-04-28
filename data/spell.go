package data

type SpellType int

const (
	Healing SpellType = iota
	CauseCondition
	DirectDamage
	Drain
	ActivateEffect
)

type Spell struct {
	Id         int
	Name       string
	Type       SpellType
	DamageType DamageType
	Variable   int
	Condition  int
	Effect     BattleEffect
	HitQuote   string
	MissQuote  string
	FailQuote  string
	ExtraQuote string
}

type BattleEffect int

const (
	NoDeath BattleEffect = iota
)

var Spells []Spell = []Spell{
	{},
	{
		Id:         1,
		Name:       "Acid Arrow",
		Type:       DirectDamage,
		DamageType: Earth,
		Variable:   2,
		Condition:  0,
		Effect:     0,
		HitQuote:   "%s cast Acid Arrow for %d Earth damage.",
		MissQuote:  "%s cast Acid Arrow but %s successfully avoided the spell!",
		FailQuote:  "%s cast Acid Arrow but %s is immune!",
		ExtraQuote: "",
	},
}
