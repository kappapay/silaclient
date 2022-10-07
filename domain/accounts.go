package domain

type PlaidSameDayAuthResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	PublicToken       string                 `json:"public_token"`
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

type GetAccountsResponse struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	AccountNumber     string `json:"account_number,omitempty"`
	RoutingNumber     string `json:"routing_number,omitempty"`
	AccountName       string `json:"account_name"`
	AccountStatus     string `json:"account_status"`
	Active            bool   `json:"active"`
	AccountLinkStatus string `json:"account_link_status"`
}

type GetAccountBalanceResponse struct {
	Success             bool                   `json:"success"`
	Reference           string                 `json:"reference"`
	Message             string                 `json:"message"`
	Status              string                 `json:"status"`
	ValidationDetails   map[string]interface{} `json:"validation_details"`
	LinkStatus          string                 `json:"link_status"`
	AvailableBalance    float64                `json:"available_balance"`
	CurrentBalance      float64                `json:"current_balance"`
	MaskedAccountNumber string                 `json:"masked_account_number"`
	RoutingNumber       string                 `json:"routing_number"`
	AccountName         string                 `json:"account_name"`
}
