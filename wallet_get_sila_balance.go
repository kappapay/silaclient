package sila

type GetSilaBalance struct {
	Address string `json:"address"`
}

type GetSilaBalanceResponse struct {
	Success           bool                   `json:"success"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	Address           string                 `json:"address"`
	SilaBalance       int64                  `json:"sila_balance"`
}

func (msg *GetSilaBalance) Do() (GetSilaBalanceResponse, error) {
	var responseBody GetSilaBalanceResponse
	err := instance.performPublicCall("/get_sila_balance", msg, &responseBody)
	return responseBody, err
}
