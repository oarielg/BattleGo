package data

import "math/rand/v2"

func GeneratePlayer() Character {
	return Players[rand.IntN(len(Players))]
}

var Players []Character = []Character{
	{
		Id:        1,
		Name:      "Barbarian",
		Might:     3,
		Power:     1,
		Speed:     3,
		Mind:      1,
		MaxHp:     30,
		CurrentHp: 30,
		Inventory: [4]int{5, 0, 0, 0},
		Spellbook: []int{27},
	},
	{
		Id:        2,
		Name:      "Cleric",
		Might:     2,
		Power:     3,
		Speed:     1,
		Mind:      2,
		MaxHp:     30,
		CurrentHp: 30,
		Inventory: [4]int{2, 0, 0, 0},
		Spellbook: []int{7, 8, 10, 24, 26},
	},
	{
		Id:        3,
		Name:      "Druid",
		Might:     2,
		Power:     3,
		Speed:     2,
		Mind:      1,
		MaxHp:     30,
		CurrentHp: 30,
		Inventory: [4]int{1, 0, 0, 0},
		Spellbook: []int{1, 12, 14, 17, 25, 27},
	},
	{
		Id:        4,
		Name:      "Mentalist",
		Might:     1,
		Power:     3,
		Speed:     1,
		Mind:      3,
		MaxHp:     30,
		CurrentHp: 30,
		Inventory: [4]int{1, 0, 0, 0},
		Spellbook: []int{4, 18, 20},
	},
	{
		Id:        5,
		Name:      "Warlock",
		Might:     2,
		Power:     3,
		Speed:     1,
		Mind:      2,
		MaxHp:     30,
		CurrentHp: 30,
		Inventory: [4]int{4, 0, 0, 0},
		Spellbook: []int{2, 3, 5, 6, 15, 19},
	},
	{
		Id:        6,
		Name:      "Wizard",
		Might:     1,
		Power:     3,
		Speed:     1,
		Mind:      3,
		MaxHp:     30,
		CurrentHp: 30,
		Inventory: [4]int{1, 0, 0, 0},
		Spellbook: []int{6, 9, 11, 13, 16, 21, 22, 23},
	},
}
