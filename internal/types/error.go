package types

type ErrorResponse struct {
	RequestID  string `json:"requestID"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

//

type ErrNotFound struct {
	msg string
}

func (e ErrNotFound) Error() string {
	return e.msg
}

func NewErrNotFound(msg string) ErrNotFound {
	return ErrNotFound{msg: msg}
}
