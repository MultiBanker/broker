package config

type AirbaPay struct {
	URL   string `long:"airba-url" env:"AIRBAPAY_URL" description:"AIRBA PAY URL" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
	Token string `long:"airba-token" env:"AIRBAPAY_TOKEN" description:"AIRBA PAY TOKEN" required:"false" default:"mongo"`
}

type Altyn struct {
	URL   string `long:"altyn-url" env:"ALTYNBANK_URL" description:"ALTYNBANK URL" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
	Token string `long:"altyn-token" env:"ALTYNBANK_TOKEN" description:"ALTYNBANK URL" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
}
type Jysan struct {
	URL   string `long:"jysan-url" env:"JYSANBANK_URL" description:"JYSANBANK URL" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
	Token string `long:"jysan-token" env:"JYSANBANK_TOKEN" description:"JYSANBANK URL" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
}
type Freedom struct {
	URL   string `long:"freedom-finance-url" env:"FREEDOMFINANCE_URL" description:"FREEDOM FINANCE URL" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
	Token string `long:"freedom-finance-token" env:"FREEDOMFINANCE_TOKEN" description:"FREEDOM FINANCE Token" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
}
type Halyk struct {
	URL   string `long:"halyk-bank-url" env:"HALYKBANK_URL" description:"HALYKBANK URL" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
	Token string `long:"halyk-bank-token" env:"HALYKBANK_TOKEN" description:"HALYKBANK TOKEN" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
}
type RBK struct {
	URL   string `long:"rbk-url" env:"RBK_URL" description:"RBK URL" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
	Token string `long:"rbk-token" env:"RBK_TOKEN" description:"RBK TOKEN" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
}
type Sber struct {
	URL   string `long:"sberbank-url" env:"SBERBANK_URL" description:"SBERBANK URL" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
	Token string `long:"sberbank-token" env:"SBERBANK_TOKEN" description:"SBERBANK TOKEN" required:"false" default:"https://sfinapi.technodom.kz/loan-app/api/v1/broker/backend"`
}
