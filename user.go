package DarkThroneApi

import (
	"fmt"
)

// CurrentUserResponse represents the response for the current user API call.
type CurrentUserResponse struct {
	Player Player `json:"player"`
}

// UserPlayersListResponse represents a list of players for the current user.
type UserPlayersListResponse []Player

// LoginRequest represents the payload for a login request.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the response from a login request.
type LoginResponse struct {
	Session struct {
		Id                  string  `json:"id"`
		Email               string  `json:"email"`
		Player_id           *string `json:"playerID"`
		Has_confirmed_email bool    `json:"hasConfirmedEmail"`
		Server_time         string  `json:"serverTime"`
	} `json:"session"`
	Token string `json:"token"`
}

// Login authenticates the user and returns a token.
// Returns the authentication token or an error if login fails.
func (d *DarkThroneApi) Login(lr LoginRequest) (string, error) {
	logger := d.config.Logger
	logger.Info("Logging in...")

	if lr.Email == "" {
		logger.Error("Email is not set in LoginRequest")
		return "", fmt.Errorf("email not set")
	}

	if lr.Password == "" {
		logger.Error("Password is not set in LoginRequest")
		return "", fmt.Errorf("password not set")
	}

	req := ApiRequest[LoginRequest, LoginResponse]{
		Method:   "POST",
		Endpoint: loginEndpoint,
		Headers:  map[string]string{},
		Body:     lr,
		Config:   d.apiConfig,
	}
	response, err := req.DoRequest()
	if err != nil {
		logger.Error("Login failed", "error", err)
		return "", err
	}
	token := response.Token
	d.token = token // Store the token in the API instance for future requests
	logger.Info("Login successful. Token acquired.")
	return token, nil
}

// RegisterRequest represents the payload for user registration.
type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Username        string `json:"username"`
}

// RegisterResponse represents the response for user registration.
type RegisterResponse struct {
	Session struct {
		Id                  string  `json:"id"`
		Email               string  `json:"email"`
		Player_id           *string `json:"playerID"`
		Has_confirmed_email bool    `json:"hasConfirmedEmail"`
		Server_time         string  `json:"serverTime"`
	} `json:"session"`
	Token string `json:"token"`
}

// Register registers a new user.
func (d *DarkThroneApi) Register(req RegisterRequest) (RegisterResponse, error) {
	logger := d.config.Logger
	logger.Info("Registering new user...")

	if req.Email == "" {
		logger.Error("Email is not set in RegisterRequest")
		return RegisterResponse{}, fmt.Errorf("email not set")
	}
	if req.Password == "" {
		logger.Error("Password is not set in RegisterRequest")
		return RegisterResponse{}, fmt.Errorf("password not set")
	}
	if req.ConfirmPassword == "" {
		logger.Error("ConfirmPassword is not set in RegisterRequest")
		return RegisterResponse{}, fmt.Errorf("confirm password not set")
	}
	if req.Username == "" {
		logger.Error("Username is not set in RegisterRequest")
		return RegisterResponse{}, fmt.Errorf("username not set")
	}
	if req.Password != req.ConfirmPassword {
		logger.Error("Password and ConfirmPassword do not match")
		return RegisterResponse{}, fmt.Errorf("passwords do not match")
	}

	apiReq := ApiRequest[RegisterRequest, RegisterResponse]{
		Method:   "POST",
		Endpoint: "auth/register",
		Headers:  map[string]string{"Content-Type": "application/json"},
		Body:     req,
		Config:   d.apiConfig,
	}
	response, err := apiReq.DoRequest()
	if err != nil {
		logger.Error("Registration failed", "error", err)
		return RegisterResponse{}, err
	}
	logger.Info("Registration successful.")
	return response, nil
}

