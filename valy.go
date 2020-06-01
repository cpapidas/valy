package valy

import (
	"encoding/json"
	"reflect"
	"strings"
)

// Validate gets two parameters the data (required) which is the struct of data to validate and the customErrors which
// is an optional parameters of map[string]string. The function will return map[string][]string object. The map key
// is the name of the struct field and the array of strings defined the errors.
//
// HOW TO USE IT
// Define and initialize the demoUser struct and call the Validate function:
// type demoUser struct {
//  	Username string `validate:"required=true,min=10,max=55"`
//  	Password string `validate:"required=true,err="Password is required"`
//  	Age      int   `validate:"required=true,min=10,max=99"`
//  }
// du := demoUser{
// 	Username: "cpapidas",
//  Password: "123",
//  Age: 5
// }
// errs := valy.Validate(du)
//
// ERROR MESSAGES
// The Valy validator give us three ways to set the errors messages:
// 	*) The defaults. Valy contains default error messages for all supported validations.
//  *) Custom error message using the property err="The error message"
//  *) Add the optional parameter customErrors map[string]string when we call the Validate function. Custom errors give
// 		us the flexibility to have translatable error messages. The map key is the name of the field and the string
// 		value is the error message.
//
// CUSTOM ERROR MESSAGES
// Define and initialize the demoUser struct, set the map of custom messages, and call the Validate function:
// type demoUser struct {
//  	Username string `validate:"required=true,min=10,max=55"`
//  	Password string `validate:"required=true`
//  	Age      int   `validate:"required=true,min=10,max=99"`
//  }
// du := demoUser{
// 	Username: "cpapidas",
//  Password: "123",
//  Age: 5
// }
// customErr := map[string]string{
//  "Username": "Username should be 55 > username > 10"
//  "Password": "Password is required"
//  "Age": "Age should be 99 > age > 10"
// }
// errs := valy.Validate(du, customErr)
func Validate(data interface{}, customErrors ...map[string]string) map[string][]string {
	return v(data, customErrors...)
}

// JValidate gets two parameters the data (required) which is the struct of data to validate and the customErrors which
// is an optional parameters of map[string]string. The function will return the errors as JSON []byte.
//
// HOW TO USE IT
// Define and initialize the demoUser struct and call the JValidate function:
// type demoUser struct {
//  	Username string `validate:"required=true,min=10,max=55"`
//  	Password string `validate:"required=true,err="Password is required"`
//  	Age      int   `validate:"required=true,min=10,max=99"`
//  }
// du := demoUser{
// 	Username: "cpapidas",
//  Password: "123",
//  Age: 5
// }
// errs := valy.Validate(du)
//
// ERROR MESSAGES
// The Valy validator give us three ways to set the errors messages:
// 	*) The defaults. Valy contains default error messages for all supported validations.
//  *) Custom error message using the property err="The error message"
//  *) Add the optional parameter customErrors map[string]string when we call the JValidate function. Custom errors give
// 		us the flexibility to have translatable error messages. The map key is the name of the field and the string
// 		value is the error message.
//
// CUSTOM ERROR MESSAGES
// Define and initialize the demoUser struct, set the map of custom messages, and call the Validate function:
// type demoUser struct {
//  	Username string `validate:"required=true,min=10,max=55"`
//  	Password string `validate:"required=true`
//  	Age      int   `validate:"required=true,min=10,max=99"`
//  }
// du := demoUser{
// 	Username: "cpapidas",
//  Password: "123",
//  Age: 5
// }
// customErr := map[string]string{
//  "Username": "Username should be 55 > username > 10"
//  "Password": "Password is required"
//  "Age": "Age should be 99 > age > 10"
// }
// errs := valy.JValidate(du, customErr)
func JValidate(data interface{}, customErrors ...map[string]string) []byte {
	errs := v(data, customErrors...)
	if len(errs) == 0 {
		return nil
	}

	s, _ := json.Marshal(errs)
	return s
}

// v is the private function which starts the validation. It gets two parameters the data which is the struct and the
// optional parameter customErrors which is a map[fieldName]errorStringMessage. It returns the errors of all fields
// with a map[string][]string
func v(data interface{}, customErrors ...map[string]string) map[string][]string {
	var ce map[string]string
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if len(customErrors) > 0 {
		ce = customErrors[0]
	}
	return parseFields(t, v, ce)
}

// parseFields parses all the struct field and field the "annotation" of validation. It get the fields and creates the
// fieldProperties object of each of them. Finally it collects all the errors from validators and return them as
// map[fieldName]errorStringMessage
func parseFields(t reflect.Type, v reflect.Value, ce map[string]string) map[string][]string {
	var errs = make(map[string][]string)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}
		fp := &fieldProperties{
			kind:         t.Field(i).Type.String(),
			value:        v.Field(i).Interface(),
			fieldName:    t.Field(i).Name,
			customErrors: ce,
		}
		fp.applyRules(strings.Split(tag, ","))
		if valErrs := fp.callValidator(); len(valErrs) > 0 {
			errs[fp.fieldName] = valErrs
		}

	}
	return errs
}

// ivalidator describes the validator providers interfaces.
type ivalidator interface {

	// validate validates the fields the set the fieldProperties errors
	validate()

	// customErrors check if there are defined any custom errors and returns them. Custom errors have bigger
	// priority than in line errors and default errors.
	customErrors() []string
}

// validator struct embeds the fileProperties and is responsible to validate the fields
type validator struct {
	fieldProperties
}

// customErrors check if there are defined any custom errors and returns them. Custom errors have bigger
// priority than in line errors and default errors.
func (val *validator) customErrors() []string {
	return val.errs
}
