package sila

import (
	"sila/domain"
)

func (client ClientImpl) GetWalletBalance(walletAddress string) GetSilaBalance {
	return &GetSilaBalanceMsg{
		Address: walletAddress,
	}
}

type GetSilaBalanceMsg struct {
	Address string `json:"address"`
}

func (msg *GetSilaBalanceMsg) Do() (domain.GetSilaBalanceResponse, error) {
	var responseBody domain.GetSilaBalanceResponse
	err := instance.performPublicCall("/get_sila_balance", msg, &responseBody)
	return responseBody, err
}
