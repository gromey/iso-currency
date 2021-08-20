package currency

import "errors"

var (
	ErrAlphaCode = errors.New("unknown alphabetic currency code")
	ErrNumCode   = errors.New("unknown numeric currency code")
)
