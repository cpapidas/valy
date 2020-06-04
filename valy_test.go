package valy_test

import (
	"encoding/json"
	"github.com/cpapidas/valy"
	"reflect"
	"testing"
)

type demoUser struct {
	Username       string  `validate:"required=true,min=10,max=23"`
	Password       string  `validate:"required=true,Err=password is required"`
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
	Married        bool
}

func TestValidate_shouldReturnErrorForInvalidUsername(t *testing.T) {
	u := demoUser{
		Username:       "cpapidas",
		Age:            0,
		FavoriteNumber: 10,
		Phone:          "1234567891011",
		CustomError:    "AbigStringAbigString",
	}
	customErrs := map[string]string{"CustomError": "It's just a custom error"}
	errs, err := valy.Validate(u, customErrs)
	if err != nil {
		t.Fatalf("expected nill err but got: %v", err)
	}

	if len(errs["Username"]) != 1 {
		t.Error("should contains Username property error for invalid username")
	}
	if errs["Username"][0] == "" {
		t.Error("should contains an error in Username property for invalid username")
	}
}

func TestValidate_shouldReturnErrorForInvalidPassword(t *testing.T) {
	u := demoUser{
		Username:       "cpapidas",
		Age:            0,
		FavoriteNumber: 10,
		Phone:          "1234567891011",
		CustomError:    "AbigStringAbigString",
	}
	customErrs := map[string]string{"CustomError": "It's just a custom error"}
	errs, err := valy.Validate(u, customErrs)
	if err != nil {
		t.Fatalf("expected nill err but got: %v", err)
	}

	if len(errs["Password"]) != 1 {
		t.Error("should contains password property error for invalid password")
	}
	if errs["Password"][0] != "the field Password should not be empty" {
		t.Error("should contains an error in Password property for invalid password")
	}
}

func TestValidate_shouldNotReturnErrorForTheEmailProperty(t *testing.T) {
	u := demoUser{
		Username:       "cpapidas",
		Age:            0,
		FavoriteNumber: 10,
		Phone:          "1234567891011",
		CustomError:    "AbigStringAbigString",
	}
	customErrs := map[string]string{"CustomError": "It's just a custom error"}
	errs, err := valy.Validate(u, customErrs)
	if err != nil {
		t.Fatalf("expected nill err but got: %v", err)
	}

	if len(errs["Email"]) > 0 {
		t.Error("email is optional should not return any error")
	}
}

func TestValidate_shouldReturnErrorForTheAgeProperty(t *testing.T) {
	u := demoUser{
		Username:       "cpapidas",
		Age:            0,
		FavoriteNumber: 10,
		Phone:          "1234567891011",
		CustomError:    "AbigStringAbigString",
	}
	customErrs := map[string]string{"CustomError": "It's just a custom error"}
	errs, err := valy.Validate(u, customErrs)
	if err != nil {
		t.Fatalf("expected nill err but got: %v", err)
	}

	if len(errs["Age"]) != 2 {
		t.Error("should return two errors for the Age field")
	}
	if errs["Age"][0] == "" {
		t.Error("error should not be nil")
	}
	if errs["Age"][1] == "" {
		t.Error("error should not be nil")
	}
}

func TestValidate_shouldNotReturnErrorForTheCountryProperty(t *testing.T) {
	u := demoUser{
		Username:       "cpapidas",
		Age:            0,
		FavoriteNumber: 10,
		Phone:          "1234567891011",
		CustomError:    "AbigStringAbigString",
	}
	customErrs := map[string]string{"CustomError": "It's just a custom error"}
	errs, err := valy.Validate(u, customErrs)
	if err != nil {
		t.Fatalf("expected nill err but got: %v", err)
	}

	if len(errs["Country"]) != 0 {
		t.Error("should not return any error")
	}
}

