package sila

import "github.com/bpancost/sila/domain"

func (client ClientImpl) DeleteRegistrationData(userHandle string) DeleteRegistrationData {
	return &DeleteRegistrationDataMsg{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

type DeleteRegistrationDataMsg struct {
	Header   *Header                     `json:"header"`
	Uuid     string                      `json:"uuid"`
	DataType domain.RegistrationDataType `json:"-"`
}

func (msg *DeleteRegistrationDataMsg) SetEmail(addressUuid string) DeleteRegistrationData {
	msg.Uuid = addressUuid
	msg.DataType = domain.EmailDataType
	return msg
}

func (msg *DeleteRegistrationDataMsg) SetPhone(phoneUuid string) DeleteRegistrationData {
	msg.Uuid = phoneUuid
	msg.DataType = domain.PhoneDataType
	return msg
}

func (msg *DeleteRegistrationDataMsg) SetIdentity(identityUuid string) DeleteRegistrationData {
	msg.Uuid = identityUuid
	msg.DataType = domain.IdentityDataType
	return msg
}

func (msg *DeleteRegistrationDataMsg) SetAddress(addressUuid string) DeleteRegistrationData {
	msg.Uuid = addressUuid
	msg.DataType = domain.IdentityDataType
	return msg
}

func (msg *DeleteRegistrationDataMsg) Do(userWalletPrivateKey string) (domain.SuccessResponse, error) {
	var responseBody domain.SuccessResponse
	err := instance.performCallWithUserAuth("/delete/"+string(msg.DataType), msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
