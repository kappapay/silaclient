package sila

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type LinkAccount struct {
	Header            *Header `json:"header"`
	Message           string  `json:"message"`
	PublicToken       string  `json:"public_token,omitempty"`
	SelectedAccountId string  `json:"selected_account_id,omitempty"`
	AccountNumber     string  `json:"account_number,omitempty"`
	RoutingNumber     string  `json:"routing_number,omitempty"`
	AccountType       string  `json:"account_type,omitempty"`
	AccountName       string  `json:"account_name,omitempty"`
}

func (msg *LinkAccount) SetRef(ref string) *LinkAccount {
	msg.Header.setRef(ref)
	return msg
}

func (msg *LinkAccount) SetPlaidLinkAccount(publicToken string, selectedAccountId string) *LinkAccount {
	msg.PublicToken = publicToken
	msg.SelectedAccountId = selectedAccountId
	return msg
}

func (msg *LinkAccount) SetDirectLinkAccount(accountNumber string, routingNumber string) *LinkAccount {
	msg.AccountNumber = accountNumber
	msg.RoutingNumber = routingNumber
	return msg
}

func (msg *LinkAccount) SetAccountType(accountType string) *LinkAccount {
	msg.AccountType = accountType
	return msg
}

func (msg *LinkAccount) SetAccountName(accountName string) *LinkAccount {
	msg.AccountName = accountName
	return msg
}

type LinkAccountResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	AccountName       string                 `json:"account_name"`
	SelectedAccountId string                 `json:"selected_account_id"`
}

func (msg *LinkAccount) Do(userWalletPrivateKey string) (LinkAccountResponse, error) {
	var responseBody LinkAccountResponse
	requestJson, err := json.Marshal(msg)
	if err != nil {
		return responseBody, nil
	}
	url := instance.environment.generateURL(instance.version, "/link_account")
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

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return responseBody, err
	}
	return responseBody, nil
}
