package main

import (
	"errors"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

// define the business logic interface
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// implements the business logic interface
// if underlying data structure isn't complex, ...
// if underlying data structure is complex, use Entity of OOP
type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {

	if strings.TrimSpace(s) == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

var ErrEmpty = errors.New("Empty string")

var _ StringService = stringService{}

// use RPC
// define request and response struct
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty""` // errors don't JSON-marshal, use string instead
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

// use endpoint provided by Go-Kit
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {

}

func main() {

}
