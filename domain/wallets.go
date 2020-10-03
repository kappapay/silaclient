package domain

type GetSilaBalanceResponse struct {
	Success           bool                   `json:"success"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	Address           string                 `json:"address"`
	SilaBalance       float64                `json:"sila_balance"`
}

type Wallet struct {
	BlockchainAddress string `json:"blockchain_address"`
	BlockchainNetwork string `json:"blockchain_network"`
	Nickname          string `json:"nickname"`
	Default           bool   `json:"default,omitempty"`
	Frozen            bool   `json:"frozen"`
}

type GetWalletResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	Wallet            Wallet                 `json:"wallet"`
	IsWhitelisted     bool                   `json:"is_whitelisted"`
	SilaBalance       float64                `json:"sila_balance"`
}

type GetWalletsResponse struct {
	Success           bool                   `json:"success"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	Wallets           []Wallet               `json:"wallets"`
	Page              int32                  `json:"page"`
	ReturnedCount     int32                  `json:"returned_count"`
	TotalCount        int32                  `json:"total_count"`
	TotalPageCount    int32                  `json:"total_page_count"`
}

type WalletSearchFilters struct {
	Page              int32  `json:"page,omitempty"`
	PerPage           int32  `json:"per_page,omitempty"`
	SortAscending     bool   `json:"sort_ascending,omitempty"`
	BlockchainNetwork string `json:"blockchain_network,omitempty"`
	BlockchainAddress string `json:"blockchain_address,omitempty"`
	Nickname          string `json:"nickname,omitempty"`
}

type RegisterWalletResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	WalletNickname    string                 `json:"wallet_nickname"`
}

type UpdateWalletResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	Wallet            Wallet                 `json:"wallet"`
	Changes           []WalletChange         `json:"changes"`
}

type WalletChange struct {
	Attribute string `json:"attribute"`
	// Either a string or a bool value depending on the field changed, matches type with NewValue
	OldValue interface{} `json:"old_value"`
	// Either a string or bool value depending on the field changed, matches type with OldValue
	NewValue interface{} `json:"new_value"`
}
