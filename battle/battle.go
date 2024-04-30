package battle

import (
	"fmt"
	"math/rand/v2"
	"strconv"

	"github.com/oarielg/BattleGo/data"
)

var battleQuotes = map[string]string{
	"attackhit":       "%s attacked with %s for %d %s damage.",
	"attackcritical":  "Critical hit! %s attacked with %s for %d %s damage.",
	"attackmiss":      "%s attacked with %s but missed.",
	"attackextrahit":  "%s's attack did %d extra %s damage.",
	"attackextrafail": "%s's attack do extra %s damage but %s is immune!",
	"attacklifesteal": "%s recovered %d HP from their attack.",
	"runsuccess":      "%s escaped!",
	"runfail":         "%s tried to escape but %s blocked the way!",
}

type BattleState int

const (
	Win BattleState = iota
	Lose
	InProgress
	Escaped
)

type ConditionCheck struct {
	IsFree    bool
	CanCast   bool
	IsCharmed bool
}

type Battle struct {
	Player     data.Character
	Monster    data.Character
	BattleText []string
	State      BattleState
	Winner     string
}

type BattleData struct {
	PlayerName  string
	PlayerHP    string
	MonsterName string
	MonsterHP   string
	BattleState string
	Winner      string
	BattleText  []string
	Spells      map[string]string
}

func NewBattle(player, monster data.Character) *Battle {
	return &Battle{
		Player:     player,
		Monster:    monster,
		BattleText: []string{},
		State:      InProgress,
		Winner:     "",
	}
}

func (b *Battle) BattleTurn(action []string) {
	b.perTurnChecks(&b.Player, true)
	b.perTurnChecks(&b.Monster, false)

	if b.State == InProgress {
		b.playerTurn(action)
	}

	if b.Monster.CurrentHp > 0 && b.State == InProgress {
		if action[0] != "" {
			b.monsterTurn()
		} else {
			if init := initiative(b.Monster.Speed, b.Player.Speed); init {
				b.monsterTurn()
			}
		}
	}

	b.checkDeath(&b.Player, true)
	b.checkDeath(&b.Monster, false)
}

func (b *Battle) GetBattleData() BattleData {
	spells := make(map[string]string)
	for _, spell := range b.Player.Spells {
		k := strconv.Itoa(spell.Id)
		spells[k] = spell.Name
	}
	return BattleData{
		PlayerName:  b.Player.Name,
		PlayerHP:    strconv.Itoa(b.Player.CurrentHp),
		MonsterName: b.Monster.Name,
		MonsterHP:   strconv.Itoa(b.Monster.CurrentHp),
		BattleState: strconv.Itoa(int(b.State)),
		Winner:      b.Winner,
		BattleText:  b.BattleText,
		Spells:      spells,
	}
}

func (b *Battle) perTurnChecks(char *data.Character, is_player bool) {
	if len(char.Conditions) > 0 {
		for i := len(char.Conditions) - 1; i >= 0; i-- {
			char.Conditions[i].Duration--
			if char.Conditions[i].Duration <= 0 {
				b.addText(fmt.Sprintf(char.Conditions[i].Condition.EndQuote, char.Name))
				if len(char.Conditions) == 1 {
					char.Conditions = []data.ActiveCondition{}
				} else {
					char.Conditions = append(char.Conditions[:i], char.Conditions[i+1:]...)
				}
			} else {
				switch char.Conditions[i].Condition.Type {
				case data.NoActions, data.NoOffensiveAction, data.NoSpells:
					b.addText(fmt.Sprintf(char.Conditions[i].Condition.ActiveQuote, char.Name))
				case data.DoDamage:
					if hasDamageImmunity(char.Immunities, char.Conditions[i].Condition.DamageType) {
						b.addText(fmt.Sprintf(char.Conditions[i].Condition.FailQuote, char.Name))
					} else {
						resistance := hasDamageResistance(char.Resistances, char.Conditions[i].Condition.DamageType)
						damage := roll(1, 3)
						if hasDamageVulnerability(char.Vulnerabilities, char.Conditions[i].Condition.DamageType) {
							damage += roll(1, 3)
						}
						damage = max((damage - resistance), 0)
						char.CurrentHp -= damage
						b.addText(fmt.Sprintf(char.Conditions[i].Condition.ActiveQuote, char.Name, damage))
						b.checkDeath(char, is_player)
					}
				}
			}
		}
	}
}

