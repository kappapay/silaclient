package sila

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

type UnlinkBusinessMemberResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	Role              string                 `json:"role"`
}

func (msg *UnlinkBusinessMemberMsg) Do(userWalletPrivateKey string, businessWalletPrivateKey string) (UnlinkBusinessMemberResponse, error) {
	var responseBody UnlinkBusinessMemberResponse
	err := instance.performCallWithUserAndBusinessAuth("/unlink_business_member", msg, &responseBody, userWalletPrivateKey, businessWalletPrivateKey)
	return responseBody, err
}
