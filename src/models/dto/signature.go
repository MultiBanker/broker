package dto

type CreateSignatureReq struct {
	IIN             string   `json:"iin"`
	Phone           string   `json:"phone"`
	SpecCode        string   `json:"code"`
	AdditionalCodes []string `json:"additional_codes"`
}

type CheckVerificationReq struct {
	SignID string `json:"-"`
	Token  string `json:"token"`
}
