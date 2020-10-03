package sila

import (
	"sila/domain"
)

func (client ClientImpl) GetWallets(userHandle string) GetWallets {
	return &GetWalletsMsg{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

type GetWalletsMsg struct {
	Header        *Header                    `json:"header"`
	SearchFilters domain.WalletSearchFilters `json:"search_filters,omitempty"`
}

func (msg *GetWalletsMsg) SetSearchFilters(filters domain.WalletSearchFilters) GetWallets {
	msg.SearchFilters = filters
	return msg
}

func (msg *GetWalletsMsg) Do(userWalletPrivateKey string) (domain.GetWalletsResponse, error) {
	var responseBody domain.GetWalletsResponse
	err := instance.performCallWithUserAuth("/get_wallets", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
