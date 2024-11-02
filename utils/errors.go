package utils

import "errors"

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrForbidden = errors.New("forbidden resource")
