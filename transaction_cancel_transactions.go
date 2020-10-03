package sila

import (
	"github.com/bpancost/sila/domain"
)

func (client ClientImpl) CancelTransaction(userHandle string, transactionId string) CancelTransactions {
	return &CancelTransactionsMsg{
		Header:        client.generateHeader().setUserHandle(userHandle),
		TransactionId: transactionId,
	}
}

type CancelTransactionsMsg struct {
	Header        *Header `json:"header"`
	TransactionId string  `json:"transaction_id"`
}

func (msg *CancelTransactionsMsg) SetRef(ref string) CancelTransactions {
	msg.Header.setRef(ref)
	return msg
}

// The wallet key passed should be registered to the user which initiated the transaction to cancel
func (msg *CancelTransactionsMsg) Do(userWalletPrivateKey string) (domain.SuccessResponse, error) {
	var responseBody domain.SuccessResponse
	err := instance.performCallWithUserAuth("/cancel_transaction", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
