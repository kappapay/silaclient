package sila

import (
	"github.com/bpancost/sila/domain"
)

func (client ClientImpl) LinkAccount(userHandle string) LinkAccount {
	return &LinkAccountMsg{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "link_account_msg",
	}
}

type LinkAccountMsg struct {
	Header            *Header `json:"header"`
	Message           string  `json:"message"`
	PublicToken       string  `json:"public_token,omitempty"`
	SelectedAccountId string  `json:"selected_account_id,omitempty"`
	AccountNumber     string  `json:"account_number,omitempty"`
	RoutingNumber     string  `json:"routing_number,omitempty"`
	AccountType       string  `json:"account_type,omitempty"`
	AccountName       string  `json:"account_name,omitempty"`
}

func (msg *LinkAccountMsg) SetRef(ref string) LinkAccount {
	msg.Header.setRef(ref)
	return msg
}

func (msg *LinkAccountMsg) SetPlaidLinkAccount(publicToken string, selectedAccountId string) LinkAccount {
	msg.PublicToken = publicToken
	msg.SelectedAccountId = selectedAccountId
	return msg
}

func (msg *LinkAccountMsg) SetDirectLinkAccount(accountNumber string, routingNumber string) LinkAccount {
	msg.AccountNumber = accountNumber
	msg.RoutingNumber = routingNumber
	return msg
}

func (msg *LinkAccountMsg) SetAccountType(accountType string) LinkAccount {
	msg.AccountType = accountType
	return msg
}

func (msg *LinkAccountMsg) SetAccountName(accountName string) LinkAccount {
	msg.AccountName = accountName
	return msg
}

func (msg *LinkAccountMsg) Do(userWalletPrivateKey string) (domain.LinkAccountResponse, error) {
	var responseBody domain.LinkAccountResponse
	err := instance.performCallWithUserAuth("/link_account", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
