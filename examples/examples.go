package main

import (
	"fmt"
	"github.com/cpapidas/valy"
)

func main() {
	fmt.Println("----- Simple Example -----")
	simpleExample()
	fmt.Println("----- JSON Example -----")
	jsonExample()
	fmt.Println("----- Custom Errors Example -----")
	customErrorsExample()
}

func simpleExample() {
	type demoUser struct {
		Username string `validate:"required=true,min=10,max=23"`
		Password string `validate:"required=true,err=password is required"`
		Age      int    `validate:"required=true,min=10,max=23"`
	}
	u := demoUser{
		Username: "cpapidas",
		Age:      0,
	}
	errs := valy.Validate(u)
	if errs != nil {
		fmt.Println(errs)
	}
}

func jsonExample() {
	type demoUser struct {
		Username string `json:"username" validate:"required=true,min=10,max=23"`
		Password string `json:"password" validate:"required=true,err=password is required"`
		Age      int    `validate:"required=true,min=10,max=23"`
	}
	u := demoUser{
		Username: "cpapidas",
		Age:      0,
	}

	errs := valy.JValidate(u)

	if errs != nil {
		// return the json object to the client
		// ...
		fmt.Println("Validation errors", string(errs))
	} else {
		// the data are valid
		fmt.Println("Valid data")
	}
}

func customErrorsExample() {
	type demoUser struct {
		Username string `json:"username" validate:"required=true,min=10,max=23"`
		Password string `json:"password" validate:"required=true,err=password is required"`
		Age      int    `validate:"required=true,min=10,max=23"`
	}
	u := demoUser{
		Username: "cpapidas",
		Age:      0,
	}

	errMess := map[string]string{
		"Username": "Username is required and should contain between 10 and 23 characters.",
	}

	jerrs := valy.JValidate(u, errMess)
	fmt.Println(string(jerrs))

	errs := valy.Validate(u, errMess)
	fmt.Println(errs)
}
