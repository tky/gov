package rules

import (
	"errors"
	"reflect"
)

var ErrValidation = errors.New("Validation Error")
var ErrIllegalParamsOnRequired = errors.New("Required should not have params")

// Validate validate target using param based on tags.
type Validate func(v interface{}, values []string) error

func isnil(x interface{}) bool {
	return (x == nil) || reflect.ValueOf(x).IsNil()
}

// Required required v return error if v is nil.
func Required(v interface{}, values []string) error {
	if len(values) != 0 {
		return ErrIllegalParamsOnRequired
	}
	if v == nil {
		return ErrValidation
	}
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		if isnil(v) {
			return ErrValidation
		}
	}

	switch v.(type) {
	case string:
		if len(v.(string)) == 0 {
			return ErrValidation
		}
	case *string:
		if len(*v.(*string)) == 0 {
			return ErrValidation
		}
	default:
		return nil
	}
	return nil
}
