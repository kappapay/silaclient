package sila

import (
	"github.com/bpancost/sila/domain"
)

func (client ClientImpl) GetBusinessRoles() GetBusinessRoles {
	return &GetBusinessRolesMsg{
		Header: client.generateHeader(),
	}
}

type GetBusinessRolesMsg struct {
	Header *Header `json:"header"`
}

func (msg *GetBusinessRolesMsg) Do() (domain.GetBusinessRolesResponse, error) {
	var responseBody domain.GetBusinessRolesResponse
	err := instance.performCall("/get_business_roles", msg, &responseBody)
	return responseBody, err
}