func (b *Battle) playerTurn(action []string) {
	conditionCheck := ConditionCheck{
		IsFree:    true,
		CanCast:   true,
		IsCharmed: false,
	}
	checkConditionAction(&conditionCheck, b.Player.Conditions)

	switch action[0] {
	case "attack":
		b.attackTurn(conditionCheck, &b.Player, &b.Monster)
	case "spell":
		id, err := strconv.Atoi(action[1])
		if err != nil {
			panic("Failed converting spell id! " + err.Error())
		}
		spell := data.Spells[id]
		b.spellTurn(conditionCheck, spell, &b.Player, &b.Monster)
	case "run":
		b.runTurn(conditionCheck)
	}
}

func (b *Battle) monsterTurn() {
	conditionCheck := ConditionCheck{
		IsFree:    true,
		CanCast:   true,
		IsCharmed: false,
	}
	checkConditionAction(&conditionCheck, b.Monster.Conditions)

	switch b.Monster.Class {
	case data.Fighter:
		b.attackTurn(conditionCheck, &b.Monster, &b.Player)
	case data.Caster:
		b.monsterSpellTurn(conditionCheck)
	case data.Hybrid:
		if roll(1, 10) <= 6 {
			b.attackTurn(conditionCheck, &b.Monster, &b.Player)
		} else {
			b.monsterSpellTurn(conditionCheck)
		}
	}
}

func (b *Battle) attackTurn(conditionCheck ConditionCheck, attacker, defender *data.Character) {
	if !conditionCheck.IsFree || conditionCheck.IsCharmed {
		return
	}
	b.attackAction(attacker.PrimarySlot, attacker, defender)
	if attacker.SecondarySlot.Name != "" {
		if attacker.SecondarySlot.Type == data.Weapon {
			b.attackAction(attacker.SecondarySlot, attacker, defender)
		}
	}
}

func (b *Battle) spellTurn(conditionCheck ConditionCheck, spell data.Spell, attacker, defender *data.Character) {
	if !conditionCheck.IsFree || !conditionCheck.CanCast {
		return
	}

	if ok, i := hasEffectOn(defender.Effects, data.Counterspell); ok {
		if spell.Type == data.DirectDamage ||
			spell.Type == data.CauseCondition ||
			spell.Type == data.Drain {
			defender.Effects = append(defender.Effects[:i], defender.Effects[i+1:]...)
			b.addText(fmt.Sprintf("%s cast a spell, but it got counterspelled by %s!", attacker.Name, defender.Name))
			return
		}
	}

	switch spell.Type {
	case data.Healing:
		if attacker.CurrentHp == attacker.MaxHp {
			b.addText(fmt.Sprintf(spell.MissQuote, attacker.Name))
		} else {
			heal := roll(1, 6) + spell.Variable + attacker.Power
			new_hp := attacker.CurrentHp + heal
			if new_hp > attacker.MaxHp {
				heal = attacker.MaxHp - attacker.CurrentHp
				new_hp = attacker.CurrentHp + heal
			}
			attacker.CurrentHp = new_hp
			b.addText(fmt.Sprintf(spell.HitQuote, attacker.Name, heal))
		}
	case data.CauseCondition:
		if conditionCheck.IsCharmed {
			return
		}
		chance := check(attacker.Power, defender.Mind)
		dice := roll(1, 20)
		if dice >= chance || dice == 20 {
			condition := data.Conditions[spell.Condition]
			if hasConditionImmunity(defender.CondImmunities, condition) {
				b.addText(fmt.Sprintf(spell.ExtraQuote, attacker.Name, defender.Name))
			} else {
				if hasCondition(defender.Conditions, condition) {
					b.addText(fmt.Sprintf(spell.FailQuote, attacker.Name, defender.Name))
				} else {
					duration := roll(condition.MinDuration, condition.MaxDuration)
					defender.Conditions = append(defender.Conditions, data.ActiveCondition{Duration: duration, Condition: condition})
					b.addText(fmt.Sprintf(spell.HitQuote, attacker.Name, defender.Name))
				}
			}
		} else {
			b.addText(fmt.Sprintf(spell.MissQuote, attacker.Name, defender.Name))
		}
	case data.DirectDamage:
		if conditionCheck.IsCharmed {
			return
		}
		chance := check(attacker.Power, defender.Speed)
		dice := roll(1, 20)
		if dice >= chance || dice == 20 {
			if hasDamageImmunity(defender.Immunities, spell.DamageType) {
				b.addText(fmt.Sprintf(spell.FailQuote, attacker.Name, defender.Name))
			} else {
				resistance := hasDamageResistance(defender.Resistances, spell.DamageType)
				damage := roll(1, 6) + attacker.Power + spell.Variable
				if hasDamageVulnerability(defender.Vulnerabilities, spell.DamageType) {
					damage += roll(1, 6)
				}
				damage = max((damage - resistance), 0)
				defender.CurrentHp -= damage
				b.addText(fmt.Sprintf(spell.HitQuote, attacker.Name, damage))
			}
		} else {
			b.addText(fmt.Sprintf(spell.MissQuote, attacker.Name, defender.Name))
		}
	case data.Drain:
		if conditionCheck.IsCharmed {
			return
		}
		chance := check(attacker.Power, defender.Speed)
		dice := roll(1, 20)
		if dice >= chance || dice == 20 {
			if hasDamageImmunity(defender.Immunities, spell.DamageType) {
				b.addText(fmt.Sprintf(spell.FailQuote, attacker.Name, defender.Name))
			} else {
				resistance := hasDamageResistance(defender.Resistances, spell.DamageType)
				damage := roll(1, 6) + attacker.Power + spell.Variable
				if hasDamageVulnerability(defender.Vulnerabilities, spell.DamageType) {
					damage += roll(1, 6)
				}
				damage = max((damage - resistance), 0)
				defender.CurrentHp -= damage
				heal := damage / 2
				new_hp := attacker.CurrentHp + heal
				if new_hp > attacker.MaxHp {
					heal = attacker.MaxHp - attacker.CurrentHp
					new_hp = attacker.CurrentHp + heal
				}
				attacker.CurrentHp = new_hp
				b.addText(fmt.Sprintf(spell.HitQuote, attacker.Name, damage, heal))
			}
		} else {
			b.addText(fmt.Sprintf(spell.MissQuote, attacker.Name, defender.Name))
		}
	case data.ActivateEffect:
		chance := check(attacker.Power, 1)
		dice := roll(1, 20)
		if dice >= chance || dice == 20 {
			if ok, _ := hasEffectOn(attacker.Effects, spell.Effect); !ok {
				attacker.Effects = append(attacker.Effects, spell.Effect)
				b.addText(fmt.Sprintf(spell.HitQuote, attacker.Name))
			} else {
				b.addText(fmt.Sprintf(spell.FailQuote, attacker.Name))
			}
		} else {
			b.addText(fmt.Sprintf(spell.MissQuote, attacker.Name))
		}
	}
}

