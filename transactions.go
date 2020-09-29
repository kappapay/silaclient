package sila

func (client Client) IssueSila(userHandle string) *IssueSila {
	return &IssueSila{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "issue_msg",
	}
}

func (client Client) TransferSila(userHandle string) *TransferSila {
	return &TransferSila{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "transfer_msg",
	}
}

func (client Client) RedeemSila(userHandle string) *RedeemSila {
	return &RedeemSila{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "redeem_msg",
	}
}

func (client Client) GetTransactions(userHandle string) *GetTransactions {
	return &GetTransactions{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "get_transactions_msg",
	}
}

func (client Client) CancelTransaction(userHandle string, transactionId string) *CancelTransactions {
	return &CancelTransactions{
		Header:        client.generateHeader().setUserHandle(userHandle),
		TransactionId: transactionId,
	}
}
