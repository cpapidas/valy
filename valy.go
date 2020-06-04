package valy

import (
	"encoding/json"
	"github.com/cpapidas/valy/field"
	"reflect"
	"strings"
)

// Validate gets two parameters the data (required) which is a struct of data to validate and the CustomErrors which
// is an optional parameters of map[string]string. The function will return a map[string][]string object. The map's key
// is the name of the property and the value is an array of strings that contains all the errors.
//
// HOW TO USE IT
// Define and initialize the demoUser struct, after it call the Validate function:
// type demoUser struct {
//  	Username string `validate:"required=true,min=10,max=55"`
//  	Password string `validate:"required=true,Err="Password is required"`
//  	Age      int    `validate:"required=true,min=10,max=99"`
//  }
// du := demoUser{
// 	Username: "cpapidas",
//  Password: "123",
//  Age: 5
// }
// errs := valy.Validate(du)
// fmt.println(errs)
//
// ERROR MESSAGES
// The Valy validator give us three ways to set the errors messages:
// 	1) The defaults. Valy contains default error messages for all supported validations.
//  2) Custom error message using the property Err="The error message" e.g. `validate:"required=true,Err="Password is required"`
//  3) Add the optional parameter CustomErrors type of map[string]string when we call the Validate function. Custom errors give
// 		us the flexibility to have translatable error messages.
//      Example of custom error messages:
//      errMess := map[string]string{
//			"Username": "Username is required and should contain between 10 and 23 characters.",
//		}
//		errs := valy.Validate(u, errMess)
func Validate(data interface{}, customErrors ...map[string]string) (map[string][]string, error) {
	return v(data, customErrors...)
}

// JValidate gets two parameters the data (required) which is a struct of data to validate and the CustomErrors which
// is an optional parameters of map[string]string. The function will return the errors as JSON []byte.
//
// HOW TO USE IT
// Define and initialize the demoUser struct and call the JValidate function:
// type demoUser struct {
//  	Username string `validate:"required=true,min=10,max=55"`
//  	Password string `validate:"required=true,Err="Password is required"`
//  	Age      int   `validate:"required=true,min=10,max=99"`
//  }
// du := demoUser{
// 	Username: "cpapidas",
//  Password: "123",
//  Age: 5
// }
// errs := valy.JValidate(du)
// fmt.println(string(errs))
//
// ERROR MESSAGES
// The Valy validator give us three ways to set the errors messages:
// 	1) The defaults. Valy contains default error messages for all supported validations.
//  2) Custom error message using the property Err="The error message" e.g. `validate:"required=true,Err="Password is required"`
//  3) Add the optional parameter CustomErrors type of map[string]string when we call the Validate function. Custom errors give
// 		us the flexibility to have translatable error messages.
//      Example of custom error messages:
//      errMess := map[string]string{
//			"Username": "Username is required and should contain between 10 and 23 characters.",
//		}
//		errs := valy.JValidate(u, errMess)
func JValidate(data interface{}, customErrors ...map[string]string) ([]byte, error) {
	validationErrs, err := v(data, customErrors...)
	if err != nil {
		return nil, err
	}
	if len(validationErrs) == 0 {
		return nil, nil
	}

	s, _ := json.Marshal(validationErrs)
	return s, nil
}

// v is the private function which starts the validation. It gets two parameters the data which is a struct and the
// optional parameter CustomErrors which is a map[FieldName]errorStringMessage.
//
// It returns the errors of all fields as a map[string][]string and if something go wrong it returns
// the nil and error.
func v(data interface{}, customErrors ...map[string]string) (map[string][]string, error) {
	var ce map[string]string
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if len(customErrors) > 0 {
		ce = customErrors[0]
	}
	return parseFields(t, v, ce)
}

// parseFields parses all the struct fields annotations. It get the fields and creates the
// fieldProperties object of each of them.
//
// Finally it collects all the errors from validator and return them as map[FieldName]errorStringMessage.
// If something go wrong it returns nil and an error message.
func parseFields(t reflect.Type, v reflect.Value, ce map[string]string) (map[string][]string, error) {
	var errs = make(map[string][]string)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}
		fp := &field.Field{
			Kind:        t.Field(i).Type.String(),
			Value:       v.Field(i).Interface(),
			FieldName:   t.Field(i).Name,
			CustomError: ce[t.Field(i).Name],
		}
		if valErrs, err := fp.CallValidator(strings.Split(tag, ",")); len(valErrs) > 0 || err != nil {
			if err != nil {
				return nil, err
			}
			errs[fp.FieldName] = valErrs
		}

	}
	return errs, nil
}
