package exception

type AuthorizationError struct {
	AuthorizationDetails string
}

func (err AuthorizationError) Error() string {
	return err.AuthorizationDetails
}
