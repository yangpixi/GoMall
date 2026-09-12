package errs

type BusinessError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *BusinessError) Error() string {
	return e.Msg
}

func New(code int, msg string) *BusinessError {
	return &BusinessError{
		Code: code,
		Msg:  msg,
	}
}

// common errors
var (
	ErrBadRequest          = New(400, "invalid request")
	ErrUnauthorized        = New(401, "authentication required")
	ErrPermissionDenied    = New(403, "permission denied")
	ErrInternalServerError = New(500, "internal server error")
)
