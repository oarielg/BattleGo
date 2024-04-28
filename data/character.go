package data

var playerData Character
var monsterData Character

type Character struct {
	Id              int
	Name            string
	Might           int
	Power           int
	Speed           int
	Mind            int
	CurrentHp       int
	MaxHp           int
	Inventory       [4]int
	PrimarySlot     Item
	SecondarySlot   Item
	Armor           Item
	Ring            Item
	Spellbook       []int
	Spells          []Spell
	Resistances     []DamageResistance
	Immunities      []DamageType
	Vulnerabilities []DamageType
	Conditions      []ActiveCondition
	CondImmunities  []int
	Effects         []BattleEffect
	Class           CharacterClass
	IsInit          bool
}

type CharacterClass int

const (
	Fighter CharacterClass = iota
	Caster
	Hybrid
)

func NewCharacter() *Character {
	return &Character{}
}

func (c *Character) InitCharacter() {
	if len(c.Spellbook) > 0 && len(c.Spells) == 0 {
		for _, s := range c.Spellbook {
			spell := Spells[s]
			c.Spells = append(c.Spells, spell)
		}
	}
	if c.Inventory[0] != 0 && c.PrimarySlot.Id == 0 {
		item := Items[c.Inventory[0]]
		c.PrimarySlot = item
	}
	if c.Inventory[1] != 0 && c.SecondarySlot.Id == 0 {
		item := Items[c.Inventory[1]]
		c.SecondarySlot = item
	}
	if c.Inventory[2] != 0 && c.Armor.Id == 0 {
		item := Items[c.Inventory[2]]
		c.Armor = item
	}
	if c.Inventory[3] != 0 && c.Ring.Id == 0 {
		item := Items[c.Inventory[3]]
		c.Ring = item
	}
	c.IsInit = true
}

func SaveCharacter(char Character, isPlayer bool) {
	if isPlayer {
		playerData = char
	} else {
		monsterData = char
	}
}

func LoadCharacter(isPlayer bool) Character {
	if isPlayer {
		if playerData.Name == "" {
			playerData = GeneratePlayer()
		}
		if !playerData.IsInit {
			playerData.InitCharacter()
		}
		return playerData
	} else {
		if monsterData.Name == "" {
			monsterData = GenerateMonster()
		}
		if !monsterData.IsInit {
			monsterData.InitCharacter()
		}
		return monsterData
	}
}

func ResetCharacterData() {
	playerData = GeneratePlayer()
	monsterData = GenerateMonster()
}
