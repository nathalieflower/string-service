package main

import (
	"testing"
)

const TestStringUpper = "HELLO WORLD"
const TestStringLower = "hello world"
const TestStringFixedLength = "fixed length string" //18chars

func TestUppercase(t *testing.T) {
	var svc StringService
	svc = stringService{}

	response, err := svc.Uppercase(TestStringLower)
	if err != nil {
		t.Error(err)
	}
	if response != TestStringUpper {
		t.Errorf("String was not in the correct format. Expecting: '%v' Actual: '%v'", TestStringUpper, response)
	}
}

func TestUppercaseEmtpyString(t *testing.T) {
	var emptyString string

	var svc StringService
	svc = stringService{}

	response, err := svc.Uppercase(emptyString)
	if response != "" {
		t.Errorf("Unexpected response. Expecting empty string, Actual: '%v'", response)
	}
	if err == nil {
		t.Error("Expecting error. No error thrown.")
	}
	if err != ErrEmpty {
		t.Errorf("Unexpected error thrown. Expecting: '%v', Actual: '%v'", ErrEmpty, err)

	}
}

func TestLowercase(t *testing.T) {
	var svc StringService
	svc = stringService{}

	response, err := svc.Lowercase(TestStringUpper)
	if err != nil {
		t.Error(err)
	}
	if response != TestStringLower {
		t.Errorf("String was not in the correct format. Expecting: '%v' Actual: '%v'", TestStringLower, response)
	}
}

func TestLowercaseEmtpyString(t *testing.T) {
	var emptyString string

	var svc StringService
	svc = stringService{}

	response, err := svc.Lowercase(emptyString)
	if response != "" {
		t.Errorf("Unexpected response. Expecting empty string, Actual: '%v'", response)
	}
	if err == nil {
		t.Error("Expecting error. No error thrown.")
	}
	if err != ErrEmpty {
		t.Errorf("Unexpected error thrown. Expecting: '%v', Actual: '%v'", ErrEmpty, err)

	}
}

func TestCount(t *testing.T) {
	var svc StringService
	svc = stringService{}

	response := svc.Count(TestStringFixedLength)

	if response != len(TestStringFixedLength) {
		t.Errorf("Incorrect value returned. Expecting: '%v' Actual: '%v'", len(TestStringFixedLength), response)
	}
}
