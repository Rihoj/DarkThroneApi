package DarkThroneApi

import (
	"testing"
)

func TestPlayer_Fields(t *testing.T) {
	p := Player{ID: "id", Name: "n", Gold: 1, Level: 2, ArmySize: 3, Units: []Unit{{UnitType: "foo", Quantity: 1}}, AttackTurns: 4}
	if p.ID != "id" || p.Name != "n" || p.Gold != 1 || p.Level != 2 || p.ArmySize != 3 || p.AttackTurns != 4 {
		t.Error("unexpected field values")
	}
	if p.Units[0].UnitType != "foo" || p.Units[0].Quantity != 1 {
		t.Error("unexpected unit values")
	}
}

func TestUnit_Fields(t *testing.T) {
	u := Unit{UnitType: "bar", Quantity: 2}
	if u.UnitType != "bar" || u.Quantity != 2 {
		t.Error("unexpected unit field values")
	}
}

// Add more tests for FetchAllPlayers, CreatePlayer, ValidatePlayerName, FetchPlayerByID, etc. as API integration or with mocks.
