package data

//Shipment defines the structure for API
type Shipment struct {
	ID                   string  `json:"ID"`
	SenderName           string  `json:"SenderName" validate:"required|maxLen:30|regex:^[^0-9]*$"`
	SenderEmail          string  `json:"SenderEmail" validate:"required|email"`
	SenderAddress        string  `json:"SenderAddress" validate:"required|maxLen:100"`
	SenderCountryCode    string  `json:"SenderCountryCode" validate:"required|len:2|regex:^[A-Z]*$"`
	RecipientName        string  `json:"RecipientName" validate:"required|maxLen:30|regex:^[^0-9]*$"`
	RecipientEmail       string  `json:"RecipientEmail" validate:"required|email"`
	RecipientAddress     string  `json:"RecipientAddress" validate:"required|maxLen:100"`
	RecipientCountryCode string  `json:"RecipientCountryCode" validate:"required|len:2|regex:^[A-Z]*$"`
	Weight               float64 `json:"Weight" validate:"required|gt:0|max:1000"`
	Price                string  `validate:"safe"`
}
