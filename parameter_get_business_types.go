package sila

type GetBusinessTypes struct {
	Header *Header `json:"header"`
}

type GetBusinessTypesResponse struct {
	Success           bool                   `json:"success"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	BusinessTypes     []BusinessType         `json:"business_types"`
}

type BusinessType struct {
	Uuid                  string `json:"uuid"`
	Name                  string `json:"name"`
	Label                 string `json:"label"`
	RequiresCertification bool   `json:"requires_certification"`
}

func (msg *GetBusinessTypes) Do() (GetBusinessTypesResponse, error) {
	var responseBody GetBusinessTypesResponse
	err := instance.performCall("/get_business_types", msg, &responseBody)
	return responseBody, err
}
