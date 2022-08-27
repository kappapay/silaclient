package sila

import (
	"github.com/kappapay/silaclient/domain"
)

func (client ClientImpl) PlaidSameDayAuth(userHandle string, accountName string) PlaidSameDayAuth {
	return &PlaidSameDayAuthMsg{
		Header:      client.generateHeader().setUserHandle(userHandle),
		AccountName: accountName,
	}
}

type PlaidSameDayAuthMsg struct {
	Header      *Header `json:"header"`
	AccountName string  `json:"account_name"`
}

func (msg *PlaidSameDayAuthMsg) Do() (domain.PlaidSameDayAuthResponse, error) {
	var responseBody domain.PlaidSameDayAuthResponse
	err := instance.performCall("/plaid_sameday_auth", msg, &responseBody)
	return responseBody, err
}
