package sila

import "github.com/kappapay/silaclient/domain"

func (client ClientImpl) CertifyBusiness(adminUserHandle string, businessHandle string) CertifyBusiness {
	return &CertifyBusinessMsg{
		Header: client.generateHeader().setUserHandle(adminUserHandle).setBusinessHandle(businessHandle),
	}
}

type CertifyBusinessMsg struct {
	Header *Header `json:"header"`
}

func (msg *CertifyBusinessMsg) Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.SuccessResponse, error) {
	var responseBody domain.SuccessResponse
	err := instance.performCallWithUserAndBusinessAuth("/certify_business", msg, &responseBody, userWalletPrivateKey, businessWalletPrivateKey)
	return responseBody, err
}
