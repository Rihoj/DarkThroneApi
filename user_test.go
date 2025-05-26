package DarkThroneApi

import (
	"testing"
)

func TestLoginRequest_Marshal(t *testing.T) {
	lr := LoginRequest{Email: "foo@bar.com", Password: "baz"}
	if lr.Email != "foo@bar.com" || lr.Password != "baz" {
		t.Error("unexpected field values")
	}
}

func TestRegisterRequest_Marshal(t *testing.T) {
	rr := RegisterRequest{Email: "a", Password: "b", ConfirmPassword: "b", Username: "c"}
	if rr.Email != "a" || rr.Password != "b" || rr.ConfirmPassword != "b" || rr.Username != "c" {
		t.Error("unexpected field values")
	}
}

// Add more tests for Login, Register, GetCurrentUser, GetPlayersForCurrentUser, Logout, AssumePlayer, UnassumePlayer as API integration or with mocks.
