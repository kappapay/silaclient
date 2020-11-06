package sila

import "github.com/bpancost/sila/domain"

type CheckHandle interface {
	// Sets a reference ID tagging the request which will also be passed back in the response
	SetRef(ref string) CheckHandle
	// Execute the Check Handle call
	Do() (domain.SuccessResponse, error)
}

type Register interface {
	// Sets a reference ID tagging the request which will also be passed back in the response
	SetRef(ref string) Register
	// Sets the user's address
	SetAddress(address domain.RegistrationAddress) Register
	// Sets the uniquely identifying information for the individual (SSN) or business (EIN). This is required to register.
	SetIdentity(identityType domain.IdentityType, identityValue string) Register
	// Sets the user's contact information.
	SetContact(contactAlias string, phone string, email string) Register
	// Sets the user's first wallet address and nickname. This is required to register. If the nickname is "default", it will be used by default for all transactions.
	SetCrypto(nickname string, address string) Register
	// Sets required entity information that is specific to an individual user.
	SetIndividualEntity(firstName string, lastName string, birthDate string) Register
	// Sets required entity information that is specific to a business.
	SetBusinessEntity(entityName string, businessType string, naicsCode int) Register
	// Sets a business website, if one exists
	SetBusinessWebsite(businessWebsite string) Register
	// Sets a business's Doing Business As name (the name they actually do business under, if different from the entity name).
	SetDoingBusinessAs(dba string) Register
	// Execute the Register call
	Do() (domain.SuccessResponse, error)
}

type RequestKyc interface {
	// Sets a reference ID tagging the request which will also be passed back in the response
	SetRef(ref string) RequestKyc
	// Sets the level of the KYC checks, if multiple levels are defined and a non-base level should be used. Not required for most use cases.
	SetKycLevel(kycLevel string) RequestKyc
	// Execute the Request KYC call, signing with one of the user's private wallet keys
	Do(userWalletPrivateKey string) (domain.RequestKycResponse, error)
}

type CheckKyc interface {
	// Sets a reference ID tagging the request which will also be passed back in the response
	SetRef(ref string) CheckKyc
	// Sets the level of the KYC checks, if multiple levels are defined and a non-base level should be checked. Not required for most use cases.
	SetKycLevel(kycLevel string) CheckKyc
	// Execute the Check KYC call, signing with one of the user's private wallet keys
	Do(userWalletPrivateKey string) (domain.CheckKycResponse, error)
}

type GetEntity interface {
	// Execute the call
	Do(userWalletPrivateKey string) (domain.GetEntityResponse, error)
}

type GetEntities interface {
	// The type of the entity (individual or business)
	SetEntityType(entityType string) GetEntities
	// Which page to fetch
	SetPage(page int32) GetEntities
	// How many entities should be returned per page
	SetPerPage(perPage int32) GetEntities
	// Execute the Get Entities call
	Do() (domain.GetEntitiesResponse, error)
}

type LinkBusinessMember interface {
	// As an admin member, set a different member's user handle as an admin of this business.
	// Admins can link others to the business on their behalf.
	SetAdminMemberAsAdmin(newMemberHandle string) LinkBusinessMember
	// Set the current member whose user handle was already provided as an admin.
	// Admins can link others to the business on their behalf.
	SetAdminMember() LinkBusinessMember
	// As an admin member, set a different member's user handle as a controlling officer of this business.
	// Controlling officers have leadership roles in the business and have the ability to sign contracts for the business.
	SetControllingOfficerMemberAsAdmin(newMemberHandle string) LinkBusinessMember
	// Set the current member whose user handle was already provided as a controlling officer.
	// Controlling officers have leadership roles in the business and have the ability to sign contracts for the business.
	SetControllingOfficerMember() LinkBusinessMember
	// As an admin member, set a different member's user handle as a beneficial owner of this business.
	// Beneficial owners have some ownership stake in the business, which should be included.
	SetBeneficialOwnerMemberAsAdmin(newMemberHandle string, ownershipStake float64) LinkBusinessMember
	// Set the current member whose user handle was already provided as a beneficial owner of this business.
	// Beneficial owners have some ownership stake in the business, which should be included.
	SetBeneficialOwnerMember(ownershipStake float64) LinkBusinessMember
	// Optionally set the description of the member being linked, which can be used to distinguish people and their roles at a later date.
	SetMemberDescription(description string) LinkBusinessMember
	// Execute the Link Business Member call, signing with one of the user's private wallet keys and one of the business's private wallet keys
	Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.LinkBusinessMemberResponse, error)
}

