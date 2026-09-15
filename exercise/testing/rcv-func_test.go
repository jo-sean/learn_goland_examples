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

package testing

import (
	"testing"
)

func maximumLimitHealthEnergyCheck(t testing.T) {
	bob := NewPlayer("Robert")

	for i := 0; i < 1; i++ {
		bob.Heal(Potion)
		if bob.health > bob.maxHealth {
			t.Errorf("%s has %d health. However, Max Health is %d", bob.name, bob.health, bob.maxHealth)

		}
	}

	for i := 0; i < 1; i++ {
		bob.RestoreMagic(Ether)
		if bob.energy > bob.maxEnergy {
			t.Errorf("%s has %d energy. However, Max Energy is %d", bob.name, bob.energy, bob.maxEnergy)
		}
	}
}

func minimumLimitHealthEnergyCheck(t testing.T) {
	bob := NewPlayer("Robert")

	for i := 0; i < 2; i++ {
		bob.Damage(Potion)
		if bob.health < 0 {
			t.Errorf("%s has %d health. That is less than zero", bob.name, bob.health)

		}
	}

	for i := 0; i < 2; i++ {
		bob.RestoreMagic(Ether)
		if bob.energy < 0 {
			t.Errorf("%s has %d energy. That is less than zero", bob.name, bob.energy)
		}
	}
}
