package sila

type RequestKyc struct {
	Header   *Header `json:"header"`
	Message  string  `json:"message"`
	KycLevel string  `json:"kyc_level,omitempty"`
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
	err := instance.performCallWithUserAuth("/request_kyc", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
