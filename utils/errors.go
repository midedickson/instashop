package utils

import "errors"

// generic errors
var ErrUserAlreadyExists = errors.New("user already exists")
var ErrResourceNotFound = errors.New("resource not found")
var ErrForbidden = errors.New("forbidden resource")
var ErrUnauthorized = errors.New("unauthorized for this action")

// order status errors
var ErrOrderAlreadyInStatus = errors.New("order already in status")
var ErrInvalidOrderTransition = errors.New("order transition is not valid")
var ErrInvalidOrderStatus = errors.New("invalid order status")
