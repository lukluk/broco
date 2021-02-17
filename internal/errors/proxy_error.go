package errors

import (
	"fmt"
)

type ProxyError struct {
	code string
	message string
}

func NewRequestError() ProxyError {
	return ProxyError{
		code:    "400",
		message: "Request Error",
	}
}

func NewGenericProxyError(message string) ProxyError {
	return ProxyError{
		code:    "901",
		message: message,
	}
}

func (p ProxyError) Error() string {
	return fmt.Sprintf("code: %s, message: %s", p.code, p.message)
}

func (p ProxyError) Code() string {
	return p.code
}

func (p ProxyError) Message() string {
	return p.message
}
