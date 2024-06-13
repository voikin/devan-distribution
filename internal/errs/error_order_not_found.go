package errs

import "fmt"

type ErrorOrderNotFound struct {
	Msg string
}

func (e *ErrorOrderNotFound) Error() string {
	return e.Msg
}

func NewErrorOrderNotFound(orderId int64) error {
	return &ErrorOrderNotFound{fmt.Sprintf("order with id %d", orderId)}
}
