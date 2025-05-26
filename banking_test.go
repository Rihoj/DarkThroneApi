package DarkThroneApi

import (
	"testing"
)

func TestBankDepositRequest_Fields(t *testing.T) {
	r := BankDepositRequest{PlayerID: "pid", Amount: 100}
	if r.PlayerID != "pid" || r.Amount != 100 {
		t.Error("unexpected field values")
	}
}

func TestBankWithdrawRequest_Fields(t *testing.T) {
	r := BankWithdrawRequest{PlayerID: "pid", Amount: 200}
	if r.PlayerID != "pid" || r.Amount != 200 {
		t.Error("unexpected field values")
	}
}

func TestBankResponse_Fields(t *testing.T) {
	r := BankResponse{Success: true, Message: "ok", Balance: 42}
	if !r.Success || r.Message != "ok" || r.Balance != 42 {
		t.Error("unexpected field values")
	}
}

// Add more tests for DepositGold and WithdrawGold as API integration or with mocks.
