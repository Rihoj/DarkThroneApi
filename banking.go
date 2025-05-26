package DarkThroneApi

// BankDepositRequest represents the payload to deposit gold.
type BankDepositRequest struct {
	PlayerID string `json:"playerId"`
	Amount   int    `json:"amount"`
}

// BankWithdrawRequest represents the payload to withdraw gold.
type BankWithdrawRequest struct {
	PlayerID string `json:"playerId"`
	Amount   int    `json:"amount"`
}

// BankResponse represents the response from a bank operation.
type BankResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Balance int    `json:"balance"`
}

// DepositGold deposits gold into the bank.
func (d *DarkThroneApi) DepositGold(req BankDepositRequest) (BankResponse, error) {
	response, err := ApiRequest[BankDepositRequest, BankResponse]{
		Method:   "POST",
		Endpoint: "bank/deposit",
		Headers:  d.getAuthHeaders(),
		Body:     req,
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		return BankResponse{}, err
	}
	return response, nil
}

// WithdrawGold withdraws gold from the bank.
func (d *DarkThroneApi) WithdrawGold(req BankWithdrawRequest) (BankResponse, error) {
	response, err := ApiRequest[BankWithdrawRequest, BankResponse]{
		Method:   "POST",
		Endpoint: "bank/withdraw",
		Headers:  d.getAuthHeaders(),
		Body:     req,
		Config:   d.apiConfig,
	}.DoRequest()
	if err != nil {
		return BankResponse{}, err
	}
	return response, nil
}
