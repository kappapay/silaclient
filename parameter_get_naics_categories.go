package sila

import (
	"github.com/kappapay/silaclient/domain"
)

func (client ClientImpl) GetNaicsCategories() GetNaicsCategories {
	return &GetNaicsCategoriesMsg{
		Header: client.generateHeader(),
	}
}

type GetNaicsCategoriesMsg struct {
	Header *Header `json:"header"`
}

func (msg *GetNaicsCategoriesMsg) Do() (domain.GetNaicsCategoriesResponse, error) {
	var responseBody domain.GetNaicsCategoriesResponse
	err := instance.performCall("/get_naics_categories", msg, &responseBody)
	return responseBody, err
}
