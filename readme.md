# Valy Validator

[![codebeat badge](https://codebeat.co/badges/5125c3a9-6723-43fe-b157-77fca39c7a77)](https://codebeat.co/projects/github-com-cpapidas-valy-master)
[![Build Status](https://travis-ci.org/cpapidas/valy.svg?branch=master)](https://travis-ci.org/cpapidas/valy)
[![codecov](https://codecov.io/gh/cpapidas/valy/branch/master/graph/badge.svg)](https://codecov.io/gh/cpapidas/valy)

Valy is a struct validator manager that helps you to validate the fields according to the predefined properties. It can
validate any struct and return the errors as types of map[string]string or JSON. Valy fully supports custom errors per field.

# Installation

`go get github.com/cpapidas/valy`

# Usage

Simple Example
```go
type user struct {
	Username string `validate:"required=true,min=10,max=23"`
	Password string `validate:"required=true,err=password is required"`
	Age      int    `validate:"required=true,min=10,max=23"`
}

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
```

JSON Example
```go
type user struct {
	Username string `json:"username" validate:"required=true,min=10,max=23"`
	Password string `json:"password" validate:"required=true,err=password is required"`
	Age      int    `validate:"required=true,min=10,max=23"`
}

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
```

Custom Errors Example
```go
type user struct {
	Username string `json:"username" validate:"required=true,min=10,max=23"`
	Password string `json:"password" validate:"required=true,err=password is required"`
	Age      int    `validate:"required=true,min=10,max=23"`
}

u := user{
    Username: "cpapidas",
    Age:      11,
}

errMess := map[string]string{
    "Username": "Username is required and should contain between 10 and 23 characters.",
}
validationErrs, err := valy.Validate(u, errMess)
if err != nil {
    fmt.Println(err)
}
if validationErrs != nil {
    fmt.Println(validationErrs)
}
```

# Supported Validators

### string

```go
type user struct {
	Field1       string  `validate:"required=true"`   
	Field2       string  `validate:"min=10"`          
	Field3       string  `validate:"max=23"`          
	Field4       string  `validate:"max=23,err=Just a custom error"`
}
```

### numeric

```go
type user struct {
	Field1       int      `validate:"required=true"`   
	Field2       float32  `validate:"min=10"`          
	Field3       uint     `validate:"max=23"`          
	Field4       unit8    `validate:"max=23,err=Just a custom error"`
}
```