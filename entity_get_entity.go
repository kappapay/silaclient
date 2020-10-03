package sila

import (
	"github.com/bpancost/sila/domain"
)

func (client ClientImpl) GetEntity(userHandle string) GetEntity {
	return &GetEntityMsg{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

type GetEntityMsg struct {
	Header *Header `json:"header"`
}

// Get wallet key should belong to the user/entity to fetch
func (msg *GetEntityMsg) Do(userWalletPrivateKey string) (domain.GetEntityResponse, error) {
	var responseBody domain.GetEntityResponse
	err := instance.performCallWithUserAuth("/get_entity", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
