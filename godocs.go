// go:build docs
//go:build docs
// +build docs

/*
Package DarkThroneApi provides a Go client for the Dark Throne Reborn MMO API.

# Overview

This package allows you to authenticate users, manage players, perform banking operations, and interact with game structures via the Dark Throne API.

# Getting Started

	import "github.com/Rihoj/DarkThroneApi"

	api := DarkThroneApi.New(&DarkThroneApi.Config{Logger: nil})

# Main Types

- type Config: Configuration for the API client (logger, etc).
- type DarkThroneApi: Main client for API operations.
- type ApiRequest: Generic API request handler.
- type BankDepositRequest, BankWithdrawRequest, BankResponse: Banking payloads.
- type Player, Unit: Player and unit data.

# Example

	api := DarkThroneApi.New(&DarkThroneApi.Config{})
	resp, err := api.DepositGold(DarkThroneApi.BankDepositRequest{PlayerID: "pid", Amount: 100})
	if err != nil {
		// handle error
	}
	fmt.Println(resp.Balance)

# License
MIT

# Disclaimer

This project is provided as-is. The author is not responsible for how this library is used. Users are solely responsible for ensuring their usage complies with the Dark Throne Reborn site's Terms and Conditions.
*/
package DarkThroneApi
