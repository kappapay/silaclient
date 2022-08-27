package sila

import "github.com/kappapay/silaclient/domain"

func (client ClientImpl) UpdateRegistrationData(userHandle string) UpdateRegistrationData {
	return &UpdateRegistrationDataMsg{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

type UpdateRegistrationDataMsg struct {
	Header          *Header                     `json:"header"`
	Uuid            string                      `json:"uuid,omitempty"`
	Email           string                      `json:"email,omitempty"`
	Phone           string                      `json:"phone,omitempty"`
	IdentityAlias   string                      `json:"identity_alias,omitempty"`
	IdentityValue   string                      `json:"identity_value,omitempty"`
	AddressAlias    string                      `json:"address_alias,omitempty"`
	StreetAddress1  string                      `json:"street_address_1,omitempty"`
	StreetAddress2  string                      `json:"street_address_2,omitempty"`
	City            string                      `json:"city,omitempty"`
	State           string                      `json:"state,omitempty"`
	Country         string                      `json:"country,omitempty"`
	PostalCode      string                      `json:"postal_code,omitempty"`
	FirstName       string                      `json:"first_name,omitempty"`
	LastName        string                      `json:"last_name,omitempty"`
	EntityName      string                      `json:"entity_name,omitempty"`
	BirthDate       string                      `json:"birth_date,omitempty"`
	BusinessType    string                      `json:"business_type,omitempty"`
	NaicsCode       int                         `json:"naics_code,omitempty"`
	DoingBusinessAs string                      `json:"doing_business_as,omitempty"`
	BusinessWebsite string                      `json:"business_website,omitempty"`
	DataType        domain.RegistrationDataType `json:"-"`
}

func (msg *UpdateRegistrationDataMsg) SetEmail(emailUuid string, email string) UpdateRegistrationData {
	msg.Uuid = emailUuid
	msg.Email = email
	msg.DataType = domain.EmailDataType
	return msg
}

func (msg *UpdateRegistrationDataMsg) SetPhone(phoneUuid string, phone string) UpdateRegistrationData {
	msg.Uuid = phoneUuid
	msg.Phone = phone
	msg.DataType = domain.PhoneDataType
	return msg
}

func (msg *UpdateRegistrationDataMsg) SetIdentity(identityUuid string, identityAlias string, identityValue string) UpdateRegistrationData {
	msg.Uuid = identityUuid
	msg.IdentityAlias = identityAlias
	msg.IdentityValue = identityValue
	msg.DataType = domain.IdentityDataType
	return msg
}

func (msg *UpdateRegistrationDataMsg) SetAddress(addressUuid string, address domain.RegistrationAddress) UpdateRegistrationData {
	msg.Uuid = addressUuid
	msg.StreetAddress1 = address.StreetAddress1
	msg.StreetAddress2 = address.StreetAddress2
	msg.City = address.City
	msg.State = address.State
	msg.Country = address.Country
	msg.PostalCode = address.PostalCode
	msg.DataType = domain.AddressDataType
	return msg
}

func (msg *UpdateRegistrationDataMsg) SetIndividualEntity(firstName string, lastName string, fullName string, birthDate string) UpdateRegistrationData {
	msg.FirstName = firstName
	msg.LastName = lastName
	msg.EntityName = fullName
	msg.BirthDate = birthDate
	msg.DataType = domain.EntityDataType
	return msg
}

func (msg *UpdateRegistrationDataMsg) SetBusinessEntity(businessName string, startDate string, businessType string, naicsCode int, doingBusinessAs string, businessWebsite string) UpdateRegistrationData {
	msg.EntityName = businessName
	msg.BirthDate = startDate
	msg.BusinessType = businessType
	msg.NaicsCode = naicsCode
	msg.DoingBusinessAs = doingBusinessAs
	msg.BusinessWebsite = businessWebsite
	msg.DataType = domain.EntityDataType
	return msg
}

func (msg *UpdateRegistrationDataMsg) Do(userWalletPrivateKey string) (domain.ModifyRegistrationDataResponse, error) {
	var responseBody domain.ModifyRegistrationDataResponse
	err := instance.performCallWithUserAuth("/update/"+string(msg.DataType), msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
