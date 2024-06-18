package errors

type Error struct {
	Code int
	Err  any
}

func newError(code int, err any) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}

func NotFoundError(err any) *Error {
	return newError(404, err)
}

func InternalServerError(err any) *Error {
	return newError(500, err)
}

func BadRequestError(err any) *Error {
	return newError(400, err)
}

func UnauthorizedError(err any) *Error {
	return newError(401, err)
}

func ValidationError(errors []string) *Error {
	return newError(400, errors)
}
