package sila

import (
	"sila/domain"
)

func (client ClientImpl) Register(userHandle string) Register {
	return &RegisterMsg{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "entity_msg",
	}
}

type RegisterMsg struct {
	Header      *Header                    `json:"header"`
	Message     string                     `json:"message"`
	Address     domain.RegistrationAddress `json:"address,omitempty"`
	Identity    RegistrationIdentity       `json:"identity,omitempty"`
	Contact     Contact                    `json:"contact,omitempty"`
	CryptoEntry CryptoEntry                `json:"crypto_entry"`
	Entity      RegistrationEntity         `json:"entity"`
}

func (msg *RegisterMsg) SetRef(ref string) Register {
	msg.Header.setRef(ref)
	return msg
}

func (msg *RegisterMsg) SetAddress(address domain.RegistrationAddress) Register {
	msg.Address = address
	return msg
}

type RegistrationIdentity struct {
	IdentityAlias domain.IdentityType `json:"identity_alias"`
	IdentityValue string              `json:"identity_value"`
}

func (msg *RegisterMsg) SetIdentity(identityType domain.IdentityType, identityValue string) Register {
	msg.Identity.IdentityAlias = identityType
	msg.Identity.IdentityValue = identityValue
	return msg
}

type Contact struct {
	Phone        string `json:"phone,omitempty"`
	ContactAlias string `json:"contact_alias"`
	Email        string `json:"email,omitempty"`
}

func (msg *RegisterMsg) SetContact(contactAlias string, phone string, email string) Register {
	msg.Contact.ContactAlias = contactAlias
	msg.Contact.Phone = phone
	msg.Contact.Email = email
	return msg
}

type CryptoEntry struct {
	CryptoAlias   string `json:"crypto_alias"`
	CryptoAddress string `json:"crypto_address"`
	CryptoCode    string `json:"crypto_code"`
}

func (msg *RegisterMsg) SetCrypto(nickname string, address string) Register {
	msg.CryptoEntry.CryptoAlias = nickname
	msg.CryptoEntry.CryptoAddress = address
	msg.CryptoEntry.CryptoCode = "ETH"
	return msg
}

type RegistrationEntity struct {
	Type            domain.EntityType `json:"type"`
	BirthDate       string            `json:"birthdate,omitempty"`
	FirstName       string            `json:"first_name,omitempty"`
	LastName        string            `json:"last_name,omitempty"`
	EntityName      string            `json:"entity_name,omitempty"`
	BusinessType    string            `json:"business_type,omitempty"`
	BusinessWebsite string            `json:"business_website,omitempty"`
	DoingBusinessAs string            `json:"doing_business_as,omitempty"`
	NaicsCode       int               `json:"naics_code,omitempty"`
}

func (msg *RegisterMsg) SetIndividualEntity(firstName string, lastName string, birthDate string) Register {
	msg.Entity.Type = domain.Individual
	msg.Entity.FirstName = firstName
	msg.Entity.LastName = lastName
	msg.Entity.BirthDate = birthDate
	return msg
}

func (msg *RegisterMsg) SetBusinessEntity(entityName string, businessType string, naicsCode int) Register {
	msg.Entity.Type = domain.Business
	msg.Entity.EntityName = entityName
	msg.Entity.BusinessType = businessType
	msg.Entity.NaicsCode = naicsCode
	return msg
}

func (msg *RegisterMsg) SetBusinessWebsite(businessWebsite string) Register {
	msg.Entity.BusinessWebsite = businessWebsite
	return msg
}

func (msg *RegisterMsg) SetDoingBusinessAs(dba string) Register {
	msg.Entity.DoingBusinessAs = dba
	return msg
}

func (msg *RegisterMsg) Do() (domain.SuccessResponse, error) {
	var responseBody domain.SuccessResponse
	err := instance.performCall("/register", msg, &responseBody)
	return responseBody, err
}
