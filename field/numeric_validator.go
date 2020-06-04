package field

import (
	"strconv"
)

// Numeric struct describes the validator for numeric values. A numeric validator is
// responsible to define the rules of validation and validate a numeric property.
type numeric struct {
	// valy.Field embedded to Numeric validator to have access to Field's properties.
	Field

	// min defines the min value.
	min float64

	// max defines the max value.
	max float64

	// require defines if the field has to be set.
	required bool

	// value is the value of the field.
	value float64
}

// NewNumeric initializes and returns a Numeric.
func newNumeric(fp *Field) *numeric {
	nv := &numeric{
		min:      -1,
		max:      -1,
		required: false,
	}
	nv.Field = *fp
	return nv
}

// Validate is responsible to validate this field. After this call
// the function will return the errors if the field is invalid.
func (n *numeric) validate() ([]string, error) {
	v := n.convertToFloat64(n.Value)
	n.value = v
	err := n.setRules(n.Rules)
	if err != nil {
		return nil, err
	}
	if n.min > -1 {
		n.minRule()
	}
	if n.max > -1 {
		n.maxRule()
	}
	if n.required {
		n.requiredRule()
	}
	return n.Errs, nil
}

// convertToFloat64 converts the field value to fload64 in order to be able to apply the rules.
func (numeric) convertToFloat64(i interface{}) float64 {
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

// setRules sets the rules for the current field.
func (n *numeric) setRules(rules map[string]string) error {
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
			return err
		}
	}
	return nil
}

// minRule checks if field is less than the min value.
func (n *numeric) minRule() {
	if n.value < n.min {
		n.Errs = append(n.Errs, "the field "+n.FieldName+" should be grater than "+
			strconv.Itoa(int(n.min)))
	}
}

// maxRule checks if field is bigger than the max value.
func (n *numeric) maxRule() {
	if n.value > n.max {
		n.Errs = append(n.Errs, "the field "+n.FieldName+" should be less than "+
			strconv.Itoa(int(n.max)))
	}
}

// requiredRule check if field is defined.
func (n *numeric) requiredRule() {
	if n.value == 0 {
		n.Errs = append(n.Errs, "the field "+n.FieldName+" should not be empty")
	}
}
