package sila

import (
	"github.com/bpancost/sila/domain"
)

func (client ClientImpl) UpdateWallet(userHandle string) UpdateWallet {
	return &UpdateWalletMsg{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

type UpdateWalletMsg struct {
	Header   *Header `json:"header"`
	Nickname string  `json:"nickname,omitempty"`
	Default  bool    `json:"default,omitempty"`
}

func (msg *UpdateWalletMsg) SetRef(ref string) UpdateWallet {
	msg.Header.setRef(ref)
	return msg
}

func (msg *UpdateWalletMsg) SetNickname(nickname string) UpdateWallet {
	msg.Nickname = nickname
	return msg
}

func (msg *UpdateWalletMsg) SetDefault(isDefault bool) UpdateWallet {
	msg.Default = isDefault
	return msg
}

// The wallet key passed in is what determines the wallet updated
func (msg *UpdateWalletMsg) Do(userWalletPrivateKey string) (domain.UpdateWalletResponse, error) {
	var responseBody domain.UpdateWalletResponse
	err := instance.performCallWithUserAuth("/update_wallet", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
