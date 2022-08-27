package sila

import (
	"github.com/kappapay/silaclient/domain"
)

func (client ClientImpl) TransferSila(userHandle string) TransferSila {
	return &TransferSilaMsg{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "transfer_msg",
	}
}

type TransferSilaMsg struct {
	Header             *Header `json:"header"`
	Message            string  `json:"message"`
	Amount             int64   `json:"amount"`
	DestinationHandle  string  `json:"destination_handle"`
	DestinationWallet  string  `json:"destination_wallet,omitempty"`
	DestinationAddress string  `json:"destination_address,omitempty"`
	Descriptor         string  `json:"descriptor,omitempty"`
	BusinessUuid       string  `json:"business_uuid,omitempty"`
}

func (msg *TransferSilaMsg) SetRef(ref string) TransferSila {
	msg.Header.setRef(ref)
	return msg
}

func (msg *TransferSilaMsg) SetAmountAndUser(amount int64, destinationHandle string) TransferSila {
	msg.Amount = amount
	msg.DestinationHandle = destinationHandle
	return msg
}

func (msg *TransferSilaMsg) SetDestinationWallet(destinationWalletName string) TransferSila {
	msg.DestinationWallet = destinationWalletName
	return msg
}

func (msg *TransferSilaMsg) SetDestinationAddress(destinationWalletAddress string) TransferSila {
	msg.DestinationAddress = destinationWalletAddress
	return msg
}

func (msg *TransferSilaMsg) SetDescriptor(descriptor string) TransferSila {
	msg.Descriptor = descriptor
	return msg
}

func (msg *TransferSilaMsg) SetBusinessUuid(businessUuid string) TransferSila {
	msg.BusinessUuid = businessUuid
	return msg
}

// The wallet key passed in is what determines the which wallet is the source wallet for the transfer
func (msg *TransferSilaMsg) Do(userWalletPrivateKey string) (domain.TransferSilaResponse, error) {
	var responseBody domain.TransferSilaResponse
	err := instance.performCallWithUserAuth("/transfer_sila", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
