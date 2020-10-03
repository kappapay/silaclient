package sila

import (
	"sila/domain"
)

func (client ClientImpl) GetAccountBalance(userHandle string, accountName string) GetAccountBalance {
	return &GetAccountBalanceMsg{
		Header:      client.generateHeader().setUserHandle(userHandle),
		AccountName: accountName,
	}
}

type GetAccountBalanceMsg struct {
	Header      *Header `json:"header"`
	AccountName string  `json:"account_name"`
}

func (msg *GetAccountBalanceMsg) SetRef(ref string) GetAccountBalance {
	msg.Header.setRef(ref)
	return msg
}

func (msg *GetAccountBalanceMsg) Do(userWalletPrivateKey string) (domain.GetAccountBalanceResponse, error) {
	var responseBody domain.GetAccountBalanceResponse
	err := instance.performCallWithUserAuth("/get_account_balance", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
