package sila

import (
	"sila/domain"
)

func (client ClientImpl) GetBusinessTypes() GetBusinessTypes {
	return &GetBusinessTypesMsg{
		Header: client.generateHeader(),
	}
}

type GetBusinessTypesMsg struct {
	Header *Header `json:"header"`
}

func (msg *GetBusinessTypesMsg) Do() (domain.GetBusinessTypesResponse, error) {
	var responseBody domain.GetBusinessTypesResponse
	err := instance.performCall("/get_business_types", msg, &responseBody)
	return responseBody, err
}
