package main

import "fmt"

type BadRequest struct {
	msg string
}

func NewBadRequest(msg string) error {
	return &BadRequest{msg: msg}
}

func (b *BadRequest) Error() string {
	return fmt.Sprintf("bad request, %s", b.msg)
}

type InternalError struct {
	msg    string
	detail error
}

func NewInternalError(msg string, err error) error {
	return &InternalError{
		msg:    msg,
		detail: err,
	}
}

func (i *InternalError) Error() string {
	return fmt.Sprintf("internal error, %s", i.msg)
}
