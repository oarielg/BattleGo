package data

type EnchantmentType int

const (
	AttributeBonus EnchantmentType = iota
	ConditionImmunity
	ResistanceDamage
	ImmunityDamage
	Imbuement
	VampiricWeapon
)

type Enchantment struct {
	Id         int
	Name       string
	Type       EnchantmentType
	Variable   int
	DamageType DamageType
}

var Enchantments []Enchantment = []Enchantment{
	{},
	{
		Id:         1,
		Name:       "Flaming",
		Type:       Imbuement,
		Variable:   0,
		DamageType: Fire,
	},
	{
		Id:         2,
		Name:       "Frozen",
		Type:       Imbuement,
		Variable:   0,
		DamageType: Cold,
	},
	{
		Id:         3,
		Name:       "Electrifying",
		Type:       Imbuement,
		Variable:   0,
		DamageType: Thunder,
	},
	{
		Id:         4,
		Name:       "Earth Attuned",
		Type:       Imbuement,
		Variable:   0,
		DamageType: Earth,
	},
	{
		Id:         5,
		Name:       "Coruscating",
		Type:       Imbuement,
		Variable:   0,
		DamageType: Radiant,
	},
	{
		Id:         6,
		Name:       "Necromantic",
		Type:       Imbuement,
		Variable:   0,
		DamageType: Necrotic,
	},
	{
		Id:         7,
		Name:       "Mystic",
		Type:       Imbuement,
		Variable:   0,
		DamageType: Psychic,
	},
	{
		Id:         8,
		Name:       "Vampiric",
		Type:       VampiricWeapon,
		Variable:   0,
		DamageType: 0,
	},
}
