package data

type Item struct {
	Id           int
	Name         string
	Type         ItemType
	Damage       int
	Armor        int
	Attack       int
	Defense      int
	Initiative   int
	Enchantments []Enchantment
}

type ItemType int

const (
	Weapon ItemType = iota
	Shield
	Armor
	Ring
)

var Items []Item = []Item{
	{},
	{
		Id:           1,
		Name:         "Dagger",
		Type:         Weapon,
		Damage:       1,
		Armor:        0,
		Attack:       1,
		Defense:      0,
		Initiative:   1,
		Enchantments: nil,
	},
}
