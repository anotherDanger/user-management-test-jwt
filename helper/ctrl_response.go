package helper

import "user-management-test/web"

func NewResponse[T any](code int, message string, data T) *web.WebResponse[T] {
	return &web.WebResponse[T]{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
