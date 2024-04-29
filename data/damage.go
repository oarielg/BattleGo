package data

type DamageType int

const (
	Bludgeoning DamageType = iota
	Slashing
	Piercing
	Fire
	Cold
	Thunder
	Earth
	Radiant
	Necrotic
	Psychic
)

var damageTypeToString = map[DamageType]string{
	Bludgeoning: "Bludgeoning",
	Slashing:    "Slashing",
	Piercing:    "Piercing",
	Fire:        "Fire",
	Cold:        "Cold",
	Thunder:     "Thunder",
	Earth:       "Earth",
	Radiant:     "Radiant",
	Necrotic:    "Necrotic",
	Psychic:     "Necrotic",
}

func (dt DamageType) String() string {
	if s, ok := damageTypeToString[dt]; ok {
		return s
	}
	return "unknown"
}

type DamageResistance struct {
	DamageType DamageType
	Resistance int
}
