package sila

// Link a business member to the specified business.
// If the user handle is an admin of the business, a different user handle should be specified to link to the business (use "AsAdmin" methods).
// If the user handle is not an admin of the business, that handle will be linked.
func (client ClientImpl) LinkBusinessMember(userHandle string, businessHandle string) LinkBusinessMember {
	return &LinkBusinessMemberMsg{
		Header: client.generateHeader().setUserHandle(userHandle).setBusinessHandle(businessHandle),
	}
}

type LinkBusinessMemberMsg struct {
	Header         *Header `json:"header"`
	MemberHandle   string  `json:"member_handle,omitempty"`
	Role           string  `json:"role"`
	OwnershipStake float64 `json:"ownership_stake,omitempty"`
	Description    string  `json:"description,omitempty"`
}

// As an admin member, set a different member's user handle as an admin of this business.
// Admins can link others to the business on their behalf.
func (msg *LinkBusinessMemberMsg) SetAdminMemberAsAdmin(newMemberHandle string) LinkBusinessMember {
	msg.MemberHandle = newMemberHandle
	msg.Role = "administrator"
	return msg
}

// Set the current member whose user handle was already provided as an admin.
// Admins can link others to the business on their behalf.
func (msg *LinkBusinessMemberMsg) SetAdminMember() LinkBusinessMember {
	msg.Role = "administrator"
	return msg
}

// As an admin member, set a different member's user handle as a controlling officer of this business.
// Controlling officers have leadership roles in the business and have the ability to sign contracts for the business.
func (msg *LinkBusinessMemberMsg) SetControllingOfficerMemberAsAdmin(newMemberHandle string) LinkBusinessMember {
	msg.MemberHandle = newMemberHandle
	msg.Role = "controlling_officer"
	return msg
}

// Set the current member whose user handle was already provided as a controlling officer.
// Controlling officers have leadership roles in the business and have the ability to sign contracts for the business.
func (msg *LinkBusinessMemberMsg) SetControllingOfficerMember() LinkBusinessMember {
	msg.Role = "controlling_officer"
	return msg
}

// As an admin member, set a different member's user handle as a beneficial owner of this business.
// Beneficial owners have some ownership stake in the business, which should be included.
func (msg *LinkBusinessMemberMsg) SetBeneficialOwnerMemberAsAdmin(newMemberHandle string, ownershipStake float64) LinkBusinessMember {
	msg.MemberHandle = newMemberHandle
	msg.Role = "beneficial_owner"
	msg.OwnershipStake = ownershipStake
	return msg
}

// Set the current member whose user handle was already provided as a beneficial owner of this business.
// Beneficial owners have some ownership stake in the business, which should be included.
func (msg *LinkBusinessMemberMsg) SetBeneficialOwnerMember(ownershipStake float64) LinkBusinessMember {
	msg.Role = "beneficial_owner"
	msg.OwnershipStake = ownershipStake
	return msg
}

// Optionally set the description of the member being linked, which can be used to distinguish people and their roles at
// a later date.
func (msg *LinkBusinessMemberMsg) SetMemberDescription(description string) LinkBusinessMember {
	msg.Description = description
	return msg
}

type LinkBusinessMemberResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	Role              string                 `json:"role"`
	Details           string                 `json:"details"`
	VerificationUuid  string                 `json:"verification_uuid"`
}

func (msg *LinkBusinessMemberMsg) Do(userWalletPrivateKey string, businessWalletPrivateKey string) (LinkBusinessMemberResponse, error) {
	var responseBody LinkBusinessMemberResponse
	err := instance.performCallWithUserAndBusinessAuth("/link_business_member", msg, &responseBody, userWalletPrivateKey, businessWalletPrivateKey)
	return responseBody, err
}
