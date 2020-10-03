package sila

import (
	"sila/domain"
)

func (client ClientImpl) RequestKyc(userHandle string) RequestKyc {
	return &RequestKycMsg{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "header_msg",
	}
}

type RequestKycMsg struct {
	Header   *Header `json:"header"`
	Message  string  `json:"message"`
	KycLevel string  `json:"kyc_level,omitempty"`
}

func (msg *RequestKycMsg) SetRef(ref string) RequestKyc {
	msg.Header.setRef(ref)
	return msg
}

func (msg *RequestKycMsg) SetKycLevel(kycLevel string) RequestKyc {
	msg.KycLevel = kycLevel
	return msg
}

func (msg *RequestKycMsg) Do(userWalletPrivateKey string) (domain.RequestKycResponse, error) {
	var responseBody domain.RequestKycResponse
	err := instance.performCallWithUserAuth("/request_kyc", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
