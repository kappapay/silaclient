package sila

import "github.com/kappapay/silaclient/domain"

func (client ClientImpl) UnlinkBusinessMember(userHandle string, businessHandle string) UnlinkBusinessMember {
	return &UnlinkBusinessMemberMsg{
		Header: client.generateHeader().setUserHandle(userHandle).setBusinessHandle(businessHandle),
	}
}

type UnlinkBusinessMemberMsg struct {
	Header *Header `json:"header"`
	Role   string  `json:"role"`
}

func (msg *UnlinkBusinessMemberMsg) SetAdminRole() UnlinkBusinessMember {
	msg.Role = "administrator"
	return msg
}

func (msg *UnlinkBusinessMemberMsg) SetBeneficialOwnerRole() UnlinkBusinessMember {
	msg.Role = "beneficial_owner"
	return msg
}

func (msg *UnlinkBusinessMemberMsg) SetControllingOfficerRole() UnlinkBusinessMember {
	msg.Role = "controlling_officer"
	return msg
}

func (msg *UnlinkBusinessMemberMsg) Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.UnlinkBusinessMemberResponse, error) {
	var responseBody domain.UnlinkBusinessMemberResponse
	err := instance.performCallWithUserAndBusinessAuth("/unlink_business_member", msg, &responseBody, userWalletPrivateKey, businessWalletPrivateKey)
	return responseBody, err
}
