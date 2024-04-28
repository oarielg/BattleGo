package data

import "math/rand/v2"

func GenerateMonster() Character {
	return Monsters[rand.IntN(len(Monsters))]
}

var Monsters []Character = []Character{
	{
		Id:        0,
		Name:      "Dummy",
		Might:     1,
		Power:     1,
		Speed:     1,
		Mind:      1,
		MaxHp:     20,
		CurrentHp: 20,
		Inventory: [4]int{1, 0, 0, 0},
		Class:     Fighter,
	},
	/* {
		Id:             1,
		Name:           "Angel Deva",
		Might:          4,
		Power:          3,
		Speed:          4,
		Mind:           3,
		MaxHp:          40,
		CurrentHp:      40,
		Inventory:      [4]int{1, 0, 0, 0},
		Spellbook:      []int{22},
		Resistances:    []DamageResistance{{DamageType: Radiant, Resistance: 2}},
		Immunities:     []DamageType{Thunder},
		CondImmunities: []int{9, 11},
		Class:          Hybrid,
	}, */
}