type UnlinkBusinessMember interface {
	// Indicates the business member to unlink is an admin
	SetAdminRole() UnlinkBusinessMember
	// Indicates the business member to unlink is a beneficial owner
	SetBeneficialOwnerRole() UnlinkBusinessMember
	// Indicates the business member to unlink is a controlling officer
	SetControllingOfficerRole() UnlinkBusinessMember
	// Execute the Unlink Business Member call, signing with one of the user to unlink's private wallet keys and one of the business's private wallet keys
	Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.UnlinkBusinessMemberResponse, error)
}

type CertifyBusiness interface {
	// Execute the Certify Business call, signing with one of the user's private wallet keys and one of the business's private wallet keys
	Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.SuccessResponse, error)
}

type CertifyBeneficialOwner interface {
	// The certification token can be fetched from calling GetEntity on a user after the individual to be certified has
	// finished KYC and the business has finished the validation portion of KYC. Calling GetEntity gives the admin user
	// a chance to review the beneficial owner's information is 100% correct.
	SetCertificationToken(userHandleToCertify string, certificationToken string) CertifyBeneficialOwner
	// Execute the Certify Beneficial Owner call, signing with one of the user's private wallet keys and one of the business's private wallet keys
	Do(userWalletPrivateKey string, businessWalletPrivateKey string) (domain.SuccessResponse, error)
}

type AddRegistrationData interface {
	// Sets the email to add
	SetEmail(email string) AddRegistrationData
	// Sets the phone number to add
	SetPhone(phone string) AddRegistrationData
	// Sets the identity information to add. Note that there can not be two of the same type of identity information (SSN or EIN).
	SetIdentity(identityAlias string, identityValue string) AddRegistrationData
	// Sets the address information to add
	SetAddress(address domain.RegistrationAddress) AddRegistrationData
	// Execute the Add Registration Data call, signing with one of the user's private wallet keys
	Do(userWalletPrivateKey string) (domain.ModifyRegistrationDataResponse, error)
}

type UpdateRegistrationData interface {
	// Sets the email information to update, and which email should be updated (via UUID).
	// The UUID can be fetched from GetEntity.
	SetEmail(emailUuid string, email string) UpdateRegistrationData
	// Sets the phone number information to update, and which email should be updated (via UUID).
	// The UUID can be fetched from GetEntity.
	SetPhone(phoneUuid string, phone string) UpdateRegistrationData
	// Sets the identity information to update, and which email should be updated (via UUID).
	// The UUID can be fetched from GetEntity.
	SetIdentity(identityUuid string, identityAlias string, identityValue string) UpdateRegistrationData
	// Sets the address information to update, and which email should be updated (via UUID).
	// The UUID can be fetched from GetEntity.
	SetAddress(addressUuid string, address domain.RegistrationAddress) UpdateRegistrationData
	// Sets the individual specific entity data to update
	SetIndividualEntity(firstName string, lastName string, fullName string, birthDate string) UpdateRegistrationData
	// Sets the business specific entity data to update
	SetBusinessEntity(businessName string, startDate string, businessType string, naicsCode int, doingBusinessAs string, businessWebsite string) UpdateRegistrationData
	// Execute the Update Registration Data call, signing with one of the user's private wallet keys
	Do(userWalletPrivateKey string) (domain.ModifyRegistrationDataResponse, error)
}

type DeleteRegistrationData interface {
	// Sets the UUID of the email data to remove from the user's registration data.
	// The UUID can be fetched from GetEntity.
	SetEmail(emailUuid string) DeleteRegistrationData
	// Sets the UUID of the phone number data to remove from the user's registration data.
	// The UUID can be fetched from GetEntity.
	SetPhone(phoneUuid string) DeleteRegistrationData
	// Sets the UUID of the identity data to remove from the user's registration data.
	// The UUID can be fetched from GetEntity.
	SetIdentity(identityUuid string) DeleteRegistrationData
	// Sets the UUID of the address data to remove from the user's registration data.
	// The UUID can be fetched from GetEntity.
	SetAddress(addressUuid string) DeleteRegistrationData
	// Execute the Delete Registration Data call, signing with one of the user's private wallet keys
	Do(userWalletPrivateKey string) (domain.SuccessResponse, error)
}
