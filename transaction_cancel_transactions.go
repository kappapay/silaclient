package sila

type CancelTransactions struct {
	Header        *Header `json:"header"`
	TransactionId string  `json:"transaction_id"`
}

func (msg *CancelTransactions) SetRef(ref string) *CancelTransactions {
	msg.Header.setRef(ref)
	return msg
}

// The wallet key passed should be registered to the user which initiated the transaction to cancel
func (msg *CancelTransactions) Do(userWalletPrivateKey string) (SuccessResponse, error) {
	var responseBody SuccessResponse
	err := instance.performCallWithUserAuth("/cancel_transaction", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
