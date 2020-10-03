package sila

import (
	"sila/domain"
)

func (client ClientImpl) GetWallet(userHandle string) GetWallet {
	return &GetWalletMsg{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

type GetWalletMsg struct {
	Header *Header `json:"header"`
}

func (msg *GetWalletMsg) SetRef(ref string) GetWallet {
	msg.Header.setRef(ref)
	return msg
}

// The wallet key passed in is what determines the wallet returned
func (msg *GetWalletMsg) Do(userWalletPrivateKey string) (domain.GetWalletResponse, error) {
	var responseBody domain.GetWalletResponse
	err := instance.performCallWithUserAuth("/get_wallet", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
