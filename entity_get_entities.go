package sila

import (
	"encoding/json"
	"strconv"
	"time"
)

type GetEntities struct {
	Header     *Header `json:"header"`
	Message    string  `json:"message"`
	EntityType string  `json:"entity_type,omitempty"`
	Page       int32   `json:"-"`
	PerPage    int32   `json:"-"`
}

func (msg *GetEntities) SetRef(ref string) *GetEntities {
	msg.Header.setRef(ref)
	return msg
}

func (msg *GetEntities) SetEntityType(entityType string) *GetEntities {
	msg.EntityType = entityType
	return msg
}

func (msg *GetEntities) SetPage(page int32) *GetEntities {
	msg.Page = page
	return msg
}

func (msg *GetEntities) SetPerPage(perPage int32) *GetEntities {
	msg.PerPage = perPage
	return msg
}

type GetEntitiesResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	Entities          Entities               `json:"entities"`
	Pagination        EntityPagination       `json:"pagination"`
}

type Entities struct {
	Individuals []IndividualEntity `json:"individuals"`
	Businesses  []BusinessEntity   `json:"businesses"`
}

type IndividualEntity struct {
	Handle              string    `json:"handle"`
	FullName            string    `json:"full_name"`
	Created             time.Time `json:"created"`
	Status              string    `json:"status"`
	BlockchainAddresses []string  `json:"blockchain_addresses"`
}

func (ie *IndividualEntity) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "handle":
			ie.Handle = value.(string)
		case "full_name":
			ie.FullName = value.(string)
		case "created":
			ie.Created = time.Unix(int64(value.(float64)), 0)
		case "status":
			ie.Status = value.(string)
		case "blockchain_addresses":
			ie.BlockchainAddresses = value.([]string)
		}
	}
	return nil
}

type BusinessEntity struct {
	Handle              string    `json:"handle"`
	FullName            string    `json:"full_name"`
	Created             time.Time `json:"created"`
	Status              string    `json:"status"`
	BlockchainAddresses []string  `json:"blockchain_addresses"`
	Uuid                string    `json:"uuid"`
	BusinessType        string    `json:"business_type"`
	Dba                 string    `json:"dba"`
}

func (be *BusinessEntity) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "handle":
			be.Handle = value.(string)
		case "full_name":
			be.FullName = value.(string)
		case "created":
			be.Created = time.Unix(int64(value.(float64)), 0)
		case "status":
			be.Status = value.(string)
		case "blockchain_addresses":
			be.BlockchainAddresses = value.([]string)
		case "uuid":
			be.Uuid = value.(string)
		case "business_type":
			be.BusinessType = value.(string)
		case "dba":
			be.Dba = value.(string)
		}
	}
	return nil
}

type EntityPagination struct {
	ReturnedCount int32 `json:"returned_count"`
	TotalCount    int64 `json:"total_count"`
	CurrentPage   int32 `json:"current_page"`
	TotalPages    int32 `json:"total_pages"`
}

func (msg *GetEntities) Do() (GetEntitiesResponse, error) {
	var responseBody GetEntitiesResponse
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
