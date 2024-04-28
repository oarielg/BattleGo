package data

import "math/rand/v2"

func GeneratePlayer() Character {
	return Players[rand.IntN(len(Players))]
}

var Players []Character = []Character{
	{
		Id:        1,
		Name:      "Wizard",
		Might:     1,
		Power:     3,
		Speed:     1,
		Mind:      3,
		MaxHp:     30,
		CurrentHp: 30,
		Inventory: [4]int{1, 0, 0, 0},
		Spellbook: []int{1},
	},
}
