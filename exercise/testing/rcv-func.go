package main

import (
	"fmt"
)

type HealthPoints uint
type EnergyPoints uint
type Level uint

const (
	LevelUpBonus              = 2
	StartValue                = 8
	FirstLevel   Level        = 1
	Potion       HealthPoints = 10
	Ether        EnergyPoints = 10
)

type Player struct {
	health, maxHealth HealthPoints
	energy, maxEnergy EnergyPoints
	name              string
	level             Level
}

func NewPlayer(playerName string) *Player {
	var startHealth HealthPoints = StartValue + LevelUpBonus
	return &Player{
		health: startHealth, maxHealth: startHealth, energy: StartValue, maxEnergy: StartValue, name: playerName, level: FirstLevel,
	}
}

func (p *Player) Damage(damage HealthPoints) {
	if damage > p.health {
		p.health = 0
		fmt.Printf("%s fainted! Please use a potion\n", p.name)
	} else {
		p.health -= damage
	}
}

func (p *Player) Heal(potion HealthPoints) {
	if potion+p.health >= p.maxHealth {
		p.health = p.maxHealth
		fmt.Printf("%s has full health!\n", p.name)
	} else {
		p.health += potion
	}
}

func (p *Player) UseMagic(used EnergyPoints) {
	if used > p.energy {
		p.energy = 0
		fmt.Printf("%s has no energy! Please use an ether\n", p.name)
	} else {
		p.energy -= used
	}
}

func (p *Player) RestoreMagic(ether EnergyPoints) {
	if ether+p.energy >= p.maxEnergy {
		p.energy = p.maxEnergy
		fmt.Printf("%s has full energy!\n", p.name)
	} else {
		p.energy += ether
	}
}

func (p *Player) LevelUp() {
	p.level += 1
	p.maxHealth += HealthPoints((LevelUpBonus * 2) * p.level)
	p.health = p.maxHealth
	p.maxEnergy += EnergyPoints(LevelUpBonus * p.level)
	p.energy = p.maxEnergy
}

func (p *Player) CurrentHP() {
	fmt.Printf("%s currently has %d/%d Health Points\n", p.name, p.health, p.maxHealth)
}

func (p *Player) CurrentEP() {
	fmt.Printf("%s currently has %d/%d Energy Points\n", p.name, p.energy, p.maxEnergy)
}

func main() {
	rae := NewPlayer("rae")
	jack := NewPlayer("jack")

	fmt.Println(*rae, *jack)

	jack.Damage(2)
	jack.Damage(2)
	jack.Damage(2)
	jack.Damage(2)
	rae.Damage(11)
	jack.CurrentHP()
	jack.CurrentEP()
	rae.CurrentHP()
	rae.CurrentEP()
	fmt.Println(*rae, *jack)

	jack.Heal(Potion)
	jack.CurrentHP()
	jack.Heal(Potion)
	rae.UseMagic(5)
	rae.UseMagic(5)
	rae.RestoreMagic(Ether)
	jack.CurrentHP()
	jack.CurrentEP()
	rae.CurrentHP()
	rae.CurrentEP()
	fmt.Println(*rae, *jack)

	jack.LevelUp()
	rae.LevelUp()
	fmt.Println(*rae, *jack)
	rae.LevelUp()
	fmt.Println(*rae, *jack)

}
