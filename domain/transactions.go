package domain

import (
	"encoding/json"
	"time"
)

type GetTransactionsResponse struct {
	Success           bool                   `json:"success"`
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
	SilaAmount         float64                `json:"sila_amount"`
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
	ErrorCode          string                 `json:"error_code"`
	ErrorMsg           string                 `json:"error_msg"`
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "user_handle":
			t.UserHandle = value.(string)
		case "reference_id":
			t.ReferenceId = value.(string)
		case "transaction_id":
			t.TransactionId = value.(string)
		case "transaction_hash":
			t.TransactionHash = value.(string)
		case "transaction_type":
			t.TransactionType = value.(string)
		case "sila_amount":
			t.SilaAmount = value.(float64)
		case "bank_account_name":
			t.BankAccountName = value.(string)
		case "status":
			t.Status = value.(string)
		case "usd_status":
			t.UsdStatus = value.(string)
		case "token_status":
			t.TokenStatus = value.(string)
		case "created":
			t.Created = value.(string)
		case "created_epoch":
			t.CreatedTime = time.Unix(int64(value.(float64)), 0)
		case "last_update":
			t.LastUpdate = value.(string)
		case "last_update_epoch":
			t.LastUpdateTime = time.Unix(int64(value.(float64)), 0)
		case "descriptor":
			t.Descriptor = value.(string)
		case "descriptor_ach":
			t.DescriptorAch = value.(string)
		case "ach_name":
			t.AchName = value.(string)
		case "processing_type":
			t.ProcessingType = value.(string)
		case "destination_address":
			t.DestinationAddress = value.(string)
		case "destination_handle":
			t.DestinationHandle = value.(string)
		case "handle_address":
			t.HandleAddress = value.(string)
		case "timeline":
			arrValue := value.([]interface{})
			convertedValue := make([]TransactionTimePoint, len(arrValue))
			for index, rawValue := range arrValue {
				mapValue := rawValue.(map[string]interface{})
				var ttp TransactionTimePoint
				for key, value := range mapValue {
					switch key {
					case "date":
						ttp.Date = value.(string)
					case "date_epoch":
						ttp.DateTime = time.Unix(int64(value.(float64)), 0)
					case "status":
						ttp.Status = value.(string)
					case "usd_status":
						ttp.UsdStatus = value.(string)
					case "token_status":
						ttp.TokenStatus = value.(string)
					}
				}
				convertedValue[index] = ttp
			}
			t.Timeline = convertedValue
		case "error_code":
			t.ErrorCode = value.(string)
		case "error_msg":
			t.ErrorMsg = value.(string)
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

func (ttp *TransactionTimePoint) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for key, value := range raw {
		switch key {
		case "date":
			ttp.Date = value.(string)
		case "date_epoch":
			ttp.DateTime = time.Unix(int64(value.(float64)), 0)
		case "status":
			ttp.Status = value.(string)
		case "usd_status":
			ttp.UsdStatus = value.(string)
		case "token_status":
			ttp.TokenStatus = value.(string)
		}
	}
	return nil
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
	var startTimeEpoch int64
	if filters.StartTime.Unix() > 0 {
		startTimeEpoch = filters.StartTime.Unix()
	}
	var endTimeEpoch int64
	if filters.EndTime.Unix() > 0 {
		endTimeEpoch = filters.EndTime.Unix()
	}
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
		StartTime:        startTimeEpoch,
		EndTime:          endTimeEpoch,
		Page:             filters.Page,
		PerPage:          filters.PerPage,
		TransactionTypes: filters.TransactionTypes,
	}

	return json.Marshal(convertedSearchFilters)
}

type IssueSilaResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	TransactionId     string                 `json:"transaction_id"`
	Descriptor        string                 `json:"descriptor,omitempty"`
}

type TransferSilaResponse struct {
	Success            bool                   `json:"success"`
	Reference          string                 `json:"reference"`
	Message            string                 `json:"message"`
	Status             string                 `json:"status"`
	ValidationDetails  map[string]interface{} `json:"validation_details"`
	DestinationAddress string                 `json:"destination_address"`
	TransactionId      string                 `json:"transaction_id"`
	Descriptor         string                 `json:"descriptor,omitempty"`
}

type RedeemSilaResponse struct {
	Success           bool                   `json:"success"`
	Reference         string                 `json:"reference"`
	Message           string                 `json:"message"`
	Status            string                 `json:"status"`
	ValidationDetails map[string]interface{} `json:"validation_details"`
	TransactionId     string                 `json:"transaction_id"`
	Descriptor        string                 `json:"descriptor,omitempty"`
}
