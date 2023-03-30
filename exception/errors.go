package exception

import "errors"

var ErrorPasswordNotMatch = errors.New("password did not match")
var ErrorUnauthorized = errors.New("unauthorized")
