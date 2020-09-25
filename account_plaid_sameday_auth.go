package sila

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type PlaidSameDayAuth struct {
	Header      *Header `json:"header"`
	AccountName string  `json:"account_name"`
}

func (msg *PlaidSameDayAuth) Do() (SuccessResponse, error) {
	var responseBody SuccessResponse
	requestJson, err := json.Marshal(msg)
	if err != nil {
		return responseBody, nil
	}
	url := instance.environment.generateURL(instance.version, "/plaid_sameday_auth")
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
