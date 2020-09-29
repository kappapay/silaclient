package sila

type TransferSila struct {
	Header             *Header `json:"header"`
	Message            string  `json:"message"`
	Amount             int64   `json:"amount"`
	DestinationHandle  string  `json:"destination_handle"`
	DestinationWallet  string  `json:"destination_wallet,omitempty"`
	DestinationAddress string  `json:"destination_address,omitempty"`
	Descriptor         string  `json:"descriptor,omitempty"`
	BusinessUuid       string  `json:"business_uuid,omitempty"`
}

func (msg *TransferSila) SetRef(ref string) *TransferSila {
	msg.Header.setRef(ref)
	return msg
}

func (msg *TransferSila) SetAmountAndUser(amount int64, destinationHandle string) *TransferSila {
	msg.Amount = amount
	msg.DestinationHandle = destinationHandle
	return msg
}

func (msg *TransferSila) SetDestinationWallet(destinationWalletName string) *TransferSila {
	msg.DestinationWallet = destinationWalletName
	return msg
}

func (msg *TransferSila) SetDestinationAddress(destinationWalletAddress string) *TransferSila {
	msg.DestinationAddress = destinationWalletAddress
	return msg
}

func (msg *TransferSila) SetDescriptor(descriptor string) *TransferSila {
	msg.Descriptor = descriptor
	return msg
}

func (msg *TransferSila) SetBusinessUuid(businessUuid string) *TransferSila {
	msg.BusinessUuid = businessUuid
	return msg
}

type TransferSilaResponse struct {
	Success            bool                   `json:"success"`
	Reference          string                 `json:"reference"`
	Message            string                 `json:"message"`
	Status             string                 `json:"status"`
	ValidationDetails  map[string]interface{} `json:"validation_details"`
	DestinationAddress string                 `json:"destination_address"`
	TransactionId      string                 `json:"transaction_id"`
	Descriptor         string                 `json:"descriptor,omitempty"`
}

// The wallet key passed in is what determines the which wallet is the source wallet for the transfer
func (msg *TransferSila) Do(userWalletPrivateKey string) (TransferSilaResponse, error) {
	var responseBody TransferSilaResponse
	err := instance.performCallWithUserAuth("/transfer_sila", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
