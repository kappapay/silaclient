package sila

type IssueSila struct {
	Header         *Header `json:"header"`
	Message        string  `json:"message"`
	Amount         int64   `json:"amount"`
	AccountName    string  `json:"account_name"`
	Descriptor     string  `json:"descriptor,omitempty"`
	BusinessUuid   string  `json:"business_uuid,omitempty"`
	ProcessingType string  `json:"processing_type,omitempty"`
}

func (msg *IssueSila) SetRef(ref string) *IssueSila {
	msg.Header.setRef(ref)
	return msg
}

func (msg *IssueSila) SetAmountToAccount(amount int64, accountName string) *IssueSila {
	msg.Amount = amount
	msg.AccountName = accountName
	return msg
}

func (msg *IssueSila) SetDescriptor(descriptor string) *IssueSila {
	msg.Descriptor = descriptor
	return msg
}

func (msg *IssueSila) SetBusinessUuid(businessUuid string) *IssueSila {
	msg.BusinessUuid = businessUuid
	return msg
}

func (msg *IssueSila) SetProcessingType(processingType string) *IssueSila {
	msg.ProcessingType = processingType
	return msg
}

type IssueSilaResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	TransactionId     string                 `json:"transaction_id"`
	Descriptor        string                 `json:"descriptor,omitempty"`
}

// The wallet key passed in is what determines the which wallet receives the Sila coin
func (msg *IssueSila) Do(userWalletPrivateKey string) (IssueSilaResponse, error) {
	var responseBody IssueSilaResponse
	err := instance.performCallWithUserAuth("/issue_sila", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
