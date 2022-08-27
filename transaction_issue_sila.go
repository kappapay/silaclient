package sila

import (
	"github.com/kappapay/silaclient/domain"
)

func (client ClientImpl) IssueSila(userHandle string) IssueSila {
	return &IssueSilaMsg{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "issue_msg",
	}
}

type IssueSilaMsg struct {
	Header         *Header `json:"header"`
	Message        string  `json:"message"`
	Amount         int64   `json:"amount"`
	AccountName    string  `json:"account_name"`
	Descriptor     string  `json:"descriptor,omitempty"`
	BusinessUuid   string  `json:"business_uuid,omitempty"`
	ProcessingType string  `json:"processing_type,omitempty"`
}

func (msg *IssueSilaMsg) SetRef(ref string) IssueSila {
	msg.Header.setRef(ref)
	return msg
}

// Sets the amount to take from the named linked account and put into the Sila wallet
func (msg *IssueSilaMsg) SetAmountFromAccount(amount int64, accountName string) IssueSila {
	msg.Amount = amount
	msg.AccountName = accountName
	return msg
}

func (msg *IssueSilaMsg) SetDescriptor(descriptor string) IssueSila {
	msg.Descriptor = descriptor
	return msg
}

func (msg *IssueSilaMsg) SetBusinessUuid(businessUuid string) IssueSila {
	msg.BusinessUuid = businessUuid
	return msg
}

func (msg *IssueSilaMsg) SetProcessingType(processingType string) IssueSila {
	msg.ProcessingType = processingType
	return msg
}

// The wallet key passed in is what determines the which wallet receives the Sila coin
func (msg *IssueSilaMsg) Do(userWalletPrivateKey string) (domain.IssueSilaResponse, error) {
	var responseBody domain.IssueSilaResponse
	err := instance.performCallWithUserAuth("/issue_sila", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
