package sila

import "github.com/kappapay/silaclient/domain"

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

func (msg *LinkBusinessMemberMsg) SetAdminMemberAsAdmin(newMemberHandle string) LinkBusinessMember {
	msg.MemberHandle = newMemberHandle
	msg.Role = "administrator"
	return msg
}

func (msg *LinkBusinessMemberMsg) SetAdminMember() LinkBusinessMember {
	msg.Role = "administrator"
	return msg
}

func (msg *LinkBusinessMemberMsg) SetControllingOfficerMemberAsAdmin(newMemberHandle string) LinkBusinessMember {
	msg.MemberHandle = newMemberHandle
	msg.Role = "controlling_officer"
	return msg
}

func (msg *LinkBusinessMemberMsg) SetControllingOfficerMember() LinkBusinessMember {
	msg.Role = "controlling_officer"
	return msg
}

func (msg *LinkBusinessMemberMsg) SetBeneficialOwnerMemberAsAdmin(newMemberHandle string, ownershipStake float64) LinkBusinessMember {
	msg.MemberHandle = newMemberHandle
	msg.Role = "beneficial_owner"
	msg.OwnershipStake = ownershipStake
	return msg
}

func (msg *LinkBusinessMemberMsg) SetBeneficialOwnerMember(ownershipStake float64) LinkBusinessMember {
	msg.Role = "beneficial_owner"
	msg.OwnershipStake = ownershipStake
	return msg
}

func (msg *LinkBusinessMemberMsg) SetMemberDescription(description string) LinkBusinessMember {
	msg.Description = description
	return msg
}

func (msg *LinkBusinessMemberMsg) Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.LinkBusinessMemberResponse, error) {
	var responseBody domain.LinkBusinessMemberResponse
	err := instance.performCallWithUserAndBusinessAuth("/link_business_member", msg, &responseBody, userWalletPrivateKey, businessWalletPrivateKey)
	return responseBody, err
}
