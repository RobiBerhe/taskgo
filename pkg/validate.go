package pkg

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// validates and returns errors if any (works based on validator lib)

func Validate(s interface{}) []string {
	validate := validator.New()

	var errMsgs []string
	if err := validate.Struct(s); err != nil {
		valErrs := err.(validator.ValidationErrors)
		for _, fieldErr := range valErrs {
			errMsgs = append(errMsgs, fmt.Sprintf("Error in field '%s' in %s:%s", fieldErr.Field(), fieldErr.Namespace(), fieldErr.ActualTag()))
		}
	}
	return errMsgs
}

func IsStructEmpty(obj interface{}) bool {
	// Reflect the object to analyze its fields
	v := reflect.ValueOf(obj)

	// Iterate over all the fields in the struct
	for i := 0; i < v.NumField(); i++ {
		// If any field is non-zero, return false
		if !v.Field(i).IsZero() {
			return false
		}
	}
	// All fields are zero
	return true
}
