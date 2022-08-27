package sila

import (
	"github.com/kappapay/silaclient/domain"
)

func (client ClientImpl) GetAccounts(userHandle string) GetAccounts {
	return &GetAccountsMsg{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "get_accounts_msg",
	}
}

type GetAccountsMsg struct {
	Header  *Header `json:"header"`
	Message string  `json:"message"`
}

func (msg *GetAccountsMsg) Do(userWalletPrivateKey string) (domain.GetAccountsResponse, error) {
	var responseBody domain.GetAccountsResponse
	err := instance.performCallWithUserAuth("/get_accounts", msg, &responseBody.Accounts, userWalletPrivateKey)
	return responseBody, err
}
