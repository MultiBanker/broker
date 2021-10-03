package dto

type SignUpPartnerRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	StoreName  string `json:"store_name"`
	Phone      string `json:"phone" validate:"required,is_phone"`
	Email      string `json:"email" validate:"required,email"`
	BIN        string `json:"bin,omitempty" validate:"omitempty"`
	Commission int    `json:"commission"`
	LogoURL    string `json:"logo_url" bson:"logo_url"`
}

type Login struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
}
