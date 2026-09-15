//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import (
	"testing"
)

func TestMaximumLimitHealthEnergyCheck(t *testing.T) {
	bob := NewPlayer("Robert")

	bob.Heal(Potion * 2)
	if bob.health > bob.maxHealth {
		t.Fatalf("%s has %d health. However, Max Health is %d", bob.name, bob.health, bob.maxHealth)

	}

	bob.RestoreMagic(Ether * 2)
	if bob.energy > bob.maxEnergy {
		t.Fatalf("%s has %d energy. However, Max Energy is %d", bob.name, bob.energy, bob.maxEnergy)
	}

}

func TestMinimumLimitHealthEnergyCheck(t *testing.T) {
	bob := NewPlayer("Robert")

	bob.Damage(Potion * 2)
	if bob.health < 0 {
		t.Fatalf("%s has %d health. That is less than zero", bob.name, bob.health)

	}
	if bob.health > bob.maxHealth {
		t.Fatalf("%s has %d health. However, Max Health is %d", bob.name, bob.health, bob.maxHealth)

	}

	bob.UseMagic(Ether * 2)
	if bob.energy < 0 {
		t.Fatalf("%s has %d energy. That is less than zero", bob.name, bob.energy)
	}
	if bob.energy > bob.maxEnergy {
		t.Fatalf("%s has %d energy. However, Max Energy is %d", bob.name, bob.energy, bob.maxEnergy)
	}

}
