package sila

type GetBusinessRoles struct {
	Header *Header `json:"header"`
}

type GetBusinessRolesResponse struct {
	Success           bool                   `json:"success"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	BusinessRoles     []BusinessRole         `json:"business_roles"`
}

type BusinessRole struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

func (msg *GetBusinessRoles) Do() (GetBusinessRolesResponse, error) {
	var responseBody GetBusinessRolesResponse
	err := instance.performCall("/get_business_roles", msg, &responseBody)
	return responseBody, err
}
