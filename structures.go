package DarkThroneApi

import "fmt"

// UpgradeStructureRequest represents the payload to upgrade a structure.
type UpgradeStructureRequest struct {
	StructureID  string `json:"structureId"`
	UpgradeLevel int    `json:"upgradeLevel"`
}

// UpgradeStructureResponse represents the response from upgrading a structure.
type UpgradeStructureResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	StructureID string `json:"structureId"`
	NewLevel    int    `json:"newLevel"`
}

// ProficiencyPointsRequest represents the payload to spend proficiency points.
type ProficiencyPointsRequest struct {
	PlayerID        string `json:"playerId"`
	PointsToSpend   int    `json:"pointsToSpend"`
	ProficiencyType string `json:"proficiencyType"`
}

// ProficiencyPointsResponse represents the response from spending proficiency points.
type ProficiencyPointsResponse struct {
	Success         bool   `json:"success"`
	Message         string `json:"message"`
	RemainingPoints int    `json:"remainingPoints"`
}

// UpgradeStructure upgrades a structure for the current player.
// Returns an error indicating the feature is not released yet. When released, this will POST to the structures/upgrade endpoint.
func (d *DarkThroneApi) UpgradeStructure(req UpgradeStructureRequest) (UpgradeStructureResponse, error) {
	return UpgradeStructureResponse{}, fmt.Errorf("structure upgrades are not released yet")
	// Uncomment below when feature is released:
	// apiReq := ApiRequest[UpgradeStructureRequest, UpgradeStructureResponse]{
	// 	Method:   "POST",
	// 	Endpoint: "structures/upgrade",
	// 	Headers:  d.getAuthHeaders(),
	// 	Body:     req,
	// 	Config:   d.apiConfig,
	// }
	// response, err := apiReq.DoRequest()
	// if err != nil {
	// 	return UpgradeStructureResponse{}, err
	// }
	// return response, nil
}

// SpendProficiencyPoints spends proficiency points for the current player.
// Returns an error indicating the feature is not released yet. When released, this will POST to the proficiency-points endpoint.
func (d *DarkThroneApi) SpendProficiencyPoints(req ProficiencyPointsRequest) (ProficiencyPointsResponse, error) {
	return ProficiencyPointsResponse{}, fmt.Errorf("proficiency points are not released yet")
	// Uncomment below when feature is released:
	// apiReq := ApiRequest[ProficiencyPointsRequest, ProficiencyPointsResponse]{
	// 	Method:   "POST",
	// 	Endpoint: "proficiency-points",
	// 	Headers:  d.getAuthHeaders(),
	// 	Body:     req,
	// 	Config:   d.apiConfig,
	// }
	// response, err := apiReq.DoRequest()
	// if err != nil {
	// 	return ProficiencyPointsResponse{}, err
	// }
	// return response, nil
}
