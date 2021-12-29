package validation

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	numErrorCode = "ERROR_NUMBER_%v"
)

const (
	numMinErrorCode     = "MIN"
	numMaxErrorCode     = "MAX"
	numBetweenErrorCode = "BETWEEN"
	numFormatErrorCode  = "FORMAT"
)

const (
	numMinErrorMessage     = "%v is smaller than %v"
	numMaxErrorMessage     = "%v is greater than %v"
	numBetweenErrorMessage = "%v is not between %v and %v"
	numFormatErrorMessage  = "%v does not conform with the format %v"
)

// NumberMin returns error if value<min. NumberMin panics if value and min have
// different types.
func NumberMin(field string, value, min interface{}) *ErrValidation {
	code := fmt.Sprintf(numErrorCode, numMinErrorCode)
	message := fmt.Sprintf(numMinErrorMessage, field, min)

	switch value.(type) {
	case int:
		v1 := value.(int)

		v2, ok := min.(int)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min int
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	case int8:
		v1 := value.(int8)

		v2, ok := min.(int8)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min int8
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	case int16:
		v1 := value.(int16)

		v2, ok := min.(int16)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min int16
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	case int32:
		v1 := value.(int32)

		v2, ok := min.(int32)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min int32
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	case int64:
		v1 := value.(int64)

		v2, ok := min.(int64)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min int64
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	case uint:
		v1 := value.(uint)

		v2, ok := min.(uint)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min uint
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	case uint8:
		v1 := value.(uint8)

		v2, ok := min.(uint8)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min uint8
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	case uint16:
		v1 := value.(uint16)

		v2, ok := min.(uint16)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min uint16
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	case uint32:
		v1 := value.(uint32)

		v2, ok := min.(uint32)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min uint32
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	case uint64:
		v1 := value.(uint64)

		v2, ok := min.(uint64)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min uint64
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	case float32:
		v1 := value.(float32)

		v2, ok := min.(float32)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min float32
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	case float64:
		v1 := value.(float64)

		v2, ok := min.(float64)

		if !ok {
			panic("value and min must have the same type")
		}

		args := struct {
			Min float64
		}{
			v2,
		}

		if v1 < v2 {
			return NewError(code, args, message, field, value)
		}
	default:
		panic("value must be a number")
	}

	return nil
}

// NumberMax returns error if value>max. NumberMax panics if value and max have
// different types.
func NumberMax(field string, value, max interface{}) *ErrValidation {
	code := fmt.Sprintf(numErrorCode, numMaxErrorCode)
	message := fmt.Sprintf(numMaxErrorMessage, field, max)

	switch value.(type) {
	case int:
		v1 := value.(int)

		v2, ok := max.(int)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max int
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	case int8:
		v1 := value.(int8)

		v2, ok := max.(int8)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max int8
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	case int16:
		v1 := value.(int16)

		v2, ok := max.(int16)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max int16
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	case int32:
		v1 := value.(int32)

		v2, ok := max.(int32)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max int32
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	case int64:
		v1 := value.(int64)

		v2, ok := max.(int64)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max int64
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	case uint:
		v1 := value.(uint)

		v2, ok := max.(uint)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max uint
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	case uint8:
		v1 := value.(uint8)

		v2, ok := max.(uint8)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max uint8
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	case uint16:
		v1 := value.(uint16)

		v2, ok := max.(uint16)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max uint16
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	case uint32:
		v1 := value.(uint32)

		v2, ok := max.(uint32)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max uint32
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	case uint64:
		v1 := value.(uint64)

		v2, ok := max.(uint64)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max uint64
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	case float32:
		v1 := value.(float32)

		v2, ok := max.(float32)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max float32
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	case float64:
		v1 := value.(float64)

		v2, ok := max.(float64)

		if !ok {
			panic("value and max must have the same type")
		}

		args := struct {
			Max float64
		}{
			v2,
		}

		if v1 > v2 {
			return NewError(code, args, message, field, value)
		}
	default:
		panic("value must be a number")
	}

	return nil
}

// NumberBetween returns error if value<min or value>max. NumberBetween panics
// if value, min, and max have different types.
func NumberBetween(field string, value, min, max interface{}) *ErrValidation {
	code := fmt.Sprintf(numErrorCode, numBetweenErrorCode)
	message := fmt.Sprintf(numBetweenErrorMessage, field, min, max)

	switch value.(type) {
	case int:
		v1 := value.(int)

		v2, ok := min.(int)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(int)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max int
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	case int8:
		v1 := value.(int8)

		v2, ok := min.(int8)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(int8)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max int8
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	case int16:
		v1 := value.(int16)

		v2, ok := min.(int16)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(int16)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max int16
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	case int32:
		v1 := value.(int32)

		v2, ok := min.(int32)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(int32)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max int32
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	case int64:
		v1 := value.(int64)

		v2, ok := min.(int64)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(int64)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max int64
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	case uint:
		v1 := value.(uint)

		v2, ok := min.(uint)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(uint)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max uint
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	case uint8:
		v1 := value.(uint8)

		v2, ok := min.(uint8)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(uint8)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max uint8
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	case uint16:
		v1 := value.(uint16)

		v2, ok := min.(uint16)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(uint16)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max uint16
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	case uint32:
		v1 := value.(uint32)

		v2, ok := min.(uint32)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(uint32)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max uint32
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	case uint64:
		v1 := value.(uint64)

		v2, ok := min.(uint64)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(uint64)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max uint64
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	case float32:
		v1 := value.(float32)

		v2, ok := min.(float32)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(float32)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max float32
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	case float64:
		v1 := value.(float64)

		v2, ok := min.(float64)

		if !ok {
			panic("value, min and max must have the same type")
		}

		v3, ok := max.(float64)

		if !ok {
			panic("value, min and max must have the same type")
		}

		args := struct {
			Min, Max float64
		}{
			v2, v3,
		}

		if v1 < v2 || v1 > v3 {
			return NewError(code, args, message, field, value)
		}
	default:
		panic("value must be a number")
	}

	return nil
}

// NumberFormat checks if value has l decimal places, where m<=l<=n, m,n from
// format.
func NumberFormat(field string, value float64, format string) *ErrValidation {
	args := struct {
		Format string
	}{
		format,
	}
	code := fmt.Sprintf(numErrorCode, numFormatErrorCode)
	message := fmt.Sprintf(numFormatErrorMessage, field, format)

	f := strings.Split(strings.TrimSpace(format), ",")

	if len(f) != 2 {
		panic("format must be in the form of m,n, where m and n are positive integers")
	}

	m, err := strconv.ParseInt(f[0], 10, 0)

	if err != nil || m < 0 {
		panic("format must be in the form of m,n, where m and n are positive integers")
	}

	n, err := strconv.ParseInt(f[1], 10, 0)

	if err != nil || n < 0 {
		panic("format must be in the form of m,n, where m and n are positive integers")
	}

	v := strconv.FormatFloat(value, 'f', -1, 64)

	vs := strings.Split(v, ".")

	// eg. value = 1.000
	if len(vs) == 1 {
		if m == 0 {
			return nil
		}

		return NewError(code, args, message, field, value)
	}

	l := int64(len(vs[1]))

	if l < m || l > n {
		return NewError(code, args, message, field, value)
	}

	return nil
}
