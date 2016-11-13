package rules

import (
	"errors"
	"reflect"
	"strconv"
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

// MinLength return error if v length is under specified parameter.
// If v is nil or nil pointer, return nil (no error) because it is responsibility of 'Required'
func MinLength(v interface{}, values []string) error {
	if len(values) != 1 {
		return errors.New("Min requires only 1 value")
	}

	var p int
	var err error
	p, err = strconv.Atoi(values[0])
	if err != nil {
		return err
	}

	if v == nil {
		return nil
	}

	if reflect.TypeOf(v).Kind() == reflect.Ptr && isnil(v) {
		return nil
	}

	switch v.(type) {
	case string:
		if len(v.(string)) < p {
			return ErrValidation
		}
	case *string:
		if len(*v.(*string)) < p {
			return ErrValidation
		}
	default:
		return errors.New("Target field of min-length requires string or string pointer")
	}
	return nil
}
