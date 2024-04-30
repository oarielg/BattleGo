package data

type Item struct {
	Id           int
	Name         string
	Type         ItemType
	DamageType   DamageType
	Damage       int
	Armor        int
	Attack       int
	Defense      int
	Initiative   int
	Enchantments []int
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
		Name:         "Staff",
		Type:         Weapon,
		DamageType:   Bludgeoning,
		Damage:       3,
		Armor:        0,
		Attack:       0,
		Defense:      0,
		Initiative:   0,
		Enchantments: nil,
	},
	{
		Id:           2,
		Name:         "Mace",
		Type:         Weapon,
		DamageType:   Bludgeoning,
		Damage:       3,
		Armor:        0,
		Attack:       0,
		Defense:      0,
		Initiative:   0,
		Enchantments: nil,
	},
	{
		Id:           3,
		Name:         "Dagger",
		Type:         Weapon,
		DamageType:   Piercing,
		Damage:       2,
		Armor:        0,
		Attack:       1,
		Defense:      0,
		Initiative:   0,
		Enchantments: nil,
	},
	{
		Id:           4,
		Name:         "Shortsword",
		Type:         Weapon,
		DamageType:   Slashing,
		Damage:       3,
		Armor:        0,
		Attack:       0,
		Defense:      0,
		Initiative:   0,
		Enchantments: nil,
	},
	{
		Id:           5,
		Name:         "Greataxe",
		Type:         Weapon,
		DamageType:   Slashing,
		Damage:       5,
		Armor:        0,
		Attack:       0,
		Defense:      0,
		Initiative:   0,
		Enchantments: nil,
	},
	{
		Id:           6,
		Name:         "Flaming Mace",
		Type:         Weapon,
		DamageType:   Bludgeoning,
		Damage:       3,
		Armor:        0,
		Attack:       0,
		Defense:      0,
		Initiative:   0,
		Enchantments: []int{1},
	},
	{
		Id:           7,
		Name:         "Flaming Greatsword",
		Type:         Weapon,
		DamageType:   Slashing,
		Damage:       5,
		Armor:        0,
		Attack:       0,
		Defense:      0,
		Initiative:   0,
		Enchantments: []int{1},
	},
	{
		Id:           8,
		Name:         "Corrupting Touch",
		Type:         Weapon,
		DamageType:   Necrotic,
		Damage:       2,
		Armor:        0,
		Attack:       1,
		Defense:      0,
		Initiative:   0,
		Enchantments: nil,
	},
	{
		Id:           9,
		Name:         "Poisonous Bite",
		Type:         Weapon,
		DamageType:   Piercing,
		Damage:       2,
		Armor:        0,
		Attack:       1,
		Defense:      0,
		Initiative:   0,
		Enchantments: []int{4},
	},
	{
		Id:           10,
		Name:         "Bite",
		Type:         Weapon,
		DamageType:   Piercing,
		Damage:       2,
		Armor:        0,
		Attack:       1,
		Defense:      0,
		Initiative:   0,
		Enchantments: nil,
	},
	{
		Id:           11,
		Name:         "Spear",
		Type:         Weapon,
		DamageType:   Piercing,
		Damage:       3,
		Armor:        0,
		Attack:       0,
		Defense:      0,
		Initiative:   0,
		Enchantments: nil,
	},
	{
		Id:           12,
		Name:         "Greatclub",
		Type:         Weapon,
		DamageType:   Bludgeoning,
		Damage:       5,
		Armor:        0,
		Attack:       0,
		Defense:      0,
		Initiative:   0,
		Enchantments: nil,
	},
	{
		Id:           13,
		Name:         "Cursed Longsword",
		Type:         Weapon,
		DamageType:   Piercing,
		Damage:       4,
		Armor:        0,
		Attack:       0,
		Defense:      0,
		Initiative:   0,
		Enchantments: []int{6},
	},
}
