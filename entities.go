package sila

import "github.com/bpancost/sila/domain"

type CheckHandle interface {
	SetRef(ref string) CheckHandle
	Do() (domain.SuccessResponse, error)
}

type Register interface {
	SetRef(ref string) Register
	SetAddress(address domain.RegistrationAddress) Register
	SetIdentity(identityType domain.IdentityType, identityValue string) Register
	SetContact(contactAlias string, phone string, email string) Register
	SetCrypto(nickname string, address string) Register
	SetIndividualEntity(firstName string, lastName string, birthDate string) Register
	SetBusinessEntity(entityName string, businessType string, naicsCode int) Register
	SetBusinessWebsite(businessWebsite string) Register
	SetDoingBusinessAs(dba string) Register
	Do() (domain.SuccessResponse, error)
}

type RequestKyc interface {
	SetRef(ref string) RequestKyc
	SetKycLevel(kycLevel string) RequestKyc
	Do(userWalletPrivateKey string) (domain.RequestKycResponse, error)
}

type CheckKyc interface {
	SetRef(ref string) CheckKyc
	SetKycLevel(kycLevel string) CheckKyc
	Do(userWalletPrivateKey string) (domain.CheckKycResponse, error)
}

type GetEntity interface {
	Do(userWalletPrivateKey string) (domain.GetEntityResponse, error)
}

type GetEntities interface {
	SetEntityType(entityType string) GetEntities
	SetPage(page int32) GetEntities
	SetPerPage(perPage int32) GetEntities
	Do() (domain.GetEntitiesResponse, error)
}
