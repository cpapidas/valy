package field_test

import (
	"github.com/cpapidas/valy/field"
	"testing"
)

func TestField_CallValidator_shouldReturnErrorForInvalidStringMin(t *testing.T) {
	f := field.Field{
		Kind: "string",
		Value: "",
		FieldName: "Username",
	}
	valsErrs, err := f.CallValidator([]string{"min=10"})
	if err != nil {
		t.Fatalf("expected not return an error but got: %v", err)
	}
	expectedErr := "the field Username should contains at least 10 characters"
	if len(valsErrs) == 0 || valsErrs[0] != expectedErr {
		t.Errorf("should return the error: %s, but got nil", expectedErr)
	}
}

func TestField_CallValidator_shouldReturnErrorForInvalidStringMax(t *testing.T) {
	f := field.Field{
		Kind: "string",
		Value: "123456",
		FieldName: "Username",
	}
	valsErrs, err := f.CallValidator([]string{"max=5"})
	if err != nil {
		t.Fatalf("expected not return an error but got: %v", err)
	}
	expectedErr := "the field Username should contains max 5 characters"
	if len(valsErrs) == 0 || valsErrs[0] != expectedErr {
		t.Errorf("should return the error: %s, but got nil", expectedErr)
	}
}

func TestField_CallValidator_shouldReturnErrorForInvalidStringRequired(t *testing.T) {
	f := field.Field{
		Kind: "string",
		Value: "",
		FieldName: "Username",
	}
	valsErrs, err := f.CallValidator([]string{"required=true"})
	if err != nil {
		t.Fatalf("expected not return an error but got: %v", err)
	}
	t.Log(valsErrs)
	expectedErr := "the field Username should not be empty"
	if len(valsErrs) == 0 || valsErrs[0] != expectedErr {
		t.Errorf("should return the error: %s, but got nil", expectedErr)
	}
}

func TestField_CallValidator_shouldReturnErrorForInvalidNumberMin(t *testing.T) {
	f := field.Field{
		Kind: "int",
		Value: 10,
		FieldName: "Age",
	}
	valsErrs, err := f.CallValidator([]string{"min=11"})
	if err != nil {
		t.Fatalf("expected not return an error but got: %v", err)
	}
	t.Log(valsErrs)
	expectedErr := "the field Age should be grater than 11"
	if len(valsErrs) == 0 || valsErrs[0] != expectedErr {
		t.Errorf("should return the error: %s, but got nil", expectedErr)
	}
}

func TestField_CallValidator_shouldReturnErrorForInvalidNumberMax(t *testing.T) {
	f := field.Field{
		Kind: "int",
		Value: 12,
		FieldName: "Age",
	}
	valsErrs, err := f.CallValidator([]string{"max=11"})
	if err != nil {
		t.Fatalf("expected not return an error but got: %v", err)
	}
	t.Log(valsErrs)
	expectedErr := "the field Age should be less than 11"
	if len(valsErrs) == 0 || valsErrs[0] != expectedErr {
		t.Errorf("should return the error: %s, but got nil", expectedErr)
	}
}

func TestField_CallValidator_shouldReturnErrorForInvalidNumberRequired(t *testing.T) {
	f := field.Field{
		Kind: "int",
		Value: 0,
		FieldName: "Age",
	}
	valsErrs, err := f.CallValidator([]string{"required=true"})
	if err != nil {
		t.Fatalf("expected not return an error but got: %v", err)
	}
	t.Log(valsErrs)
	expectedErr := "the field Age should not be empty"
	if len(valsErrs) == 0 || valsErrs[0] != expectedErr {
		t.Errorf("should return the error: %s, but got nil", expectedErr)
	}
}

func TestField_CallValidator_shouldReturnCustomError(t *testing.T) {
	f := field.Field{
		Kind: "int",
		Value: 0,
		FieldName: "Age",
		CustomError: "This is a custom error",
	}
	valsErrs, err := f.CallValidator([]string{"required=true"})
	if err != nil {
		t.Fatalf("expected not return an error but got: %v", err)
	}
	t.Log(valsErrs)
	expectedErr := f.CustomError
	if len(valsErrs) == 0 || valsErrs[0] != expectedErr {
		t.Errorf("should return the error: %s, but got nil", expectedErr)
	}
}

func TestField_CallValidator_shouldReturnFailedErrorOnStringForInvalidRules(t *testing.T) {
	f := field.Field{
		Kind: "string",
		Value: "this is a string",
		FieldName: "Username",
	}
	valsErrs, err := f.CallValidator([]string{"min=rule"})
	if err == nil {
		t.Error("expected to return an error but got nil")
	}
	if valsErrs != nil {
		t.Error("expected to return nil results")
	}
}

func TestField_CallValidator_shouldReturnFailedErrorOnNumericForInvalidRules(t *testing.T) {
	f := field.Field{
		Kind: "int",
		Value: 1,
		FieldName: "Age",
	}
	valsErrs, err := f.CallValidator([]string{"min=rule"})
	if err == nil {
		t.Error("expected to return an error but got nil")
	}
	if valsErrs != nil {
		t.Error("expected to return nil results")
	}
}