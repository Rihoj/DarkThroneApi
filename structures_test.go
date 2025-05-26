package DarkThroneApi

import (
	"testing"
)

func TestUpgradeStructureRequest_Fields(t *testing.T) {
	r := UpgradeStructureRequest{StructureID: "sid", UpgradeLevel: 2}
	if r.StructureID != "sid" || r.UpgradeLevel != 2 {
		t.Error("unexpected field values")
	}
}

func TestUpgradeStructureResponse_Fields(t *testing.T) {
	r := UpgradeStructureResponse{Success: true, Message: "msg", StructureID: "sid", NewLevel: 3}
	if !r.Success || r.Message != "msg" || r.StructureID != "sid" || r.NewLevel != 3 {
		t.Error("unexpected field values")
	}
}

func TestProficiencyPointsRequest_Fields(t *testing.T) {
	r := ProficiencyPointsRequest{PlayerID: "pid", PointsToSpend: 5, ProficiencyType: "foo"}
	if r.PlayerID != "pid" || r.PointsToSpend != 5 || r.ProficiencyType != "foo" {
		t.Error("unexpected field values")
	}
}

func TestProficiencyPointsResponse_Fields(t *testing.T) {
	r := ProficiencyPointsResponse{Success: true, Message: "bar", RemainingPoints: 7}
	if !r.Success || r.Message != "bar" || r.RemainingPoints != 7 {
		t.Error("unexpected field values")
	}
}

// Add more tests for UpgradeStructure and SpendProficiencyPoints as API integration or with mocks.
