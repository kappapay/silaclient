package sila

type RegisterWallet struct {
	Header                      *Header `json:"header"`
	WalletVerificationSignature string  `json:"wallet_verification_signature"`
	Wallet                      Wallet  `json:"wallet"`
}

type Wallet struct {
	BlockchainAddress string `json:"blockchain_address"`
	BlockchainNetwork string `json:"blockchain_network"`
	Nickname          string `json:"nickname"`
	Default           bool   `json:"default,omitempty"`
	Frozen            bool   `json:"frozen"`
}

func (msg *RegisterWallet) SetRef(ref string) *RegisterWallet {
	msg.Header.setRef(ref)
	return msg
}

// Sets wallet information to register.
// The new wallet signature should be generated from signing the new wallet address with the GenerateWalletSignature function.
func (msg *RegisterWallet) SetWallet(nickname string, address string, newWalletSignature string) *RegisterWallet {
	msg.Wallet.BlockchainAddress = address
	msg.Wallet.BlockchainNetwork = "ETH"
	msg.Wallet.Nickname = nickname

	msg.WalletVerificationSignature = newWalletSignature
	return msg
}

type RegisterWalletResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	WalletNickname    string                 `json:"wallet_nickname"`
}

func (msg *RegisterWallet) Do(userWalletPrivateKey string) (RegisterWalletResponse, error) {
	var responseBody RegisterWalletResponse
	err := instance.performCallWithUserAuth("/register_wallet", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
