package exception

type Exception struct {
	Code    int
	Err     error
	Message string
	Details interface{}
}

func (ex *Exception) Error() string {
	return ex.Message
}

func New(code int, err error, msg string, details interface{}) *Exception {
	return &Exception{
		Code:    code,
		Err:     err,
		Message: msg,
		Details: details,
	}
}
