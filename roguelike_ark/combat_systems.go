package main

import (
	"fmt"
)

func AttackSystem(g *Game, attackerPosition *Position, defenderPosition *Position) {
	var attacker *ActorInstance = nil
	var defender *ActorInstance = nil

	// Position, Armor, Health, Name, MeleeWeapon, UserMessage
	query := g.Context.Combatants.Query()

	for query.Next() {
		//pos, armor, health, name, weapon, msg := query.Get()

		combatant := g.CombatantInstance(query.Entity())
		if combatant == nil {
			continue
		}

		if combatant.Pos.IsEqual(attackerPosition) {
			attacker = combatant
		} else if combatant.Pos.IsEqual(defenderPosition) {
			defender = combatant
		}
	}

	if attacker == nil || defender == nil {
		return
	}

	defenderArmor := defender.Armor
	defenderHealth := defender.Health
	defenderName := defender.Name.Label
	defenderMessage := defender.Message

	attackerWeapon := attacker.Weapon
	attackerName := attacker.Name.Label
	attackerMessage := attacker.Message

	if attacker.Health.CurrentHealth <= 0 {
		return
	}

	toHitRoll := GetDiceRoll(10)

	if toHitRoll+attackerWeapon.ToHitBonus >= defenderArmor.ArmorClass {
		damageRoll := GetRandomBetween(attackerWeapon.MinDamage, attackerWeapon.MaxDamage)
		damageDone := damageRoll - defenderArmor.Defense

		if damageDone < 0 {
			damageDone = 0
		}
		defenderHealth.CurrentHealth -= damageDone

		attackerMessage.AttackMessage = fmt.Sprintf("%s swings %s at %s and hits for %d health.\n", attackerName, attackerWeapon.Name, defenderName, damageDone)

		if defenderHealth.CurrentHealth <= 0 {
			defenderMessage.DeadMessage = fmt.Sprintf("%s has died!\n", defenderName)
			if defenderName == "Player" {
				defenderMessage.GameStateMessage = "Game Over!\n"
				g.Turn = GameOver
			}
			//g.World.DisposeEntities(defender.Entity)
		}
	} else {
		attackerMessage.AttackMessage = fmt.Sprintf("%s swings %s at %s and misses. \n", attackerName, attackerWeapon.Name, defenderName)
	}
}
