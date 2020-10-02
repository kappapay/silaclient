package sila

type RedeemSila struct {
	Header         *Header `json:"header"`
	Message        string  `json:"message"`
	Amount         int64   `json:"amount"`
	AccountName    string  `json:"account_name"`
	Descriptor     string  `json:"descriptor,omitempty"`
	BusinessUuid   string  `json:"business_uuid,omitempty"`
	ProcessingType string  `json:"processing_type,omitempty"`
}

func (msg *RedeemSila) SetRef(ref string) *RedeemSila {
	msg.Header.setRef(ref)
	return msg
}

// Sets the amount to take from the Sila wallet and put into the named linked account
func (msg *RedeemSila) SetAmountToAccount(amount int64, accountName string) *RedeemSila {
	msg.Amount = amount
	msg.AccountName = accountName
	return msg
}

func (msg *RedeemSila) SetDescriptor(descriptor string) *RedeemSila {
	msg.Descriptor = descriptor
	return msg
}

func (msg *RedeemSila) SetBusinessUuid(businessUuid string) *RedeemSila {
	msg.BusinessUuid = businessUuid
	return msg
}

func (msg *RedeemSila) SetProcessingType(processingType string) *RedeemSila {
	msg.ProcessingType = processingType
	return msg
}

type RedeemSilaResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	TransactionId     string                 `json:"transaction_id"`
	Descriptor        string                 `json:"descriptor,omitempty"`
}

// The wallet key passed in is what determines the which wallet redeems the Sila coin
func (msg *RedeemSila) Do(userWalletPrivateKey string) (RedeemSilaResponse, error) {
	var responseBody RedeemSilaResponse
	err := instance.performCallWithUserAuth("/redeem_sila", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
