package sila

import (
	"sila/domain"
)

func (client ClientImpl) CheckKyc(userHandle string) CheckKyc {
	return &CheckKycMsg{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "header_msg",
	}
}

type CheckKycMsg struct {
	Header   *Header `json:"header"`
	Message  string  `json:"message"`
	KycLevel string  `json:"kyc_level,omitempty"`
}

func (msg *CheckKycMsg) SetRef(ref string) CheckKyc {
	msg.Header.setRef(ref)
	return msg
}

func (msg *CheckKycMsg) SetKycLevel(kycLevel string) CheckKyc {
	msg.KycLevel = kycLevel
	return msg
}

func (msg *CheckKycMsg) Do(userWalletPrivateKey string) (domain.CheckKycResponse, error) {
	var responseBody domain.CheckKycResponse
	err := instance.performCallWithUserAuth("/check_kyc", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
