//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import (
	"fmt"
)

//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

type HealthPoints uint
type EnergyPoints uint
type Level uint

const (
	levelUpBonus              = 2
	startValue                = 8
	firstLevel   Level        = 1
	potion       HealthPoints = 10
	ether        EnergyPoints = 10
)

type Player struct {
	health, maxHealth HealthPoints
	energy, maxEnergy EnergyPoints
	name              string
	level             Level
}

func newPlayer(playerName string) *Player {
	var startHealth HealthPoints = startValue + levelUpBonus
	return &Player{
		health: startHealth, maxHealth: startHealth, energy: startValue, maxEnergy: startValue, name: playerName, level: firstLevel,
	}
}

func (p *Player) damage(damage HealthPoints) {
	if damage > p.health {
		p.health = 0
		fmt.Printf("%s fainted! Please use a potion\n", p.name)
	} else {
		p.health -= damage
	}
}

func (p *Player) heal(potion HealthPoints) {
	if potion+p.health >= p.maxHealth {
		p.health = p.maxHealth
		fmt.Printf("%s has full health!\n", p.name)
	} else {
		p.health += potion
	}
}

func (p *Player) useMagic(used EnergyPoints) {
	if used > p.energy {
		p.energy = 0
		fmt.Printf("%s has no energy! Please use an ether\n", p.name)
	} else {
		p.energy -= used
	}
}

func (p *Player) restoreMagic(ether EnergyPoints) {
	if ether+p.energy >= p.maxEnergy {
		p.energy = p.maxEnergy
		fmt.Printf("%s has full energy!\n", p.name)
	} else {
		p.energy += ether
	}
}

func (p *Player) levelUp() {
	p.level += 1
	p.maxHealth += HealthPoints((levelUpBonus * 2) * p.level)
	p.health = p.maxHealth
	p.maxEnergy += EnergyPoints(levelUpBonus * p.level)
	p.energy = p.maxEnergy
}

func (p *Player) currentHP() {
	fmt.Printf("%s currently has %d/%d Health Points\n", p.name, p.health, p.maxHealth)
}

func (p *Player) currentEP() {
	fmt.Printf("%s currently has %d/%d Energy Points\n", p.name, p.energy, p.maxEnergy)
}

func main() {
	rae := newPlayer("rae")
	jack := newPlayer("jack")

	fmt.Println(*rae, *jack)

	jack.damage(2)
	jack.damage(2)
	jack.damage(2)
	jack.damage(2)
	rae.damage(11)
	jack.currentHP()
	jack.currentEP()
	rae.currentHP()
	rae.currentEP()
	fmt.Println(*rae, *jack)

	jack.heal(potion)
	jack.currentHP()
	jack.heal(potion)
	rae.useMagic(5)
	rae.useMagic(5)
	rae.restoreMagic(ether)
	jack.currentHP()
	jack.currentEP()
	rae.currentHP()
	rae.currentEP()
	fmt.Println(*rae, *jack)

	jack.levelUp()
	rae.levelUp()
	fmt.Println(*rae, *jack)
	rae.levelUp()
	fmt.Println(*rae, *jack)

}
