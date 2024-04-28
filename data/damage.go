package data

type DamageType int

const (
	Slashing DamageType = iota
	Piercing
	Bludgeoning
	Fire
	Cold
	Thunder
	Earth
	Radiant
	Necrotic
	Psychic
)

type DamageResistance struct {
	DamageType DamageType
	Resistance int
}
