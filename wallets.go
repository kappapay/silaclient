package sila

import "sila/domain"

type RegisterWallet interface {
	SetRef(ref string) RegisterWallet
	SetWallet(nickname string, address string, newWalletSignature string) RegisterWallet
	Do(userWalletPrivateKey string) (domain.RegisterWalletResponse, error)
}

type GetWallet interface {
	SetRef(ref string) GetWallet
	Do(userWalletPrivateKey string) (domain.GetWalletResponse, error)
}

type GetWallets interface {
	SetSearchFilters(filters domain.WalletSearchFilters) GetWallets
	Do(userWalletPrivateKey string) (domain.GetWalletsResponse, error)
}

type UpdateWallet interface {
	SetRef(ref string) UpdateWallet
	SetNickname(nickname string) UpdateWallet
	SetDefault(isDefault bool) UpdateWallet
	Do(userWalletPrivateKey string) (domain.UpdateWalletResponse, error)
}

type GetSilaBalance interface {
	Do() (domain.GetSilaBalanceResponse, error)
}

type DeleteWallet interface {
	SetRef(ref string) DeleteWallet
	Do(userWalletPrivateKey string) (domain.SuccessResponse, error)
}
