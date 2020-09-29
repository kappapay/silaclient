package sila

import (
	"encoding/json"
	"time"
)

type GetTransactions struct {
	Header        *Header                  `json:"header"`
	Message       string                   `json:"message"`
	SearchFilters TransactionSearchFilters `json:"search_filters,omitempty"`
}

func (msg *GetTransactions) SetRef(ref string) *GetTransactions {
	msg.Header.setRef(ref)
	return msg
}

type TransactionSearchFilters struct {
	TransactionId    string    `json:"transaction_id,omitempty"`
	ReferenceId      string    `json:"reference_id,omitempty"`
	ShowTimelines    bool      `json:"show_timelines,omitempty"`
	SortAscending    bool      `json:"sort_ascending,omitempty"`
	MaxSilaAmount    int64     `json:"max_sila_amount,omitempty"`
	MinSilaAmount    int64     `json:"min_sila_amount,omitempty"`
	Statuses         []string  `json:"statuses,omitempty"`
	StartTime        time.Time `json:"start_epoch,omitempty"`
	EndTime          time.Time `json:"end_epoch,omitempty"`
	Page             int32     `json:"page,omitempty"`
	PerPage          int32     `json:"per_page,omitempty"`
	TransactionTypes []string  `json:"transaction_types,omitempty"`
}

func (filters *TransactionSearchFilters) MarshalJSON() ([]byte, error) {
	convertedSearchFilters := struct {
		TransactionId    string   `json:"transaction_id,omitempty"`
		ReferenceId      string   `json:"reference_id,omitempty"`
		ShowTimelines    bool     `json:"show_timelines,omitempty"`
		SortAscending    bool     `json:"sort_ascending,omitempty"`
		MaxSilaAmount    int64    `json:"max_sila_amount,omitempty"`
		MinSilaAmount    int64    `json:"min_sila_amount,omitempty"`
		Statuses         []string `json:"statuses,omitempty"`
		StartTime        int64    `json:"start_epoch,omitempty"`
		EndTime          int64    `json:"end_epoch,omitempty"`
		Page             int32    `json:"page,omitempty"`
		PerPage          int32    `json:"per_page,omitempty"`
		TransactionTypes []string `json:"transaction_types,omitempty"`
	}{
		TransactionId:    filters.TransactionId,
		ReferenceId:      filters.ReferenceId,
		ShowTimelines:    filters.ShowTimelines,
		SortAscending:    filters.SortAscending,
		MaxSilaAmount:    filters.MaxSilaAmount,
		MinSilaAmount:    filters.MinSilaAmount,
		Statuses:         filters.Statuses,
		StartTime:        filters.StartTime.Unix(),
		EndTime:          filters.EndTime.Unix(),
		Page:             filters.Page,
		PerPage:          filters.PerPage,
		TransactionTypes: filters.TransactionTypes,
	}

	return json.Marshal(convertedSearchFilters)
}

func (msg *GetTransactions) SetSearchFilters(searchFilters TransactionSearchFilters) *GetTransactions {
	msg.SearchFilters = searchFilters
	return msg
}

type GetTransactionsResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	Page              int32                  `json:"page"`
	ReturnedCount     int32                  `json:"returned_count"`
	TotalCount        int64                  `json:"total_count"`
	Transactions      []Transaction          `json:"transactions"`
}

type Transaction struct {
	UserHandle         string                 `json:"user_handle"`
	ReferenceId        string                 `json:"reference_id"`
	TransactionId      string                 `json:"transaction_id"`
	TransactionHash    string                 `json:"transaction_hash"`
	TransactionType    string                 `json:"transaction_type"`
	SilaAmount         int64                  `json:"sila_amount"`
	BankAccountName    string                 `json:"bank_account_name"`
	Status             string                 `json:"status"`
	UsdStatus          string                 `json:"usd_status"`
	TokenStatus        string                 `json:"token_status"`
	Created            string                 `json:"created"`
	CreatedTime        time.Time              `json:"created_epoch"`
	LastUpdate         string                 `json:"last_update"`
	LastUpdateTime     time.Time              `json:"last_update_epoch"`
	Descriptor         string                 `json:"descriptor"`
	DescriptorAch      string                 `json:"descriptor_ach"`
	AchName            string                 `json:"ach_name"`
	ProcessingType     string                 `json:"processing_type"`
	DestinationAddress string                 `json:"destination_address"`
	DestinationHandle  string                 `json:"destination_handle"`
	HandleAddress      string                 `json:"handle_address"`
	Timeline           []TransactionTimePoint `json:"timeline"`
}

func (ch *Transaction) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "user_handle":
			ch.UserHandle = value.(string)
		case "reference_id":
			ch.ReferenceId = value.(string)
		case "transaction_id":
			ch.TransactionId = value.(string)
		case "transaction_hash":
			ch.TransactionHash = value.(string)
		case "transaction_type":
			ch.TransactionType = value.(string)
		case "sila_amount":
			ch.SilaAmount = value.(int64)
		case "bank_account_name":
			ch.BankAccountName = value.(string)
		case "status":
			ch.Status = value.(string)
		case "usd_status":
			ch.UsdStatus = value.(string)
		case "token_status":
			ch.TokenStatus = value.(string)
		case "created":
			ch.Created = value.(string)
		case "created_epoch":
			ch.CreatedTime = time.Unix(int64(value.(float64)), 0)
		case "last_update":
			ch.LastUpdate = value.(string)
		case "last_update_epoch":
			ch.LastUpdateTime = time.Unix(int64(value.(float64)), 0)
		case "descriptor":
			ch.Descriptor = value.(string)
		case "descriptor_ach":
			ch.DescriptorAch = value.(string)
		case "ach_name":
			ch.AchName = value.(string)
		case "processing_type":
			ch.ProcessingType = value.(string)
		case "destination_address":
			ch.DestinationAddress = value.(string)
		case "destination_handle":
			ch.DestinationHandle = value.(string)
		case "handle_address":
			ch.HandleAddress = value.(string)
		case "timeline":
			ch.Timeline = value.([]TransactionTimePoint)
		}

	}
	return nil
}

type TransactionTimePoint struct {
	Date        string    `json:"date"`
	DateTime    time.Time `json:"date_epoch"`
	Status      string    `json:"status"`
	UsdStatus   string    `json:"usd_status"`
	TokenStatus string    `json:"token_status"`
}

func (ch *TransactionTimePoint) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "date":
			ch.Date = value.(string)
		case "date_epoch":
			ch.DateTime = time.Unix(int64(value.(float64)), 0)
		case "status":
			ch.Status = value.(string)
		case "usd_status":
			ch.UsdStatus = value.(string)
		case "token_status":
			ch.TokenStatus = value.(string)
		}
	}
	return nil
}

// The wallet key passed should be registered to the user which initiated the transaction to cancel
func (msg *GetTransactions) Do(userWalletPrivateKey string) (GetTransactionsResponse, error) {
	var responseBody GetTransactionsResponse
	err := instance.performCallWithUserAuth("/get_transactions", msg, &responseBody, userWalletPrivateKey)
	return responseBody, err
}
