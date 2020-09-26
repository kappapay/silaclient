package sila

type GetWallet struct {
	Header *Header `json:"header"`
}

func (msg *GetWallet) SetRef(ref string) *GetWallet {
	msg.Header.setRef(ref)
	return msg
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

// The wallet key passed in is what determines the wallet returned
func (msg *GetWallet) Do(userWalletPrivateKey string) (GetWalletResponse, error) {
	var responseBody GetWalletResponse
	err := instance.performCallWithUserAuth("/get_wallet", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
