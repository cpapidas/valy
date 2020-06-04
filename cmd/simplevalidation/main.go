package main

import (
	"fmt"
	"github.com/cpapidas/valy"
)

type user struct {
	Username string `validate:"required=true,min=10,max=23"`
	Password string `validate:"required=true,err=password is required"`
	Age      int    `validate:"required=true,min=10,max=23"`
}

func main() {
	u := user{
		Username: "cpapidas",
		Age:      9,
	}
	validationErrs, err := valy.Validate(u)
	if err != nil {
		fmt.Println(err)
	}
	if validationErrs != nil {
		fmt.Println(validationErrs)
	}
}
