package sila

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type CheckKyc struct {
	Header   *Header `json:"header"`
	Message  string  `json:"message"`
	KycLevel string  `json:"kyc_level,omitempty"`
}

func (msg *CheckKyc) SetRef(ref string) *CheckKyc {
	msg.Header.setRef(ref)
	return msg
}

func (msg *CheckKyc) SetKycLevel(kycLevel string) *CheckKyc {
	msg.KycLevel = kycLevel
	return msg
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
			vh.Reasons = value.([]string)
		case "tags":
			vh.Tags = value.([]string)
		case "score":
			vh.Score = value.(float64)
		case "valid_kyc_levels":
			vh.ValidKycLevels = value.([]string)
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
			ch.BeneficialOwnerCertifications = value.([]string)
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

func (msg *CheckKyc) Do(userWalletPrivateKey string) (CheckKycResponse, error) {
	var responseBody CheckKycResponse
	requestJson, err := json.Marshal(msg)
	if err != nil {
		return responseBody, nil
	}
	url := instance.environment.generateURL(instance.version, "/request_kyc")
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		return responseBody, err
	}
	request.Header.Set("Content-type", "application/json")
	authSignature, err := instance.GenerateAuthSignature(requestJson)
	if err != nil {
		return responseBody, errors.Errorf("failed to generate auth signature: %v", err)
	}
	request.Header.Set("authsignature", authSignature)
	userSignature, err := instance.GenerateUserSignature(requestJson, userWalletPrivateKey)
	if err != nil {
		return responseBody, errors.Errorf("failed to generate user signature: %v", err)
	}
	request.Header.Set("usersignature", userSignature)
	httpClient := http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		return responseBody, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return responseBody, err
	}
	return responseBody, nil
}
