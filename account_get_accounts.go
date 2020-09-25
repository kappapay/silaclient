package sila

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type GetAccounts struct {
	Header  *Header `json:"header"`
	Message string  `json:"message"`
}

type GetAccountsResponse struct {
	Accounts []Account
}

type Account struct {
	AccountNumber     string `json:"account_number,omitempty"`
	RoutingNumber     string `json:"routing_number,omitempty"`
	AccountName       string `json:"account_name"`
	AccountStatus     string `json:"account_status"`
	Active            bool   `json:"active"`
	AccountLinkStatus string `json:"account_link_status"`
}

func (msg *GetAccounts) Do(userWalletPrivateKey string) (GetAccountsResponse, error) {
	var responseBody GetAccountsResponse
	responseBody.Accounts = make([]Account, 0)
	requestJson, err := json.Marshal(msg)
	if err != nil {
		return responseBody, nil
	}
	url := instance.environment.generateURL(instance.version, "/get_accounts")
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		return responseBody, err
	}
	request.Header.Set("Content-type", "application/json")
	authSignature, err := instance.GenerateAuthSignature(requestJson)
	if err != nil {
		return responseBody, errors.Errorf("failed to generate auth signature: %v", err)
	}
	request.Header.Set("authsignature", authSignature)
	userSignature, err := instance.GenerateUserSignature(requestJson, userWalletPrivateKey)
	if err != nil {
		return responseBody, errors.Errorf("failed to generate user signature: %v", err)
	}
	request.Header.Set("usersignature", userSignature)
	httpClient := http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		return responseBody, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseBody.Accounts)
	if err != nil {
		return responseBody, err
	}
	return responseBody, nil
}
