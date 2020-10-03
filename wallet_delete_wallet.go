package sila

import (
	"sila/domain"
)

func (client ClientImpl) DeleteWallet(userHandle string) DeleteWallet {
	return &DeleteWalletMsg{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

type DeleteWalletMsg struct {
	Header *Header `json:"header"`
}

func (msg *DeleteWalletMsg) SetRef(ref string) DeleteWallet {
	msg.Header.setRef(ref)
	return msg
}

// The wallet key passed in is what determines the wallet deleted
func (msg *DeleteWalletMsg) Do(userWalletPrivateKey string) (domain.SuccessResponse, error) {
	var responseBody domain.SuccessResponse
	err := instance.performCallWithUserAuth("/delete_wallet", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
