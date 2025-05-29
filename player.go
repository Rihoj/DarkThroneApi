package DarkThroneApi

import (
	"errors"
	"fmt"
)

// Player represents a player in the Dark Throne game.
type Player struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Gold        int    `json:"gold"`
	Level       int    `json:"level"`
	ArmySize    int    `json:"armySize"`
	Units       []Unit `json:"units"`
	AttackTurns int    `json:"attackTurns"`
}

// Unit represents a unit in a player's army.
type Unit struct {
	UnitType string `json:"unitType"`
	Quantity int    `json:"quantity"`
}

// PlayersListResponse represents a paginated list of players.
type PlayersListResponse struct {
	Items []Player `json:"items"`
}

// AttackResponse represents the result of an attack action.
type AttackResponse struct {
	IsAttackerVictor bool `json:"isAttackerVictor"`
}

// CreatePlayerRequest represents the payload to create a player.
type CreatePlayerRequest struct {
	Name     string `json:"name"`
	Race     string `json:"race"`
	Password string `json:"password"`
}

// TrainUnitsRequest represents the payload to train units.
type TrainUnitsRequest struct {
	PlayerID string        `json:"playerId"`
	Units    []UnitRequest `json:"units"`
}

// TrainUnitsResponse represents the response from training units.
type TrainUnitsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UntrainUnitsRequest represents the payload to untrain units.
type UntrainUnitsRequest struct {
	PlayerID string        `json:"playerId"`
	Units    []UnitRequest `json:"units"`
}

// UntrainUnitsResponse represents the response from untraining units.
type UntrainUnitsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UnitRequest represents a unit and quantity for training/untraining.
type UnitRequest struct {
	UnitType string `json:"unitType"`
	Quantity int    `json:"quantity"`
}

// WarHistory represents a war history record.
type WarHistory struct {
	ID        string `json:"id"`
	PlayerID  string `json:"playerId"`
	Opponent  string `json:"opponent"`
	Result    string `json:"result"`
	Timestamp string `json:"timestamp"`
}

// GetPlayerByIndex retrieves a player by index from the user's player list and assumes that player.
// If the index is out of range, it returns an error.
func (d *DarkThroneApi) GetPlayerByIndex(index int) (Player, error) {
	logger := d.config.Logger
	logger.Debug("Fetching player list for selection...")
	if d.token == "" {
		logger.Error("Token is not set. Please ensure login() is called before making requests.")
		return Player{}, errors.New("token is not set")
	}

	players, err := d.getPlayersListAPI()
	if err != nil || len(players) == 0 {
		logger.Error("No players found in the response from auth/current-user/players")
		return Player{}, errors.New("no players found")
	}
	logger.Debug("Players list response", "players", players)

	if index < 0 || index >= len(players) {
		logger.Error("Player index out of range", "index", index, "player_count", len(players))
		return Player{}, fmt.Errorf("player index %d out of range (found %d players)", index, len(players))
	}

	playerID := players[index].ID
	if playerID == "" {
		logger.Error("Failed to set player_id from the players list")
		return Player{}, errors.New("failed to set player_id from the players list")
	}

	player, err := d.assumePlayerAPI(playerID)
	if err != nil {
		logger.Error("Failed to assume player", "error", err)
		return Player{}, fmt.Errorf("failed to assume player: %w", err)
	}
	return player, nil
}

// getPlayersListAPI fetches the list of players for the current user.
// It returns a slice of Player and an error if the request fails.
func (d *DarkThroneApi) getPlayersListAPI() ([]Player, error) {
	playersReq := ApiRequest[struct{}, UserPlayersListResponse]{
		Method:   "GET",
		Endpoint: playersListEndpoint,
		Headers:  d.getAuthHeaders(),
		Body:     struct{}{},
		Config:   d.apiConfig,
	}
	return playersReq.DoRequest()
}

// assumePlayerAPI assumes the given player ID and returns the Player.
// It sends a POST request to the assume player endpoint and returns the assumed Player or an error.
func (d *DarkThroneApi) assumePlayerAPI(playerID string) (Player, error) {
	payload := map[string]string{"playerID": playerID}
	assumeReq := ApiRequest[map[string]string, CurrentUserResponse]{
		Method:   "POST",
		Endpoint: assumePlayerEndpoint,
		Headers:  d.getAuthHeaders(),
		Body:     payload,
		Config:   d.apiConfig,
	}
	assumeResp, err := assumeReq.DoRequest()
	if err != nil {
		return Player{}, err
	}
	return assumeResp.Player, nil
}

