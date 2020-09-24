package sila

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type RequestKyc struct {
	Header   *Header `json:"header"`
	Message  string  `json:"message"`
	KycLevel string  `json:"kyc_level"`
}

func (msg *RequestKyc) SetRef(ref string) *RequestKyc {
	msg.Header.setRef(ref)
	return msg
}

func (msg *RequestKyc) SetKycLevel(kycLevel string) *RequestKyc {
	msg.KycLevel = kycLevel
	return msg
}

type RequestKycResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	VerificationUuid  string                 `json:"verification_uuid"`
}

func (msg *RequestKyc) Do(userWalletPrivateKey string) (RequestKycResponse, error) {
	var responseBody RequestKycResponse
	requestJson, err := json.Marshal(msg)
	if err != nil {
		return responseBody, nil
	}
	url := instance.environment.generateURL(instance.version, "/request_kyc")
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
