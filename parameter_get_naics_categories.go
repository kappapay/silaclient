package sila

type GetNaicsCategories struct {
	Header *Header `json:"header"`
}

type GetNaicsCategoriesResponse struct {
	Success           bool                     `json:"success"`
	Message           string                   `json:"message"`
	Status            string                   `json:"status"`
	ValidationDetails map[string]interface{}   `json:"validation_details"`
	NaicsCategories   map[string]NaicsCategory `json:"naics_categories"`
}

type NaicsCategory struct {
	Code        string `json:"code"`
	Subcategory string `json:"subcategory"`
}

func (msg *GetNaicsCategories) Do() (GetNaicsCategoriesResponse, error) {
	var responseBody GetNaicsCategoriesResponse
	err := instance.performCall("/get_naics_categories", msg, &responseBody)
	return responseBody, err
}
