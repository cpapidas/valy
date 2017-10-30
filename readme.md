# Valy Validator

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