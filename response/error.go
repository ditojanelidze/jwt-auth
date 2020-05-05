package response

type Error struct {
	Message string
	Code int
	Status string
}

func BadRequestError(message string) *Error{
	return &Error{
		Message: message,
		Code: 400,
		Status: "bad_request",
	}
}

func UnautorizedError(message string) * Error {
	return &Error{
		Message: message,
		Code: 401,
		Status: "unauthorized",
	}
}
