package auth

import "strings"

func NormName(s string) string {
	return strings.Title(strings.ToLower(s))
}
