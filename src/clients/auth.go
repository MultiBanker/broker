package clients

import (
	"context"
	"fmt"
)

var ErrPartnerUnknown = fmt.Errorf("[ERROR] Err partner is unknown")

type PartnerAuth interface {
	Auth(ctx context.Context) (string, error)
}

func PartnerDetect(partnerCode string) (PartnerAuth, error) {
	switch partnerCode {
	case "airba_pay":
		return Airba{}, nil
	}

	return nil, ErrPartnerUnknown
}

type Airba struct {
}

func (a Airba) Auth(_ context.Context) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzeXN0ZW1fY29kZSI6IlREX0JST0tFUiJ9.Fa4wL9KID3A_-8fYmvhZKXi68K5GRMlLsYK0y6PASI4", nil
}

type Halyk struct {
}

type FreedomFinance struct {
}

type Alfa struct {
}