func (b *Battle) runTurn(conditionCheck ConditionCheck) {
	if !conditionCheck.IsFree || conditionCheck.IsCharmed {
		return
	}
	to_run := roll(1, 20) + b.Player.Speed + b.Player.Armor.Initiative
	to_block := roll(1, 20) + b.Monster.Speed + b.Monster.Armor.Initiative
	if to_run >= to_block {
		b.addText(fmt.Sprintf(battleQuotes["runsuccess"], b.Player.Name))
		b.endBattle(Escaped, "")
	} else {
		b.addText(fmt.Sprintf(battleQuotes["runfail"], b.Player.Name, b.Monster.Name))
	}
}

func (b *Battle) monsterSpellTurn(conditionCheck ConditionCheck) {
	if !conditionCheck.CanCast {
		b.attackTurn(conditionCheck, &b.Monster, &b.Player)
	} else {
		chosenSpell := pickSpell(b.Monster.Spells)
		if chosenSpell.Name != "" {
			b.spellTurn(conditionCheck, chosenSpell, &b.Monster, &b.Player)
		} else {
			b.attackTurn(conditionCheck, &b.Monster, &b.Player)
		}
	}
}

func (b *Battle) attackAction(weapon data.Item, attacker, defender *data.Character) {
	attack := attacker.Might
	damage := roll(1, 6) + attacker.Might
	weapon_name := "Fists"

	if weapon.Name != "" {
		attack += weapon.Attack
		damage += weapon.Damage
		weapon_name = weapon.Name
	}

	defense := defender.Speed
	armor := 0

	if defender.SecondarySlot.Name != "" {
		if defender.SecondarySlot.Type == data.Shield {
			defense += defender.SecondarySlot.Defense
			armor += defender.SecondarySlot.Armor
		}
	}

	if defender.Armor.Name != "" {
		defense += defender.Armor.Defense
		armor += defender.Armor.Armor
	}

	check := check(attack, defense)
	dice := roll(1, 20)
	if dice >= check || dice == 20 {
		if dice == 20 {
			damage *= 2
			b.addText(fmt.Sprintf(battleQuotes["attackcritical"], attacker.Name, weapon_name, damage, weapon.DamageType.String()))
		} else {
			b.addText(fmt.Sprintf(battleQuotes["attackhit"], attacker.Name, weapon_name, damage, weapon.DamageType.String()))
		}
		damage = max(damage-armor, 0)
		defender.CurrentHp -= damage

		if len(weapon.Enchantments) > 0 {
			for _, e := range weapon.Enchantments {
				en := data.Enchantments[e]
				if en.Type == data.Imbuement {
					if hasDamageImmunity(defender.Immunities, en.DamageType) {
						b.addText(fmt.Sprintf(battleQuotes["attackextrafail"], attacker.Name, en.DamageType.String(), defender.Name))
					} else {
						resistance := hasDamageResistance(defender.Resistances, en.DamageType)
						damage := roll(1, 3)
						if hasDamageVulnerability(defender.Vulnerabilities, en.DamageType) {
							damage += roll(1, 3)
						}
						damage = max((damage - resistance), 0)
						defender.CurrentHp -= damage
						b.addText(fmt.Sprintf(battleQuotes["attackextrahit"], attacker.Name, damage, en.DamageType.String()))
					}
				}
			}
		}
	} else {
		b.addText(fmt.Sprintf(battleQuotes["attackmiss"], attacker.Name, weapon_name))
	}
}

