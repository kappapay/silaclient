package domain

import (
	"encoding/json"
	"time"
)

type EntityType string

const (
	Individual EntityType = "individual"
	Business              = "business"
)

type RegistrationAddress struct {
	AddressAlias   string `json:"address_alias,omitempty"`
	StreetAddress1 string `json:"street_address_1,omitempty"`
	StreetAddress2 string `json:"street_address_2,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	Country        string `json:"country,omitempty"`
	PostalCode     string `json:"postal_code,omitempty"`
}

type IdentityType string

const (
	Unknown IdentityType = "UNKNOWN"
	Ssn                  = "SSN"
	Ein                  = "EIN"
)

func GetIdentityType(identityType string) IdentityType {
	switch identityType {
	case Ssn:
		return Ssn
	case Ein:
		return Ein
	default:
		return Unknown
	}
}

type CheckKycResponse struct {
	Success              bool                   `json:"success"`
	Reference            string                 `json:"reference"`
	Message              string                 `json:"message"`
	Status               string                 `json:"status"`
	ValidationDetails    map[string]interface{} `json:"validation_details"`
	EntityType           string                 `json:"entity_type"`
	VerificationStatus   string                 `json:"verification_status"`
	VerificationHistory  []VerificationHistory  `json:"verification_history"`
	ValidKycLevels       []string               `json:"valid_kyc_levels"`
	CertificationStatus  string                 `json:"certification_status,omitempty"`
	CertificationHistory []CertificationHistory `json:"certification_history,omitempty"`
	Members              []Member               `json:"members"`
}

type VerificationHistory struct {
	VerificationId     string    `json:"verification_id"`
	VerificationStatus string    `json:"verification_status"`
	KycLevel           string    `json:"kyc_level"`
	RequestedAt        time.Time `json:"requested_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Reasons            []string  `json:"reasons"`
	Tags               []string  `json:"tags"`
	Score              float64   `json:"score,omitempty"`
	ValidKycLevels     []string  `json:"valid_kyc_levels"`
}

func (vh *VerificationHistory) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "verification_id":
			vh.VerificationId = value.(string)
		case "verification_status":
			vh.VerificationStatus = value.(string)
		case "kyc_level":
			vh.KycLevel = value.(string)
		case "requested_at":
			vh.RequestedAt = time.Unix(int64(value.(float64)), 0)
		case "updated_at":
			vh.UpdatedAt = time.Unix(int64(value.(float64)), 0)
		case "reasons":
			arrValue := value.([]interface{})
			convertedValue := make([]string, len(arrValue))
			for index, rawValue := range arrValue {
				convertedValue[index] = rawValue.(string)
			}
			vh.Reasons = convertedValue
		case "tags":
			arrValue := value.([]interface{})
			convertedValue := make([]string, len(arrValue))
			for index, rawValue := range arrValue {
				convertedValue[index] = rawValue.(string)
			}
			vh.Tags = convertedValue
		case "score":
			if value != nil {
				vh.Score = value.(float64)
			}
		case "valid_kyc_levels":

			arrValue := value.([]interface{})
			convertedValue := make([]string, len(arrValue))
			for index, rawValue := range arrValue {
				convertedValue[index] = rawValue.(string)
			}
			vh.ValidKycLevels = convertedValue
		}
	}
	return nil
}

type CertificationHistory struct {
	AdministratorUserHandle       string    `json:"administrator_user_handle"`
	Created                       string    `json:"created"`
	CreatedTime                   time.Time `json:"created_epoch"`
	ExpiresAfter                  string    `json:"expires_after"`
	ExpiresAfterTime              time.Time `json:"expires_after_epoch"`
	BeneficialOwnerCertifications []string  `json:"beneficial_owner_certifications"`
}

func (ch *CertificationHistory) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "administrator_user_handle":
			ch.AdministratorUserHandle = value.(string)
		case "created":
			ch.Created = value.(string)
		case "created_epoch":
			ch.CreatedTime = time.Unix(int64(value.(float64)), 0)
		case "expires_after":
			ch.ExpiresAfter = value.(string)
		case "expires_after_epoch":
			ch.ExpiresAfterTime = time.Unix(int64(value.(float64)), 0)
		case "beneficial_owner_certifications":
			arrValue := value.([]interface{})
			convertedValue := make([]string, len(arrValue))
			for index, rawValue := range arrValue {
				convertedValue[index] = rawValue.(string)
			}
			ch.BeneficialOwnerCertifications = convertedValue
		}
	}
	return nil
}

type Member struct {
	UserHandle                         string  `json:"user_handle"`
	FirstName                          string  `json:"first_name"`
	LastName                           string  `json:"last_name"`
	Role                               string  `json:"role"`
	Details                            string  `json:"details"`
	OwnershipStake                     float64 `json:"ownership_stake"`
	VerificationStatus                 string  `json:"verification_status"`
	VerificationRequired               bool    `json:"verification_required"`
	VerificationId                     string  `json:"verification_id"`
	BeneficialOwnerCertificationStatus string  `json:"beneficial_owner_certification_status"`
	BusinessCertificationStatus        string  `json:"business_certification_status"`
}

type GetEntitiesResponse struct {
	Success           bool                   `json:"success"`
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
			if value != nil {
				ie.FullName = value.(string)
			}
		case "created":
			ie.Created = time.Unix(int64(value.(float64)), 0)
		case "status":
			ie.Status = value.(string)
		case "blockchain_addresses":
			arrValue := value.([]interface{})
			convertedValue := make([]string, len(arrValue))
			for index, rawValue := range arrValue {
				convertedValue[index] = rawValue.(string)
			}
			ie.BlockchainAddresses = convertedValue
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
			arrValue := value.([]interface{})
			convertedValue := make([]string, len(arrValue))
			for index, rawValue := range arrValue {
				convertedValue[index] = rawValue.(string)
			}
			be.BlockchainAddresses = convertedValue
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

type RequestKycResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	VerificationUuid  string                 `json:"verification_uuid"`
}
