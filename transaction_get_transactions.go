package sila

import (
	"sila/domain"
)

func (client ClientImpl) GetTransactions(userHandle string) GetTransactions {
	return &GetTransactionsMsg{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "get_transactions_msg",
	}
}

type GetTransactionsMsg struct {
	Header        *Header                         `json:"header"`
	Message       string                          `json:"message"`
	SearchFilters domain.TransactionSearchFilters `json:"search_filters,omitempty"`
}

func (msg *GetTransactionsMsg) SetSearchFilters(searchFilters domain.TransactionSearchFilters) GetTransactions {
	msg.SearchFilters = searchFilters
	return msg
}

// The wallet key passed should be registered to the user which initiated the transaction to cancel
func (msg *GetTransactionsMsg) Do(userWalletPrivateKey string) (domain.GetTransactionsResponse, error) {
	var responseBody domain.GetTransactionsResponse
	err := instance.performCallWithUserAuth("/get_transactions", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
