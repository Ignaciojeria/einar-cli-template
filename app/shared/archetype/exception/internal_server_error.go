package exception

type InternalServerError struct {
}

func (e InternalServerError) Error() string {
	return "INTERNAL_SERVER_ERROR"
}
