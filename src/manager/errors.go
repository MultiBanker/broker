package manager

import "fmt"

var ErrUnauthorized = fmt.Errorf("[ERROR] User unauthorized")
var ErrAuthorization = fmt.Errorf("[ERROR] Wrong Username and Password")