func TestValidate_shouldReturnErrorForTheFavoriteNumberProperty(t *testing.T) {
	u := demoUser{
		Username:       "cpapidas",
		Age:            0,
		FavoriteNumber: 10,
		Phone:          "1234567891011",
		CustomError:    "AbigStringAbigString",
	}
	customErrs := map[string]string{"CustomError": "It's just a custom error"}
	errs, err := valy.Validate(u, customErrs)
	if err != nil {
		t.Fatalf("expected nill err but got: %v", err)
	}

	if len(errs["FavoriteNumber"]) != 1 {
		t.Error("should return an error about invalid FavoriteNumber")
	}
	if errs["FavoriteNumber"][0] == "" {
		t.Error("should contains an error in FavoriteNumber property for invalid favorite number")
	}
}

func TestValidate_shouldReturnErrorForThePhoneProperty(t *testing.T) {
	u := demoUser{
		Username:       "cpapidas",
		Age:            0,
		FavoriteNumber: 10,
		Phone:          "1234567891011",
		CustomError:    "AbigStringAbigString",
	}
	customErrs := map[string]string{"CustomError": "It's just a custom error"}
	errs, err := valy.Validate(u, customErrs)
	if err != nil {
		t.Fatalf("expected nill err but got: %v", err)
	}

	if len(errs["Phone"]) != 1 {
		t.Error("should return an error for invalid Phone")
	}
	if errs["Phone"][0] == "" {
		t.Error("should contains an error in Phone property for invalid phone")
	}
}

func TestValidate_shouldReturnCustomError(t *testing.T) {
	u := demoUser{
		Username:       "cpapidas",
		Age:            0,
		FavoriteNumber: 10,
		Phone:          "1234567891011",
		CustomError:    "AbigStringAbigString",
	}
	customErrs := map[string]string{"CustomError": "It's just a custom error"}
	errs, err := valy.Validate(u, customErrs)
	if err != nil {
		t.Fatalf("expected nill err but got: %v", err)
	}

	if len(errs["CustomError"]) != 1 {
		t.Error("should return an error for invalid CustomError")
	}
	if errs["CustomError"][0] != "It's just a custom error" {
		t.Errorf("expected `It's just a custom error`, but got: `%s`", errs["CustomError"][0])
	}
}

type demoInvalidNumericValidation struct {
	Demo int16 `validate:"min=invalid min number"`
}

func TestValidate_invalidNumericValidator(t *testing.T) {
	u := demoInvalidNumericValidation{Demo: 0}
	_, err := valy.Validate(u)
	if err == nil {
		t.Error("error should not be nil")
	}
}

type demoInvalidStringValidation struct {
	Demo string `validate:"min=invalid min number"`
}

func TestValidate_invalidStringValidator(t *testing.T) {
	u := demoInvalidStringValidation{Demo: ""}
	_, err := valy.Validate(u)
	if err == nil {
		t.Error("error should not be nil")
	}
}

type demoJsonValidation struct {
	Demo string `json:"demo" validate:"min=10"`
}

func TestValidate_jsonValidation(t *testing.T) {
	d := demoJsonValidation{Demo: "123"}
	errs, err := valy.JValidate(d)
	if err != nil {
		t.Fatalf("expected nill err but got: %v", err)
	}
	if string(errs) != `{"Demo":["the field Demo should contains at least 10 characters"]}` {
		t.Log("Should return a valid error message")
		t.Fail()
	}
	d = demoJsonValidation{Demo: "12345678900"}
	errs, err = valy.JValidate(d)
	if err != nil {
		t.Fatalf("expected nill err but got: %v", err)
	}
	if errs != nil {
		t.Log("Should return nil")
		t.Fail()
	}
}

func TestValidate_json(t *testing.T) {
	var djvO demoJsonValidation
	func(djv interface{}) {
		djvJson := []byte(`{"demo":"12312312311123"}`)
		err := json.Unmarshal(djvJson, djv)
		if err != nil {
			t.Log("unmarshal error", err)
			t.Fail()
		}
		t.Log("struct", djv)
		k := reflect.Indirect(reflect.ValueOf(djv)).Interface()
		errs, err := valy.JValidate(k)
		if err != nil {
			t.Fatalf("expected nill err but got: %v", err)
		}
		if errs != nil {
			t.Log("validation should not returns any errors", string(errs))
			t.Fail()
		}
	}(&djvO)
}
