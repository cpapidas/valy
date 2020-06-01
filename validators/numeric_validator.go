package valy

import (
	"strconv"
)

// numericValidator describe the number validator provider.
type numericValidator struct {
	// validator embedded validator in order to have access to fieldProperties and customErrors function.
	validator

	// min defines the min value.
	min float64

	// max defines the max value.
	max float64

	// require defines if the field has to be set.
	required bool

	// numValue is the value of the field.
	numValue float64
}

// newNumericValidator initializes and returns a numericValidator.
func newNumericValidator(fp *fieldProperties) *numericValidator {
	nv := &numericValidator{
		min:      -1,
		max:      -1,
		required: false,
	}
	nv.validator.fieldProperties = *fp
	return nv
}

// validate validate the field. The errors have been set in fieldProperties.errs fields.
func (n *numericValidator) validate() {
	v := n.convertToFloat64(n.value)
	n.numValue = v
	n.setRules(n.rules)
	n.validateNumeric()
}

// convertToFloat64 converts the field value to fload64 in order to be able to apply the rules.
func (numericValidator) convertToFloat64(i interface{}) float64 {
	switch s := i.(type) {
	case float32:
		return float64(s)
	case int:
		return float64(s)
	case int64:
		return float64(s)
	case int32:
		return float64(s)
	case int16:
		return float64(s)
	case int8:
		return float64(s)
	case uint:
		return float64(s)
	case uint64:
		return float64(s)
	case uint32:
		return float64(s)
	case uint16:
		return float64(s)
	case uint8:
		return float64(s)
	default:
		return s.(float64)
	}
}

// setRules sets the rules for current field.
func (n *numericValidator) setRules(rules map[string]string) {
	var err error
	for k, v := range rules {
		switch k {
		case "min":
			n.min, err = strconv.ParseFloat(v, 64)
		case "max":
			n.max, err = strconv.ParseFloat(v, 64)
		case "required":
			n.required, err = strconv.ParseBool(v)
		}
		if err != nil {
			panic(err)
		}
	}
}

// validateNumeric calls the validate functions for each rule.
func (n *numericValidator) validateNumeric() {
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

// minRule checks if field is less than min value.
func (n *numericValidator) minRule() {
	if n.numValue < n.min {
		n.errs = append(n.errs, "the field "+n.fieldName+" should be grater than "+
			strconv.Itoa(int(n.min)))
	}
}

// maxRule checks if field is bigger than max value.
func (n *numericValidator) maxRule() {
	if n.numValue > n.max {
		n.errs = append(n.errs, "the field "+n.fieldName+" should be less than "+
			strconv.Itoa(int(n.max)))
	}
}

// requiredRule check if field is defined.
func (n *numericValidator) requiredRule() {
	if n.numValue == 0 {
		n.errs = append(n.errs, "the field "+n.fieldName+" should not be empty")
	}
}
