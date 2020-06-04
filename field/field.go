package field

import (
	"errors"
	"strings"
)

// Validator describes the validator providers interfaces.
//
// This interface can have many provider according to properties types.
//
// After calling the Validate function the provider should validate
// the property and return any errors.
//
// An example of a Validator provider could be the string validator
// which is responsible to validate a string property.
type validator interface {
	// validate validates the current property and returns any errors.
	validate() ([]string, error)
}

// Field struct describes the properties of validation annotation.
//
// For example, if you have the following struct:
//
// struct User {
//   Username string `validate:"required=true,min=10,max=23"`
// }
// u := &User{"cpapidas"}
//
// The properties in the Field struct will have the followings values:
//
// Kind = "string"
// Value = "cpapidas"
// FieldName = "Username"
// Rules = {"required": "true", "min": "10", "max": "23"}
type Field struct {
	// Kind is the field's type e.g. Kind="string"
	Kind string

	// Value is the actual value of the field
	Value interface{}

	// FieldName is the name of the field e.g. for the property
	// struct {Username string `validate:"required=true,min=10,max=23"`}
	// the FieldName = "Username"
	FieldName string

	// Rules describes the validation rules.
	// e.g. for the annotation `validate:"required=true,min=10,max=23"`
	// Rules = {"required": "true", "min": "10", "max": "23"}
	Rules map[string]string

	// Err property contains the annotation's error.
	// For example for the annotation: `validate:"required=true,err=password is required"`
	// the Err property will have the value Field.Err = "password is required"
	Err string

	// Errs contains the Field's errors after the validation.
	Errs []string

	// CustomError property contains the custom error for this field.
	// By setting this property all the default Errs and Err will be overridden
	// This property can be set:
	// errMess := map[string]string{
	// 		"Username": "Username is required and should contain between 10 and 23 characters.",
	// }
	// errs := valy.Validate(u, errMess)
	CustomError string
}

// callValidator is responsible to identify which validator to call according to Field's kind field.
// For example if the field is a string then we want to call the Validators.str.
// If the type of the struct property is not supported then we will return an error.
func (fp *Field) CallValidator(validations []string) ([]string, error) {
	fp.applyRules(validations)
	var errs []string
	var v validator
	var err error
	if fp.Kind == "string" {
		v = newString(fp)
	} else if isNumeric(fp.Kind) {
		v = newNumeric(fp)
	} else {
		return nil, errors.New("Cannot support " + fp.Kind + " field type")
	}
	validateErrs, err := v.validate()
	if err != nil {
		return nil, err
	}
	if fp.CustomError != "" {
		return append(errs, fp.CustomError), nil
	}
	return append(errs, validateErrs...), nil
}

// applyRules is responsible to apply the annotation rules to Rule property.
// Each rule is described as a map[string]string property.
// For example the rule max=23 from `validate:"required=true,min=10,max=23"`
// will have the value Rule = {"max": "23"}
func (fp *Field) applyRules(validations []string) {
	var rules = make(map[string]string)
	for _, v := range validations {
		// Split the rule in order to get the key and the Value (e.g max=32 max->key 32->Value)
		f := strings.Split(v, "=")
		if f[0] == "Err" {
			fp.Err = f[1]
		} else {
			rules[f[0]] = f[1]
		}

	}
	fp.Rules = rules
}

// isNumeric it checks if a field is numeric type in order to run the defined validator.
func isNumeric(s string) bool {
	numericType := []string{"int", "int8", "int16", "int32", "int64", "float", "float32", "float64", "uint", "uint",
		"uint8", "uint16", "uint32", "uint64"}
	for _, a := range numericType {
		if a == s {
			return true
		}
	}
	return false
}
