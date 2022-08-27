package sila

import "github.com/kappapay/silaclient/domain"

func (client ClientImpl) CertifyBeneficialOwner(adminUserHandle string, businessHandle string) CertifyBeneficialOwner {
	return &CertifyBeneficialOwnerMsg{
		Header: client.generateHeader().setUserHandle(adminUserHandle).setBusinessHandle(businessHandle),
	}
}

type CertifyBeneficialOwnerMsg struct {
	Header             *Header `json:"header"`
	MemberHandle       string  `json:"member_handle"`
	CertificationToken string  `json:"certification_token"`
}

func (msg *CertifyBeneficialOwnerMsg) SetCertificationToken(userHandleToCertify string, certificationToken string) CertifyBeneficialOwner {
	msg.MemberHandle = userHandleToCertify
	msg.CertificationToken = certificationToken
	return msg
}

func (msg *CertifyBeneficialOwnerMsg) Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.SuccessResponse, error) {
	var responseBody domain.SuccessResponse
	err := instance.performCallWithUserAndBusinessAuth("/certify_beneficial_owner", msg, &responseBody, userWalletPrivateKey, businessWalletPrivateKey)
	return responseBody, err
}
