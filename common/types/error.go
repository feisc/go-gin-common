package types

// api error
type APIError struct {
	From string
	Err  error
}

func (e APIError) Error() string {
	return e.Message()
}

func (e APIError) Message() string {
	return e.Err.Error()
}
