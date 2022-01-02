package currency

import "errors"

var (
	InvalidAlphabeticCode = errors.New("invalid alphabetic currency code")
	UnknownAlphabeticCode = errors.New("unknown alphabetic currency code")
	InvalidNumericCode    = errors.New("invalid numeric currency code")
	UnknownNumericCode    = errors.New("unknown numeric currency code")
)
