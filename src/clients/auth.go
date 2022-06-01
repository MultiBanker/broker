package clients

import (
	"context"
	"fmt"
)

var ErrPartnerUnknown = fmt.Errorf("[ERROR] Err partner is unknown")

type PartnerAuth interface {
	Auth(ctx context.Context) (string, error)
}

func PartnerDetect(_ string) (PartnerAuth, error) {
	return nil, ErrPartnerUnknown
}

type Halyk struct {}

type Eurasian struct {}

type Jysan struct {}
