package common

type CommonError struct {
	Message string
}

func (c CommonError) Error() string {
	return c.Message
}
