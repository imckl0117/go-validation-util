package validation

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	strErrorCode = "ERROR_STRING_%v"
)

const (
	strLenErrorCode              = "LENGTH_%v"
	strLenMinErrorCode           = "LENGTH_MIN_%v"
	strLenMaxErrorCode           = "LENGTH_MAX_%v"
	strLenBetweenErrorCode       = "LENGTH_BETWEEN_%v_%v"
	strOnlyASCIIErrorCode        = "ONLY_ASCII"
	strOnlyAlphanumericErrorCode = "ONLY_ALPHANUMERIC"
	strOnlyNumericErrorCode      = "ONLY_NUMERIC"
	strInErrorCode               = "IN"
	strNoDuplicateErrorCode      = "NO_DUPLICATE"
)

const (
	strLenErrorMessage              = "length of %v is not %v"
	strLenMinErrorMessage           = "length of %v is smaller than %v"
	strLenMaxErrorMessage           = "length of %v is greater than %v"
	strLenBetweenErrorMessage       = "length of %v is not between %v and %v"
	strOnlyASCIIErrorMessage        = "%v contains non-ASCII character(s)"
	strOnlyAlphanumericErrorMessage = "%v contains non-alphanumeric character(s)"
	strOnlyNumericErrorMessage      = "%v contains non-numeric character(s)"
	strInErrorMessage               = "%v has no match in %v"
	strNoDuplicateErrorMessage      = "%v has duplicated values"
)

// StringLen returns error if len(value)!=length, otherwise nil.
func StringLen(field, value string, length int) *ErrValidation {
	if len(value) != length {
		code := fmt.Sprintf(strErrorCode, fmt.Sprintf(strLenErrorCode, length))
		message := fmt.Sprintf(strLenErrorMessage, field, length)

		return NewError(code, message, field, value)
	}

	return nil
}

// StringLenMin returns error if len(value)<min, otherwise nil.
func StringLenMin(field, value string, min int) *ErrValidation {
	if len(value) < min {
		code := fmt.Sprintf(strErrorCode, fmt.Sprintf(strLenMinErrorCode, min))
		message := fmt.Sprintf(strLenMinErrorMessage, field, min)

		return NewError(code, message, field, value)
	}

	return nil
}

// StringLenMax returns error if len(value)>max, otherwise nil.
func StringLenMax(field, value string, max int) *ErrValidation {
	if len(value) > max {
		code := fmt.Sprintf(strErrorCode, fmt.Sprintf(strLenMaxErrorCode, max))
		message := fmt.Sprintf(strLenMaxErrorMessage, field, max)

		return NewError(code, message, field, value)
	}

	return nil
}

// StringLenBetween returns error if len(value)<min or len(value)>max,
// otherwise nil.
func StringLenBetween(field, value string, min, max int) *ErrValidation {
	if len(value) < min || len(value) > max {
		code := fmt.Sprintf(strErrorCode, fmt.Sprintf(strLenBetweenErrorCode, min, max))
		message := fmt.Sprintf(strLenBetweenErrorMessage, field, min, max)

		return NewError(code, message, field, value)
	}

	return nil
}

// StringOnlyASCII returns error if value contains non-ASCII characters, such
// as non-ASCII unicode characters, otherwise nil.
func StringOnlyASCII(field, value string) *ErrValidation {
	for _, c := range value {
		if c > unicode.MaxASCII {
			code := fmt.Sprintf(strErrorCode, strOnlyASCIIErrorCode)
			message := fmt.Sprintf(strOnlyASCIIErrorMessage, field)

			return NewError(code, message, field, value)
		}
	}

	return nil
}

// StringOnlyAlphanumeric returns error if value contains non-alphanumeric
// characters, such as special symbols, otherwise nil.
func StringOnlyAlphanumeric(field, value string) *ErrValidation {
	for _, c := range value {
		if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
			code := fmt.Sprintf(strErrorCode, strOnlyAlphanumericErrorCode)
			message := fmt.Sprintf(strOnlyAlphanumericErrorMessage, field)

			return NewError(code, message, field, value)
		}
	}

	return nil
}

// StringOnlyNumeric returns error if value contains non-numeric characters.
func StringOnlyNumeric(field, value string) *ErrValidation {
	for _, c := range value {
		if !unicode.IsDigit(c) {
			code := fmt.Sprintf(strErrorCode, strOnlyNumericErrorCode)
			message := fmt.Sprintf(strOnlyNumericErrorMessage, field)

			return NewError(code, message, field, value)
		}
	}

	return nil
}

// StringIn returns error if value has no match in values, otherwise nil.
// Comparison is done case-sensitively.
func StringIn(field, value string, values []string) *ErrValidation {
	for _, v := range values {
		if v == value {
			return nil
		}
	}

	code := fmt.Sprintf(strErrorCode, strInErrorCode)
	message := fmt.Sprintf(strInErrorMessage, field, values)

	return NewError(code, message, field, value)
}

// StringInIgnoreCase returns error if value has no match in values, otherwise
// nil. Comparison is done case-insensitively.
func StringInIgnoreCase(field, value string, values []string) *ErrValidation {
	for _, v := range values {
		if strings.ToLower(v) == strings.ToLower(value) {
			return nil
		}
	}

	code := fmt.Sprintf(strErrorCode, strInErrorCode)
	message := fmt.Sprintf(strInErrorMessage, field, values)

	return NewError(code, message, field, value)
}

// StringNoDuplicate returns error if values contain duplicated value,
// otherwise nil. Comparison is done case-sensitively.
func StringNoDuplicate(field string, values []string) *ErrValidation {
	m := make(map[string]struct{})

	for _, v := range values {
		if _, ok := m[v]; ok {
			code := fmt.Sprintf(strErrorCode, strNoDuplicateErrorCode)
			message := fmt.Sprintf(strNoDuplicateErrorMessage, field)

			return NewError(code, message, field, nil)
		}

		m[v] = struct{}{}
	}

	return nil
}

// StringNoDuplicateIgnoreCase returns error if values contain duplicated
// value, otherwise nil. Comparison is done case-insensitively.
func StringNoDuplicateIgnoreCase(field string, values []string) *ErrValidation {
	m := make(map[string]struct{})

	for _, v := range values {
		w := strings.ToLower(v)

		if _, ok := m[w]; ok {
			code := fmt.Sprintf(strErrorCode, strNoDuplicateErrorCode)
			message := fmt.Sprintf(strNoDuplicateErrorMessage, field)

			return NewError(code, message, field, nil)
		}

		m[w] = struct{}{}
	}

	return nil
}
