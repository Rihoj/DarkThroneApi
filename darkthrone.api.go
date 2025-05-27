package DarkThroneApi

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

const (
	loginEndpoint        = "auth/login"
	currentUserEndpoint  = "auth/current-user"
	playersListEndpoint  = "auth/current-user/players"
	assumePlayerEndpoint = "auth/assume-player"

	min_attack_turns = 10
	page_size        = 100
)

var (
	instance *DarkThroneApi
	once     sync.Once
)

// Config holds configuration for the DarkThroneApi client, such as the logger.
type Config struct {
	Logger *slog.Logger
}

// DarkThroneApi is the main client for interacting with the Dark Throne API.
type DarkThroneApi struct {
	config    *Config
	token     string
	apiConfig *ApiRequestConfig
}

// New creates a new instance of DarkThroneApi with the provided configuration.
func New(config *Config) *DarkThroneApi {
	once.Do(func() {
		instance = &DarkThroneApi{
			config: config,
			apiConfig: &ApiRequestConfig{
				BaseURL: "https://api.darkthronereborn.com",
				Logger:  config.Logger,
			},
		}
	})
	return instance
}

// GetInstance returns the singleton instance of DarkThroneApi. Panics if not initialized.
func GetInstance() *DarkThroneApi {
	if instance == nil {
		panic("DarkThroneApi instance is not initialized. Call New() first.")
	}
	return instance
}

// getAuthHeaders returns a map of headers including the Authorization header if the token is set.
func (d *DarkThroneApi) getAuthHeaders() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	if d.token != "" {
		headers["Authorization"] = "Bearer " + d.token
	}
	return headers
}

// Ping checks if the Dark Throne API server can be reached by making a HEAD request to the base URL.
// It returns the latency in milliseconds if successful, or an error if not.
func (d *DarkThroneApi) Ping() (latencyMs int64, err error) {
	url := d.apiConfig.BaseURL
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		if d.config != nil && d.config.Logger != nil {
			d.config.Logger.Error("Ping request creation failed", "error", err)
		}
		return 0, err
	}
	start := time.Now()
	client := &http.Client{}
	resp, err := client.Do(req)
	latency := time.Since(start).Milliseconds()
	if err != nil {
		if d.config != nil && d.config.Logger != nil {
			d.config.Logger.Error("Ping failed", "error", err)
		}
		return latency, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		err = fmt.Errorf("ping failed with status: %s", resp.Status)
		if d.config != nil && d.config.Logger != nil {
			d.config.Logger.Error("Ping failed", "status", resp.Status)
		}
		return latency, err
	}
	if d.config != nil && d.config.Logger != nil {
		d.config.Logger.Info("Ping successful", "url", url, "status", resp.Status, "latency_ms", latency)
	}
	return latency, nil
}
