package field

import (
	"strconv"
)

// str struct describes the string validator. A string validator is
// responsible to define the rules of validation and validate a string property.
type str struct {
	// valy.Field embedded to str validator to have access to Field's properties.
	Field

	// min defines the min character that can contains this field.
	min int

	// max defines the max character that can contains this field.
	max int

	// require defines if the field has to be set.
	required bool

	// value is the value of the field.
	value string
}

// newString initializes and returns a str.
func newString(fp *Field) *str {
	nv := &str{
		min:      -1,
		max:      -1,
		required: false,
	}
	nv.Field = *fp
	return nv
}

// validate is responsible to validate this field. After this call
// the function will return the errors if the field is invalid.
func (n *str) validate() ([]string, error) {
	v := n.Field.Value.(string)
	n.value = v
	err := n.setRules(n.Field.Rules)
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

// setRules sets the rules for the current field.
func (n *str) setRules(rules map[string]string) error {
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
			return err
		}
	}
	return nil
}

// minRule checks if field contains less than X characters.
func (n *str) minRule() {
	if len(n.value) < n.min {
		n.Field.Errs = append(n.Field.Errs, "the field "+n.Field.FieldName+" should contains at least "+
			strconv.Itoa(n.min)+" characters")
	}
}

// maxRule checks if field contains more than X characters.
func (n *str) maxRule() {
	if len(n.value) > n.max {
		n.Field.Errs = append(n.Field.Errs, "the field "+n.Field.FieldName+" should contains max "+
			strconv.Itoa(n.max)+" characters")
	}
}

// requiredRule check if field is defined.
func (n *str) requiredRule() {
	if n.value == "" {
		n.Field.Errs = append(n.Field.Errs, "the field "+n.Field.FieldName+" should not be empty")
	}
}
