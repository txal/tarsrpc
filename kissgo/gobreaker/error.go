// Package errors is an interface for defining detailed errors
package gobreaker

import (
	"encoding/json"
)

const (
	Success = 0
	Failure = 9999
)

// Errors provide a way to return detailed information
// for an RPC request error. The error is normally
// JSON encoded.
type Error struct {
	Id     string `json:"id"`
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func Parse(err string) *Error {
	e := new(Error)
	errr := json.Unmarshal([]byte(err), e)
	if errr != nil {
		e.Code = Failure
		e.Detail = err
	}
	return e
}
