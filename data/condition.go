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
		ActiveQuote: "<b>%s</b> is Asleep!",
		EndQuote:    "<b>%s</b> woke up.",
	},
}
