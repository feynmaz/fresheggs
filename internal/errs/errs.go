package errs

type ErrNoData struct {
	msg string
}

func (e ErrNoData) Error() string {
	return e.msg
}

func NewErrNoData(msg string) ErrNoData {
	return ErrNoData{msg: msg}
}

type ErrBadRequest struct {
	msg string
}

func (e ErrBadRequest) Error() string {
	return e.msg
}

func NewErrBadRequest(msg string) ErrBadRequest {
	return ErrBadRequest{msg: msg}
}
