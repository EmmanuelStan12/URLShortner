package errors

type Error struct {
	code int
	err  error
}

func newError(code int, err error) *Error {
	return &Error{
		code: code,
		err:  err,
	}
}

func NotFoundError(err error) *Error {
	return newError(404, err)
}

func InternalServerError(err error) *Error {
	return newError(500, err)
}

func BadRequestError(err error) *Error {
	return newError(404, err)
}

func UnauthorizedError(err error) *Error {
	return newError(401, err)
}
