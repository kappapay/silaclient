package sila

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type CheckHandle struct {
	Header  Header `json:"header"`
	Message string `json:"message"`
}

type CheckHandleResponse struct {
	client    Client
	Success   bool   `json:"success"`
	Reference string `json:"reference"`
	Message   string `json:"message"`
	Status    string `json:"status"`
}

func (client Client) CheckHandle(userHandle string) *CheckHandle {
	return &CheckHandle{
		Header:  client.generateHeader(userHandle, ""),
		Message: "header_msg",
	}
}

func (msg *CheckHandle) Ref(ref string) *CheckHandle {
	msg.Header.Ref(ref)
	return msg
}

func (msg *CheckHandle) Do() (CheckHandleResponse, error) {
	var responseBody CheckHandleResponse
	requestJson, err := json.Marshal(msg)
	if err != nil {
		return responseBody, nil
	}
	url := instance.environment.generateURL(instance.version, "/check_handle")
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		return responseBody, err
	}
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("authsignature", instance.GenerateAuthSignature(requestJson))
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
