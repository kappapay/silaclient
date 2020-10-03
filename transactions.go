package sila

import "sila/domain"

type IssueSila interface {
	SetRef(ref string) IssueSila
	SetAmountFromAccount(amount int64, accountName string) IssueSila
	SetDescriptor(descriptor string) IssueSila
	SetBusinessUuid(businessUuid string) IssueSila
	SetProcessingType(processingType string) IssueSila
	Do(userWalletPrivateKey string) (domain.IssueSilaResponse, error)
}
type TransferSila interface {
	SetRef(ref string) TransferSila
	SetAmountAndUser(amount int64, destinationHandle string) TransferSila
	SetDestinationWallet(destinationWalletName string) TransferSila
	SetDestinationAddress(destinationWalletAddress string) TransferSila
	SetDescriptor(descriptor string) TransferSila
	SetBusinessUuid(businessUuid string) TransferSila
	Do(userWalletPrivateKey string) (domain.TransferSilaResponse, error)
}
type RedeemSila interface {
	SetRef(ref string) RedeemSila
	SetAmountToAccount(amount int64, accountName string) RedeemSila
	SetDescriptor(descriptor string) RedeemSila
	SetBusinessUuid(businessUuid string) RedeemSila
	SetProcessingType(processingType string) RedeemSila
	Do(userWalletPrivateKey string) (domain.RedeemSilaResponse, error)
}
type GetTransactions interface {
	SetSearchFilters(searchFilters domain.TransactionSearchFilters) GetTransactions
	Do(userWalletPrivateKey string) (domain.GetTransactionsResponse, error)
}
type CancelTransactions interface {
	SetRef(ref string) CancelTransactions
	Do(userWalletPrivateKey string) (domain.SuccessResponse, error)
}