// FetchAllPlayers fetches all players (paginated).
func (d *DarkThroneApi) FetchAllPlayers(page, pageSize int) ([]Player, error) {
	endpoint := fmt.Sprintf("players?page=%d&pageSize=%d", page, pageSize)
	response, err := ApiRequest[struct{}, PlayersListResponse]{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  d.getAuthHeaders(),
		Body:     struct{}{},
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}

// CreatePlayer creates a new player.
func (d *DarkThroneApi) CreatePlayer(req CreatePlayerRequest) (Player, error) {
	response, err := ApiRequest[CreatePlayerRequest, Player]{
		Method:   "POST",
		Endpoint: "players",
		Headers:  d.getAuthHeaders(),
		Body:     req,
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		return Player{}, err
	}
	return response, nil
}

// ValidatePlayerName validates a player name.
func (d *DarkThroneApi) ValidatePlayerName(name string) (bool, error) {
	payload := map[string]string{"name": name}
	response, err := ApiRequest[map[string]string, struct {
		Valid bool `json:"valid"`
	}]{
		Method:   "POST",
		Endpoint: "players/validate-name",
		Headers:  d.getAuthHeaders(),
		Body:     payload,
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		return false, err
	}
	return response.Valid, nil
}

// FetchPlayerByID fetches a player by ID.
func (d *DarkThroneApi) FetchPlayerByID(id string) (Player, error) {
	endpoint := fmt.Sprintf("players/%s", id)
	response, err := ApiRequest[struct{}, Player]{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  d.getAuthHeaders(),
		Body:     struct{}{},
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		return Player{}, err
	}
	return response, nil
}

// FetchAllMatchingIDs fetches all matching player IDs.
func (d *DarkThroneApi) FetchAllMatchingIDs(ids []string) ([]Player, error) {
	payload := map[string][]string{"ids": ids}
	response, err := ApiRequest[map[string][]string, struct {
		Players []Player `json:"players"`
	}]{
		Method:   "POST",
		Endpoint: "players/matching-ids",
		Headers:  d.getAuthHeaders(),
		Body:     payload,
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		return nil, err
	}
	return response.Players, nil
}

// FetchWarHistoryByID fetches war history by ID.
func (d *DarkThroneApi) FetchWarHistoryByID(id string) (WarHistory, error) {
	endpoint := fmt.Sprintf("war-history/%s", id)
	response, err := ApiRequest[struct{}, WarHistory]{
		Method:   "GET",
		Endpoint: endpoint,
		Headers:  d.getAuthHeaders(),
		Body:     struct{}{},
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		return WarHistory{}, err
	}
	return response, nil
}

// FetchAllWarHistory fetches all war history.
func (d *DarkThroneApi) FetchAllWarHistory() ([]WarHistory, error) {
	response, err := ApiRequest[struct{}, struct {
		Items []WarHistory `json:"items"`
	}]{
		Method:   "GET",
		Endpoint: "war-history",
		Headers:  d.getAuthHeaders(),
		Body:     struct{}{},
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}

// TrainUnits trains units for the current player.
func (d *DarkThroneApi) TrainUnits(req TrainUnitsRequest) (TrainUnitsResponse, error) {
	response, err := ApiRequest[TrainUnitsRequest, TrainUnitsResponse]{
		Method:   "POST",
		Endpoint: "training/train",
		Headers:  d.getAuthHeaders(),
		Body:     req,
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		return TrainUnitsResponse{}, err
	}
	return response, nil
}

// UntrainUnits untrains units for the current player.
func (d *DarkThroneApi) UntrainUnits(req UntrainUnitsRequest) (UntrainUnitsResponse, error) {
	response, err := ApiRequest[UntrainUnitsRequest, UntrainUnitsResponse]{
		Method:   "POST",
		Endpoint: "training/untrain",
		Headers:  d.getAuthHeaders(),
		Body:     req,
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		return UntrainUnitsResponse{}, err
	}
	return response, nil
}

// AttackPlayer attacks a player by ID.
// It returns true if the attack was successful, or false and an error otherwise.
func (d *DarkThroneApi) AttackPlayer(targetID string) (bool, error) {
	logger := d.config.Logger
	logger.Warn("Attacking player", "target_id", targetID)
	payload := map[string]any{
		"targetID":    targetID,
		"attackTurns": min_attack_turns,
	}
	response, err := ApiRequest[map[string]any, AttackResponse]{
		Method:   "POST",
		Endpoint: "attack",
		Headers:  d.getAuthHeaders(),
		Body:     payload,
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		logger.Error("Attack request failed", "error", err)
		return false, err
	}
	if response.IsAttackerVictor {
		logger.Warn("Attack successful", "target_id", targetID)
	} else {
		logger.Warn("Attack failed", "target_id", targetID)
	}
	return response.IsAttackerVictor, nil
}
