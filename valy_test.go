package valy_test

import (
	"github.com/cpapidas/valy"
	"testing"
	"fmt"
)

type demoUser struct {
	Username       string  `validate:"required=true,min=10,max=23"`
	Password       string  `validate:"required=true,err=password is required"`
	Email          string  `validate:"required=false"`
	Age            int     `validate:"required=true,min=10,max=23"`
	Country        string  `validate:""`
	FavoriteNumber uint8   `validate:"max=9"`
	PostCode1      uint16  `validate:"required=true"`
	PostCode2      uint32  `validate:"required=true"`
	PostCode3      uint64  `validate:"required=true"`
	PostCode4      int8    `validate:"required=true"`
	PostCode5      int16   `validate:"required=true"`
	PostCode6      int32   `validate:"required=true"`
	PostCode7      int64   `validate:"required=true"`
	PostCode8      uint    `validate:"required=true"`
	Results1       float32 `validate:"required=true"`
	Results2       float64 `validate:"required=true"`
	Phone          string  `validate:"max=10"`
	CustomError    string  `validate:"max=10"`
	Married        bool    `validate:"required=true"`
}

func TestValy(t *testing.T) {
	u := demoUser{
		Username:       "cpapidas",
		Age:            0,
		FavoriteNumber: 10,
		Phone:          "1234567891011",
		CustomError:    "AbigStringAbigString",
	}
	customErrs := map[string]string{"CustomError": "It's just a custom error"}
	errs := valy.Validate(u, customErrs)
	fmt.Println(errs["Rate"])
	if len(errs["Username"]) != 1 {
		t.Log("Should return errors for invalid username.")
		t.Fail()
	}
	if errs["Username"][0] == "" {
		t.Log("Error should not be nil.")
		t.Fail()
	}
	if len(errs["Password"]) != 1 {
		t.Log("Should return errors for invalid password.")
		t.Fail()
	}
	if errs["Password"][0] != "password is required" {
		t.Log("Should return the error 'password is required'.")
		t.Fail()
	}
	if len(errs["Email"]) > 0 {
		t.Log("Email is options should not return any error.")
		t.Fail()
	}
	if len(errs["Age"]) != 2 {
		t.Log("Should return two error about Age field.")
		t.Fail()
	}
	if errs["Age"][0] == "" {
		t.Log("Error should not be nil.")
		t.Fail()
	}
	if errs["Age"][1] == "" {
		t.Log("Error should not be nil.")
		t.Fail()
	}
	if len(errs["Country"]) != 0 {
		t.Log("Should not return any error.")
		t.Fail()
	}
	if len(errs["FavoriteNumber"]) != 1 {
		t.Log("Should return an error about invalid FavoriteNumber.")
		t.Fail()
	}
	if len(errs["Phone"]) != 1 {
		t.Log("Should return an error about invalid Phone.")
		t.Fail()
	}
	if errs["Phone"][0] == "" {
		t.Log("Error should not be nil.")
		t.Fail()
	}
	if len(errs["CustomError"]) != 1 {
		t.Log("Should return an error about invalid CustomError.")
		t.Fail()
	}
	if errs["CustomError"][0] != "It's just a custom error" {
		t.Log("Error should be `It's just a custom error`.")
		t.Fail()
	}
}

type demoInvalidNumericValidation struct {
	Demo int16 `validate:"min=invalid min number"`
}

func TestValy_invalidNumericValidator(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	u := demoInvalidNumericValidation{Demo: 0}
	valy.Validate(u)
}

type demoInvalidStringValidation struct {
	Demo string `validate:"min=invalid min number"`
}

func TestValy_invalidStringValidator(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	u := demoInvalidStringValidation{Demo: ""}
	valy.Validate(u)
}

type demoJsonValidation struct {
	Demo string `json:"demo" validate:"min=10"`
}

func TestValy_jsonValidation(t *testing.T) {
	d := demoJsonValidation{Demo: "123"}
	errs := valy.JValidate(d)
	if string(errs) != `{"Demo":["the field Demo should contains at least 10 characters"]}` {
		t.Log("Should return a valid error message")
		t.Fail()
	}
	d = demoJsonValidation{Demo: "12345678900"}
	errs = valy.JValidate(d)
	if errs != nil {
		t.Log("Should return nil")
		t.Fail()
	}
}

func TestValy_invalidJsonValidation(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	valy.JValidate(1)
}