// GetCurrentUserAPI fetches the current user (not player) from the API.
// Returns the CurrentUserResponse or an error if the request fails.
func (d *DarkThroneApi) GetCurrentUserAPI() (CurrentUserResponse, error) {
	currentUserReq := ApiRequest[struct{}, CurrentUserResponse]{
		Method:   "GET",
		Endpoint: currentUserEndpoint,
		Headers:  map[string]string{},
		Body:     struct{}{},
		Config:   d.apiConfig,
	}
	return currentUserReq.DoRequest()
}

// GetCurrentUser fetches the current authenticated user.
// TODO: Move implementation from darkthrone.api.go and remove from there.
func (d *DarkThroneApi) GetCurrentUser() (CurrentUserResponse, error) {
	logger := d.config.Logger
	logger.Info("Fetching current authenticated user...")
	apiReq := ApiRequest[struct{}, CurrentUserResponse]{
		Method:   "GET",
		Endpoint: currentUserEndpoint,
		Headers:  d.getAuthHeaders(),
		Body:     struct{}{},
		Config:   d.apiConfig,
	}
	response, err := apiReq.DoRequest()
	if err != nil {
		logger.Error("Failed to fetch current user", "error", err)
		return CurrentUserResponse{}, err
	}
	return response, nil
}

// GetPlayersForCurrentUser fetches the list of players for the current user.
// It returns a slice of Player and an error if the request fails.
func (d *DarkThroneApi) GetPlayersForCurrentUser() ([]Player, error) {
	logger := d.config.Logger
	logger.Info("Fetching players for current user...")
	apiReq := ApiRequest[struct{}, UserPlayersListResponse]{
		Method:   "GET",
		Endpoint: playersListEndpoint,
		Headers:  d.getAuthHeaders(),
		Body:     struct{}{},
		Config:   d.apiConfig,
	}
	response, err := apiReq.DoRequest()
	if err != nil {
		logger.Error("Failed to fetch players for user", "error", err)
		return nil, err
	}
	return response, nil
}

// Logout logs out the current user.
// It clears the authentication token and returns an error if the logout fails.
func (d *DarkThroneApi) Logout() error {
	logger := d.config.Logger
	logger.Info("Logging out current user...")
	apiReq := ApiRequest[struct{}, struct{}]{
		Method:   "POST",
		Endpoint: "auth/logout",
		Headers:  d.getAuthHeaders(),
		Body:     struct{}{},
		Config:   d.apiConfig,
	}
	_, err := apiReq.DoRequest()
	if err != nil {
		logger.Error("Logout failed", "error", err)
		return err
	}
	d.token = "" // Clear token on logout
	logger.Info("Logout successful.")
	return nil
}

// AssumePlayer assumes the given player ID and returns the Player.
// It sends a POST request to the assume player endpoint and returns the assumed Player or an error.
func (d *DarkThroneApi) AssumePlayer(playerID string) (Player, error) {
	logger := d.config.Logger
	logger.Info("Assuming player", "playerID", playerID)
	payload := map[string]string{"playerID": playerID}
	apiReq := ApiRequest[map[string]string, CurrentUserResponse]{
		Method:   "POST",
		Endpoint: "auth/assume-player",
		Headers:  d.getAuthHeaders(),
		Body:     payload,
		Config:   d.apiConfig,
	}
	response, err := apiReq.DoRequest()
	if err != nil {
		logger.Error("Failed to assume player", "error", err)
		return Player{}, err
	}
	logger.Info("Player assumed successfully.")
	return response.Player, nil
}

// UnassumePlayer unassumes the current player.
// It sends a POST request to the unassume player endpoint and returns an error if the operation fails.
func (d *DarkThroneApi) UnassumePlayer() error {
	logger := d.config.Logger
	logger.Info("Unassuming current player...")
	apiReq := ApiRequest[struct{}, struct{}]{
		Method:   "POST",
		Endpoint: "auth/unassume-player",
		Headers:  d.getAuthHeaders(),
		Body:     struct{}{},
		Config:   d.apiConfig,
	}
	_, err := apiReq.DoRequest()
	if err != nil {
		logger.Error("Unassume player failed", "error", err)
		return err
	}
	logger.Info("Player unassumed successfully.")
	return nil
}
