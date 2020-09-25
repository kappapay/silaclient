package sila

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type GetAccountBalance struct {
	Header      *Header `json:"header"`
	AccountName string  `json:"account_name"`
}

func (msg *GetAccountBalance) SetRef(ref string) *GetAccountBalance {
	msg.Header.setRef(ref)
	return msg
}

type GetAccountBalanceResponse struct {
	Success             bool                   `json:"success"`
	Reference           string                 `json:"reference"`
	Message             string                 `json:"message"`
	Status              string                 `json:"status"`
	ValidationDetails   map[string]interface{} `json:"validation_details"`
	LinkStatus          string                 `json:"link_status"`
	AvailableBalance    int64                  `json:"available_balance"`
	CurrentBalance      int64                  `json:"current_balance"`
	MaskedAccountNumber string                 `json:"masked_account_number"`
	RoutingNumber       int64                  `json:"routing_number"`
	AccountName         string                 `json:"account_name"`
}

func (msg *GetAccountBalance) Do(userWalletPrivateKey string) (GetAccountBalanceResponse, error) {
	var responseBody GetAccountBalanceResponse
	requestJson, err := json.Marshal(msg)
	if err != nil {
		return responseBody, nil
	}
	url := instance.environment.generateURL(instance.version, "/get_account_balance")
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
