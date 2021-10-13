package middleware

import (
	"os"

	"github.com/rs/zerolog"
)

func logger() zerolog.Logger {
	return zerolog.New(os.Stdout)
}
