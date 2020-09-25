package sila

type PlaidSameDayAuth struct {
	Header      *Header `json:"header"`
	AccountName string  `json:"account_name"`
}

func (msg *PlaidSameDayAuth) Do() (SuccessResponse, error) {
	var responseBody SuccessResponse
	err := instance.performCall("/plaid_sameday_auth", msg, &responseBody)
	return responseBody, err
}
