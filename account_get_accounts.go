package sila

type GetAccounts struct {
	Header  *Header `json:"header"`
	Message string  `json:"message"`
}

type GetAccountsResponse struct {
	Accounts []Account
}

type Account struct {
	AccountNumber     string `json:"account_number,omitempty"`
	RoutingNumber     string `json:"routing_number,omitempty"`
	AccountName       string `json:"account_name"`
	AccountStatus     string `json:"account_status"`
	Active            bool   `json:"active"`
	AccountLinkStatus string `json:"account_link_status"`
}

func (msg *GetAccounts) Do(userWalletPrivateKey string) (GetAccountsResponse, error) {
	var responseBody GetAccountsResponse
	err := instance.performCallWithUserAuth("/get_accounts", msg, &responseBody.Accounts, userWalletPrivateKey)
	return responseBody, err
}
