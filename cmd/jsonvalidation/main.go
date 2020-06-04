package main

import (
	"fmt"
	"github.com/cpapidas/valy"
)

type user struct {
	Username string `json:"username" validate:"required=true,min=10,max=23"`
	Password string `json:"password" validate:"required=true,err=password is required"`
	Age      int    `validate:"required=true,min=10,max=23"`
}

func main() {
	u := user{
		Username: "cpapidas",
		Age:      3,
	}

	validationErrs, err := valy.JValidate(u)
	if err != nil {
		fmt.Println(err)
	}
	if validationErrs != nil {
		// return the json object to the client
		fmt.Println("Validation errors", string(validationErrs))
	}
}
