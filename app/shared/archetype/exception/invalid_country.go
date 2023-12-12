package exception

type InvalidCountry struct {
}

func (e InvalidCountry) Error() string {
	return "INVALID_COUNTRY"
}
