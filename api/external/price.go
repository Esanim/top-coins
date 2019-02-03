package apiext

// Status is the status structure
type Status struct {
	Timestamp    string  `json:"timestamp"`
	ErrorCode    int     `json:"error_code"`
	ErrorMessage *string `json:"error_message"`
}

// Response is the response structure
type Response struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

// PriceListing is the listing structure
type PriceListing struct {
	ID     float64           `json:"id"`
	Name   string            `json:"name"`
	Symbol string            `json:"symbol"`
	Quote  map[string]*Quote `json:"quote"`
}

// Quote is the quote structure
type Quote struct {
	Price float64 `json:"price"`
}
