package sila

type PlaidSameDayAuth struct {
	Header      *Header `json:"header"`
	AccountName string  `json:"account_name"`
}

type PlaidSameDayAuthResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	PublicToken       string                 `json:"public_token"`
}

func (msg *PlaidSameDayAuth) Do() (PlaidSameDayAuthResponse, error) {
	var responseBody PlaidSameDayAuthResponse
	err := instance.performCall("/plaid_sameday_auth", msg, &responseBody)
	return responseBody, err
}
