package http

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NotFoundErr(err error) AppError {
	return AppError{
		Code:    http.StatusNotFound,
		Message: err.Error(),
	}
}

func ServerError(err error) AppError {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
}