func (b *Battle) checkDeath(char *data.Character, is_player bool) {
	if char.CurrentHp <= 0 {
		if ok, i := hasEffectOn(char.Effects, data.NoDeath); ok {
			char.CurrentHp = 1
			char.Effects = append(char.Effects[:i], char.Effects[i+1:]...)
			b.addText(fmt.Sprintf("The Reaper spared %s's life this once!", char.Name))
		} else {
			state := InProgress
			winner := ""
			if is_player {
				state = Lose
				winner = b.Monster.Name
			} else {
				state = Win
				winner = b.Player.Name
			}
			b.endBattle(state, winner)
		}
	}
}

func (b *Battle) endBattle(state BattleState, winner string) {
	b.State = state
	b.Winner = winner
}

func (b *Battle) addText(text string) {
	b.BattleText = append(b.BattleText, text)
}

func pickSpell(spells []data.Spell) data.Spell {
	if len(spells) == 0 {
		return data.Spell{}
	}
	return spells[rand.IntN(len(spells))]
}

func hasConditionImmunity(condImmunities []int, condition data.Condition) bool {
	if len(condImmunities) == 0 {
		return false
	}
	for _, c := range condImmunities {
		if c == condition.Id {
			return true
		}
	}
	return false
}

func hasCondition(activeConditions []data.ActiveCondition, condition data.Condition) bool {
	if len(activeConditions) == 0 {
		return false
	}
	for _, c := range activeConditions {
		if c.Condition.Id == condition.Id {
			return true
		}
	}
	return false
}

func hasDamageImmunity(immunities []data.DamageType, damageType data.DamageType) bool {
	if len(immunities) == 0 {
		return false
	}
	for _, i := range immunities {
		if i == damageType {
			return true
		}
	}
	return false
}

func hasDamageResistance(resistances []data.DamageResistance, damageType data.DamageType) int {
	if len(resistances) == 0 {
		return 0
	}
	for _, r := range resistances {
		if r.DamageType == damageType {
			return r.Resistance
		}
	}
	return 0
}

func hasDamageVulnerability(vulnerabilities []data.DamageType, damageType data.DamageType) bool {
	if len(vulnerabilities) == 0 {
		return false
	}
	for _, v := range vulnerabilities {
		if v == damageType {
			return true
		}
	}
	return false
}

func hasEffectOn(effects []data.BattleEffect, effect data.BattleEffect) (bool, int) {
	if len(effects) == 0 {
		return false, 0
	}
	for i, e := range effects {
		if e == effect {
			return true, i
		}
	}
	return false, 0
}

func checkConditionAction(conditionCheck *ConditionCheck, conditions []data.ActiveCondition) {
	if len(conditions) == 0 {
		return
	}
	for _, c := range conditions {
		if c.Condition.Type == data.NoActions {
			conditionCheck.IsFree = false
		}
		if c.Condition.Type == data.NoSpells {
			conditionCheck.CanCast = false
		}
		if c.Condition.Type == data.NoOffensiveAction {
			conditionCheck.IsCharmed = true
		}
	}
}

func check(active, passive int) int {
	return min(20, max(1, (8-active)+passive))
}

func roll(min, max int) int {
	return rand.IntN(max+1-min) + min
}

func initiative(active, passive int) bool {
	active = roll(1, 20) + active
	passive = roll(1, 20) + passive
	return active >= passive
}
