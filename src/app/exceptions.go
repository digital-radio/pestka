//Package app implements http server api.
package app

//AppError is an error thrown by app with code and message that should be returned in http response
type AppError struct {
	Err     error
	Message string
	Code    int
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Error() string { return e.Message }
