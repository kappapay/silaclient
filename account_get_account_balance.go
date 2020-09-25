package sila

type GetAccountBalance struct {
	Header      *Header `json:"header"`
	AccountName string  `json:"account_name"`
}

func (msg *GetAccountBalance) SetRef(ref string) *GetAccountBalance {
	msg.Header.setRef(ref)
	return msg
}

type GetAccountBalanceResponse struct {
	Success             bool                   `json:"success"`
	Reference           string                 `json:"reference"`
	Message             string                 `json:"message"`
	Status              string                 `json:"status"`
	ValidationDetails   map[string]interface{} `json:"validation_details"`
	LinkStatus          string                 `json:"link_status"`
	AvailableBalance    int64                  `json:"available_balance"`
	CurrentBalance      int64                  `json:"current_balance"`
	MaskedAccountNumber string                 `json:"masked_account_number"`
	RoutingNumber       int64                  `json:"routing_number"`
	AccountName         string                 `json:"account_name"`
}

func (msg *GetAccountBalance) Do(userWalletPrivateKey string) (GetAccountBalanceResponse, error) {
	var responseBody GetAccountBalanceResponse
	err := instance.performCallWithUserAuth("/get_account_balance", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
