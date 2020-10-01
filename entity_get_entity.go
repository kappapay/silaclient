package sila

import (
	"encoding/json"
	"time"
)

type GetEntity struct {
	Header *Header `json:"header"`
}

type GetEntityResponse struct {
	Success           bool                   `json:"success"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	UserHandle        string                 `json:"user_handle"`
	EntityType        EntityType             `json:"entity_type"`
	Entity            Entity                 `json:"entity"`
	Addresses         []EntityAddress        `json:"addresses"`
	Identities        []EntityIdentity       `json:"identities"`
	Emails            []EntityEmail          `json:"emails"`
	Phones            []EntityPhone          `json:"phones"`
	Memberships       []EntityMembership     `json:"memberships"`
}

type Entity struct {
	CreatedTime time.Time `json:"created_epoch"`
	EntityName  string    `json:"entity_name"`
	Birthdate   string    `json:"birthdate"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
}

func (e *Entity) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "created_epoch":
			e.CreatedTime = time.Unix(int64(value.(float64)), 0)
		case "entity_name":
			e.EntityName = value.(string)
		case "birthdate":
			e.Birthdate = value.(string)
		case "first_name":
			e.FirstName = value.(string)
		case "last_name":
			e.LastName = value.(string)
		}
	}
	return nil
}

type EntityAddress struct {
	AddedTime      time.Time `json:"added_epoch"`
	ModifiedTime   time.Time `json:"modified_epoch"`
	Uuid           string    `json:"uuid"`
	Nickname       string    `json:"nickname"`
	StreetAddress1 string    `json:"street_address_1"`
	StreetAddress2 string    `json:"street_address_2"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	Country        string    `json:"country"`
	PostalCode     string    `json:"postal_code"`
}

func (ea *EntityAddress) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "added_epoch":
			ea.AddedTime = time.Unix(int64(value.(float64)), 0)
		case "modified_epoch":
			ea.ModifiedTime = time.Unix(int64(value.(float64)), 0)
		case "uuid":
			ea.Uuid = value.(string)
		case "nickname":
			ea.Nickname = value.(string)
		case "street_address_1":
			ea.StreetAddress1 = value.(string)
		case "street_address_2":
			ea.StreetAddress2 = value.(string)
		case "city":
			ea.City = value.(string)
		case "state":
			ea.State = value.(string)
		case "country":
			ea.Country = value.(string)
		case "postal_code":
			ea.PostalCode = value.(string)
		}
	}
	return nil
}

type EntityIdentity struct {
	AddedTime    time.Time    `json:"added_epoch"`
	ModifiedTime time.Time    `json:"modified_epoch"`
	Uuid         string       `json:"uuid"`
	IdentityType IdentityType `json:"identity_type"`
	Identity     string       `json:"identity"`
}

func (ei *EntityIdentity) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "added_epoch":
			ei.AddedTime = time.Unix(int64(value.(float64)), 0)
		case "modified_epoch":
			ei.ModifiedTime = time.Unix(int64(value.(float64)), 0)
		case "uuid":
			ei.Uuid = value.(string)
		case "identity_type":
			ei.IdentityType = GetIdentityType(value.(string))
		case "identity":
			ei.Identity = value.(string)
		}
	}
	return nil
}

type EntityEmail struct {
	AddedTime    time.Time `json:"added_epoch"`
	ModifiedTime time.Time `json:"modified_epoch"`
	Uuid         string    `json:"uuid"`
	Email        string    `json:"email"`
}

func (ee *EntityEmail) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "added_epoch":
			ee.AddedTime = time.Unix(int64(value.(float64)), 0)
		case "modified_epoch":
			ee.ModifiedTime = time.Unix(int64(value.(float64)), 0)
		case "uuid":
			ee.Uuid = value.(string)
		case "email":
			ee.Email = value.(string)
		}
	}
	return nil
}

type EntityPhone struct {
	AddedTime    time.Time `json:"added_epoch"`
	ModifiedTime time.Time `json:"modified_epoch"`
	Uuid         string    `json:"uuid"`
	Phone        string    `json:"phone"`
}

func (ep *EntityPhone) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "added_epoch":
			ep.AddedTime = time.Unix(int64(value.(float64)), 0)
		case "modified_epoch":
			ep.ModifiedTime = time.Unix(int64(value.(float64)), 0)
		case "uuid":
			ep.Uuid = value.(string)
		case "phone":
			ep.Phone = value.(string)
		}
	}
	return nil
}

type EntityMembership struct {
	BusinessHandle     string  `json:"business_handle"`
	EntityName         string  `json:"entity_name"`
	Role               string  `json:"role"`
	Details            string  `json:"details"`
	OwnershipStake     float64 `json:"ownership_stake"`
	CertificationToken string  `json:"certification_token"`
}

// Get wallet key should belong to the user/entity to fetch
func (msg *GetEntity) Do(userWalletPrivateKey string) (GetEntityResponse, error) {
	var responseBody GetEntityResponse
	err := instance.performCallWithUserAuth("/get_entity", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
