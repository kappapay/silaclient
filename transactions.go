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
