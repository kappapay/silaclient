package sila

type UpdateWallet struct {
	Header   *Header `json:"header"`
	Nickname string  `json:"nickname,omitempty"`
	Default  bool    `json:"default,omitempty"`
}

func (msg *UpdateWallet) SetRef(ref string) *UpdateWallet {
	msg.Header.setRef(ref)
	return msg
}

func (msg *UpdateWallet) SetNickname(nickname string) *UpdateWallet {
	msg.Nickname = nickname
	return msg
}

func (msg *UpdateWallet) SetDefault(isDefault bool) *UpdateWallet {
	msg.Default = isDefault
	return msg
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

// The wallet key passed in is what determines the wallet updated
func (msg *UpdateWallet) Do(userWalletPrivateKey string) (UpdateWalletResponse, error) {
	var responseBody UpdateWalletResponse
	err := instance.performCallWithUserAuth("/update_wallet", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
