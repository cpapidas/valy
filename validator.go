package valy

// ivalidator describes the validator providers interfaces.
type ivalidator interface {

	// validate validates the fields the set the fieldProperties errors
	validate()

	// checkCustomErrors check if there are defined any custom errors and returns them. Custom errors have bigger
	// priority than in line errors and default errors.
	checkCustomErrors() []string
}

// validator struct embeds the fileProperties and is responsible to validate the fields
type validator struct {
	fieldProperties
}

// checkCustomErrors check if there are defined any custom errors and returns them. Custom errors have bigger
// priority than in line errors and default errors.
func (val *validator) checkCustomErrors() []string {
	if v, ok := val.customErrors[val.fieldName]; ok {
		return []string{v}
	} else if val.err != "" {
		return []string{val.err}
	}
	return val.errs
}
