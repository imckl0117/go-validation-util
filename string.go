package validation

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

const (
	stringLenError              = "%v,ERROR_STRING_LENGTH_%v"
	stringLenMinError           = "%v,ERROR_STRING_LENGTH_MIN_%v"
	stringLenMaxError           = "%v,ERROR_STRING_LENGTH_MAX_%v"
	stringLenBetweenError       = "%v,ERROR_STRING_LENGTH_BETWEEN_%v_%v"
	stringOnlyASCIIError        = "%v,ERROR_STRING_ONLY_ASCII"
	stringOnlyAlphanumericError = "%v,ERROR_STRING_ONLY_ALPHANUMERIC"
	stringOnlyNumericError      = "%v,ERROR_STRING_ONLY_NUMERIC"
	stringInError               = "%v,ERROR_STRING_IN"
)

// StringLen returns error if len(value)!=length, otherwise nil.
func StringLen(field, value string, length int) error {
	if len(value) != length {
		return errors.New(fmt.Sprintf(stringLenError, field, length))
	}

	return nil
}

// StringLenMin returns error if len(value)<min, otherwise nil.
func StringLenMin(field, value string, min int) error {
	if len(value) < min {
		return errors.New(fmt.Sprintf(stringLenMinError, field, min))
	}

	return nil
}

// StringLenMax returns error if len(value)>max, otherwise nil.
func StringLenMax(field, value string, max int) error {
	if len(value) > max {
		return errors.New(fmt.Sprintf(stringLenMaxError, field, max))
	}

	return nil
}

// StringLenBetween returns error if len(value)<min or len(value)>max,
// otherwise nil.
func StringLenBetween(field, value string, min, max int) error {
	if len(value) < min || len(value) > max {
		return errors.New(fmt.Sprintf(stringLenBetweenError, field, min, max))
	}

	return nil
}

// StringOnlyASCII returns error if value contains non-ASCII characters, such
// as non-ASCII unicode characters, otherwise nil.
func StringOnlyASCII(field, value string) error {
	for _, c := range value {
		if c > unicode.MaxASCII {
			return errors.New(fmt.Sprintf(stringOnlyASCIIError, field))
		}
	}

	return nil
}

// StringOnlyAlphanumeric returns error if value contains non-alphanumeric
// characters, such as special symbols, otherwise nil.
func StringOnlyAlphanumeric(field, value string) error {
	for _, c := range value {
		if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
			return errors.New(fmt.Sprintf(stringOnlyAlphanumericError, field))
		}
	}

	return nil
}

// StringOnlyNumeric returns error if value contains non-numeric characters.
func StringOnlyNumeric(field, value string) error {
	for _, c := range value {
		if !unicode.IsDigit(c) {
			return errors.New(fmt.Sprintf(stringOnlyNumericError, field))
		}
	}

	return nil
}

// StringIn returns error if value has no match in values, otherwise nil.
// Comparison is done case-sensitively.
func StringIn(field, value string, values []string) error {
	for _, v := range values {
		if v == value {
			return nil
		}
	}

	return errors.New(fmt.Sprintf(stringInError, field))
}

// StringInIgnoreCase returns error if value has no match in values, otherwise
// nil. Comparison is done case-insensitively.
func StringInIgnoreCase(field, value string, values []string) error {
	for _, v := range values {
		if strings.ToLower(v) == strings.ToLower(value) {
			return nil
		}
	}

	return errors.New(fmt.Sprintf(stringInError, field))
}
