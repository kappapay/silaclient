package sila

import "github.com/kappapay/silaclient/domain"

func (client ClientImpl) AddRegistrationData(userHandle string) AddRegistrationData {
	return &AddRegistrationDataMsg{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

type AddRegistrationDataMsg struct {
	Header         *Header                     `json:"header"`
	Email          string                      `json:"email,omitempty"`
	Phone          string                      `json:"phone,omitempty"`
	IdentityAlias  string                      `json:"identity_alias,omitempty"`
	IdentityValue  string                      `json:"identity_value,omitempty"`
	AddressAlias   string                      `json:"address_alias,omitempty"`
	StreetAddress1 string                      `json:"street_address_1,omitempty"`
	StreetAddress2 string                      `json:"street_address_2,omitempty"`
	City           string                      `json:"city,omitempty"`
	State          string                      `json:"state,omitempty"`
	Country        string                      `json:"country,omitempty"`
	PostalCode     string                      `json:"postal_code,omitempty"`
	DataType       domain.RegistrationDataType `json:"-"`
}

func (msg *AddRegistrationDataMsg) SetEmail(email string) AddRegistrationData {
	msg.Email = email
	msg.DataType = domain.EmailDataType
	return msg
}

func (msg *AddRegistrationDataMsg) SetPhone(phone string) AddRegistrationData {
	msg.Phone = phone
	msg.DataType = domain.PhoneDataType
	return msg
}

func (msg *AddRegistrationDataMsg) SetIdentity(identityAlias string, identityValue string) AddRegistrationData {
	msg.IdentityAlias = identityAlias
	msg.IdentityValue = identityValue
	msg.DataType = domain.IdentityDataType
	return msg
}

func (msg *AddRegistrationDataMsg) SetAddress(address domain.RegistrationAddress) AddRegistrationData {
	msg.StreetAddress1 = address.StreetAddress1
	msg.StreetAddress2 = address.StreetAddress2
	msg.City = address.City
	msg.State = address.State
	msg.Country = address.Country
	msg.PostalCode = address.PostalCode
	msg.DataType = domain.IdentityDataType
	return msg
}

func (msg *AddRegistrationDataMsg) Do(userWalletPrivateKey string) (domain.ModifyRegistrationDataResponse, error) {
	var responseBody domain.ModifyRegistrationDataResponse
	err := instance.performCallWithUserAuth("/add/"+string(msg.DataType), msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
