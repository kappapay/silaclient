package sila

type DeleteWallet struct {
	Header *Header `json:"header"`
}

func (msg *DeleteWallet) SetRef(ref string) *DeleteWallet {
	msg.Header.setRef(ref)
	return msg
}

// The wallet key passed in is what determines the wallet deleted
func (msg *DeleteWallet) Do(userWalletPrivateKey string) (SuccessResponse, error) {
	var responseBody SuccessResponse
	err := instance.performCallWithUserAuth("/delete_wallet", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
