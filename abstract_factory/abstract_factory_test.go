package abstract_factory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	elfKingdam := NewElfKingdomFactory()
	elfCastle  := elfKingdam.createCastle()
	elfKing    := elfKingdam.createKing()
	elfArmy    := elfKingdam.createArmy()

	elfCastleExpect := "This is the Elven Castle!"
	if elfCastle.getDescription() != elfCastleExpect {
		t.Errorf("Expect output to %s, but %s\n", elfCastleExpect, elfCastle.getDescription())
	}

	elfKingExpect := "This is the Eleven King!"
	if elfKing.getDescription() != elfKingExpect {
		t.Errorf("Expect output to %s, but %s\n", elfKingExpect, elfKing.getDescription())
	}

	elfArmyExpect := "This is the Elven Army!"
	if elfArmy.getDescription() != elfArmyExpect {
		t.Errorf("Expect output to %s, but %s\n", elfArmyExpect, elfArmy.getDescription())
	}

	orkKingdam := NewOrcKingdomFactory()
	orkCastle  := orkKingdam.createCastle()
	orkKing    := orkKingdam.createKing()
	orkArmy    := orkKingdam.createArmy()
	orkCastleExpect := "This is the Orc Castle!"
	if orkCastle.getDescription() != orkCastleExpect {
		t.Errorf("Expect output to %s, but %s\n", orkCastleExpect, orkCastle.getDescription())
	}

	orkKingExpect := "This is the Orc King!"
	if orkKing.getDescription() != orkKingExpect {
		t.Errorf("Expect output to %s, but %s\n", orkKingExpect, orkKing.getDescription())
	}

	orkArmyExpect := "This is the Orc Army!"
	if orkArmy.getDescription() != orkArmyExpect {
		t.Errorf("Expect output to %s, but %s\n", orkArmyExpect, orkArmy.getDescription())
	}
}