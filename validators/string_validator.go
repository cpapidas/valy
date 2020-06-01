package valy

import (
	"strconv"
)

// stringValidator describe the string validator provider.
type stringValidator struct {
	// validator embedded validator in order to have access to fieldProperties and customErrors function.
	validator

	// min defines the min character that can contain the string.
	min int

	// max defines the max character that can contain the string.
	max int

	// require defines if the field has to be set.
	required bool

	// stringValue is the value of the field.
	stringValue string
}

// newStringValidator initializes and returns a stringValidator.
func newStringValidator(fp *fieldProperties) *stringValidator {
	nv := &stringValidator{
		min:      -1,
		max:      -1,
		required: false,
	}
	nv.validator.fieldProperties = *fp
	return nv
}

// validate validate the field. The errors have been set in fieldProperties.errs fields.
func (n *stringValidator) validate() {
	v := n.value.(string)
	n.stringValue = v
	n.setRules(n.rules)
	n.validateString()
}

// setRules sets the rules for current field.
func (n *stringValidator) setRules(rules map[string]string) {
	var err error
	for k, v := range rules {
		switch k {
		case "min":
			n.min, err = strconv.Atoi(v)
		case "max":
			n.max, err = strconv.Atoi(v)
		case "required":
			n.required, err = strconv.ParseBool(v)
		}
		if err != nil {
			panic(err)
		}
	}
}

// validateString calls the validate functions for each rule.
func (n *stringValidator) validateString() {
	if n.min > -1 {
		n.minRule()
	}
	if n.max > -1 {
		n.maxRule()
	}
	if n.required {
		n.requiredRule()
	}
}

// minRule checks if field contains less than X characters.
func (n *stringValidator) minRule() {
	if len(n.stringValue) < n.min {
		n.errs = append(n.errs, "the field "+n.fieldName+" should contains at least "+
			strconv.Itoa(n.min)+" characters")
	}
}

// maxRule checks if field contains more than X characters.
func (n *stringValidator) maxRule() {
	if len(n.stringValue) > n.max {
		n.errs = append(n.errs, "the field "+n.fieldName+" should contains max "+
			strconv.Itoa(n.max)+" characters")
	}
}

// requiredRule check if field is defined.
func (n *stringValidator) requiredRule() {
	if n.stringValue == "" {
		n.errs = append(n.errs, "the field "+n.fieldName+" should not be empty")
	}
}
