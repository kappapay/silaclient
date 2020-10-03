package sila

import (
	"github.com/bpancost/sila/domain"
)

type LinkAccount interface {
	SetRef(ref string) LinkAccount
	SetPlaidLinkAccount(publicToken string, selectedAccountId string) LinkAccount
	SetDirectLinkAccount(accountNumber string, routingNumber string) LinkAccount
	SetAccountType(accountType string) LinkAccount
	SetAccountName(accountName string) LinkAccount
	Do(userWalletPrivateKey string) (domain.LinkAccountResponse, error)
}

type PlaidSameDayAuth interface {
	Do() (domain.PlaidSameDayAuthResponse, error)
}

type GetAccounts interface {
	Do(userWalletPrivateKey string) (domain.GetAccountsResponse, error)
}

type GetAccountBalance interface {
	SetRef(ref string) GetAccountBalance
	Do(userWalletPrivateKey string) (domain.GetAccountBalanceResponse, error)
}
