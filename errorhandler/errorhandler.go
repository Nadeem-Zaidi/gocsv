package errorhandler

type CustomError struct {
	Statuscode int
	Err        error
}

func (e *CustomError) Error() string {
	return e.Err.Error()

}
