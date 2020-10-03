package sila

import (
	"github.com/bpancost/sila/domain"
)

func (client ClientImpl) RegisterWallet(userHandle string) RegisterWallet {
	return &RegisterWalletMsg{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

type RegisterWalletMsg struct {
	Header                      *Header       `json:"header"`
	WalletVerificationSignature string        `json:"wallet_verification_signature"`
	Wallet                      domain.Wallet `json:"wallet"`
}

func (msg *RegisterWalletMsg) SetRef(ref string) RegisterWallet {
	msg.Header.setRef(ref)
	return msg
}

// Sets wallet information to register.
// The new wallet signature should be generated from signing the new wallet address with the GenerateWalletSignature function.
func (msg *RegisterWalletMsg) SetWallet(nickname string, address string, newWalletSignature string) RegisterWallet {
	msg.Wallet.BlockchainAddress = address
	msg.Wallet.BlockchainNetwork = "ETH"
	msg.Wallet.Nickname = nickname

	msg.WalletVerificationSignature = newWalletSignature
	return msg
}

func (msg *RegisterWalletMsg) Do(userWalletPrivateKey string) (domain.RegisterWalletResponse, error) {
	var responseBody domain.RegisterWalletResponse
	err := instance.performCallWithUserAuth("/register_wallet", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
