package sila

type GetWallets struct {
	Header        *Header             `json:"header"`
	SearchFilters WalletSearchFilters `json:"search_filters,omitempty"`
}

type WalletSearchFilters struct {
	Page              int32  `json:"page,omitempty"`
	PerPage           int32  `json:"per_page,omitempty"`
	SortAscending     bool   `json:"sort_ascending,omitempty"`
	BlockchainNetwork string `json:"blockchain_network,omitempty"`
	BlockchainAddress string `json:"blockchain_address,omitempty"`
	Nickname          string `json:"nickname,omitempty"`
}

func (msg *GetWallets) SetRef(ref string) *GetWallets {
	msg.Header.setRef(ref)
	return msg
}

func (msg *GetWallets) SetSearchFilters(filters WalletSearchFilters) *GetWallets {
	msg.SearchFilters = filters
	return msg
}

type GetWalletsResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	Wallets           []Wallet               `json:"wallets"`
	Page              int32                  `json:"page"`
	ReturnedCount     int32                  `json:"returned_count"`
	TotalCount        int32                  `json:"total_count"`
	TotalPageCount    int32                  `json:"total_page_count"`
}

func (msg *GetWallets) Do(userWalletPrivateKey string) (GetWalletsResponse, error) {
	var responseBody GetWalletsResponse
	err := instance.performCallWithUserAuth("/get_wallets", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
