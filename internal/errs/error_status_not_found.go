package errs

import "fmt"

type ErrorStatusNotFound struct {
	Msg string
}

func (e *ErrorStatusNotFound) Error() string {
	return e.Msg
}

func NewErrorStatusNotFound(orderId int64) error {
	return &ErrorStatusNotFound{fmt.Sprintf("order with id %d", orderId)}
}
