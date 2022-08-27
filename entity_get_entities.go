package sila

import (
	"strconv"

	"github.com/kappapay/silaclient/domain"
)

func (client ClientImpl) GetEntities() GetEntities {
	return &GetEntitiesMsg{
		Header:  client.generateHeader(),
		Message: "header_msg",
	}
}

type GetEntitiesMsg struct {
	Header     *Header `json:"header"`
	Message    string  `json:"message"`
	EntityType string  `json:"entity_type,omitempty"`
	Page       int32   `json:"-"`
	PerPage    int32   `json:"-"`
}

func (msg *GetEntitiesMsg) SetEntityType(entityType string) GetEntities {
	msg.EntityType = entityType
	return msg
}

func (msg *GetEntitiesMsg) SetPage(page int32) GetEntities {
	msg.Page = page
	return msg
}

func (msg *GetEntitiesMsg) SetPerPage(perPage int32) GetEntities {
	msg.PerPage = perPage
	return msg
}

func (msg *GetEntitiesMsg) Do() (domain.GetEntitiesResponse, error) {
	var responseBody domain.GetEntitiesResponse
	path := "/get_entities"
	var params string
	if msg.Page > 0 {
		params = "?page=" + strconv.FormatInt(int64(msg.Page), 10)
	}
	if msg.PerPage > 0 {
		if len(params) > 0 {
			params += "&per_page=" + strconv.FormatInt(int64(msg.PerPage), 10)
		} else {
			params = "?per_page=" + strconv.FormatInt(int64(msg.PerPage), 10)
		}
	}
	err := instance.performCall(path+params, msg, &responseBody)
	return responseBody, err
}
