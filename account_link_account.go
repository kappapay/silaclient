package sila

type LinkAccount struct {
	Header            *Header `json:"header"`
	Message           string  `json:"message"`
	PublicToken       string  `json:"public_token,omitempty"`
	SelectedAccountId string  `json:"selected_account_id,omitempty"`
	AccountNumber     string  `json:"account_number,omitempty"`
	RoutingNumber     string  `json:"routing_number,omitempty"`
	AccountType       string  `json:"account_type,omitempty"`
	AccountName       string  `json:"account_name,omitempty"`
}

func (msg *LinkAccount) SetRef(ref string) *LinkAccount {
	msg.Header.setRef(ref)
	return msg
}

func (msg *LinkAccount) SetPlaidLinkAccount(publicToken string, selectedAccountId string) *LinkAccount {
	msg.PublicToken = publicToken
	msg.SelectedAccountId = selectedAccountId
	return msg
}

func (msg *LinkAccount) SetDirectLinkAccount(accountNumber string, routingNumber string) *LinkAccount {
	msg.AccountNumber = accountNumber
	msg.RoutingNumber = routingNumber
	return msg
}

func (msg *LinkAccount) SetAccountType(accountType string) *LinkAccount {
	msg.AccountType = accountType
	return msg
}

func (msg *LinkAccount) SetAccountName(accountName string) *LinkAccount {
	msg.AccountName = accountName
	return msg
}

type LinkAccountResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	AccountName       string                 `json:"account_name"`
	SelectedAccountId string                 `json:"selected_account_id"`
}

func (msg *LinkAccount) Do(userWalletPrivateKey string) (LinkAccountResponse, error) {
	var responseBody LinkAccountResponse
	err := instance.performCallWithUserAuth("/link_account", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
