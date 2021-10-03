package clients

type OrderResponse struct {
	Status      string `json:"status"`
	Code        string `json:"code"`
	RedirectURL string `json:"redirectUrl"`
	RequestUUID string `json:"requestUuid"`
	Message     string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type OrderUpdate struct {
	ReferenceID string `json:"referenceId"`
	State       string `json:"state"`
	StateTitle  string `json:"stateTitle"`
	ProductCode string `json:"productCode"`
	LoanLength  string `json:"loanLength"`
	Reason      string `json:"reason"`
}
