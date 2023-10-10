package exceptions

type InvalidArgumentError struct {
	Message string
}

func (e *InvalidArgumentError) Error() string {
	return e.Message
}
