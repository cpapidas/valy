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

	errs := valy.JValidate(u)
	if errs != nil {
		// return the json object to the client
		fmt.Println("Validation errors", string(errs))
	}
}
