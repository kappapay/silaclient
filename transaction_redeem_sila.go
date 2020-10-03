package sila

import (
	"sila/domain"
)

func (client ClientImpl) RedeemSila(userHandle string) RedeemSila {
	return &RedeemSilaMsg{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "redeem_msg",
	}
}

type RedeemSilaMsg struct {
	Header         *Header `json:"header"`
	Message        string  `json:"message"`
	Amount         int64   `json:"amount"`
	AccountName    string  `json:"account_name"`
	Descriptor     string  `json:"descriptor,omitempty"`
	BusinessUuid   string  `json:"business_uuid,omitempty"`
	ProcessingType string  `json:"processing_type,omitempty"`
}

func (msg *RedeemSilaMsg) SetRef(ref string) RedeemSila {
	msg.Header.setRef(ref)
	return msg
}

// Sets the amount to take from the Sila wallet and put into the named linked account
func (msg *RedeemSilaMsg) SetAmountToAccount(amount int64, accountName string) RedeemSila {
	msg.Amount = amount
	msg.AccountName = accountName
	return msg
}

func (msg *RedeemSilaMsg) SetDescriptor(descriptor string) RedeemSila {
	msg.Descriptor = descriptor
	return msg
}

func (msg *RedeemSilaMsg) SetBusinessUuid(businessUuid string) RedeemSila {
	msg.BusinessUuid = businessUuid
	return msg
}

func (msg *RedeemSilaMsg) SetProcessingType(processingType string) RedeemSila {
	msg.ProcessingType = processingType
	return msg
}

// The wallet key passed in is what determines the which wallet redeems the Sila coin
func (msg *RedeemSilaMsg) Do(userWalletPrivateKey string) (domain.RedeemSilaResponse, error) {
	var responseBody domain.RedeemSilaResponse
	err := instance.performCallWithUserAuth("/redeem_sila", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
