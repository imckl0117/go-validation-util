/*
Package validation exposes some utility functions to validate strings and
numbers. Validation error messages are formatted in the form of
'fieldName,errorCode'. errorCode may contain additional information, such as
min/max length, etc.
*/
package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	numberMinError     = "%v,ERROR_NUMBER_MIN_%v"
	numberMaxError     = "%v,ERROR_NUMBER_MAX_%v"
	numberBetweenError = "%v,ERROR_NUMBER_BETWEEN_%v_%v"
	numberFormatError  = "%v,ERROR_NUMBER_FORMAT_%v"
)

func NumberMin(field string, value, min interface{}) error {
	switch value.(type) {
	case int:
		v1 := value.(int)

		v2, ok := min.(int)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	case int8:
		v1 := value.(int8)

		v2, ok := min.(int8)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	case int16:
		v1 := value.(int16)

		v2, ok := min.(int16)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	case int32:
		v1 := value.(int32)

		v2, ok := min.(int32)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	case int64:
		v1 := value.(int64)

		v2, ok := min.(int64)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	case uint:
		v1 := value.(uint)

		v2, ok := min.(uint)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	case uint8:
		v1 := value.(uint8)

		v2, ok := min.(uint8)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	case uint16:
		v1 := value.(uint16)

		v2, ok := min.(uint16)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	case uint32:
		v1 := value.(uint32)

		v2, ok := min.(uint32)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	case uint64:
		v1 := value.(uint64)

		v2, ok := min.(uint64)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	case float32:
		v1 := value.(float32)

		v2, ok := min.(float32)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	case float64:
		v1 := value.(float64)

		v2, ok := min.(float64)

		if !ok {
			panic("value and min must have the same type")
		}

		if v1 < v2 {
			return errors.New(fmt.Sprintf(numberMinError, field, min))
		}
	default:
		panic("value must be a number")
	}

	return nil
}

func NumberMax(field string, value, max interface{}) error {
	switch value.(type) {
	case int:
		v1 := value.(int)

		v2, ok := max.(int)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	case int8:
		v1 := value.(int8)

		v2, ok := max.(int8)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	case int16:
		v1 := value.(int16)

		v2, ok := max.(int16)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	case int32:
		v1 := value.(int32)

		v2, ok := max.(int32)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	case int64:
		v1 := value.(int64)

		v2, ok := max.(int64)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	case uint:
		v1 := value.(uint)

		v2, ok := max.(uint)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	case uint8:
		v1 := value.(uint8)

		v2, ok := max.(uint8)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	case uint16:
		v1 := value.(uint16)

		v2, ok := max.(uint16)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	case uint32:
		v1 := value.(uint32)

		v2, ok := max.(uint32)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	case uint64:
		v1 := value.(uint64)

		v2, ok := max.(uint64)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	case float32:
		v1 := value.(float32)

		v2, ok := max.(float32)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	case float64:
		v1 := value.(float64)

		v2, ok := max.(float64)

		if !ok {
			panic("value and max must have the same type")
		}

		if v1 > v2 {
			return errors.New(fmt.Sprintf(numberMaxError, field, max))
		}
	default:
		panic("value must be a number")
	}

	return nil
}

func NumberBetween(field string, value, min, max interface{}) error {
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
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

		if v1 < v2 || v1 > v3 {
			return errors.New(fmt.Sprintf(numberBetweenError, field, min, max))
		}
	default:
		panic("value must be a number")
	}

	return nil
}

func NumberFormat(field string, value float64, format string) error {
	f := strings.Split(strings.TrimSpace(format), ",")

	if len(f) != 2 {
		panic("format must be in the form of m,n, where m and n are positive integers")
	}

	ndp, err := strconv.ParseInt(f[1], 10, 0)

	if err != nil || ndp < 0 {
		panic("format must be in the form of m,n, where m and n are positive integers")
	}

	v := strings.Split(strconv.FormatFloat(value, 'f', -1, 64), ".")

	// eg. value = 1.000
	if len(v) == 1 {
		return nil
	}

	// eg. value = 1.234
	if int64(len(v[1])) > ndp {
		return errors.New(fmt.Sprintf(numberFormatError, field, format))
	}

	return nil
}