/*
Main string service definition file
*/
package main

import (
	"errors"
	"strings"
)

/*StringService ...*/
type StringService interface {
	Uppercase(string) (string, error)
	Lowercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

/*Function implementations*/
func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}

	return strings.ToUpper(s), nil
}

func (stringService) Lowercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}

	return strings.ToLower(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty ...
var ErrEmpty = errors.New("Empty String")

// ServiceMiddleware ...
type ServiceMiddleware func(StringService) StringService
