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
	Id        int
	Name      string
	Type      EnchantmentType
	Variable  int
	Variable2 int
}
