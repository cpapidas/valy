package valy

import (
	"strings"
)

// fieldProperties describes a field of the given struct object.
type fieldProperties struct {
	kind         string
	value        interface{}
	fieldName    string
	rules        map[string]string
	err          string
	errs         []string
	customErrors map[string]string
}

// applyRules applies the defined rules for the current field. It gets the validation rules as array string.
func (fp *fieldProperties) applyRules(validations []string) {
	var rules = make(map[string]string)
	for _, v := range validations {
		// Split the rule in order to get the key and the value (e.g max=32 max->key 32->value)
		f := strings.Split(v, "=")
		if f[0] == "err" {
			fp.err = f[1]
		} else {
			rules[f[0]] = f[1]
		}

	}
	fp.rules = rules
}

// callValidator calls the validator for the current field. It returns the errors as array of string.
func (fp *fieldProperties) callValidator() []string {
	var errs []string
	var v ivalidator
	if fp.kind == "string" {
		v = newStringValidator(fp)
		v.validate()
		errs = append(errs, v.customErrors()...)
	} else if isNumeric(fp.kind) {
		v = newNumericValidator(fp)
		v.validate()
		errs = append(errs, v.customErrors()...)
	}
	if len(fp.customErrors) != 0 && fp.customErrors[fp.fieldName] != "" && len(errs) != 0 {
		return []string{fp.customErrors[fp.fieldName]}
	}
	return errs
}

// isNumeric it checks if a field is numeric type in order to run the defined validator.
func isNumeric(s string) bool {
	numericType := []string{"int","int8",  "int16", "int32", "int64", "float", "float32", "float64", "uint", "uint",
		"uint8", "uint16", "uint32", "uint64"}
	for _, a := range numericType {
		if a == s {
			return true
		}
	}
	return false
}