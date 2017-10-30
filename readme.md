# Valy Validator

[![codebeat badge](https://codebeat.co/badges/5125c3a9-6723-43fe-b157-77fca39c7a77)](https://codebeat.co/projects/github-com-cpapidas-valy-master)
[![Build Status](https://travis-ci.org/cpapidas/valy.svg?branch=master)](https://travis-ci.org/cpapidas/valy)
[![codecov](https://codecov.io/gh/cpapidas/valy/branch/master/graph/badge.svg)](https://codecov.io/gh/cpapidas/valy)

Valy Validator is a struct validator manager which validates the fields according to predefined properties. It can
validate the JSON struct and returns the errors as JSON string. It also supports custom errors message.

# Installation

`go get github.com/cpapidas/valy`

# Usage

Simple Example
```go
type demoUser struct {
	Username       string  `validate:"required=true,min=10,max=23"`
	Password       string  `validate:"required=true,err=password is required"`
	Age            int     `validate:"required=true,min=10,max=23"`
}
u := demoUser{
    Username:       "cpapidas",
    Age:            0,
}
errs := valy.Validate(u)
if errs != nil {
    fmt.Println(errs)
}
```

JSON Example
```go
type demoUser struct {
	Username       string  `json:"username" validate:"required=true,min=10,max=23"`
	Password       string  `json:"password" validate:"required=true,err=password is required"`
	Age            int     `validate:"required=true,min=10,max=23"`
}
u := demoUser{
    Username:       "cpapidas",
    Age:            0,
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
```

Custom Errors Example
```go
type demoUser struct {
	Username       string  `json:"username" validate:"required=true,min=10,max=23"`
	Password       string  `json:"password" validate:"required=true,err=password is required"`
	Age            int     `validate:"required=true,min=10,max=23"`
}
u := demoUser{
    Username:       "cpapidas",
    Age:            0,
}

errMess := map[string]string {
    "Username": "Username is required and should contain between 10 and 23 characters.",
}

jerrs := valy.JValidate(u, errMess)
fmt.Println(string(jerrs))

errs := valy.Validate(u, errMess)
fmt.Println(errs)
```

You can run all the example above by
`$ go run examples/examples.go`

# Supported Validators

### string

```go
type demoUser struct {
	Field1       string  `validate:"required=true"`   // Checks if the field presents.
	Field2       string  `validate:"min=10"`          // Checks if the string length is less than min value.
	Field3       int     `validate:"max=23"`          // Checks if the string length is bigger than max value.
	Field4       int     `validate:"max=23,err=Just a custom error"`
}
```

### numeric

```go
type demoUser struct {
	Field1       int      `validate:"required=true"`   // Checks if the field is != 0.
	Field2       float32  `validate:"min=10"`          // Checks if field is less than 10.
	Field3       uint     `validate:"max=23"`          // Checks if the field is grater than 23.
	Field4       unit8    `validate:"max=23,err=Just a custom error"`
}
```