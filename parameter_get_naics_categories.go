package sila

import (
	"encoding/json"
)

type GetNaicsCategories struct {
	Header *Header `json:"header"`
}

type GetNaicsCategoriesResponse struct {
	Success           bool                       `json:"success"`
	Message           string                     `json:"message"`
	Status            string                     `json:"status"`
	ValidationDetails map[string]interface{}     `json:"validation_details"`
	NaicsCategories   map[string][]NaicsCategory `json:"naics_categories"`
}

func (gncr *GetNaicsCategoriesResponse) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "success":
			gncr.Success = value.(bool)
		case "message":
			gncr.Message = value.(string)
		case "status":
			gncr.Status = value.(string)
		case "validation_details":
			gncr.ValidationDetails = value.(map[string]interface{})
		case "naics_categories":
			naicsCategories := make(map[string][]NaicsCategory)
			rawCategories := value.(map[string]interface{})
			for category, rawSubCategories := range rawCategories {
				subCategories := rawSubCategories.([]interface{})
				naicsSubCategories := make([]NaicsCategory, len(subCategories))
				for index, subCategory := range subCategories {
					naicsSubCategory := subCategory.(map[string]interface{})
					naicsSubCategories[index] = NaicsCategory{
						Code:        naicsSubCategory["code"].(float64),
						Subcategory: naicsSubCategory["subcategory"].(string),
					}
				}
				naicsCategories[category] = naicsSubCategories
			}
			gncr.NaicsCategories = naicsCategories
		}
	}
	return nil
}

type NaicsCategory struct {
	Code        float64 `json:"code"`
	Subcategory string  `json:"subcategory"`
}

func (msg *GetNaicsCategories) Do() (GetNaicsCategoriesResponse, error) {
	var responseBody GetNaicsCategoriesResponse
	err := instance.performCall("/get_naics_categories", msg, &responseBody)
	return responseBody, err
}
