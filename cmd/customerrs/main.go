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
		Age:      11,
	}

	errMess := map[string]string{
		"Username": "Username is required and should contain between 10 and 23 characters.",
	}
	errs := valy.Validate(u, errMess)
	if errs != nil {
		fmt.Println(errs)
	}
}
